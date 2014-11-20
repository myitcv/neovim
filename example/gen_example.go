package example

import (
	"github.com/juju/errors"
	"github.com/myitcv/neovim"
	"github.com/vmihailenco/msgpack"
)

func (e *Example) newBufCreateSub() (chan *BufCreate, neovim.AsyncDecoder) {
	ch := make(chan *BufCreate)
	res := &bufCreateDecoder{ch: ch}
	return ch, res
}

func (e *Example) newGetANumberDecoder() neovim.SyncDecoder {
	res := &getANumberDecoder{Example: e}
	return res
}

type bufCreateDecoder struct {
	ch chan *BufCreate
}

type bufCreateRunner struct {
	ch chan *BufCreate
	v  *BufCreate
}

func (b *bufCreateDecoder) Decode(dec *msgpack.Decoder) (neovim.AsyncRunner, error) {
	val := &BufCreate{}

	l, err := dec.DecodeSliceLen()
	if err != nil {
		return nil, errors.Annotatef(err, "Could not decode slice len")
	}

	if l != 0 {
		return nil, errors.Errorf("Expected 0 arguments, not %v", l)
	}

	res := &bufCreateRunner{}
	res.v = val
	res.ch = b.ch

	return res, nil
}

func (b *bufCreateRunner) Run() error {
	b.ch <- b.v

	return nil
}

type getANumberDecoder struct {
	*Example
}

type getANumberRunner struct {
	*Example
}

type getANumberEncoder struct {
	i int
}

func (g *getANumberDecoder) Decode(dec *msgpack.Decoder) (neovim.SyncRunner, error) {
	l, err := dec.DecodeSliceLen()
	if err != nil {
		return nil, err
	}

	if l != 1 {
		return nil, errors.Errorf("Expected 1 argument, not %v", l)
	}

	l, err = dec.DecodeSliceLen()
	if err != nil {
		return nil, err
	}

	if l != 0 {
		return nil, errors.Errorf("Expected 0 argument, not %v", l)
	}

	return &getANumberRunner{Example: g.Example}, nil
}

func (g *getANumberRunner) Run() (neovim.SyncEncoder, error, error) {
	res := &getANumberEncoder{}

	i, mErr, err := g.Example.GetANumber()

	if err != nil || mErr != nil {
		return nil, mErr, err
	}

	res.i = i

	return res, nil, nil
}

func (g *getANumberEncoder) Encode(enc *msgpack.Encoder) error {
	err := enc.EncodeInt(g.i)
	if err != nil {
		return err
	}

	return nil
}
