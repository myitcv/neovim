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

	client, err := neovim.NewCmdClient(cmd)
	if err != nil {
		log.Fatalf("Could not create new client: %v", errgo.Details(err))
	}

	topic := "topic1" // corresponds to the topic used in Neovim's send_event()
	respChan := make(chan neovim.SubscriptionEvent)
	errChan := make(chan error)

	client.SubChan <- neovim.Subscription{
		Topic:  topic,
		Events: respChan,
		Error:  errChan,
	}
	err = <-errChan
	if err != nil {
		log.Fatalf("Could not register subscription handler: %v", errgo.Details(err))
	}

	err = client.Subscribe(topic)
	if err != nil {
		log.Fatalf("Could not subscribe to topic %v: %v", topic, errgo.Details(err))
	}

	// Now wait to receive a notification on respChan
	// resp := <-respChan

	err = client.Close()
	if err != nil {
		log.Fatalf("Could not close client: %v\n", err)
	}

	// Output:
}

func ExampleClient_GetCurrentBuffer() {
	cmd := exec.Command(os.Getenv("NEOVIM_BIN"), "-u", "/dev/null")
	cmd.Dir = "/tmp"

	client, err := neovim.NewCmdClient(cmd)
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
