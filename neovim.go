package neovim

import (
	"log"
	"net"
	"sync"
	"sync/atomic"

	"github.com/juju/errgo"
	"github.com/juju/errors"
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
	return atomic.AddUint32(&c.next_req, 1)
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
	res.SubChan = make(chan Subscription)
	res.UnsubChan = make(chan Subscription)
	go res.doListen()
	return res, nil
}

func (c *Client) doSubscriptionManager(se chan SubscriptionEvent) {
	subs := make(map[string]map[chan SubscriptionEvent]struct{})

	send_or_close := func(c chan error, e error) {
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
				for k, _ := range chans {
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
				send_or_close(sub.Error, errors.Errorf("Already have subscription for topic %v on this channel", sub.Topic))
			}
			m[sub.Events] = struct{}{}
			send_or_close(sub.Error, nil)
		case unsub := <-c.UnsubChan:
			m, ok := subs[unsub.Topic]
			if !ok {
				send_or_close(unsub.Error, errors.Errorf("We don't have any subscriptions for topic %v", unsub.Topic))
			}
			if _, ok := m[unsub.Events]; !ok {
				send_or_close(unsub.Error, errors.Errorf("We don't have a subscription on topic %v on this channel", unsub.Topic))
			}
			delete(m, unsub.Events)
			send_or_close(unsub.Error, nil)
		}
	}
}

func (c *Client) doListen() {
	// TODO need kill channel

	// TODO look at the semantics of making this buffered...
	sub_events := make(chan SubscriptionEvent, 10)
	go c.doSubscriptionManager(sub_events)

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
			res, err := rh.dec()
			if err != nil {
				log.Fatalf("Could not decode response: %v", err)
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

			sub_events <- ev
		default:
			log.Fatalf("Unexpected type of message: %v\n", t)
		}
	}
}

func (c *Client) makeCall(req_meth_id uint32, e Encoder, d Decoder) (chan *response, error) {
	req_type := 0
	req_id := c.nextReqId()
	enc := c.enc

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

	err = e()
	if err != nil {
		return nil, errgo.NoteMask(err, "Could not encode method args ")
	}

	// TODO need a flush here?

	return res, nil
}

func (c *Client) API() (*API, error) {
	enc := func() (_err error) {
		_err = c.enc.EncodeSliceLen(0)
		if _err != nil {
			return
		}

		return
	}
	resp_chan, err := c.makeCall(0, enc, c.decodeAPI)
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
