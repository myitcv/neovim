// Copyright 2014 Paul Jolly <paul@myitcv.org.uk>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package neovim

import (
	"io"
	"sync"

	"github.com/vmihailenco/msgpack"
	"gopkg.in/tomb.v2"
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
	t       tomb.Tomb

	// PanicOnError can be set to have the Client panic when an error would
	// otherwise have been returned via an API method. Note: any attempt to
	// change this option during concurrent use of the Client will be racey.
	// This is useful for debugging.
	PanicOnError bool
	KillChannel  chan struct{}
	log          Logger
}

type Plugin interface {
	Init(*Client, Logger) error
	Shutdown() error
}

type subTask int

type RequestHandler func([]interface{}) ([]interface{}, error)

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

type Logger interface {
	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
	Fatalln(v ...interface{})
	Flags() int
	Output(calldepth int, s string) error
	Panic(v ...interface{})
	Panicf(format string, v ...interface{})
	Panicln(v ...interface{})
	Prefix() string
	Print(v ...interface{})
	Printf(format string, v ...interface{})
	Println(v ...interface{})
	SetFlags(flag int)
	SetPrefix(prefix string)
}

type responseHolder struct {
	dec decoder
	ch  chan *response
}

type response struct {
	obj interface{}
	err error
}

type StdWrapper struct {
	Stdin  io.WriteCloser
	Stdout io.ReadCloser
}

func (s *StdWrapper) Read(p []byte) (n int, err error) {
	return s.Stdout.Read(p)
}

func (s *StdWrapper) Write(p []byte) (n int, err error) {
	return s.Stdin.Write(p)
}

func (s *StdWrapper) Close() error {
	return s.Stdin.Close()
}

type encoder func() error
type decoder func() (interface{}, error)
