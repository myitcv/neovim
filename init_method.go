package neovim

import (
	"github.com/juju/errors"
	"github.com/vmihailenco/msgpack"
)

type InitMethodWrapper struct {
	InitMethod
	*InitMethodArgs
	*InitMethodRetVals
}

type InitMethodArgs struct {
}

type InitMethodRetVals struct {
	InitMethod
}

func (i *InitMethodArgs) DecodeMsg(dec *msgpack.Decoder) error {
	l, err := dec.DecodeSliceLen()
	if err != nil {
		return err
	}

	if l != 0 {
		return errors.Errorf("Expected 0 arguments, not %v", l)
	}

	return nil
}

func (i *InitMethodRetVals) EncodeMsg(enc *msgpack.Encoder) error {
	err := enc.EncodeNil()
	if err != nil {
		return err
	}

	return nil
}

func (i *InitMethodWrapper) Run() (error, error) {
	return nil, i.InitMethod()
}
