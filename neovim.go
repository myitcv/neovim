// Copyright 2014 Paul Jolly <paul@myitcv.org.uk>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package neovim implements support for writing Neovim plugins in Go. It also
implements a tool for generating the MSGPACK-based API against a Neovim instance.

All API methods are supported, as are notifications. See Subscription for an example
of how to register a subscription on a given topic.

Client

Everything starts from Client:

	_, err := neovim.NewUnixClient("unix", nil, &net.UnixAddr{Name: "/tmp/neovim"})
	if err != nil {
		log.Fatalf("Could not create new Unix client: %v", errgo.Details(err))
	}

See the examples for further usage patterns.

Concurrency

A single Client may safely be used by multiple goroutines. Calls to API methods are blocking
by design.

Generating the API

See the github repo for details on re-generating the API.

Compatibility

There are currently no checks to verify a connected Neovim instance exposes the same API
against which the neovim package was generated. This is future work (and probably needs
some work on the Neovim side).

Errors

Errors returned by this package are created using errgo at http://godoc.org/github.com/juju/errgo.
Hence errors may be inspected using functions like errgo.Details for example:

	_, err := client.GetCurrentBuffer()
	if err != nil {
		log.Fatalf("Could not get current buffer: %v", errgo.Details(err))
	}
*/
package neovim

import (
	"log"
	"net"
	"sync/atomic"

	"github.com/juju/errgo"
	"github.com/juju/errors"
	"github.com/vmihailenco/msgpack"
)

// NewUnixClient is a convenience method for creating a new *Client. Method signature matches
// that of net.DialUnix
func NewUnixClient(_net string, laddr, raddr *net.UnixAddr) (*Client, error) {
	c, err := net.DialUnix(_net, laddr, raddr)
	if err != nil {
		return nil, errgo.Notef(err, "Could not establish connection to Neovim, _net %v, laddr %v, %v", _net, laddr, raddr)
	}
	return NewClient(c)
}

// NewClient creates a new Client
func NewClient(c net.Conn) (*Client, error) {
	res := &Client{conn: c}
	res.respMap = newSyncMap()
	res.dec = msgpack.NewDecoder(c)
	res.enc = msgpack.NewEncoder(c)
	res.SubChan = make(chan Subscription)
	res.UnsubChan = make(chan Subscription)
	go res.doListen()
	return res, nil
}

func (c *Client) doListen() {
	// TODO need kill channel

	// TODO look at the semantics of making this buffered...
	subEvents := make(chan SubscriptionEvent, 10)
	go c.doSubscriptionManager(subEvents)

	dec := c.dec
	for {
		_, err := dec.DecodeSliceLen()
		if err != nil {
			log.Fatalf("Could not decode message slice length: %v", err)
		}

		t, err := dec.DecodeInt()
		if err != nil {
			log.Fatalf("Could not decode message type: %v", err)
		}

		switch t {
		case 1:
			// handle response
			reqID, err := dec.DecodeUint32()
			if err != nil {
				log.Fatalf("Could not decode request id: %v", err)
			}

			// do we have an error?
			re, err := dec.DecodeInterface()
			if err != nil {
				log.Fatalf("Could not decode response error: %v", err)
			}
			if re != nil {
				log.Fatalf("Got a response error: %v", re)
			}

			// no, carry on
			rh, err := c.respMap.Get(reqID)
			if err != nil {
				log.Fatalf("Could not get response holder for %v: %v", reqID, err)
			}

			// we have a valid response, dispatch to our decoder for the response
			res, err := rh.dec()
			if err != nil {
				log.Fatalf("Could not decode response: %v\n", err)
			}

			resp := &response{obj: res, err: nil}
			rh.ch <- resp
		case 2:
			// handle notification
			topic, err := dec.DecodeString()
			if err != nil {
				log.Fatalf("Could not decode topic: %v", err)
			}

			// TODO this could be more efficient?
			obj, err := dec.DecodeInterface()
			if err != nil {
				log.Fatalf("Could not decode obj payload: %v", err)
			}

			ev := SubscriptionEvent{
				Topic: topic,
				Value: obj,
			}

			subEvents <- ev
		default:
			log.Fatalf("Unexpected type of message: %v\n", t)
		}
	}
}

func (c *Client) doSubscriptionManager(se chan SubscriptionEvent) {
	subs := make(map[string]map[chan SubscriptionEvent]struct{})

	sendOrClose := func(c chan error, e error) {
		if c != nil {
			if e != nil {
				c <- e
			} else {
				close(c)
			}
		}
	}

	for {
		select {
		case event := <-se:
			// TODO should we really swallow events on topics for which we have no subs?
			if chans, ok := subs[event.Topic]; ok {
				for k := range chans {
					k <- event
				}
			} else {
				log.Printf("Got an event for which we have no subs on topic %v\n", event.Topic)
			}
		case sub := <-c.SubChan:
			m, ok := subs[sub.Topic]
			if !ok {
				m = make(map[chan SubscriptionEvent]struct{})
				subs[sub.Topic] = m
			}
			if _, ok := m[sub.Events]; ok {
				sendOrClose(sub.Error, errors.Errorf("Already have subscription for topic %v on this channel", sub.Topic))
			}
			m[sub.Events] = struct{}{}
			sendOrClose(sub.Error, nil)
		case unsub := <-c.UnsubChan:
			m, ok := subs[unsub.Topic]
			if !ok {
				sendOrClose(unsub.Error, errors.Errorf("We don't have any subscriptions for topic %v", unsub.Topic))
			}
			if _, ok := m[unsub.Events]; !ok {
				sendOrClose(unsub.Error, errors.Errorf("We don't have a subscription on topic %v on this channel", unsub.Topic))
			}
			delete(m, unsub.Events)
			sendOrClose(unsub.Error, nil)
		}
	}
}

func (c *Client) makeCall(reqMethID neovimMethodID, e encoder, d decoder) (chan *response, error) {
	reqType := 0
	reqID := c.nextReqID()
	enc := c.enc

	res := make(chan *response)
	rh := &responseHolder{dec: d, ch: res}
	err := c.respMap.Put(reqID, rh)
	if err != nil {
		return nil, errgo.NoteMask(err, "Could not store response holder")
	}

	err = enc.EncodeSliceLen(4)
	if err != nil {
		return nil, errgo.NoteMask(err, "Could not encode request length")
	}

	err = enc.EncodeInt(reqType)
	if err != nil {
		return nil, errgo.NoteMask(err, "Could not encode request type")
	}

	err = enc.EncodeUint32(reqID)
	if err != nil {
		return nil, errgo.NoteMask(err, "Could not encode request ID")
	}

	err = enc.EncodeUint32(uint32(reqMethID))
	if err != nil {
		return nil, errgo.NoteMask(err, "Could not encode request method ID")
	}

	err = e()
	if err != nil {
		return nil, errgo.NoteMask(err, "Could not encode method args ")
	}

	// TODO need a flush here?

	return res, nil
}

func (c *Client) nextReqID() uint32 {
	return atomic.AddUint32(&c.nextReq, 1)
}

func (c *Client) panicOrReturn(e error) error {
	if e != nil && c.PanicOnError {
		panic(e)
	}
	return e
}
