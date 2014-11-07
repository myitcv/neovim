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
	"sync/atomic"

	"testing"

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
	startListen chan chan struct{}
}

func Test(t *testing.T) { TestingT(t) }

var _ = Suite(&NeovimTest{})

func (t *NeovimTest) SetUpTest(c *C) {
	// now start the process and wait for the socket file to be created
	t.nvim = exec.Command(os.Getenv("NEOVIM_BIN"), "-u", "/dev/null")
	t.nvim.Dir = "/tmp"

	underlying := log.New(os.Stdout, "", 0)
	logger := neovim.NewStackLogger(underlying)

	// now we can create a new client
	client, err := neovim.NewCmdClient(t.nvim, logger)
	if err != nil {
		log.Fatalf("Could not setup client: %v", errgo.Details(err))
	}

	// TODO need to handle nvim subprocess bombing out...

	// this is important; all tests below ignore errors...
	client.PanicOnError = true
	t.client = client
}

func (t *NeovimTest) TearDownTest(c *C) {
	err := t.client.Close()
	if err != nil {
		log.Fatalf("Could not close client: %v\n", err)
	}
	state, err := t.nvim.Process.Wait()
	if err != nil {
		log.Fatalf("Process did not exit cleanly: %v, %v\n", err, state)
	}
}

func (t *NeovimTest) TestClientGetBuffers(c *C) {
	ba, _ := t.client.GetBuffers()
	c.Assert(ba, NotNil)
}

func (t *NeovimTest) TestClientGetAPIInfo(c *C) {
	chanID, api, _ := t.client.GetAPIInfo()
	c.Assert(chanID > 0, Equals, true)
	c.Assert(api, NotNil)
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
	l, _ := b.LineCount()
	c.Assert(l, Equals, 1)
}

func (t *NeovimTest) TestBufferGetLine(c *C) {
	// See TestBufferSetLine
}

func (t *NeovimTest) TestBufferSetLine(c *C) {
	b, _ := t.client.GetCurrentBuffer()

	line1 := "This is line 1"
	_ = b.SetLine(0, line1)
	l, _ := b.GetLine(0)
	length, _ := b.LineCount()

	c.Assert(length, Equals, 1)
	c.Assert(l, Equals, line1)

	_ = b.DelLine(0)
	l, _ = b.GetLine(0)
	length, _ = b.LineCount()

	c.Assert(length, Equals, 1)
	c.Assert(l, Equals, "")
}

func (t *NeovimTest) TestEval(c *C) {
	res, _ := t.client.Eval(`4`)
	resInt64 := res.(int64)
	c.Assert(resInt64 > 0, Equals, true)
}

func (t *NeovimTest) TestClientSubscribe(c *C) {
	topic := "event1"
	val := []int64{1, 2, 3}

	vals := make([]string, len(val))
	for i := range val {
		vals[i] = fmt.Sprintf("%v", val[i])
	}

	sub, _ := t.client.Subscribe(topic)
	command := fmt.Sprintf(`call rpcnotify(0, "%v", %v)`, topic, strings.Join(vals, ","))
	_ = t.client.Command(command)
	resp := <-sub.Events
	c.Assert(len(val), Equals, len(resp.Value))
	for i := range resp.Value {
		c.Assert(resp.Value[i], Equals, val[i])
	}
	_ = t.client.Unsubscribe(sub)
	_ = t.client.Command(command)

	// try and resubsubscribe; if there is an unhandled notification this will block
	// forever and fail the tests
	sub, _ = t.client.Subscribe(topic)
}

func (t *NeovimTest) TestGetSetLine(c *C) {
	cl := "This is our line"
	t.client.SetCurrentLine(cl)
	c_cl, _ := t.client.GetCurrentLine()
	c.Assert(c_cl, Equals, cl)
}

func (t *NeovimTest) TestGetLineSlice(c *C) {
	cb, _ := t.client.GetCurrentBuffer()
	lc, _ := cb.LineCount()
	c.Assert(lc, Equals, 1)

	new_lines := []string{"This is", "a test"}

	cb.SetLineSlice(0, -1, true, true, new_lines)
	lines, _ := cb.GetLineSlice(0, -1, true, true)
	c.Assert(lines, NotNil)
	c.Assert(len(lines), Equals, len(new_lines))
	for i := range new_lines {
		c.Assert(lines[i], Equals, new_lines[i])
	}

	lc, _ = cb.LineCount()
	c.Assert(lc, Equals, 2)

}

