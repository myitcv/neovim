// Copyright 2014 Paul Jolly <paul@myitcv.org.uk>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package neovim_test

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
	"sync"
	"testing"

	"github.com/go-fsnotify/fsnotify"
	"github.com/juju/errgo"
	"github.com/myitcv/neovim"

	. "gopkg.in/check.v1"
	// "github.com/vmihailenco/msgpack"
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

// func (t *NeovimTest) SetUpSuite(c *C) {
// 	watcher, err := fsnotify.NewWatcher()
// 	if err != nil {
// 		log.Fatalf("Could not create a new watcher: %v", err)
// 	}
// 	t.watcher = watcher

// 	t.startListen = make(chan chan struct{})

// 	go func() {
// 		var respChan chan struct{}
// 		for {
// 			select {
// 			case respChan = <-t.startListen:
// 				respChan <- struct{}{}
// 			case event := <-t.watcher.Events:
// 				fmt.Println("Got an event")
// 				if respChan == nil {
// 					log.Fatalf("Got unexpected event; not listening but got %v\n", event)
// 				}
// 				if event.Op&fsnotify.Create == fsnotify.Create {
// 					respChan <- struct{}{}
// 					respChan = nil
// 				}
// 			case err := <-t.watcher.Errors:
// 				log.Fatalf("Got an error in the watcher: %v\n", err)
// 			}
// 		}
// 	}()
// }

// func (t *NeovimTest) TearDownSuite(c *C) {
// 	err := t.watcher.Close()
// 	if err != nil {
// 		log.Fatalf("Could not cleanly shut down watcher: %v\n", err)
// 	}
// }

func (t *NeovimTest) SetUpTest(c *C) {
	la := os.Getenv("NEOVIM_LISTEN_ADDRESS")

	// cur_t := time.Now()
	// la := fmt.Sprintf("/tmp/neovim.%v%v", cur_t.Unix(), cur_t.Nanosecond())
	// e_la := "NEOVIM_LISTEN_ADDRESS=" + la

	// now start the process and wait for the socket file to be created
	// t.nvim = exec.Command("nvim", "-u /dev/null")
	// new_env := os.Environ()

	// found := false
	// for i, _ := range new_env {
	// 	if strings.HasPrefix(new_env[i], "NEOVIM_LISTEN_ADDRESS=") {
	// 		found = true
	// 		new_env[i] = e_la
	// 	}
	// }
	// if !found {
	// 	new_env = append(new_env, e_la)
	// }
	// t.nvim.Env = new_env

	// done_chan := make(chan struct{})
	// t.startListen <- done_chan
	// <-done_chan
	// t.watcher.Add(la)
	// err := t.nvim.Start()
	// if err != nil {
	// 	log.Fatalf("Could not start nvim instance: %v\n", err)
	// }
	// <-done_chan
	// t.watcher.Remove(la)

	// fmt.Println("Starting test")

	// now we can create a new client
	client, err := neovim.NewUnixClient("unix", nil, &net.UnixAddr{Name: la})
	if err != nil {
		log.Fatalf("Could not setup client: %v", errgo.Details(err))
	}
	t.client = client
}

// func (t *NeovimTest) TearDownTest(c *C) {
// 	err := t.nvim.Process.Kill()
// 	if err != nil {
// 		log.Fatalf("Could not kill nvim instance: %v\n", err)
// 	}
// }

func (t *NeovimTest) TestClientGetBuffers(c *C) {
	ba, err := t.client.GetBuffers()
	c.Assert(err, IsNil)
	c.Assert(ba, NotNil)
}

func (t *NeovimTest) TestConcurrentClientGetBuffers(c *C) {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ba, err := t.client.GetBuffers()
			c.Assert(err, IsNil)
			c.Assert(ba, NotNil)
		}()
	}
	wg.Wait()
}

func (t *NeovimTest) TestClientGetCurrentBuffer(c *C) {
	_, err := t.client.GetCurrentBuffer()
	c.Assert(err, IsNil)
}

func (t *NeovimTest) TestBufferGetLength(c *C) {
	b, err := t.client.GetCurrentBuffer()
	c.Assert(err, IsNil)
	l, err := b.GetLength()
	c.Assert(err, IsNil)
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
