package neovim

import (
	"bytes"

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

func decodeBufferSlice(d *msgpack.Decoder) (interface{}, error) {
	l, err := d.DecodeSliceLen()
	if err != nil {
		return nil, errgo.NoteMask(err, "Could not decode slice length")
	}

	res := make([]Buffer, l)

	for i := 0; i < l; i++ {
		b, err := d.DecodeUint32()
		if err != nil {
			return nil, errgo.Notef(err, "Could not decode Buffer int at index %v\n", i)
		}
		res[i] = Buffer{Id: b}
	}

	return res, nil
}

func decodeAPI(d *msgpack.Decoder) (interface{}, error) {
	sl, err := d.DecodeSliceLen()
	if err != nil {
		return nil, errgo.NoteMask(err, "Could not decode slice len")
	}

	_, err = d.DecodeInt()
	if err != nil {
		return nil, errgo.NoteMask(err, "Could not decode channel id")
	}

	if sl != 2 {
		return nil, errgo.New("Expected slice to be lenght 2")
	}

	mb, err := d.DecodeBytes()
	if err != nil {
		return nil, errgo.NoteMask(err, "Could not decode bytes for api map")
	}

	br := bytes.NewReader(mb)

	ad := msgpack.NewDecoder(br)

	ml, err := ad.DecodeMapLen()
	if err != nil {
		return nil, errgo.NoteMask(err, "Could not decode map length")
	}

	resp := &API{}

	for i := 0; i < ml; i++ {
		k, err := ad.DecodeString()
		if err != nil {
			return nil, errgo.NoteMask(err, "Could not decode key of top level api map")
		}

		switch k {
		case "classes":
			classes, err := decodeAPIClassSlice(ad)
			if err != nil {
				return nil, errgo.NoteMask(err, "Could not decode class slice")
			}
			resp.Classes = classes
		case "functions":
			functions, err := decodeAPIFunctionSlice(ad)
			if err != nil {
				return nil, errgo.NoteMask(err, "Could not decode function slice")
			}
			resp.Functions = functions
		}
	}

	return resp, nil
}

func decodeAPIClass(d *msgpack.Decoder) (APIClass, error) {
	resp := APIClass{}
	cn, err := d.DecodeString()
	if err != nil {
		return resp, errgo.NoteMask(err, "Could not decode class name")
	}

	resp.Name = cn
	return resp, nil
}

func decodeAPIClassSlice(d *msgpack.Decoder) ([]APIClass, error) {
	sl, err := d.DecodeSliceLen()
	if err != nil {
		return nil, errgo.NoteMask(err, "Could not decode slice length")
	}

	resp := make([]APIClass, sl)

	for i := 0; i < sl; i++ {
		nvc, err := decodeAPIClass(d)
		if err != nil {
			return nil, errgo.Notef(err, "Could not decode class at index %v", i)
		}
		resp[i] = nvc
	}
	return resp, nil
}

func decodeAPIFunction(d *msgpack.Decoder) (APIFunction, error) {
	resp := APIFunction{}
	ml, err := d.DecodeMapLen()
	if err != nil {
		return resp, errgo.NoteMask(err, "Could not decode map length")
	}

	for i := 0; i < ml; i++ {
		k, err := d.DecodeString()
		if err != nil {
			return resp, errgo.NoteMask(err, "Could not decode function property key")
		}

		switch k {
		case "name":
			s, err := d.DecodeString()
			if err != nil {
				return resp, errgo.NoteMask(err, "Could not decode function name")
			}
			resp.Name = s
		case "receives_channel_id":
			b, err := d.DecodeBool()
			if err != nil {
				return resp, errgo.NoteMask(err, "Could not decode function receives_channel_id")
			}
			resp.ReceivesChannelId = b
		case "can_fail":
			b, err := d.DecodeBool()
			if err != nil {
				return resp, errgo.NoteMask(err, "Could not decode function can_fail")
			}
			resp.CanFail = b
		case "return_type":
			s, err := d.DecodeString()
			if err != nil {
				return resp, errgo.NoteMask(err, "Could not decode function return type")
			}
			resp.ReturnType = s
		case "id":
			i, err := d.DecodeUint32()
			if err != nil {
				return resp, errgo.NoteMask(err, "Could not decode function id")
			}
			resp.Id = i
		case "parameters":
			ps, err := decodeAPIFunctionParameterSlice(d)
			if err != nil {
				return resp, errgo.NoteMask(err, "Could not decode function parameters")
			}
			resp.Parameters = ps
		default:
			return resp, errgo.Newf("Unknown function property %v", k)
		}
	}

	return resp, nil
}

func decodeAPIFunctionSlice(d *msgpack.Decoder) ([]APIFunction, error) {
	sl, err := d.DecodeSliceLen()
	if err != nil {
		return nil, errgo.NoteMask(err, "Could not decode slice length")
	}

	resp := make([]APIFunction, sl)

	for i := 0; i < sl; i++ {
		nvc, err := decodeAPIFunction(d)
		if err != nil {
			return nil, errgo.Notef(err, "Could not decode function at index %v", i)
		}
		resp[i] = nvc
	}
	return resp, nil
}

func decodeAPIFunctionParameter(d *msgpack.Decoder) (APIFunctionParameter, error) {
	resp := APIFunctionParameter{}

	// we should have a slice of length 2
	sl, err := d.DecodeSliceLen()
	if err != nil {
		return resp, errgo.NoteMask(err, "Could not decode slice length")
	}

	if sl != 2 {
		return resp, errgo.Newf("Expected lenght to be 2; got %v", sl)
	}

	pt, err := d.DecodeString()
	if err != nil {
		return resp, errgo.NoteMask(err, "Could not decode class name")
	}
	resp.Type = pt
	pn, err := d.DecodeString()
	if err != nil {
		return resp, errgo.NoteMask(err, "Could not decode class name")
	}
	resp.Name = pn
	return resp, nil
}

func decodeAPIFunctionParameterSlice(d *msgpack.Decoder) ([]APIFunctionParameter, error) {
	sl, err := d.DecodeSliceLen()
	if err != nil {
		return nil, errgo.NoteMask(err, "Could not decode slice length")
	}

	resp := make([]APIFunctionParameter, sl)

	for i := 0; i < sl; i++ {
		nvc, err := decodeAPIFunctionParameter(d)
		if err != nil {
			return nil, errgo.Notef(err, "Could not decode function parameter at index %v", i)
		}
		resp[i] = nvc
	}
	return resp, nil
}

// here are some neovim.Decoder wrappers around msgpack functions
func decodeInterface(d *msgpack.Decoder) (interface{}, error) {
	return d.DecodeInterface()
}

func decodeSlice(d *msgpack.Decoder) (interface{}, error) {
	return d.DecodeSlice()
}
