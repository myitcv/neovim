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
	"fmt"
	"io"
	_log "log"
	"net"
	"os"
	"os/exec"
	"sync/atomic"

	"github.com/juju/errgo"
	"github.com/vmihailenco/msgpack"
)

// NewUnixClient is a convenience method for creating a new *Client. Method signature matches
// that of net.DialUnix
func NewUnixClient(_net string, laddr, raddr *net.UnixAddr, l Logger) (*Client, error) {
	_log.Fatal("Test")
	c, err := net.DialUnix(_net, laddr, raddr)
	if err != nil {
		return nil, errgo.Notef(err, "Could not establish connection to Neovim, _net %v, laddr %v, %v", _net, laddr, raddr)
	}
	return NewClient(c, l)
}

// NewCmdClient creates a new Client that is linked via stdin/stdout to the
// supplied exec.Cmd, which is assumed to launch Neovim. The Neovim flag
// --embedded-mode is added if it is missing, and the exec.Cmd is started
// as part of creating the client. Calling Close() will close stdin on the
// embedded Neovim instance, thereby ending the process
func NewCmdClient(c *exec.Cmd, log Logger) (*Client, error) {
	stdin, err := c.StdinPipe()
	if err != nil {
		log.Fatalf("Could not get a stdin pipe to embedded nvim: %v\n", err)
	}
	stdout, err := c.StdoutPipe()
	if err != nil {
		log.Fatalf("Could not get a stdout pipe to embedded nvim: %v\n", err)
	}
	wrap := &StdWrapper{Stdin: stdin, Stdout: stdout}

	// ensure that we have --embedded-mode
	found := false
	for i := range c.Args {
		if c.Args[i] == "--embedded-mode" {
			found = true
		}
	}

	if !found {
		c.Args = append(c.Args, "--embedded-mode")
	}

	err = c.Start()
	if err != nil {
		log.Fatalf("Could not start the cmd: %v\n", err)
	}

	return NewClient(wrap, log)
}

// NewClient creates a new Client
func NewClient(c io.ReadWriteCloser, log Logger) (*Client, error) {
	res := &Client{rw: c}
	res.respMap = newSyncMap()
	res.dec = msgpack.NewDecoder(c)
	res.enc = msgpack.NewEncoder(c)
	res.subChan = make(chan subWrapper)

	if log != nil {
		res.log = log
	} else {
		res.log = _log.New(os.Stderr, "neovim ", _log.Llongfile|_log.Ldate|_log.Ltime)
	}

	// do not need to put this in the tomb because
	// the closing of the the reader will handle the exit
	// from this go routine
	go res.doListen()

	return res, nil
}

func (c *Client) RegisterProvider(m string, r RequestHandler) error {
	return nil
}

// Subscribe subscribes to a topic of events from Neovim. The
// *Subscription.Events channel will receive SubscriptionEvent's
// Unsubscribe needs to be called on a different goroutine to
// the goroutine that handles these SubscriptionEvent's
func (c *Client) Subscribe(topic string) (*Subscription, error) {
	respChan := make(chan *SubscriptionEvent)
	errChan := make(chan error)

	res := &Subscription{
		Topic:  topic,
		Events: respChan,
	}

	c.subChan <- subWrapper{
		sub:     res,
		errChan: errChan,
		task:    _Sub,
	}

	err := <-errChan
	if err != nil {
		return nil, c.panicOrReturn(errgo.NoteMask(err, "Could not register subscription"))
	}

	return res, nil
}

// Unsubscribe unsubscribes from a topic of events from Neovim.
// This needs to be called on a different goroutine to that which
// is handling the SubscriptionEvent's
func (c *Client) Unsubscribe(sub *Subscription) error {
	errChan := make(chan error)
	c.subChan <- subWrapper{
		sub:     sub,
		errChan: errChan,
		task:    _Unsub,
	}

	err := <-errChan
	if err != nil {
		return c.panicOrReturn(errgo.NoteMask(err, "Could not register unsubscribe"))
	}

	return nil
}

