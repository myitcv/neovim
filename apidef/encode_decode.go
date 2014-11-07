package apidef

import (
	"github.com/juju/errgo"
	"github.com/vmihailenco/msgpack"
)

// An API represents the API as advertised by Neovim
type API struct {
	Types      []APIClass
	Functions  []APIFunction
	ErrorTypes []APIClass
	Features   []APIFeature
}

// An APIClass represents a class as defined as part of the API
type APIClass struct {
	Name string
	Id   int
}

// An APIFunction represents a class as defined as part of the API
type APIFunction struct {
	Name              string
	ReturnType        string
	ID                uint32
	CanFail           bool
	Deferred          bool
	ReceivesChannelID bool
	Parameters        []APIFunctionParameter
}

type APIFeature struct {
	Name    string
	Methods []string
}

// An APIFunctionParameter represents a function parameters as defined by an APIFunction
type APIFunctionParameter struct {
	Type, Name string
}

func GetAPI(ad *msgpack.Decoder) (*API, error) {

	ml, err := ad.DecodeMapLen()
	if err != nil {
		return nil, errgo.NoteMask(err, "Could not decode map length")

	}

	resp := &API{}

	for i := 0; i < ml; i++ {
		k, err := ad.DecodeBytes()
		if err != nil {
			return nil, errgo.NoteMask(err, "Could not decode key of top level api map")
		}

		switch string(k) {
		case "types":
			classes, err := decodeAPIClassSlice(ad)
			if err != nil {
				return nil, errgo.NoteMask(err, "Could not decode class slice")
			}
			resp.Types = classes
		case "functions":
			functions, err := decodeAPIFunctionSlice(ad)
			if err != nil {
				return nil, errgo.NoteMask(err, "Could not decode function slice")
			}
			resp.Functions = functions
		case "features":
			ml, err := ad.DecodeMapLen()
			if err != nil {
				return nil, errgo.NoteMask(err, "Could not decode map length")
			}

			features := make([]APIFeature, ml)

			for i := range features {
				fn, err := ad.DecodeBytes()
				if err != nil {
					return nil, errgo.NoteMask(err, "Could not decode feature name")
				}

				features[i].Name = string(fn)

				meths, err := ad.DecodeSliceLen()
				if err != nil {
					return nil, errgo.NoteMask(err, "Could not decode length of features methods slice")
				}

				features[i].Methods = make([]string, meths)

				for j := range features[i].Methods {
					mn, err := ad.DecodeBytes()
					if err != nil {
						return nil, errgo.Notef(err, "Could not decode feature method name at index %v", j)
					}
					features[i].Methods[j] = string(mn)
				}
			}
		case "error_types":
			ets, err := decodeAPIClassSlice(ad)
			if err != nil {
				return nil, errgo.NoteMask(err, "Could not decode error_type slice")
			}
			resp.ErrorTypes = ets
		}
	}

	return resp, nil
}

func decodeAPIClass(d *msgpack.Decoder) (APIClass, error) {
	resp := APIClass{}
	cn, err := d.DecodeBytes()
	if err != nil {
		return resp, errgo.NoteMask(err, "Could not decode class name")
	}

	ml, err := d.DecodeMapLen()
	if err != nil {
		return resp, errgo.NoteMask(err, "Could not decode map length")
	}

	if ml != 1 {
		return resp, errgo.Newf("Expected map length of 1; got %v", ml)
	}

	mk, err := d.DecodeBytes()
	if err != nil {
		return resp, errgo.NoteMask(err, "Could not decode ID key")
	}

	if string(mk) != "id" {
		return resp, errgo.Newf("Expected single key to be 'id'; got %v", mk)
	}

	id, err := d.DecodeInt()
	if err != nil {
		return resp, errgo.NoteMask(err, "Could not decode ID value")
	}

	resp.Name = string(cn)
	resp.Id = id
	return resp, nil
}

func decodeAPIClassSlice(d *msgpack.Decoder) ([]APIClass, error) {
	sl, err := d.DecodeMapLen()
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
			resp.ReceivesChannelID = b
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
			resp.ID = i
		case "parameters":
			ps, err := decodeAPIFunctionParameterSlice(d)
			if err != nil {
				return resp, errgo.NoteMask(err, "Could not decode function parameters")
			}
			resp.Parameters = ps
		case "deferred":
			b, err := d.DecodeBool()
			if err != nil {
				return resp, errgo.NoteMask(err, "Could not decode deferred property")
			}
			resp.Deferred = b
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
