// Copyright 2014 Paul Jolly <paul@myitcv.org.uk>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package neovim_test

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"

	"testing"

	"github.com/go-fsnotify/fsnotify"
	"github.com/juju/errgo"
	"github.com/myitcv/neovim"

	. "gopkg.in/check.v1"
)

type mpResponse struct {
	t     int
	msgID uint32
	err   interface{}
}

type NeovimTest struct {
	client      *neovim.Client
	nvim        *exec.Cmd
	watcher     *fsnotify.Watcher
	startListen chan chan struct{}
}

func Test(t *testing.T) { TestingT(t) }

var _ = Suite(&NeovimTest{})

func (t *NeovimTest) SetUpTest(c *C) {
	// now start the process and wait for the socket file to be created
	t.nvim = exec.Command(os.Getenv("NEOVIM_BIN"), "-u", "/dev/null")
	t.nvim.Dir = "/tmp"

	// now we can create a new client
	client, err := neovim.NewCmdClient(t.nvim)
	if err != nil {
		log.Fatalf("Could not setup client: %v", errgo.Details(err))
	}

	// TODO need to handle nvim subprocess bombing out...

	// this is important; all tests below ignore errors...
	client.PanicOnError = true
	t.client = client
}

func (t *NeovimTest) TearDownTest(c *C) {
	done := make(chan struct{})
	go func() {
		err := t.nvim.Wait()
		if err != nil {
			log.Fatalf("Process did not exit cleanly: %v\n", err)
		}
		done <- struct{}{}
	}()
	err := t.client.Close()
	if err != nil {
		log.Fatalf("Could not close client: %v\n", err)
	}
	<-done
}

func (t *NeovimTest) TestClientGetBuffers(c *C) {
	ba, _ := t.client.GetBuffers()
	c.Assert(ba, NotNil)
}

func (t *NeovimTest) TestConcurrentClientGetBuffers(c *C) {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ba, _ := t.client.GetBuffers()
			c.Assert(ba, NotNil)
		}()
	}
	wg.Wait()
}

func (t *NeovimTest) TestClientGetCurrentBuffer(c *C) {
	cb, _ := t.client.GetCurrentBuffer()
	c.Assert(c, NotNil)
	c.Assert(cb.ID > 0, Equals, true)
}

func (t *NeovimTest) TestBufferGetLength(c *C) {
	b, _ := t.client.GetCurrentBuffer()
	l, _ := b.GetLength()
	c.Assert(l, Equals, 1)
}

func (t *NeovimTest) TestBufferSetGetLine(c *C) {
	b, err := t.client.GetCurrentBuffer()
	c.Assert(err, IsNil)
	val := "This is line 1"
	err = b.SetLine(0, val)
	c.Assert(err, IsNil)
	l, err := b.GetLine(0)
	c.Assert(err, IsNil)
	c.Assert(l, Equals, val)
	err = b.DelLine(0)
	c.Assert(err, IsNil)
	length, err := b.GetLength()
	c.Assert(err, IsNil)
	c.Assert(length, Equals, 1)
}

func (t *NeovimTest) TestEval(c *C) {
	res, err := t.client.Eval(`4`)
	c.Assert(err, IsNil)
	resInt64 := res.(int64)
	c.Assert(resInt64 > 0, Equals, true)
}

func (t *NeovimTest) TestClientSubscribe(c *C) {
	respChan := make(chan neovim.SubscriptionEvent)
	errChan := make(chan error)
	topic := "event1"
	val := []interface{}{1, 2, 3}

	vals := make([]string, len(val))
	for i := range val {
		vals[i] = fmt.Sprintf("%v", val[i])
	}

	t.client.SubChan <- neovim.Subscription{
		Topic:  topic,
		Events: respChan,
		Error:  errChan,
	}
	err := <-errChan
	c.Assert(err, IsNil)
	err = t.client.Subscribe(topic)
	c.Assert(err, IsNil)
	command := fmt.Sprintf(`call send_event(0, "%v", [%v])`, topic, strings.Join(vals, ","))
	err = t.client.Command(command)
	c.Assert(err, IsNil)
	resp := <-respChan
	c.Assert(resp, NotNil)
}

func (t *NeovimTest) TestGetSlice(c *C) {
	cb, err := t.client.GetCurrentBuffer()
	c.Assert(err, IsNil)
	lines, err := cb.GetSlice(0, -1, true, true)
	c.Assert(err, IsNil)
	c.Assert(lines, NotNil)
}

func (t *NeovimTest) TestNumberEval(c *C) {
	_, err := t.client.Eval("127")
	c.Assert(err, IsNil)
}

func (t *NeovimTest) TestArrayEval(c *C) {
	err := t.client.Command("let x=1 | let y=2")
	_v, err := t.client.Eval("[x,y]")
	v := _v.([]interface{})
	c.Assert(err, IsNil)
	comp := []int64{1, 2}
	c.Assert(len(v), Equals, len(comp))
	for i := range v {
		c.Assert(comp[i], Equals, v[i].(int64))
	}
}

func (t *NeovimTest) BenchmarkCommandAndEval(c *C) {
	for i := 0; i < c.N; i++ {
		err := t.client.Command(fmt.Sprintf("let x=%v", i))
		c.Assert(err, IsNil)
		v, err := t.client.Eval("x")
		c.Assert(err, IsNil)
		switch v.(type) {
		case int64:
			c.Assert(v, Equals, int64(i))
		case uint64:
			c.Assert(v, Equals, uint64(i))
		default:
			panic("Unkown type")
		}
	}
}

func (t *NeovimTest) BenchmarkMatchAddEmptyBuffer(c *C) {
	for i := 0; i < c.N; i++ {
		id, err := t.client.Eval(fmt.Sprintf("matchadd('String', '\\%%%vl\\%%2c\\_.\\{8\\}')", i))
		c.Assert(err, IsNil)
		c.Assert(id, NotNil)
	}
}

func (t *NeovimTest) BenchmarkGetBufferContents(c *C) {
	// TODO this needs to first fill the buffer with suitable contents
	cb, _ := t.client.GetCurrentBuffer()
	c.ResetTimer()
	for i := 0; i < c.N; i++ {
		_, _ = cb.GetSlice(0, -1, true, true)
	}
}
