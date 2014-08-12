// Copyright 2014 Paul Jolly <paul@myitcv.org.uk>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package neovim

import (
	"net"

	"github.com/vmihailenco/msgpack"
)

type neovimMethodId uint32

// A Client represents a connection to a single Neovim instance
type Client struct {
	conn      net.Conn
	dec       *msgpack.Decoder
	enc       *msgpack.Encoder
	next_req  uint32
	resp_map  *sync_map
	SubChan   chan Subscription
	UnsubChan chan Subscription
}

// A Subscription is used to register/unregister interest in a topic
// in the form of a SubscriptionEvent channel (can be viewed as the
// handler)
//
// This needs to be used in conjunction with Client.Subscribe and
// Client.Unsubscribe
type Subscription struct {
	Topic  string
	Error  chan error
	Events chan SubscriptionEvent
}

// A SubscriptionEvent contains the value Value announced via a notification
// on topic Topic
type SubscriptionEvent struct {
	Topic string
	Value interface{}
}

// Represents a Neovim Buffer
//
// Multiple goroutines may invoke methods on a Buffer simultaneously
type Buffer struct {
	Id     uint32
	client *Client
}

// Represents a Neovim Window
//
// Multiple goroutines may invoke methods on a Window simultaneously
type Window struct {
	Id     uint32
	client *Client
}

// Represents a Neovim Tabpage
//
// Multiple goroutines may invoke methods on a Tabpage simultaneously
type Tabpage struct {
	Id     uint32
	client *Client
}

type response_holder struct {
	dec decoder
	ch  chan *response
}

type response struct {
	obj interface{}
	err error
}

type encoder func() error
type decoder func() (interface{}, error)
