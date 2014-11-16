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
	rw           io.ReadWriteCloser
	dec          *msgpack.Decoder
	enc          *msgpack.Encoder
	nextReq      uint32
	respMap      *syncRespMap
	syncProvMap  *syncProviderMap
	asyncProvMap *asyncProviderMap
	lock         sync.Mutex
	subChan      chan subWrapper
	t            tomb.Tomb

	// PanicOnError can be set to have the Client panic when an error would
	// otherwise have been returned via an API method. Note: any attempt to
	// change this option during concurrent use of the Client will be racey.
	// This is useful for debugging.
	PanicOnError bool
	KillChannel  chan struct{}
	log          Logger
}

type InitMethod func() error

func NullInitMethod() error { return nil }

// Plugin is the interface implemented by writers of Neovim plugins using the
// neovim package
type Plugin interface {
	Init(*Client, Logger) error
	Shutdown() error
}

type subTask int

// RequestHandler is the type signature of callback handlers used in
// RegisterRequestHandler
type RequestHandler func([]interface{}) ([]interface{}, error)

const (
	_MethodInit string = "plugin_load"
)

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
	Topic string
	AsyncDecoder
}

// A SubscriptionEvent contains the value Value announced via a notification
// on topic Topic
type SubscriptionEvent struct {
	Topic string
	Value []interface{}
}

type Runner func() (Encoder, error)
type Encoder func(*msgpack.Encoder) error

// Use for async notifications
// Here the error would simply be reported to the log
// (because there is nothing to return)
type AsyncDecoder func(*msgpack.Decoder) error

type SyncDecoder func(*msgpack.Decoder) (Runner, error)

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

// Logger is a local definition of the inteface effectively exposed by
// http://godoc.org/log
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

// StdWrapper is a wrapper around two io.WriterCloser and
// io.ReadCloser instances that exposes itself as an
// io.ReadWriteCloser. Typically used with os.Stdin and
// os.Stdout or their pipe equivalents
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
