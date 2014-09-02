// Copyright 2014 Paul Jolly <paul@myitcv.org.uk>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package neovim

import (
	"io"
	"sync"

	"github.com/vmihailenco/msgpack"
)

type neovimMethodID string

// A Client represents a connection to a single Neovim instance
type Client struct {
	rw      io.ReadWriteCloser
	dec     *msgpack.Decoder
	enc     *msgpack.Encoder
	nextReq uint32
	respMap *syncMap

	// SubChan is the channel on which subscription requests are registered
	SubChan chan Subscription

	// UnsubChan is the channel on which subscription requests are unregistered
	UnsubChan chan Subscription

	// PanicOnError can be set to have the Client panic when an error would
	// otherwise have been returned via an API method. Note: any attempt to
	// change this option during concurrent use of the Client will be racey
	PanicOnError bool
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

// Buffer represents a Neovim Buffer
//
// Multiple goroutines may invoke methods on a Buffer simultaneously
type Buffer struct {
	ID     uint32
	client *Client
}

// Window represents a Neovim Window
//
// Multiple goroutines may invoke methods on a Window simultaneously
type Window struct {
	ID     uint32
	client *Client
}

// Tabpage represents a Neovim Tabpage
//
// Multiple goroutines may invoke methods on a Tabpage simultaneously
type Tabpage struct {
	ID     uint32
	client *Client
}

type responseHolder struct {
	dec decoder
	ch  chan *response
}

type response struct {
	obj interface{}
	err error
}

type stdWrapper struct {
	lock   sync.Mutex
	stdin  io.WriteCloser
	stdout io.ReadCloser
}

func (s *stdWrapper) Read(p []byte) (n int, err error) {
	return s.stdout.Read(p)
}

func (s *stdWrapper) Write(p []byte) (n int, err error) {
	// TODO this should be unncessary, but leave in for now
	// while Neovim concurrent test fails
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.stdin.Write(p)
}

func (s *stdWrapper) Close() error {
	return s.stdin.Close()
}

type encoder func() error
type decoder func() (interface{}, error)
