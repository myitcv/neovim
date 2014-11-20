// Copyright 2014 Paul Jolly <paul@myitcv.org.uk>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package neovim implements support for writing Neovim plugins in Go. It also
implements a tool for generating the MSGPACK-based API against a Neovim instance.

All API methods are supported, as are notifications. See Subscription for an example
of how to register a subscription on a given topic.

Example Plugin

For an example plugin see http://godoc.org/github.com/myitcv/neovim/example

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
	"io"
	_log "log"
	"net"
	"os"
	"os/exec"
	"sync/atomic"

	"github.com/juju/errgo"
	"github.com/myitcv/neovim/apidef"
	"github.com/vmihailenco/msgpack"
)

// NewUnixClient is a convenience method for creating a new *Client. Method signature matches
// that of net.DialUnix
func NewUnixClient(im InitMethod, _net string, laddr, raddr *net.UnixAddr, log Logger) (*Client, error) {
	c, err := net.DialUnix(_net, laddr, raddr)
	if err != nil {
		return nil, errgo.Notef(err, "Could not establish connection to Neovim, _net %v, laddr %v, %v", _net, laddr, raddr)
	}
	return NewClient(im, c, log)
}

// NewCmdClient creates a new Client that is linked via stdin/stdout to the
// supplied exec.Cmd, which is assumed to launch Neovim. The Neovim flag
// --embed is added if it is missing, and the exec.Cmd is started
// as part of creating the client. Calling Close() will close stdin on the
// embedded Neovim instance, thereby ending the process
func NewCmdClient(im InitMethod, c *exec.Cmd, log Logger) (*Client, error) {
	stdin, err := c.StdinPipe()
	if err != nil {
		return nil, errgo.Notef(err, "Could not get a stdin pipe to embedded nvim")
	}
	stdout, err := c.StdoutPipe()
	if err != nil {
		return nil, errgo.Notef(err, "Could not get a stdout pipe to embedded nvim")
	}
	wrap := &StdWrapper{Stdin: stdin, Stdout: stdout}

	// ensure that we have --embed
	found := false
	for i := range c.Args {
		if c.Args[i] == "--embed" {
			found = true
		}
	}

	if !found {
		c.Args = append(c.Args, "--embed")
	}

	err = c.Start()
	if err != nil {
		return nil, errgo.Notef(err, "Could not start the embedded nvim")
	}

	return NewClient(im, wrap, log)
}

func loggerOrStderr(log Logger) (res Logger) {
	return
}

// NewClient creates a new Client
func NewClient(im InitMethod, c io.ReadWriteCloser, log Logger) (*Client, error) {
	if log == nil {
		log = _log.New(os.Stderr, "neovim ", _log.Llongfile|_log.Ldate|_log.Ltime)
	}

	res := &Client{rw: c}
	res.respMap = newSyncRespMap()
	res.syncProvMap = newSyncProviderMap()
	res.asyncProvMap = newAsyncProviderMap()
	res.dec = msgpack.NewDecoder(c)
	res.enc = msgpack.NewEncoder(c)
	res.log = log
	res.subChan = make(chan subWrapper)
	res.KillChannel = make(chan struct{})

	err := res.syncProvMap.Put(_MethodInit, &initMethodDecoder{InitMethod: im})
	if err != nil {
		return nil, errgo.Notef(err, "Could not add init method handler")
	}

	return res, nil
}

func (c *Client) Run() {
	go c.doListen()
}

type initMethodDecoder struct {
	InitMethod
}

type initMethodRunner struct {
	InitMethod
}

type initMethodEncoder struct{}

func (i *initMethodDecoder) Decode(dec *msgpack.Decoder) (Runner, error) {
	l, err := dec.DecodeSliceLen()
	if err != nil {
		return nil, err
	}

	if l != 0 {
		return nil, errgo.Newf("Expected 0 arguments, not %v", l)
	}

	res := &initMethodRunner{InitMethod: i.InitMethod}

	return res, nil
}

func (i *initMethodEncoder) Encode(enc *msgpack.Encoder) error {
	err := enc.EncodeNil()
	if err != nil {
		return err
	}

	return nil
}

func (i *initMethodRunner) Run() (Encoder, error, error) {
	return &initMethodEncoder{}, nil, i.InitMethod()
}

func (c *Client) RegisterSyncRequestHandler(m string, d SyncDecoder) error {
	if m == _MethodInit {
		return errgo.Newf("Cannot register a provider with the protected method %v", _MethodInit)
	}
	err := c.syncProvMap.Put(m, d)
	if err != nil {
		return errgo.Notef(err, "Could not store RequestHanlder in provider map")
	}

	return nil
}

