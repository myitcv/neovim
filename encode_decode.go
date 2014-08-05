package neovim

import (
	"github.com/juju/errgo"
	"github.com/vmihailenco/msgpack"
)

type Encoder func(*msgpack.Encoder, interface{}) error
type Decoder func(*msgpack.Decoder) (interface{}, error)

func encodeNoArgs(e *msgpack.Encoder, args interface{}) error {
	err := e.Encode([]string{})
	if err != nil {
		return errgo.NoteMask(err, "Could not encode nil")
	}
	return nil
}

func decodeBufferArray(d *msgpack.Decoder) (interface{}, error) {
	l, err := d.DecodeSliceLen()
	if err != nil {
		return nil, errgo.NoteMask(err, "Could not decode array length")
	}

	res := make([]Buffer, l)

	for i := 0; i < l; i++ {
		b, err := d.DecodeUint32()
		if err != nil {
			return nil, errgo.Notef(err, "Could not decode Buffer int at index %v\n", i)
		}
		res[i] = Buffer{id: b}
	}

	return res, nil
}