// Close cleanly kills the client connection to Neovim
func (c *Client) Close() error {
	c.lock.Lock()
	defer c.lock.Unlock()
	err := c.rw.Close()
	if err != nil {
		return c.panicOrReturn(errgo.Notef(err, "Could not cleanly close client"))
	}
	c.t.Kill(nil)
	c.t.Wait()
	return nil
}

func (c *Client) doListen() error {
	subEvents := make(chan *SubscriptionEvent, 10)
	c.doSubscriptionManager(subEvents)

	dec := c.dec
	for {
		c.log.Println("Listening for request")
		_, err := dec.DecodeSliceLen()
		if err == io.EOF {
			break
		} else if err != nil {
			c.log.Fatalf("Could not decode message slice length: %v", err)
		}

		t, err := dec.DecodeInt()
		if err == io.EOF {
			break
		} else if err != nil {
			c.log.Fatalf("Could not decode message type: %v", err)
		}

		switch t {
		// TODO implement support for handling requests in a Go client, i.e.
		// Neovim making a request to the Go client, and the Go client sending
		// a response
		case 0:
			reqID, err := dec.DecodeUint32()
			if err == io.EOF {
				break
			} else if err != nil {
				log.Fatalf("Could not decode request id: %v", err)
			}

			reqMeth, err := dec.DecodeString()
			if err == io.EOF {
				break
			} else if err != nil {
				log.Fatalf("Could not decode request method name: %v", err)
			}

			reqArgs, err := dec.DecodeSlice()
			if err == io.EOF {
				break
			} else if err != nil {
				log.Fatalf("Could not decode request method args: %v", err)
			}

			fmt.Printf("We got a request, id %v, for method %v, with args: %v\n", reqID, reqMeth, reqArgs)
			go c.sendResponse(reqID, nil, nil)
		case 1:
			// handle response
			reqID, err := dec.DecodeUint32()
			if err == io.EOF {
				break
			} else if err != nil {
				c.log.Fatalf("Could not decode request id: %v", err)
			}

			// do we have an error?
			re, err := dec.DecodeInterface()
			if err == io.EOF {
				break
			} else if err != nil {
				c.log.Fatalf("Could not decode response error: %v", err)
			}
			if re != nil {
				c.log.Fatalf("Got a response error for request %v: %v", reqID, re)
			}

			// no, carry on
			rh, err := c.respMap.Get(reqID)
			if err == io.EOF {
				break
			} else if err != nil {
				c.log.Fatalf("Could not get response holder for %v: %v", reqID, err)
			}

			res, err := rh.dec()
			if err == io.EOF {
				break
			} else if err != nil {
				c.log.Fatalf("Could not decode response: %v\n", err)
			}

			resp := &response{obj: res, err: nil}
			rh.ch <- resp
		case 2:
			// handle notification
			topic, err := dec.DecodeString()
			if err == io.EOF {
				break
			} else if err != nil {
				c.log.Fatalf("Could not decode topic: %v", err)
			}

			// TODO we could make a decode part of the subscription
			// interface to avoid this reflection based decoding
			obj, err := dec.DecodeSlice()
			if err == io.EOF {
				break
			} else if err != nil {
				c.log.Fatalf("Could not decode obj payload: %v", err)
			}

			ev := &SubscriptionEvent{
				Topic: topic,
				Value: obj,
			}

			subEvents <- ev
		default:
			c.log.Fatalf("Unexpected type of message: %v\n", t)
		}
	}

	// TODO we could improve this, instead of log.Fatal above
	// return the actual error
	return nil
}

