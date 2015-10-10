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

func (n *doSomethingAsyncWrapper) Eval() msgp.Decodable {
	return nil
}

func (n *doSomethingAsyncWrapper) Params() *neovim.MethodOptionParams {
	return nil
}

type doSomethingAsyncWrapper struct {
	*Example
	args *DoSomethingAsyncArgs
}

//msgp:tuple DoSomethingAsyncArgs
type DoSomethingAsyncArgs struct {
	Arg0 string
}

func (g *doSomethingAsyncWrapper) Run() error {
	err := g.Example.DoSomethingAsync(nil, string(g.args.Arg0))
	return err
}

// **************************
// GetTwoNumbers
func (n *Example) newGetTwoNumbersResponder() neovim.SyncDecoder {
	return &getTwoNumbersWrapper{
		Example: n,
		args:    &GetTwoNumbersArgs{},
		results: &GetTwoNumbersResults{},
		eval:    new(MyEvalResult),
		params:  new(neovim.MethodOptionParams),
	}
}

func (n *getTwoNumbersWrapper) Args() msgp.Decodable {
	return n.args
}

func (n *getTwoNumbersWrapper) Eval() msgp.Decodable {
	return n.eval
}

func (n *getTwoNumbersWrapper) Params() *neovim.MethodOptionParams {
	return n.params
}

func (n *getTwoNumbersWrapper) Results() msgp.Encodable {
	return n.results
}

type getTwoNumbersWrapper struct {
	*Example
	params  *neovim.MethodOptionParams
	args    *GetTwoNumbersArgs
	results *GetTwoNumbersResults
	eval    *MyEvalResult
}

//msgp:tuple GetTwoNumbersArgs
type GetTwoNumbersArgs struct {
	Arg0 int64
}

//msgp:tuple GetTwoNumbersResults
type GetTwoNumbersResults struct {
	Ret0 int64
	Ret1 string
}

func (g *getTwoNumbersWrapper) Run() (error, error) {
	res := &GetTwoNumbersResults{}

	// TODO method option params

	retVal0, retVal1, mErr, err := g.Example.GetTwoNumbers(g.Params(), int(g.args.Arg0), g.eval)

	if err != nil || mErr != nil {
		return mErr, err
	}

	res.Ret0 = int64(retVal0)
	res.Ret1 = string(retVal1)

	g.results = res

	return nil, nil
}
