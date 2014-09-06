package example

import (
	"fmt"
	"log"
	"os"

	// import the neovim package
	"github.com/myitcv/neovim"
)

type Example struct {
	client *neovim.Client
	log    neovim.Logger
}

func (n *Example) Init(c *neovim.Client, l neovim.Logger) error {
	n.client = c

	// Tell Neovim to broadcast TextChanged*
	topic := "text_changed"
	com := fmt.Sprintf(`au TextChanged,TextChangedI <buffer> call send_event(0, "%v", [])`, topic)
	_ = c.Command(com)

	// Setup a subscription
	sub, _ := c.Subscribe(topic)
	go n.subLoop(sub.Events)

	// Handle an RPC request from Neovim
	n.client.RegisterProvider("get_a_number", n.getANumber)

	return nil
}

func (n *Example) getANumber(args []interface{}) ([]interface{}, error) {
	log.Printf("Got a request to getANumber: %v\n", args)
	return []interface{}{42}, nil
}

func (n *Example) subLoop(events chan *neovim.SubscriptionEvent) {
	for {
		select {
		case <-n.client.KillChannel:
			return
		case <-events:
			// Make an API request
			cb, _ := n.client.GetCurrentBuffer()
			bc, _ := cb.GetSlice(0, -1, true, true)

			// in practice we would use bc to do something useful
			// just log for now
			fmt.Fprintf(os.Stderr, "Buffer is: %v\n", bc)
		}
	}
}

func (n *Example) Shutdown() error {
	return nil
}

type Banana struct{}

func (b Banana) Init(c *neovim.Client) error {
	return nil
}
func (b Banana) Shutdown() error {
	return nil
}
