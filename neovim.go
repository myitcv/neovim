// Copyright 2014 Paul Jolly <paul@myitcv.org.uk>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package neovim implements support for writing Neovim plugins in Go. Communication
with Neovim is via MSGPACK:

https://github.com/msgpack/msgpack/blob/master/spec.md

All Neovim API methods are supported. In addition there is support for handling synchronous
method calls or asynchronous notifications from
Neovim by registering handlers in your plugin. See the MSGPACK RPC spec for further details
on these two types of callback:

https://github.com/msgpack-rpc/msgpack-rpc/blob/master/spec.md

Via a plugin manifest (details to follow) your plugin can bind these handlers to either
Autocmds, Functions or Commands in Neovim.

Status

This project is still in alpha.

Example Plugin

For a complete example, see the example.Example plugin http://godoc.org/github.com/myitcv/neovim/example

Writing plugins

Plugins implement the Plugin interface. Plugins are initialised with a Client that is passed
to a plugin via the Init method they implement. The Client is used to communicate with a
Neovim instance.

Concurrency

A single Client may safely be used by multiple goroutines. Calls to API methods are blocking
by design.

Compatibility

There are currently no checks to verify a connected Neovim instance exposes the same API
against which the neovim package was generated. This is future work (and probably needs
some work on the Neovim side).

See also

The tool for generating the API
The Neovim Go plugin manager
The code generator used by plugin writers

Errors

Errors returned by this package are created using errors at http://godoc.org/github.com/juju/errors.
Hence errors may be inspected using functions like errors.Details for example:

	_, err := client.GetCurrentBuffer()
	if err != nil {
		log.Fatalf("Could not get current buffer: %v", errors.Details(err))
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

	"github.com/juju/errors"
	"github.com/myitcv/neovim/apidef"
	"github.com/vmihailenco/msgpack"
)

// NewUnixClient is a convenience method for creating a new *Client. Method signature matches
// that of net.DialUnix
func NewUnixClient(im InitMethod, _net string, laddr, raddr *net.UnixAddr, log Logger) (*Client, error) {
	c, err := net.DialUnix(_net, laddr, raddr)
	if err != nil {
		return nil, errors.Annotatef(err, "could not establish connection to Neovim, _net %v, laddr %v, %v", _net, laddr, raddr)
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
		return nil, errors.Annotatef(err, "could not get a stdin pipe to embedded nvim")
	}
	stdout, err := c.StdoutPipe()
	if err != nil {
		return nil, errors.Annotatef(err, "could not get a stdout pipe to embedded nvim")
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
		c.Args = append(c.Args, "--embed", "-N")
	}

	err = c.Start()
	if err != nil {
		return nil, errors.Annotatef(err, "could not start the embedded nvim")
	}

	return NewClient(im, wrap, log)
}

// NewClient creates a new Client
func NewClient(im InitMethod, c io.ReadWriteCloser, log Logger) (*Client, error) {
	if log == nil {
		log = _log.New(os.Stderr, "neovim ", _log.Llongfile|_log.Ldate|_log.Ltime)
	}

	res := &Client{rw: c}
	res.respMap = newrespSyncMap()
	res.syncProvMap = newsyncProvSyncMap()
	res.asyncProvMap = newasyncProvSyncMap()
	res.dec = msgpack.NewDecoder(c)
	res.enc = msgpack.NewEncoder(c)
	res.log = log
	res.KillChannel = make(chan struct{})

	err := res.syncProvMap.Put(_MethodInit, func() SyncDecoder {
		res := &InitMethodWrapper{InitMethod: im}
		return res
	})
	if err != nil {
		return nil, errors.Annotatef(err, "Could not add init method handler")
	}

	return res, nil
}

func (c *Client) Run() {
	go c.doListen()
}

func (c *Client) RegisterSyncRequestHandler(m string, d NewSyncDecoder) error {
	if m == _MethodInit {
		return errors.Errorf("Cannot register a provider with the protected method %v", _MethodInit)
	}
	err := c.syncProvMap.Put(m, d)
	if err != nil {
		return errors.Annotatef(err, "Could not store RequestHanlder in provider map")
	}

	return nil
}

func (c *Client) RegisterAsyncRequestHandler(m string, d NewAsyncDecoder) error {
	err := c.asyncProvMap.Put(m, d)
	if err != nil {
		return errors.Annotatef(err, "Could not store RequestHanlder in provider map")
	}

	return nil
}

