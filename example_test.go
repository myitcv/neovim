// Copyright 2014 Paul Jolly <paul@myitcv.org.uk>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package neovim_test

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/juju/errors"
	"github.com/myitcv/neovim"
)

func ExampleClient_GetCurrentBuffer() {
	cmd := exec.Command(os.Getenv("NEOVIM_BIN"), "-u", "/dev/null")
	cmd.Dir = "/tmp"

	client, err := neovim.NewCmdClient(neovim.NullInitMethod, cmd, nil)
	if err != nil {
		log.Fatalf("Could not create new client: %v", errors.Details(err))
	}
	client.Run()

	b, err := client.GetCurrentBuffer()
	if err != nil {
		log.Fatalf("Could not get current buffer: %v", errors.Details(err))
	}
	n, err := b.GetName()
	if err != nil {
		log.Fatalf("Could not get name for buffer %v: %v", b, errors.Details(err))
	}
	fmt.Printf("Current buffer is: %v %v\n", b.ID, n)

	err = client.Close()
	if err != nil {
		log.Fatalf("Could not close client: %v\n", err)
	}

	// Output:
	// Current buffer is: 2
}
