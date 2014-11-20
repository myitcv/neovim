// Copyright 2014 Paul Jolly <paul@myitcv.org.uk>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package neovim_test

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"

	"testing"

	"github.com/juju/errors"
	"github.com/myitcv/neovim"
	"github.com/vmihailenco/msgpack"

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

var _ = Suite(NewNeovimTest())

func NewNeovimTest() neovimTester {
	return &NeovimTest{}
}

func (t *NeovimTest) SetUpTest(c *C) {
	// now start the process and wait for the socket file to be created
	t.nvim = exec.Command(os.Getenv("NEOVIM_BIN"), "-u", "/dev/null")
	t.nvim.Dir = "/tmp"

	// dev_null, err := os.OpenFile("/dev/null", os.O_WRONLY, 0644)
	// if err != nil {
	// 	log.Fatalf("Could not open /dev/null: %v\n", err)
	// }
	// underlying := log.New(dev_null, "", 0)
	underlying := log.New(os.Stdout, "", 0)
	logger := newStackLogger(underlying)

	// now we can create a new client
	client, err := neovim.NewCmdClient(neovim.NullInitMethod, t.nvim, logger)
	if err != nil {
		log.Fatalf("Could not setup client: %v", errors.Details(err))
	}
	client.PanicOnError = true
	t.client = client
	client.Run()
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

// func (t *NeovimTest) TestClientGetAPIInfo(c *C) {
// 	chanID, api, _ := t.client.GetAPIInfo()
// 	c.Assert(chanID > 0, Equals, true)
// 	c.Assert(api, NotNil)
// }

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

func (t *NeovimTest) BenchmarkBufferGetLine(c *C) {
	// TODO
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

// func (t *NeovimTest) TestClientSubscribe(c *C) {
// 	topic := "event1"
// 	val := []int64{1, 2, 3}

// 	vals := make([]string, len(val))
// 	for i := range val {
// 		vals[i] = fmt.Sprintf("%v", val[i])
// 	}

// 	sub, _ := t.client.Subscribe(topic)

// 	// This is intentionally notifying on the broadcast channel
// 	// Doing so exercises the effect of the underlying call to
// 	// vim_subscribe which notifies on all channels that have subscribed
// 	// to a given topic
// 	command := fmt.Sprintf(`call rpcnotify(0, "%v", %v)`, topic, strings.Join(vals, ","))
// 	_ = t.client.Command(command)
// 	resp := <-sub.Events
// 	c.Assert(len(val), Equals, len(resp.Value))
// 	for i := range resp.Value {
// 		c.Assert(resp.Value[i], Equals, val[i])
// 	}
// 	_ = t.client.Unsubscribe(sub)
// 	_ = t.client.Command(command)

// 	// try and resubsubscribe; if there is an unhandled notification this will block
// 	// forever and fail the tests
// 	sub, _ = t.client.Subscribe(topic)
// }

func newGetANumberDecoder() neovim.SyncDecoder {
	res := &getANumberDecoder{}
	return res
}

type getANumberDecoder struct{}

type getANumberRunner struct{}

type getANumberEncoder struct {
	i int
}

func (g *getANumberDecoder) Decode(dec *msgpack.Decoder) (neovim.SyncRunner, error) {
	l, err := dec.DecodeSliceLen()
	if err != nil {
		return nil, err
	}

	if l != 1 {
		return nil, errors.Errorf("Expected 1 argument, not %v", l)
	}

	l, err = dec.DecodeSliceLen()
	if err != nil {
		return nil, err
	}

	if l != 0 {
		return nil, errors.Errorf("Expected 0 argument, not %v", l)
	}

	return &getANumberRunner{}, nil
}

func (g *getANumberRunner) Run() (neovim.Encoder, error, error) {
	res := &getANumberEncoder{}

	i, mErr, err := getANumber()

	if err != nil || mErr != nil {
		return nil, mErr, err
	}

	res.i = i

	return res, nil, nil
}

func (g *getANumberEncoder) Encode(enc *msgpack.Encoder) error {
	err := enc.EncodeInt(g.i)
	if err != nil {
		return err
	}

	return nil
}

func getANumber() (int, error, error) {
	return 42, nil, nil
}

func (t *NeovimTest) TestFunctionOnChannel(c *C) {
	t.client.RegisterSyncRequestHandler("GetANumber", newGetANumberDecoder())
	topic := "GetANumber"
	commandDef := fmt.Sprintf(`call rpc#define#FunctionOnChannel(1, "%v", 1, "%v", {})`, topic, topic)
	_ = t.client.Command(commandDef)
	res, _ := t.client.Eval(`GetANumber()`)
	c.Assert(res, Equals, int64(42))
}

// func (t *NeovimTest) TestAutocmdOnChannel(c *C) {
// 	cb, _ := t.client.GetCurrentBuffer()
// 	cbn, _ := cb.GetNumber()
// 	topic := fmt.Sprintf("Buffer[%v].TextChanged", cb.ID)
// 	sub, _ := t.client.Subscribe(topic)
// 	commandDef := fmt.Sprintf(`call rpc#define#AutocmdOnChannel(0, "%v", 0, "TextChanged", {"pattern": "<buffer=%v>"})`, topic, cbn)
// 	t.client.Command(commandDef)
// 	cb.Insert(0, []string{"This is a test"})
// 	_ = <-sub.Events
// }

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
	_ = t.client.Command("let x=0 | let y=2")
	_v, _ := t.client.Eval("[x,y]")
	v := _v.([]interface{})
	comp := []int64{0, 2}
	c.Assert(len(v), Equals, len(comp))
	for i := range v {
		c.Assert(comp[i], Equals, v[i].(int64))
	}
}

// func (t *NeovimTest) TestRegisterRequestHandler(c *C) {
// 	err := t.client.RegisterRequestHandler("my_first_method", func(args []interface{}) ([]interface{}, error) {
// 		return []interface{}{5}, nil
// 	})
// 	c.Assert(err, IsNil)
// 	res, err := t.client.Eval(fmt.Sprintf("rpcrequest(%v, 'my_first_method')", t.client.ChannelID))
// 	c.Assert(err, IsNil)
// 	c.Assert(res, Equals, int64(5))
// }

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

// func (t *NeovimTest) TestMultiClientSubscribe(c *C) {
// 	topic := "event1"
// 	var subDone, unsubDone, doneDone sync.WaitGroup
// 	var check int64

// 	number := 1000

// 	for i := 1; i <= number; i++ {
// 		subDone.Add(1)
// 		if i%2 == 1 {
// 			unsubDone.Add(1)
// 		}
// 		doneDone.Add(1)
// 		go func(topic string, n int, check *int64) {
// 			sub, _ := t.client.Subscribe(topic)
// 			subDone.Done()
// 			resp := <-sub.Events
// 			val := resp.Value[0].(int64)
// 			atomic.AddInt64(check, val)
// 			if n%2 == 0 {
// 				// listen again
// 				resp := <-sub.Events
// 				val := resp.Value[0].(int64)
// 				atomic.AddInt64(check, val)
// 			} else {
// 				// unsubscribe
// 				t.client.Unsubscribe(sub)
// 				unsubDone.Done()
// 			}
// 			doneDone.Done()
// 		}(topic, i, &check)
// 	}

// 	subDone.Wait()

// 	command := fmt.Sprintf(`call rpcnotify(0, "%v", 1)`, topic)
// 	_ = t.client.Command(command)

// 	unsubDone.Wait()

// 	_ = t.client.Command(command)

// 	doneDone.Wait()

// 	c.Assert(atomic.LoadInt64(&check), Equals, int64(number+number/2))
// }

func (t *NeovimTest) TestBufferDelLine(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkBufferDelLine(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestBufferGetLineSlice(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkBufferGetLineSlice(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestBufferGetMark(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkBufferGetMark(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestBufferGetName(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkBufferGetName(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestBufferGetNumber(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkBufferGetNumber(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestBufferGetOption(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkBufferGetOption(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestBufferGetVar(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkBufferGetVar(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkBufferInsert(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestBufferIsValid(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkBufferIsValid(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestBufferLineCount(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkBufferLineCount(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkBufferSetLine(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestBufferSetLineSlice(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkBufferSetLineSlice(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestBufferSetName(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkBufferSetName(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestBufferSetOption(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkBufferSetOption(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestBufferSetVar(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkBufferSetVar(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestTabpageGetVar(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkTabpageGetVar(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestTabpageGetWindow(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkTabpageGetWindow(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestTabpageGetWindows(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkTabpageGetWindows(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestTabpageIsValid(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkTabpageIsValid(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestTabpageSetVar(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkTabpageSetVar(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestClientChangeDirectory(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkClientChangeDirectory(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestClientCommand(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkClientCommand(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestClientCommandOutput(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkClientCommandOutput(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestClientDelCurrentLine(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkClientDelCurrentLine(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestClientErrWrite(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkClientErrWrite(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestClientEval(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkClientEval(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestClientFeedkeys(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkClientFeedkeys(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkClientGetBuffers(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkClientGetCurrentBuffer(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestClientGetCurrentLine(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkClientGetCurrentLine(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestClientGetCurrentTabpage(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkClientGetCurrentTabpage(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestClientGetCurrentWindow(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkClientGetCurrentWindow(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestClientGetOption(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkClientGetOption(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestClientGetTabpages(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkClientGetTabpages(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestClientGetVar(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkClientGetVar(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestClientGetVvar(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkClientGetVvar(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestClientGetWindows(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkClientGetWindows(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestClientInput(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkClientInput(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestClientListRuntimePaths(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkClientListRuntimePaths(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestClientOutWrite(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkClientOutWrite(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestClientReplaceTermcodes(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkClientReplaceTermcodes(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestClientReportError(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkClientReportError(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestClientSetCurrentBuffer(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkClientSetCurrentBuffer(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestClientSetCurrentLine(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkClientSetCurrentLine(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestClientSetCurrentTabpage(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkClientSetCurrentTabpage(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestClientSetCurrentWindow(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkClientSetCurrentWindow(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestClientSetOption(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkClientSetOption(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestClientSetVar(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkClientSetVar(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestClientStrwidth(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkClientStrwidth(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestClientsubscribe(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkClientsubscribe(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestClientunsubscribe(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkClientunsubscribe(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestWindowGetBuffer(c *C) {
	cw, _ := t.client.GetCurrentWindow()
	cw.GetBuffer()
}
func (t *NeovimTest) BenchmarkWindowGetBuffer(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestWindowGetCursor(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkWindowGetCursor(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestWindowGetHeight(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkWindowGetHeight(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestWindowGetOption(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkWindowGetOption(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestWindowGetPosition(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkWindowGetPosition(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestWindowGetTabpage(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkWindowGetTabpage(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestWindowGetVar(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkWindowGetVar(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestWindowGetWidth(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkWindowGetWidth(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestWindowIsValid(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkWindowIsValid(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestWindowSetCursor(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkWindowSetCursor(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestWindowSetHeight(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkWindowSetHeight(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestWindowSetOption(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkWindowSetOption(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestWindowSetVar(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkWindowSetVar(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestWindowSetWidth(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkWindowSetWidth(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestClientUnsubscribe(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkClientUnsubscribe(c *C) {
	// implementation pending
}
func (t *NeovimTest) TestClientSubscribe(c *C) {
	// implementation pending
}
func (t *NeovimTest) BenchmarkClientSubscribe(c *C) {
	// implementation pending
}
