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
	l.Println("Example Init")

	g := n.NewGetANumberDecoder()
	err := n.client.RegisterSyncRequestHandler("GetANumber", g)
	if err != nil {
		n.log.Fatalf("Could not register sync request handler: %v\n", err)
	}

	ch, d := n.NewBufCreateSub()
	err = n.client.RegisterAsyncRequestHandler("BufCreate", d)
	if err != nil {
		n.log.Fatalf("Could not register async request handler: %v\n", err)
	}
	go n.subLoop(ch)

	l.Println("Example Init done")
	return nil
}

func (n *Example) Shutdown() error {
	return nil
}

type BufCreate struct {
	BufNumber int
}

func (n *Example) GetANumber() (int, error) {
	log.Printf("Got a request to getANumber\n")
	return 42, nil
}

func (n *Example) subLoop(ch chan BufCreate) {
	for {
		select {
		case <-n.client.KillChannel:
			return
		case v := <-ch:
			n.log.Printf("Got a %v event\n", v)
		}
	}
}
