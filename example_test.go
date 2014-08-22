// Copyright 2014 Paul Jolly <paul@myitcv.org.uk>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package neovim_test

import (
	"fmt"
	"log"
	"net"

	"github.com/juju/errgo"
	"github.com/myitcv/neovim"
)

func ExampleNewUnixClient() {
	_, err := neovim.NewUnixClient("unix", nil, &net.UnixAddr{Name: "/tmp/neovim"})
	if err != nil {
		log.Fatalf("Could not create new Unix client: %v", errgo.Details(err))
	}
	// Output:
}

func ExampleSubscription() {
	client, err := neovim.NewUnixClient("unix", nil, &net.UnixAddr{Name: "/tmp/neovim"})
	if err != nil {
		log.Fatalf("Could not create new Unix client: %v", errgo.Details(err))
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
	// Output:
}

func ExampleClient_GetCurrentBuffer() {
	client, err := neovim.NewUnixClient("unix", nil, &net.UnixAddr{Name: "/tmp/neovim"})
	if err != nil {
		log.Fatalf("Could not create new Unix client: %v", errgo.Details(err))
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
	// Output:
	// Current buffer is: 2
}
