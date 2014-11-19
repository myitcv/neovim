package example

import (
	"log"

	// import the neovim package
	"github.com/myitcv/neovim"
)

type Example struct {
	client *neovim.Client
	log    neovim.Logger
}

func (n *Example) Init(c *neovim.Client, l neovim.Logger) error {
	n.client = c
	n.log = l

	g := n.NewGetANumberDecoder()
	err := n.client.RegisterSyncRequestHandler("GetANumber", g)
	if err != nil {
		n.log.Fatalf("Could not register sync request handler: %v\n", err)
	}

	l.Println("**************************")

	ch, d := n.NewBufCreateSub()
	err = n.client.RegisterAsyncRequestHandler(":autocmd:BufCreate:*", d)
	if err != nil {
		n.log.Fatalf("Could not register async request handler: %v\n", err)
	}
	go n.subLoop(ch)

	return nil
}

func (n *Example) Shutdown() error {
	return nil
}

type BufCreate struct {
	BufNumber int
}

func (n *Example) GetANumber() (int, error, error) {
	log.Printf("Got a request to getANumber\n")
	return 42, nil, nil
}

func (n *Example) subLoop(ch chan *BufCreate) {
	for {
		select {
		case <-n.client.KillChannel:
			return
		case v := <-ch:
			n.log.Printf("Got a %v event\n", v)
		}
	}
}
