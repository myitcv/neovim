// Copyright 2014 Paul Jolly <paul@myitcv.org.uk>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package neovim_test

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/juju/errgo"
	"github.com/myitcv/neovim"
)

func ExampleSubscription() {
	cmd := exec.Command(os.Getenv("NEOVIM_BIN"), "-u", "/dev/null")
	cmd.Dir = "/tmp"

	client, err := neovim.NewCmdClient(cmd, nil)
	if err != nil {
		log.Fatalf("Could not create new client: %v", errgo.Details(err))
	}
	client.PanicOnError = true

	topic := "topic1"
	sub, err := client.Subscribe(topic)
	if err != nil {
		log.Fatalf("Could not subscribe to topic %v: %v", topic, errgo.Details(err))
	}

	unsubbed := make(chan struct{})
	done := make(chan struct{})
	received := make(chan struct{})

	go func() {
	ForLoop:
		for {
			select {
			case e := <-sub.Events:
				if e == nil {
					break ForLoop
				}
				fmt.Printf("We got %v\n", e.Value)
				received <- struct{}{}
			}
		}
		done <- struct{}{}
	}()

	command := fmt.Sprintf(`call rpcnotify(0, "%v", 1)`, topic)
	_ = client.Command(command)

	<-received

	go func() {

		_ = client.Unsubscribe(sub)
		unsubbed <- struct{}{}
	}()

	<-done

	<-unsubbed

	_ = client.Close()

	// Output:
	// We got [1]
}

func ExampleClient_GetCurrentBuffer() {
	cmd := exec.Command(os.Getenv("NEOVIM_BIN"), "-u", "/dev/null")
	cmd.Dir = "/tmp"

	client, err := neovim.NewCmdClient(cmd, nil)
	if err != nil {
		log.Fatalf("Could not create new client: %v", errgo.Details(err))
	}
	b, err := client.GetCurrentBuffer()
	if err != nil {
		log.Fatalf("Could not get current buffer: %v", errgo.Details(err))
	}
	n, err := b.GetName()
	if err != nil {
		log.Fatalf("Could not get name for buffer %v: %v", b, errgo.Details(err))
	}
	fmt.Printf("Current buffer is: %v %v\n", b.ID, n)

	err = client.Close()
	if err != nil {
		log.Fatalf("Could not close client: %v\n", err)
	}

	// Output:
	// Current buffer is: 2
}