func (c *Client) RegisterAsyncRequestHandler(m string, d AsyncDecoder) error {
	err := c.asyncProvMap.Put(m, d)
	if err != nil {
		return errgo.Notef(err, "Could not store RequestHanlder in provider map")
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
	// c.t.Kill(nil)
	// c.t.Wait()
	return nil
}

func (c *Client) doListen() error {
	// subEvents := make(chan *SubscriptionEvent, 10)
	// c.doSubscriptionManager(subEvents)

	dec := c.dec
	for {
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
				c.log.Fatalf("Could not decode request id: %v", err)
			}

			reqMeth, err := dec.DecodeString()
			if err == io.EOF {
				break
			} else if err != nil {
				c.log.Fatalf("Could not decode request method name: %v", err)
			}
			c.log.Printf("Got a request for %v\n", reqMeth)

			decoder, err := c.syncProvMap.Get(reqMeth)
			if err != nil {
				c.log.Fatalf("Could not find RequestHandler for method [%v]: %v\n", reqMeth, err)
			}

			runner, err := decoder.Decode(dec)
			if err == io.EOF {
				break
			} else if err != nil {
				c.log.Fatalf("Could not decode request method args: %v", err)
			}

			go func(rr Runner) {
				encoder, mErr, err := rr.Run()
				if err != nil {
					c.log.Fatalf("Could not run method: %v\n", err)
				}
				err = c.sendResponse(reqID, mErr, encoder)
				if err != nil {
					c.log.Fatalf("Could not send response: %v\n", err)
				}
			}(runner)
		case 1:
			// handle response
			reqID, err := dec.DecodeUint32()
			if err == io.EOF {
				break
			} else if err != nil {
				c.log.Fatalf("Could not decode request id: %v", err)
			}

			re, err := dec.DecodeInterface()
			if err == io.EOF {
				break
			} else if err != nil {
				c.log.Fatalf("Could not decode response error: %v", err)
			}
			if re != nil {
				re := re.([]interface{})
				// b := re[1].([]byte)
				errBytes := re[1].([]uint8)
				errString := make([]byte, len(errBytes))
				for i := range errBytes {
					errString[i] = byte(errBytes[i])
				}
				c.log.Fatalf("Got a response error for request %v: %v", reqID, string(errString))
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

			decoder, err := c.asyncProvMap.Get(topic)
			if err != nil {
				c.log.Fatalf("Could not find async handler for topic [%v]: %v\n", topic, err)
			}

			runner, err := decoder.Decode(dec)
			if err == io.EOF {
				break
			} else if err != nil {
				c.log.Fatalf("Could not decode request method args: %v", err)
			}

			// TODO make async?
			_, _, err = runner.Run()
			if err != nil {
				c.log.Fatalf("Could not run async notification")
			}

		default:
			c.log.Fatalf("Unexpected type of message: %v\n", t)
		}
	}

	// TODO we could improve this, instead of log.Fatal above
	// return the actual error
	return nil
}

func (c *Client) sendResponse(reqID uint32, respErr error, e Encoder) error {
	if e == nil {
		c.log.Fatalf("Need to send an encoder...")
	}

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

	// TODO support for response errors with the respErr passed in
	err = enc.EncodeNil()
	if err != nil {
		return errgo.NoteMask(err, "Could not encode response error")
	}

	err = e.Encode(enc)
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

	err = enc.EncodeBytes([]byte(reqMethID))
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
		c.log.Panic(e)
	}
	return e
}

