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
	"github.com/tinylib/msgp/msgp"
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
	res.respMap = newRespSyncMap()
	res.syncProvMap = newSyncProvSyncMap()
	res.asyncProvMap = newAsyncProvSyncMap()
	res.dec = msgp.NewReader(c)
	res.enc = msgp.NewWriter(c)
	res.log = log
	res.KillChannel = make(chan struct{})

	err := res.syncProvMap.Put(_MethodInit, NewSyncDecoderOptions{NewSyncDecoder: func() SyncDecoder {
		res := &InitMethodWrapper{InitMethod: im, Client: res, args: &InitMethodArgs{}, results: &InitMethodRetVals{}}
		return res
	}, MethodOptions: &MethodOptions{Type: FUNCTION}})
	if err != nil {
		return nil, errors.Annotatef(err, "Could not add init method handler")
	}

	return res, nil
}

func (c *Client) Run() {
	go c.doListen()
}

func (c *Client) RegisterSyncFunction(m string, d NewSyncDecoder, rangeBased bool, eval bool) error {
	return c.RegisterSyncRequestHandler(m, d, &MethodOptions{
		Type:  FUNCTION,
		Range: rangeBased,
		Eval:  eval,
	})
}

func (c *Client) RegisterAsyncFunction(m string, d NewAsyncDecoder, rangeBased bool, eval bool) error {
	return c.RegisterAsyncRequestHandler(m, d, &MethodOptions{
		Type:  FUNCTION,
		Range: rangeBased,
		Eval:  eval,
	})
}

func (c *Client) RegisterSyncRequestHandler(m string, d NewSyncDecoder, o *MethodOptions) error {
	if m == _MethodInit {
		return errors.Errorf("Cannot register a provider with the protected method %v", _MethodInit)
	}
	err := c.syncProvMap.Put(m, NewSyncDecoderOptions{d, o})
	if err != nil {
		return errors.Annotatef(err, "Could not store RequestHanlder in provider map")
	}

	return nil
}

