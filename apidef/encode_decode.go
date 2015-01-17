package apidef

import (
	"github.com/juju/errors"
	"github.com/tinylib/msgp/msgp"
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
	Async             bool
}

type APIFeature struct {
	Name    string
	Methods []string
}

// An APIFunctionParameter represents a function parameters as defined by an APIFunction
type APIFunctionParameter struct {
	Type, Name string
}

func GetAPI(ad *msgp.Reader) (*API, error) {

	ml, err := ad.ReadMapHeader()
	if err != nil {
		return nil, errors.Annotate(err, "Could not decode map length")

	}

	resp := &API{}

	for i := uint32(0); i < ml; i++ {
		k, err := ad.ReadBytes(nil)
		if err != nil {
			return nil, errors.Annotate(err, "Could not decode key of top level api map")
		}

		switch string(k) {
		case "types":
			classes, err := decodeAPIClassSlice(ad)
			if err != nil {
				return nil, errors.Annotate(err, "Could not decode class slice")
			}
			resp.Types = classes
		case "functions":
			functions, err := decodeAPIFunctionSlice(ad)
			if err != nil {
				return nil, errors.Annotate(err, "Could not decode function slice")
			}
			resp.Functions = functions
		case "features":
			ml, err := ad.ReadMapHeader()
			if err != nil {
				return nil, errors.Annotate(err, "Could not decode map length")
			}

			features := make([]APIFeature, ml)

			for i := range features {
				fn, err := ad.ReadBytes(nil)
				if err != nil {
					return nil, errors.Annotate(err, "Could not decode feature name")
				}

				features[i].Name = string(fn)

				meths, err := ad.ReadArrayHeader()
				if err != nil {
					return nil, errors.Annotate(err, "Could not decode length of features methods slice")
				}

				features[i].Methods = make([]string, meths)

				for j := range features[i].Methods {
					mn, err := ad.ReadBytes(nil)
					if err != nil {
						return nil, errors.Annotatef(err, "Could not decode feature method name at index %v", j)
					}
					features[i].Methods[j] = string(mn)
				}
			}
		case "error_types":
			ets, err := decodeAPIClassSlice(ad)
			if err != nil {
				return nil, errors.Annotate(err, "Could not decode error_type slice")
			}
			resp.ErrorTypes = ets
		}
	}

	return resp, nil
}

func decodeAPIClass(d *msgp.Reader) (APIClass, error) {
	resp := APIClass{}
	cn, err := d.ReadBytes(nil)
	if err != nil {
		return resp, errors.Annotate(err, "Could not decode class name")
	}

	ml, err := d.ReadMapHeader()
	if err != nil {
		return resp, errors.Annotate(err, "Could not decode map length")
	}

	if ml != 1 {
		return resp, errors.Errorf("Expected map length of 1; got %v", ml)
	}

	mk, err := d.ReadBytes(nil)
	if err != nil {
		return resp, errors.Annotate(err, "Could not decode ID key")
	}

	if string(mk) != "id" {
		return resp, errors.Errorf("Expected single key to be 'id'; got %v", mk)
	}

	id, err := d.ReadInt()
	if err != nil {
		return resp, errors.Annotate(err, "Could not decode ID value")
	}

	resp.Name = string(cn)
	resp.Id = id
	return resp, nil
}

func decodeAPIClassSlice(d *msgp.Reader) ([]APIClass, error) {
	sl, err := d.ReadMapHeader()
	if err != nil {
		return nil, errors.Annotate(err, "Could not decode slice length")
	}

	resp := make([]APIClass, sl)

	for i := uint32(0); i < sl; i++ {
		nvc, err := decodeAPIClass(d)
		if err != nil {
			return nil, errors.Annotatef(err, "Could not decode class at index %v", i)
		}
		resp[i] = nvc
	}
	return resp, nil
}