func (c *Client) getAPIInfo() (uint8, *apidef.API, error) {
	var retChanID uint8
	var retAPI *apidef.API
	enc := func() (_err error) {
		_err = c.enc.EncodeSliceLen(0)
		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		l, _err := c.dec.DecodeSliceLen()
		if _err != nil {
			return
		}

		if l != 2 {
			return nil, errgo.Newf("Expected slice len to be 2; got %v", l)
		}

		chanID, _err := c.dec.DecodeUint8()
		if _err != nil {
			return
		}

		api, _err := apidef.GetAPI(c.dec)
		if _err != nil {
			return
		}

		_i = []interface{}{chanID, api}

		return
	}
	respChan, err := c.makeCall("vim_get_api_info", enc, dec)
	if err != nil {
		return retChanID, retAPI, c.panicOrReturn(errgo.NoteMask(err, "Could not make call to Client.GetBuffers"))
	}
	resp := <-respChan
	if resp == nil {
		return retChanID, retAPI, c.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retChanID, retAPI, c.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	retVal := resp.obj.([]interface{})
	retChanID = retVal[0].(uint8)
	retAPI = retVal[1].(*apidef.API)
	return retChanID, retAPI, nil

}

// // Subscribe subscribes to a topic of events from Neovim. The
// // *Subscription.Events channel will receive SubscriptionEvent's
// // Unsubscribe needs to be called on a different goroutine to
// // the goroutine that handles these SubscriptionEvent's
// func (c *Client) Subscribe(topic string, dec Decoder) (*Subscription, error) {
// 	errChan := make(chan error)

// 	res := &Subscription{
// 		Topic:   topic,
// 		Decoder: dec,
// 	}

// 	c.subChan <- subWrapper{
// 		sub:     res,
// 		errChan: errChan,
// 		task:    _Sub,
// 	}

// 	err := <-errChan
// 	if err != nil {
// 		return nil, c.panicOrReturn(errgo.NoteMask(err, "Could not register subscription"))
// 	}

// 	return res, nil
// }

// // Unsubscribe unsubscribes from a topic of events from Neovim.
// // This needs to be called on a different goroutine to that which
// // is handling the SubscriptionEvent's
// func (c *Client) Unsubscribe(sub *Subscription) error {
// 	errChan := make(chan error)
// 	c.subChan <- subWrapper{
// 		sub:     sub,
// 		errChan: errChan,
// 		task:    _Unsub,
// 	}

// 	err := <-errChan
// 	if err != nil {
// 		return c.panicOrReturn(errgo.NoteMask(err, "Could not register unsubscribe"))
// 	}

// 	return nil
// }

// func (c *Client) doSubscriptionManager(se chan *SubscriptionEvent) {
// 	c.t.Go(func() error {
// 		// this map keeps track of subscriptions
// 		// subs["topic"] == nil indicates no subscriptions on a topic
// 		// subs["topic"] otherwise contains a map, where the keys are the
// 		// Events channels that have subscribed
// 		subs := make(map[string]map[Subscription]struct{})

// 		// a goroutine that is responsible for handling subscribe/unsubscribe calls
// 		// on a separate goroutine because the calls to sub/unsub are blocking
// 		subTasks := make(chan subWrapper, 10)
// 		c.t.Go(func() error {
// 			for {
// 				select {
// 				// case t := <-subTasks:
// 				case <-subTasks:
// 					// if t.task == _Sub {
// 					// 	err := c.subscribe(t.sub.Topic)
// 					// 	if err != nil {
// 					// 		t.errChan <- errgo.NoteMask(err, "Could not subscribe")
// 					// 	} else {
// 					// 		close(t.errChan)
// 					// 	}
// 					// } else if t.task == _Unsub {
// 					// 	err := c.unsubscribe(t.sub.Topic)
// 					// 	if err != nil {
// 					// 		t.errChan <- errgo.NoteMask(err, "Could not unsubscribe")
// 					// 	} else {
// 					// 		close(t.errChan)
// 					// 	}
// 					// }
// 				case <-c.t.Dying():
// 					return nil
// 				}
// 			}
// 		})

// 		for {
// 			select {
// 			case <-c.t.Dying():
// 				return nil
// 			case event := <-se:
// 				// receive from the main doListen goroutine
// 				if chans, ok := subs[event.Topic]; ok {
// 					for k := range chans {
// 						k <- event
// 					}
// 				} else {
// 					c.log.Fatalf("Got an event for which we have no subs on topic %v\n", event.Topic)
// 				}
// 			case w := <-c.subChan:
// 				if w.task == _Sub {
// 					sub := w.sub
// 					m, ok := subs[sub.Topic]
// 					if !ok {
// 						// we have no subscriptions on this topic
// 						// the handling of the sub task will close
// 						// the error channel
// 						// subTasks <- w
// 						m = make(map[chan *SubscriptionEvent]struct{})
// 						subs[sub.Topic] = m
// 						close(w.errChan)
// 					} else if _, ok := m[sub.Events]; ok {
// 						// fatal error if we already have subscribed
// 						// using this channel
// 						w.errChan <- errgo.Newf("Already have subscription for topic %v with this channel: %v", sub.Topic, sub.Events)
// 					} else {
// 						// we are simply going to add to an existing
// 						// subscription. Close the error channel
// 						close(w.errChan)
// 					}
// 					m[sub.Events] = struct{}{}
// 				} else if w.task == _Unsub {
// 					unsub := w.sub
// 					close(unsub.Events)
// 					m, ok := subs[unsub.Topic]
// 					if !ok {
// 						w.errChan <- errgo.Newf("We don't have any subscriptions for topic %v", unsub.Topic)
// 					}
// 					if _, ok := m[unsub.Events]; !ok {
// 						w.errChan <- errgo.Newf("We don't have a subscription on topic %v for this channel %v", unsub.Topic, unsub.Events)
// 					}
// 					delete(m, unsub.Events)
// 					if len(m) == 0 {
// 						// we are back down to 0 again; unsubscribe
// 						subTasks <- w
// 						delete(subs, unsub.Topic)
// 					} else {
// 						close(w.errChan)
// 					}
// 				}
// 			}
// 		}
// 	})
// }
