package example

import

// import the neovim package
"github.com/myitcv/neovim"

type Example struct {
	client *neovim.Client
	log    neovim.Logger
}

func (n *Example) Init(c *neovim.Client, l neovim.Logger) error {
	n.client = c
	n.log = l

	err := n.client.RegisterSyncRequestHandler(":function:GetANumber", n.newGetANumberResponder)
	if err != nil {
		n.log.Fatalf("Could not register sync request handler: %v\n", err)
	}

	ch, d := n.newBufCreateChanHandler()
	err = n.client.RegisterAsyncRequestHandler(":autocmd:BufCreate:*", d)
	if err != nil {
		n.log.Fatalf("Could not register async request handler: %v\n", err)
	}

	n.log.Println("*****************")
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