func (c *Client) RegisterAsyncRequestHandler(m string, d NewAsyncDecoder, o *MethodOptions) error {
	err := c.asyncProvMap.Put(m, NewAsyncDecoderOptions{d, o})
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
		_, err := dec.ReadArrayHeader()
		if err == io.EOF {
			break
		} else if err != nil {
			c.log.Fatalf("Could not decode message slice length: %v", err)
		}

		t, err := dec.ReadInt()
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
			reqID, err := dec.ReadUint32()
			if err == io.EOF {
				break
			} else if err != nil {
				c.log.Fatalf("Could not decode request id: %v", err)
			}

			reqMeth, err := dec.ReadString()
			if err == io.EOF {
				break
			} else if err != nil {
				c.log.Fatalf("Could not decode request method name: %v", err)
			}

			newdecoderoptions, err := c.syncProvMap.Get(string(reqMeth))
			if err != nil {
				c.log.Fatalf("Could not find RequestHandler for method [%v]: %v\n", string(reqMeth), err)
			}

			opts := newdecoderoptions.MethodOptions

			// now we have an array of values
			// the length of the array is determined by the method options
			expArgs := opts.ArgsLength()
			actArgs, err := dec.ReadArrayHeader()
			if err != nil {
				c.log.Fatalf("Could not decode request method args length: %v", err)
			}

			if expArgs != actArgs {
				c.log.Fatalf("Expected args length %v, but got %v", expArgs, actArgs)
			}

			dre := newdecoderoptions.NewSyncDecoder()

			// TODO need to add support for command and autocommand here
			if opts.Type == FUNCTION {
				err = dre.Args().DecodeMsg(c.dec)
				if err == io.EOF {
					break
				} else if err != nil {
					c.log.Fatalf("Could not decode request method args: %v", err)
				}
			} else {
				c.log.Fatalf("We can't support anything other than functions right now")
			}

			// now we read the options
			if dre.Params() != nil {
				dre.Params().DecodeParams(opts, dec)
			}

			// now we read the eval result if there is one
			if opts.Eval {
				dre.Eval().DecodeMsg(dec)
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
			reqID, err := dec.ReadUint32()
			if err == io.EOF {
				break
			} else if err != nil {
				c.log.Fatalf("Could not decode request id: %v", err)
			}

			re, err := dec.ReadIntf()
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
			topic, err := dec.ReadString()
			if err == io.EOF {
				break
			} else if err != nil {
				c.log.Fatalf("Could not decode topic: %v", err)
			}

			// TODO lots of code repetition here with case 1

			newdecoderoptions, err := c.asyncProvMap.Get(string(topic))
			if err != nil {
				c.log.Fatalf("Could not find RequestHandler for method [%v]: %v\n", string(topic), err)
			}

			opts := newdecoderoptions.MethodOptions

			// now we have an array of values
			// the length of the array is determined by the method options
			expArgs := opts.ArgsLength()
			actArgs, err := dec.ReadArrayHeader()
			if err != nil {
				c.log.Fatalf("Could not decode request method args length: %v", err)
			}

			if expArgs != actArgs {
				c.log.Fatalf("Expected args length %v, but got %v", expArgs, actArgs)
			}

			dr := newdecoderoptions.NewAsyncDecoder()

			// TODO need to add support for command and autocommand here
			if opts.Type == FUNCTION {
				err = dr.Args().DecodeMsg(c.dec)
				if err == io.EOF {
					break
				} else if err != nil {
					c.log.Fatalf("Could not decode request method args: %v", err)
				}
			} else {
				c.log.Fatalf("We can't support anything other than functions right now")
			}

			// now we read the options
			if dr.Params() != nil {
				dr.Params().DecodeParams(opts, dec)
			}

			// now we read the eval result if there is one
			if opts.Eval {
				dr.Eval().DecodeMsg(dec)
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

	err := enc.WriteArrayHeader(4)
	if err != nil {
		return errors.Annotate(err, "Could not encode request length")
	}

	err = enc.WriteInt(reqType)
	if err != nil {
		return errors.Annotate(err, "Could not encode response type")
	}

	err = enc.WriteUint32(reqID)
	if err != nil {
		return errors.Annotate(err, "Could not encode reqID")
	}

	// TODO support for response errors with the respErr passed in
	err = enc.WriteNil()
	if err != nil {
		return errors.Annotate(err, "Could not encode response error")
	}

	err = e.Results().EncodeMsg(c.enc)
	if err != nil {
		return errors.Annotate(err, "Could not encode response vals")
	}

	err = enc.Flush()
	if err != nil {
		return errors.Annotate(err, "Could not flush response")
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

	err = enc.WriteArrayHeader(4)
	if err != nil {
		return nil, errors.Annotate(err, "Could not encode request length")
	}

	err = enc.WriteInt(reqType)
	if err != nil {
		return nil, errors.Annotate(err, "Could not encode request type")
	}

	err = enc.WriteUint32(reqID)
	if err != nil {
		return nil, errors.Annotate(err, "Could not encode request ID")
	}

	// err = enc.EncodeBytes([]byte(reqMethID))
	err = enc.WriteString(string(reqMethID))
	if err != nil {
		return nil, errors.Annotate(err, "Could not encode request method ID")
	}

	err = e()
	if err != nil {
		return nil, errors.Annotate(err, "Could not encode method args ")
	}
	err = enc.Flush()
	if err != nil {
		return nil, errors.Annotate(err, "Could not flush encoder")
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
		_err = c.enc.WriteArrayHeader(0)
		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		l, _err := c.dec.ReadArrayHeader()
		if _err != nil {
			return
		}

		if l != 2 {
			return nil, errors.Errorf("Expected slice len to be 2; got %v", l)
		}

		chanID, _err := c.dec.ReadUint8()
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
