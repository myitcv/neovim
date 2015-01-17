package example

import "github.com/myitcv/neovim"

type Example struct {
	client                *neovim.Client
	log                   neovim.Logger
	doSomethingAsyncChans []chan string
}

func (n *Example) Init(c *neovim.Client, l neovim.Logger) error {
	n.client = c
	n.log = l

	n.client.RegisterSyncRequestHandler("GetTwoNumbers", n.newGetTwoNumbersResponder)
	n.client.RegisterAsyncRequestHandler("DoSomethingAsync", n.newDoSomethingAsyncResponder)

	ch := make(chan string)
	n.AddDoSomethingAsyncChan(ch)
	go n.subLoop(ch)

	return nil
}

func (n *Example) Shutdown() error {
	return nil
}

type theThing struct {
	i int
}

// a synchronous method that returns two numbers
func (n *Example) GetTwoNumbers(i int) (int, string, error, error) {
	return i + 42, "42", nil, nil
}

func (n *Example) AddDoSomethingAsyncChan(c chan string) {
	// TODO clearly this is not thread safe
	n.doSomethingAsyncChans = append(n.doSomethingAsyncChans, c)
}

// an async method defines no return values
func (n *Example) DoSomethingAsync(s string) error {
	// TODO clearly this is not thread safe
	for _, c := range n.doSomethingAsyncChans {
		c <- s
	}
	return nil
}

func (n *Example) subLoop(ch chan string) {
	for {
		select {
		case <-n.client.KillChannel:
			return
		case v := <-ch:
			n.log.Printf("Got an event: %v\n", v)
		}
	}
}
