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
	lock    sync.Mutex
	subChan chan subWrapper

	// PanicOnError can be set to have the Client panic when an error would
	// otherwise have been returned via an API method. Note: any attempt to
	// change this option during concurrent use of the Client will be racey.
	// This is useful for debugging.
	PanicOnError bool
}

type subTask int

const (
	_Sub subTask = iota
	_Unsub
)

type subWrapper struct {
	sub     *Subscription
	errChan chan error
	task    subTask
}

// A Subscription represents a subscription to a Neovim event on a particular
// topic.
type Subscription struct {
	Topic  string
	Events chan *SubscriptionEvent
}

// A SubscriptionEvent contains the value Value announced via a notification
// on topic Topic
type SubscriptionEvent struct {
	Topic string
	Value []interface{}
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
	stdin  io.WriteCloser
	stdout io.ReadCloser
}

func (s *stdWrapper) Read(p []byte) (n int, err error) {
	return s.stdout.Read(p)
}

func (s *stdWrapper) Write(p []byte) (n int, err error) {
	return s.stdin.Write(p)
}

func (s *stdWrapper) Close() error {
	return s.stdin.Close()
}

type encoder func() error
type decoder func() (interface{}, error)
