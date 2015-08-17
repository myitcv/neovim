package example

import (
	"github.com/juju/errors"
	"github.com/myitcv/neovim"
	"gopkg.in/vmihailenco/msgpack.v2"
)

func (e *Example) newBufCreateChanHandler() (chan *BufCreate, neovim.NewAsyncDecoder) {
	ch := make(chan *BufCreate)
	res := func() neovim.AsyncDecoder {
		return &bufCreateSubWrapper{ch: ch}
	}
	return ch, res
}

type bufCreateSubWrapper struct {
	ch chan *BufCreate
	*BufCreate
}

func (b *BufCreate) DecodeMsg(dec *msgpack.Decoder) error {
	l, err := dec.DecodeSliceLen()
	if err != nil {
		return errors.Annotatef(err, "Could not decode slice len")
	}

	if l != 0 {
		return errors.Errorf("Expected 0 arguments, not %v", l)
	}

	return nil
}

func (b *bufCreateSubWrapper) Run() error {
	b.ch <- b.BufCreate

	return nil
}

func (e *Example) newGetANumberResponder() neovim.SyncDecoder {
	res := &getANumberWrapper{Example: e}
	return res
}

type getANumberWrapper struct {
	*Example
	*getANumberArgs
	*getANumberRetVals
}

type getANumberArgs struct{}

type getANumberRetVals struct {
	i int
}

func (g *getANumberArgs) DecodeMsg(dec *msgpack.Decoder) error {
	l, err := dec.DecodeSliceLen()
	if err != nil {
		return err
	}

	if l != 1 {
		return errors.Errorf("Expected 1 argument, not %v", l)
	}

	l, err = dec.DecodeSliceLen()
	if err != nil {
		return err
	}

	if l != 0 {
		return errors.Errorf("Expected 0 argument, not %v", l)
	}

	return nil
}

func (g *getANumberWrapper) Run() (error, error) {
	res := &getANumberRetVals{}

	i, mErr, err := g.Example.GetANumber()

	if err != nil || mErr != nil {
		return mErr, err
	}

	res.i = i

	g.getANumberRetVals = res

	return nil, nil
}

func (g *getANumberRetVals) EncodeMsg(enc *msgpack.Encoder) error {
	err := enc.EncodeInt(g.i)
	if err != nil {
		return err
	}

	return nil
}
