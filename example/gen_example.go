package example

import (
	"github.com/juju/errgo"
	"github.com/myitcv/neovim"
	"github.com/vmihailenco/msgpack"
)

func (e *Example) NewBufCreateSub() (chan BufCreate, neovim.AsyncDecoder) {
	ch := make(chan BufCreate)
	dec := func(dec *msgpack.Decoder) (_err error) {
		val := BufCreate{}

		l, _err := dec.DecodeSliceLen()
		if _err != nil {
			return
		}

		if l != 0 {
			return errgo.Newf("Expected 0 arguments, not %v", l)
		}

		ch <- val

		return nil
	}
	return ch, dec
}

func (e *Example) NewGetANumberDecoder() neovim.SyncDecoder {
	res := func(dec *msgpack.Decoder) (neovim.Runner, error) {

		l, err := dec.DecodeSliceLen()
		if err != nil {
			return nil, err
		}

		if l != 1 {
			return nil, errgo.Newf("Expected 1 argument, not %v", l)
		}

		l, err = dec.DecodeSliceLen()
		if err != nil {
			return nil, err
		}

		if l != 0 {
			return nil, errgo.Newf("Expected 0 argument, not %v", l)
		}

		runner := func() (neovim.Encoder, error) {
			i, err := e.GetANumber()
			if err != nil {
				return nil, err
			}

			encoder := func(enc *msgpack.Encoder) error {
				err := enc.EncodeInt(i)
				if err != nil {
					return err
				}

				return nil
			}

			return encoder, nil
		}

		return runner, nil
	}
	return res
}
