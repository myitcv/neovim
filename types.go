// Copyright 2014 Paul Jolly <paul@myitcv.org.uk>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package neovim

import (
	"io"
	"sync"

	"github.com/juju/errors"
	"github.com/tinylib/msgp/msgp"
)

//go:generate gotemplate "github.com/myitcv/neovim/template/syncmap" "respSyncMap(uint32, *responseHolder)"
//go:generate gotemplate "github.com/myitcv/neovim/template/syncmap" "syncProvSyncMap(string, NewSyncDecoderOptions)"
//go:generate gotemplate "github.com/myitcv/neovim/template/syncmap" "asyncProvSyncMap(string, NewAsyncDecoderOptions)"

type neovimMethodID string

// A Client represents a connection to a single Neovim instance
type Client struct {
	rw           io.ReadWriteCloser
	dec          *msgp.Reader
	enc          *msgp.Writer
	nextReq      uint32
	respMap      *respSyncMap
	syncProvMap  *syncProvSyncMap
	asyncProvMap *asyncProvSyncMap

	// used to prevent a race between Close and send
	// TODO but maybe that's unnecessary?
	lock sync.Mutex

	// PanicOnError can be set to have the Client panic when an error would
	// otherwise have been returned via an API method. Note: any attempt to
	// change this option during concurrent use of the Client will be racey.
	// This is useful for debugging.
	PanicOnError bool
	KillChannel  chan struct{}

	// TODO remove this
	HostName string
	log      Logger
}

type NilDeocdable struct{}

func (n *NilDeocdable) DecodeMsg(r *msgp.Reader) error {
	i, err := r.ReadArrayHeader()
	if err != nil {
		return err
	}

	if i != 0 {
		return errors.Errorf("Expected zero arguments, got %v", i)
	}

	return nil
}

type InitMethod func() error

type ChannelID uint8

func NullInitMethod() error { return nil }

type MethodOptionParams struct {
	Range *Range
}

func (m *MethodOptionParams) DecodeParams(o *MethodOptions, reader *msgp.Reader) error {
	if o.Range {
		r := new(Range)
		err := r.DecodeMsg(reader)
		if err != nil {
			return errors.Annotatef(err, "could not read a range")
		}
		m.Range = r
	}
	return nil
}

type MethodType uint

const (
	FUNCTION MethodType = iota
	COMMAND
	AUTOCOMMAND
)

type MethodOptions struct {
	Type  MethodType
	Range bool
	Eval  bool

	// TODO
	// Bang bool
	// Pattern string
	// Complete bool
	// NArgs int
	// Count int
	// Register
}

func (m *MethodOptions) ArgsLength() (res uint32) {
	switch m.Type {
	case FUNCTION:
		// args are always sent for a function, even if there are none
		// in which case they will be sent as []
		res = 1
		if m.Range {
			res += 1
		}
		if m.Eval {
			res += 1
		}
	}
	return res
}

// Plugin is the interface implemented by writers of Neovim plugins using the
// neovim package
type Plugin interface {
	Init(*Client, Logger) error
	Shutdown() error
}

// RequestHandler is the type signature of callback handlers used in
// RegisterRequestHandler
type RequestHandler func([]interface{}) ([]interface{}, error)

const (
	_MethodInit string = "plugin_load"
)

type SyncRunner interface {
	Run() (error, error)
}

type AsyncRunner interface {
	Run() error
}

type Encoder interface {
	EncodeMsg(*msgp.Writer) error
}

type NewSyncDecoder func() SyncDecoder
type NewAsyncDecoder func() AsyncDecoder

type NewSyncDecoderOptions struct {
	NewSyncDecoder
	*MethodOptions
}

type NewAsyncDecoderOptions struct {
	NewAsyncDecoder
	*MethodOptions
}

// Use for async notifications
// Here the error would simply be reported to the log
// (because there is nothing to return)
type Decoder interface {
	DecodeMsg(*msgp.Reader) error
}

type CallArgs interface {
	Eval() msgp.Decodable
	Params() *MethodOptionParams
	Args() msgp.Decodable
}

type SyncDecoder interface {
	CallArgs
	SyncRunner
	Results() msgp.Encodable
}

type AsyncDecoder interface {
	CallArgs
	AsyncRunner
}

// Buffer represents a Neovim Buffer
//
// Multiple goroutines may invoke methods on a Buffer simultaneously
type Buffer struct {
	ID     uint8
	client *Client
}

func (b *Buffer) ExtensionType() int8 {
	return typeBuffer
}

func (b *Buffer) Len() int {
	return 1
}

func (b *Buffer) MarshalBinaryTo(buf []byte) error {
	buf[0] = b.ID
	return nil
}

func (b *Buffer) UnmarshalBinary(buf []byte) error {
	b.ID = buf[0]
	return nil
}

// Window represents a Neovim Window
//
// Multiple goroutines may invoke methods on a Window simultaneously
type Window struct {
	ID     uint8
	client *Client
}

func (b *Window) ExtensionType() int8 {
	return typeWindow
}

func (b *Window) Len() int {
	return 1
}

func (b *Window) MarshalBinaryTo(buf []byte) error {
	buf[0] = b.ID
	return nil
}

func (b *Window) UnmarshalBinary(buf []byte) error {
	b.ID = buf[0]
	return nil
}

// Tabpage represents a Neovim Tabpage
//
// Multiple goroutines may invoke methods on a Tabpage simultaneously
type Tabpage struct {
	ID     uint8
	client *Client
}

func (b *Tabpage) ExtensionType() int8 {
	return typeTabpage
}

func (b *Tabpage) Len() int {
	return 1
}

func (b *Tabpage) MarshalBinaryTo(buf []byte) error {
	buf[0] = b.ID
	return nil
}

func (b *Tabpage) UnmarshalBinary(buf []byte) error {
	b.ID = buf[0]
	return nil
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