func decodeAPIFunction(d *msgp.Reader) (APIFunction, error) {
	resp := APIFunction{}
	ml, err := d.ReadMapHeader()
	if err != nil {
		return resp, errors.Annotate(err, "Could not decode map length")
	}

	for i := uint32(0); i < ml; i++ {
		k, err := d.ReadBytes(nil)
		if err != nil {
			return resp, errors.Annotate(err, "Could not decode function property key")
		}

		switch string(k) {
		case "name":
			s, err := d.ReadBytes(nil)
			if err != nil {
				return resp, errors.Annotate(err, "Could not decode function name")
			}
			resp.Name = string(s)
		case "receives_channel_id":
			b, err := d.ReadBool()
			if err != nil {
				return resp, errors.Annotate(err, "Could not decode function receives_channel_id")
			}
			resp.ReceivesChannelID = b
		case "can_fail":
			b, err := d.ReadBool()
			if err != nil {
				return resp, errors.Annotate(err, "Could not decode function can_fail")
			}
			resp.CanFail = b
		case "async":
			b, err := d.ReadBool()
			if err != nil {
				return resp, errors.Annotate(err, "Could not decode function can_fail")
			}
			resp.Async = b
		case "return_type":
			s, err := d.ReadBytes(nil)
			if err != nil {
				return resp, errors.Annotate(err, "Could not decode function return type")
			}
			resp.ReturnType = string(s)
		case "id":
			i, err := d.ReadUint32()
			if err != nil {
				return resp, errors.Annotate(err, "Could not decode function id")
			}
			resp.ID = i
		case "parameters":
			ps, err := decodeAPIFunctionParameterSlice(d)
			if err != nil {
				return resp, errors.Annotate(err, "Could not decode function parameters")
			}
			resp.Parameters = ps
		case "deferred":
			b, err := d.ReadBool()
			if err != nil {
				return resp, errors.Annotate(err, "Could not decode deferred property")
			}
			resp.Deferred = b
		default:
			return resp, errors.Errorf("Unknown function property %v", string(k))
		}
	}

	return resp, nil
}

func decodeAPIFunctionSlice(d *msgp.Reader) ([]APIFunction, error) {
	sl, err := d.ReadArrayHeader()
	if err != nil {
		return nil, errors.Annotate(err, "Could not decode slice length")
	}

	resp := make([]APIFunction, sl)

	for i := uint32(0); i < sl; i++ {
		nvc, err := decodeAPIFunction(d)
		if err != nil {
			return nil, errors.Annotatef(err, "Could not decode function at index %v", i)
		}
		resp[i] = nvc
	}
	return resp, nil
}

func decodeAPIFunctionParameter(d *msgp.Reader) (APIFunctionParameter, error) {
	resp := APIFunctionParameter{}

	// we should have a slice of length 2
	sl, err := d.ReadArrayHeader()
	if err != nil {
		return resp, errors.Annotate(err, "Could not decode slice length")
	}

	if sl != 2 {
		return resp, errors.Errorf("Expected lenght to be 2; got %v", sl)
	}

	pt, err := d.ReadBytes(nil)
	if err != nil {
		return resp, errors.Annotate(err, "Could not decode class name")
	}
	resp.Type = string(pt)
	pn, err := d.ReadBytes(nil)
	if err != nil {
		return resp, errors.Annotate(err, "Could not decode class name")
	}
	resp.Name = string(pn)
	return resp, nil
}

func decodeAPIFunctionParameterSlice(d *msgp.Reader) ([]APIFunctionParameter, error) {
	sl, err := d.ReadArrayHeader()
	if err != nil {
		return nil, errors.Annotate(err, "Could not decode slice length")
	}

	resp := make([]APIFunctionParameter, sl)

	for i := uint32(0); i < sl; i++ {
		nvc, err := decodeAPIFunctionParameter(d)
		if err != nil {
			return nil, errors.Annotatef(err, "Could not decode function parameter at index %v", i)
		}
		resp[i] = nvc
	}
	return resp, nil
}