func (t *NeovimTest) TestBufferInsert(c *C) {
	// append the lines to the end of the buffer
	// cb.SetLineSlice(3, -1, true, true, new_lines)
	// lc, _ = cb.LineCount()
	// lines, _ = cb.GetLineSlice(0, -1, true, true)
	// fmt.Println(lines)
	// c.Assert(lc, Equals, 4)
}

func (t *NeovimTest) TestNumberEval(c *C) {
	i, _ := t.client.Eval("1")
	// according to the Neovim API all numbers are int64
	c.Assert(i, Equals, int64(1))
}

func (t *NeovimTest) TestArrayEval(c *C) {
	_ = t.client.Command("let x=1 | let y=2")
	_v, _ := t.client.Eval("[x,y]")
	v := _v.([]interface{})
	comp := []int64{1, 2}
	c.Assert(len(v), Equals, len(comp))
	for i := range v {
		c.Assert(comp[i], Equals, v[i].(int64))
	}
}

func (t *NeovimTest) TestRegisterRequestHandler(c *C) {
	err := t.client.RegisterRequestHandler("my_first_method", func(args []interface{}) ([]interface{}, error) {
		return []interface{}{5}, nil
	})
	c.Assert(err, IsNil)
	res, err := t.client.Eval(fmt.Sprintf("rpcrequest(%v, 'my_first_method')", t.client.ChannelID))
	c.Assert(err, IsNil)
	c.Assert(res, Equals, int64(5))
}

// func (t *NeovimTest) TestRegisterProvider(c *C) {
// 	err := t.client.RegisterProvider("my_first_method", func(args []interface{}) ([]interface{}, error) {
// 		return nil, nil
// 	})
// 	c.Assert(err, IsNil)
// 	// _ = t.client.Command("call provider_call('my_first_method')")
// }

func (t *NeovimTest) BenchmarkCommandAndEval(c *C) {
	for i := 0; i < c.N; i++ {
		_ = t.client.Command(fmt.Sprintf("let x=%v", i))
		v, _ := t.client.Eval("x")
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
		id, _ := t.client.Eval(fmt.Sprintf("matchadd('String', '\\%%%vl\\%%2c\\_.\\{8\\}')", i))
		c.Assert(id, NotNil)
	}
}

func (t *NeovimTest) BenchmarkGetBufferContents(c *C) {
	// TODO this needs to first fill the buffer with suitable contents
	cb, _ := t.client.GetCurrentBuffer()
	c.ResetTimer()
	for i := 0; i < c.N; i++ {
		_, _ = cb.GetLineSlice(0, -1, true, true)
	}
}

func (t *NeovimTest) TestMultiClientSubscribe(c *C) {
	topic := "event1"
	var subDone, unsubDone, doneDone sync.WaitGroup
	var check int64

	number := 1000

	for i := 1; i <= number; i++ {
		subDone.Add(1)
		if i%2 == 1 {
			unsubDone.Add(1)
		}
		doneDone.Add(1)
		go func(topic string, n int, check *int64) {
			sub, _ := t.client.Subscribe(topic)
			subDone.Done()
			resp := <-sub.Events
			val := resp.Value[0].(int64)
			atomic.AddInt64(check, val)
			if n%2 == 0 {
				// listen again
				resp := <-sub.Events
				val := resp.Value[0].(int64)
				atomic.AddInt64(check, val)
			} else {
				// unsubscribe
				t.client.Unsubscribe(sub)
				unsubDone.Done()
			}
			doneDone.Done()
		}(topic, i, &check)
	}

	subDone.Wait()

	command := fmt.Sprintf(`call rpcnotify(0, "%v", 1)`, topic)
	_ = t.client.Command(command)

	unsubDone.Wait()

	_ = t.client.Command(command)

	doneDone.Wait()

	c.Assert(atomic.LoadInt64(&check), Equals, int64(number+number/2))
}