// Close cleanly kills the client connection to Neovim
func (c *Client) Close() error {
	c.lock.Lock()
	defer c.lock.Unlock()
	err := c.rw.Close()
	if err != nil {
		return c.panicOrReturn(errors.Annotatef(err, "Could not cleanly close client"))
	}
	return nil
}

// TODO split this beast apart
func (c *Client) doListen() error {
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

			newdecoder, err := c.syncProvMap.Get(reqMeth)
			if err != nil {
				c.log.Fatalf("Could not find RequestHandler for method [%v]: %v\n", reqMeth, err)
			}

			dre := newdecoder()

			err = dre.DecodeMsg(c.dec)
			if err == io.EOF {
				break
			} else if err != nil {
				c.log.Fatalf("Could not decode request method args: %v", err)
			}

			go func(reqID uint32, dre SyncDecoder) {
				mErr, err := dre.Run()
				if err != nil {
					c.log.Fatalf("Could not run method: %v\n", err)
				}
				err = c.sendResponse(reqID, mErr, dre)
				if err != nil {
					c.log.Fatalf("Could not send response: %v\n", err)
				}
			}(reqID, dre)
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

			newDecoder, err := c.asyncProvMap.Get(topic)
			if err != nil {
				c.log.Fatalf("Could not find async handler for topic [%v]: %v\n", topic, err)
			}

			dr := newDecoder()
			err = dr.DecodeMsg(c.dec)
			if err == io.EOF {
				break
			} else if err != nil {
				c.log.Fatalf("Could not decode request method args: %v", err)
			}

			go func(dr AsyncDecoder) {
				err = dr.Run()
				if err != nil {
					c.log.Fatalf("Could not run async notification")
				}
			}(dr)

		default:
			c.log.Fatalf("Unexpected type of message: %v\n", t)
		}
	}

	// TODO we could improve this, instead of log.Fatal above
	// return the actual error
	return nil
}

func (c *Client) sendResponse(reqID uint32, respErr error, e SyncDecoder) error {
	if e == nil {
		c.log.Fatalf("Need to send an encoder...")
	}

	c.lock.Lock()
	defer c.lock.Unlock()

	reqType := 1
	enc := c.enc

	err := enc.EncodeSliceLen(4)
	if err != nil {
		return errors.Annotate(err, "Could not encode request length")
	}

	err = enc.EncodeInt(reqType)
	if err != nil {
		return errors.Annotate(err, "Could not encode response type")
	}

	err = enc.EncodeUint32(reqID)
	if err != nil {
		return errors.Annotate(err, "Could not encode reqID")
	}

	// TODO support for response errors with the respErr passed in
	err = enc.EncodeNil()
	if err != nil {
		return errors.Annotate(err, "Could not encode response error")
	}

	err = e.EncodeMsg(c.enc)
	if err != nil {
		return errors.Annotate(err, "Could not encode response vals")
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
		return nil, errors.Annotate(err, "Could not store response holder")
	}

	err = enc.EncodeSliceLen(4)
	if err != nil {
		return nil, errors.Annotate(err, "Could not encode request length")
	}

	err = enc.EncodeInt(reqType)
	if err != nil {
		return nil, errors.Annotate(err, "Could not encode request type")
	}

	err = enc.EncodeUint32(reqID)
	if err != nil {
		return nil, errors.Annotate(err, "Could not encode request ID")
	}

	// err = enc.EncodeBytes([]byte(reqMethID))
	err = enc.EncodeString(string(reqMethID))
	if err != nil {
		return nil, errors.Annotate(err, "Could not encode request method ID")
	}

	err = e()
	if err != nil {
		return nil, errors.Annotate(err, "Could not encode method args ")
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

func (c *Client) GetAPIInfo() (ChannelID, *apidef.API, error) {
	var retChanID ChannelID
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
			return nil, errors.Errorf("Expected slice len to be 2; got %v", l)
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
		return retChanID, retAPI, c.panicOrReturn(errors.Annotate(err, "Could not make call to Client.GetBuffers"))
	}
	resp := <-respChan
	if resp == nil {
		return retChanID, retAPI, c.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retChanID, retAPI, c.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	retVal := resp.obj.([]interface{})
	retChanID = ChannelID(retVal[0].(uint8))
	retAPI = retVal[1].(*apidef.API)
	return retChanID, retAPI, nil

}
