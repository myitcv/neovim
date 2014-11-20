package neovim

import (
	"github.com/juju/errors"
	"github.com/vmihailenco/msgpack"
)

type initMethodDecoder struct {
	InitMethod
}

type initMethodRunner struct {
	InitMethod
}

type initMethodEncoder struct{}

func (i *initMethodDecoder) Decode(dec *msgpack.Decoder) (SyncRunner, error) {
	l, err := dec.DecodeSliceLen()
	if err != nil {
		return nil, err
	}

	if l != 0 {
		return nil, errors.Errorf("Expected 0 arguments, not %v", l)
	}

	res := &initMethodRunner{InitMethod: i.InitMethod}

	return res, nil
}

func (i *initMethodEncoder) Encode(enc *msgpack.Encoder) error {
	err := enc.EncodeNil()
	if err != nil {
		return err
	}

	return nil
}

func (i *initMethodRunner) Run() (SyncEncoder, error, error) {
	return &initMethodEncoder{}, nil, i.InitMethod()
}
