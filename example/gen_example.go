//go:generate msgp
package example

import (
	"github.com/myitcv/neovim"
	"github.com/tinylib/msgp/msgp"
)

// **************************
// DoSomethingAsync
func (n *Example) newDoSomethingAsyncResponder() neovim.AsyncDecoder {
	return &doSomethingAsyncWrapper{
		Example: n,
		args:    &DoSomethingAsyncArgs{},
	}
}

func (n *doSomethingAsyncWrapper) Args() msgp.Decodable {
	return n.args
}

type doSomethingAsyncWrapper struct {
	*Example
	args *DoSomethingAsyncArgs
}

//msgp:tuple DoSomethingAsyncArgs
type DoSomethingAsyncArgs struct {
	FunctionArgs DoSomethingAsyncFunctionArgs
}

//msgp:tuple DoSomethingAsyncFunctionArgs
type DoSomethingAsyncFunctionArgs struct {
	Arg0 []byte
}

func (g *doSomethingAsyncWrapper) Run() error {
	err := g.Example.DoSomethingAsync(string(g.args.FunctionArgs.Arg0))
	return err
}

// **************************
// GetTwoNumbers
func (n *Example) newGetTwoNumbersResponder() neovim.SyncDecoder {
	return &getTwoNumbersWrapper{
		Example: n,
		args:    &GetTwoNumbersArgs{},
		results: &GetTwoNumbersResults{},
	}
}

func (n *getTwoNumbersWrapper) Args() msgp.Decodable {
	return n.args
}

func (n *getTwoNumbersWrapper) Results() msgp.Encodable {
	return n.results
}

type getTwoNumbersWrapper struct {
	*Example
	args    *GetTwoNumbersArgs
	results *GetTwoNumbersResults
}

//msgp:tuple GetTwoNumbersArgs
type GetTwoNumbersArgs struct {
	FunctionArgs GetTwoNumbersFunctionArgs
}

//msgp:tuple GetTwoNumbersFunctionArgs
type GetTwoNumbersFunctionArgs struct {
	Arg0 int64
}

//msgp:tuple GetTwoNumbersResults
type GetTwoNumbersResults struct {
	Ret0 int64
	Ret1 []byte
}

func (g *getTwoNumbersWrapper) Run() (error, error) {
	res := &GetTwoNumbersResults{}

	retVal0, retVal1, mErr, err := g.Example.GetTwoNumbers(int(g.args.FunctionArgs.Arg0))

	if err != nil || mErr != nil {
		return mErr, err
	}

	res.Ret0 = int64(retVal0)
	res.Ret1 = []byte(retVal1)

	g.results = res

	return nil, nil
}