func (c *Client) doSubscriptionManager(se chan *SubscriptionEvent) {
	c.t.Go(func() error {
		// this map keeps track of subscriptions
		// subs["topic"] == nil indicates no subscriptions on a topic
		// subs["topic"] otherwise contains a map, where the keys are the
		// Events channels that have subscribed
		subs := make(map[string]map[chan *SubscriptionEvent]struct{})

		// a goroutine that is responsible for handling subscribe/unsubscribe calls
		// on a separate goroutine because the calls to sub/unsub are blocking
		subTasks := make(chan subWrapper, 10)
		c.t.Go(func() error {
			for {
				select {
				case t := <-subTasks:
					if t.task == _Sub {
						err := c.subscribe(t.sub.Topic)
						if err != nil {
							t.errChan <- errgo.NoteMask(err, "Could not subscribe")
						} else {
							close(t.errChan)
						}
					} else if t.task == _Unsub {
						err := c.unsubscribe(t.sub.Topic)
						if err != nil {
							t.errChan <- errgo.NoteMask(err, "Could not unsubscribe")
						} else {
							close(t.errChan)
						}
					}
				case <-c.t.Dying():
					return nil
				}
			}

			// TODO we could improve this, instead of log.Fatal above
			// return the actual error
			return nil
		})

		for {
			select {
			case <-c.t.Dying():
				return nil
			case event := <-se:
				// receive from the main doListen goroutine
				if chans, ok := subs[event.Topic]; ok {
					for k := range chans {
						k <- event
					}
				} else {
					c.log.Fatalf("Got an event for which we have no subs on topic %v\n", event.Topic)
				}
			case w := <-c.subChan:
				if w.task == _Sub {
					sub := w.sub
					m, ok := subs[sub.Topic]
					if !ok {
						// we have no subscriptions on this topic
						// the handling of the sub task will close
						// the error channel
						subTasks <- w
						m = make(map[chan *SubscriptionEvent]struct{})
						subs[sub.Topic] = m
					} else if _, ok := m[sub.Events]; ok {
						// fatal error if we already have subscribed
						// using this channel
						w.errChan <- errgo.Newf("Already have subscription for topic %v with this channel: %v", sub.Topic, sub.Events)
					} else {
						// we are simply going to add to an existing
						// subscription. Close the error channel
						close(w.errChan)
					}
					m[sub.Events] = struct{}{}
				} else if w.task == _Unsub {
					unsub := w.sub
					close(unsub.Events)
					m, ok := subs[unsub.Topic]
					if !ok {
						w.errChan <- errgo.Newf("We don't have any subscriptions for topic %v", unsub.Topic)
					}
					if _, ok := m[unsub.Events]; !ok {
						w.errChan <- errgo.Newf("We don't have a subscription on topic %v for this channel %v", unsub.Topic, unsub.Events)
					}
					delete(m, unsub.Events)
					if len(m) == 0 {
						// we are back down to 0 again; unsubscribe
						subTasks <- w
						delete(subs, unsub.Topic)
					} else {
						close(w.errChan)
					}
				}
			}
		}

		// TODO we could improve this, instead of log.Fatal above
		// return the actual error
		return nil
	})
}

func (c *Client) sendResponse(reqID uint32, respErr error, e encoder) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	reqType := 1
	enc := c.enc

	err := enc.EncodeSliceLen(4)
	if err != nil {
		return errgo.NoteMask(err, "Could not encode request length")
	}

	err = enc.EncodeInt(reqType)
	if err != nil {
		return errgo.NoteMask(err, "Could not encode response type")
	}

	err = enc.EncodeUint32(reqID)
	if err != nil {
		return errgo.NoteMask(err, "Could not encode reqID")
	}

	// TODO support for response errors
	err = enc.EncodeNil()
	if err != nil {
		return errgo.NoteMask(err, "Could not encode response error")
	}

	// TODO actually encode the response vals
	// err = e()
	err = enc.Encode([]interface{}{})
	if err != nil {
		return errgo.NoteMask(err, "Could not encode response vals")
	}

	return nil
}

func (c *Client) makeCall(reqMethID neovimMethodID, e encoder, d decoder) (chan *response, error) {
	// TODO implement support for handling requests in a Go client, i.e.
	// Neovim making a request to the Go client, and the Go client sending
	// a response
	c.lock.Lock()
	defer c.lock.Unlock()

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

	err = enc.EncodeString(string(reqMethID))
	if err != nil {
		return nil, errgo.NoteMask(err, "Could not encode request method ID")
	}

	err = e()
	if err != nil {
		return nil, errgo.NoteMask(err, "Could not encode method args ")
	}

	return res, nil
}

func (c *Client) nextReqID() uint32 {
	// TODO this is no longer necessary... see makeCall
	return atomic.AddUint32(&c.nextReq, 1)
}

func (c *Client) panicOrReturn(e error) error {
	if e != nil && c.PanicOnError {
		panic(e)
	}
	return e
}
