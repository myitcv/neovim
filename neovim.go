package neovim

import (
	"log"
	"net"
	"sync"

	"github.com/juju/errgo"
	"github.com/vmihailenco/msgpack"
)

// TODO the uniqueness of the request ID is specific to this
// instance of neovim. How does this work with multiple plugins
// making requests?

type sync_map struct {
	lock    *sync.Mutex
	the_map map[uint32]*response_holder
}

func (c *Client) nextReqId() uint32 {
	return sync.AddUnint(&c.next_req, 1)
}

func newSyncMap() *sync_map {
	return &sync_map{
		lock:    new(sync.Mutex),
		the_map: make(map[uint32]*response_holder),
	}
}

func (this *sync_map) Put(k uint32, v *response_holder) error {
	this.lock.Lock()
	defer this.lock.Unlock()

	if _, present := this.the_map[k]; present {
		return errgo.Newf("Key already exists for key %v", k)
	}

	this.the_map[k] = v
	return nil
}

func (this *sync_map) Get(k uint32) (*response_holder, error) {
	this.lock.Lock()
	defer this.lock.Unlock()

	if res, present := this.the_map[k]; !present {
		return nil, errgo.Newf("Key does not exist for %v", k)
	} else {
		delete(this.the_map, k)
		return res, nil
	}
}

type response_holder struct {
	dec Decoder
	ch  chan *response
}

type response struct {
	obj interface{}
	err error
}

func NewUnixClient(ua_name, ua_net string) (*Client, error) {
	a := &net.UnixAddr{Name: ua_name, Net: ua_net}
	c, err := net.DialUnix(a.Net, nil, a)
	if err != nil {
		return nil, errgo.NoteMask(err, "Could not establish connection")
	}
	res := &Client{conn: c}
	res.resp_map = newSyncMap()
	res.dec = msgpack.NewDecoder(c)
	res.enc = msgpack.NewEncoder(c)
	res.lock = new(sync.Mutex)
	go res.doListen()
	return res, nil
}

func (c *Client) doListen() {
	// TODO need kill channel
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

		if t == 0 {
			log.Fatalln("Got a request on the listen channel")
		}

		if t == 2 {
			log.Fatalln("Got a notification on the listen channel; don't know about this yet")
		}

		if t != 1 {
			log.Fatalf("Got %v, expected 1", t)
		}

		// we have a response - get the req_id

		req_id, err := dec.DecodeUint32()
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
		rh, err := c.resp_map.Get(req_id)
		if err != nil {
			log.Fatalf("Could not get response holder for %v: %v", req_id, err)
		}

		// we have a valid response, dispatch to our decoder for the response
		res, err := rh.dec(dec)
		if err != nil {
			log.Fatalf("Could not decode response: %v", err)
		}

		resp := &response{obj: res, err: nil}
		rh.ch <- resp
	}

}

func (c *Client) makeCall(req_meth_id uint32, args []interface{}, encoders []Encoder, d Decoder) (chan *response, error) {
	req_type := 0
	req_id := c.nextReqId()
	enc := c.enc

	if len(args) != len(encoders) || len(args) == 0 {
		return nil, errgo.Newf("args and encoders not the same length > 0: %v and %v", args, encoders)
	}

	res := make(chan *response)
	rh := &response_holder{dec: d, ch: res}
	err := c.resp_map.Put(req_id, rh)
	if err != nil {
		return nil, errgo.NoteMask(err, "Could not store response holder")
	}

	err = enc.EncodeSliceLen(4)
	if err != nil {
		return nil, errgo.NoteMask(err, "Could not encode request length")
	}

	err = enc.EncodeInt(req_type)
	if err != nil {
		return nil, errgo.NoteMask(err, "Could not encode request type")
	}

	err = enc.EncodeUint32(req_id)
	if err != nil {
		return nil, errgo.NoteMask(err, "Could not encode request ID")
	}

	err = enc.EncodeUint32(req_meth_id)
	if err != nil {
		return nil, errgo.NoteMask(err, "Could not encode request method ID")
	}

	for i, _ := range encoders {
		err := encoders[i](enc, args[i])
		if err != nil {
			return nil, errgo.Notef(err, "Could not encode method args at index %v with args %v", i, args[i])
		}
	}

	err = w.Flush()
	if err != nil {
		return nil, errgo.NoteMask(err, "Could not flush writer")
	}

	return res, nil
}

func (c *Client) API() (*API, error) {
	resp_chan, err := c.makeCall(0, []interface{}{nil}, []Encoder{encodeNoArgs}, decodeAPI)
	if err != nil {
		return nil, errgo.NoteMask(err, "Could not make call")
	}
	resp := <-resp_chan
	if resp == nil {
		return nil, errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return nil, errgo.NoteMask(err, "We got a non-nil error in our response")
	}
	ba := resp.obj
	return ba.(*API), nil
}
