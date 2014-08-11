// **** THIS FILE IS GENERATED - DO NOT EDIT BY HAND
package neovim

import "github.com/juju/errgo"

// methods on the API

func (recv *Tabpage) GetWindows() (ret_val []Window, ret_err error) {
	enc := func() (_err error) {
		_err = recv.client.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = recv.client.encodeTabpage(*recv)
		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = recv.client.decodeWindowSlice()

		return
	}
	resp_chan, err := recv.client.makeCall(1, enc, dec)
	if err != nil {
		return ret_val, errgo.NoteMask(err, "Could not make call to Tabpage.GetWindows")
	}
	resp := <-resp_chan
	if resp == nil {
		return ret_val, errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return ret_val, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	ret_val = resp.obj.([]Window)
	return ret_val, ret_err

}

func (recv *Tabpage) GetVar(i_name string) (ret_val interface{}, ret_err error) {
	enc := func() (_err error) {
		_err = recv.client.enc.EncodeSliceLen(2)
		if _err != nil {
			return
		}

		_err = recv.client.encodeTabpage(*recv)
		if _err != nil {
			return
		}

		_err = recv.client.enc.EncodeString(i_name)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = recv.client.dec.DecodeInterface()

		return
	}
	resp_chan, err := recv.client.makeCall(2, enc, dec)
	if err != nil {
		return ret_val, errgo.NoteMask(err, "Could not make call to Tabpage.GetVar")
	}
	resp := <-resp_chan
	if resp == nil {
		return ret_val, errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return ret_val, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	ret_val = resp.obj.(interface{})
	return ret_val, ret_err

}

func (recv *Tabpage) SetVar(i_name string, i_value interface{}) (ret_val interface{}, ret_err error) {
	enc := func() (_err error) {
		_err = recv.client.enc.EncodeSliceLen(3)
		if _err != nil {
			return
		}

		_err = recv.client.encodeTabpage(*recv)
		if _err != nil {
			return
		}

		_err = recv.client.enc.EncodeString(i_name)

		if _err != nil {
			return
		}

		_err = recv.client.enc.Encode(i_value)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = recv.client.dec.DecodeInterface()

		return
	}
	resp_chan, err := recv.client.makeCall(3, enc, dec)
	if err != nil {
		return ret_val, errgo.NoteMask(err, "Could not make call to Tabpage.SetVar")
	}
	resp := <-resp_chan
	if resp == nil {
		return ret_val, errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return ret_val, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	ret_val = resp.obj.(interface{})
	return ret_val, ret_err

}

func (recv *Tabpage) GetWindow() (ret_val Window, ret_err error) {
	enc := func() (_err error) {
		_err = recv.client.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = recv.client.encodeTabpage(*recv)
		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = recv.client.decodeWindow()

		return
	}
	resp_chan, err := recv.client.makeCall(4, enc, dec)
	if err != nil {
		return ret_val, errgo.NoteMask(err, "Could not make call to Tabpage.GetWindow")
	}
	resp := <-resp_chan
	if resp == nil {
		return ret_val, errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return ret_val, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	ret_val = resp.obj.(Window)
	return ret_val, ret_err

}

func (recv *Tabpage) IsValid() (ret_val bool, ret_err error) {
	enc := func() (_err error) {
		_err = recv.client.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = recv.client.encodeTabpage(*recv)
		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = recv.client.dec.DecodeBool()

		return
	}
	resp_chan, err := recv.client.makeCall(5, enc, dec)
	if err != nil {
		return ret_val, errgo.NoteMask(err, "Could not make call to Tabpage.IsValid")
	}
	resp := <-resp_chan
	if resp == nil {
		return ret_val, errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return ret_val, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	ret_val = resp.obj.(bool)
	return ret_val, ret_err

}

func (recv *Buffer) GetLength() (ret_val int, ret_err error) {
	enc := func() (_err error) {
		_err = recv.client.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = recv.client.encodeBuffer(*recv)
		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = recv.client.dec.DecodeInt()

		return
	}
	resp_chan, err := recv.client.makeCall(6, enc, dec)
	if err != nil {
		return ret_val, errgo.NoteMask(err, "Could not make call to Buffer.GetLength")
	}
	resp := <-resp_chan
	if resp == nil {
		return ret_val, errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return ret_val, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	ret_val = resp.obj.(int)
	return ret_val, ret_err

}

func (recv *Buffer) GetLine(i_index int) (ret_val string, ret_err error) {
	enc := func() (_err error) {
		_err = recv.client.enc.EncodeSliceLen(2)
		if _err != nil {
			return
		}

		_err = recv.client.encodeBuffer(*recv)
		if _err != nil {
			return
		}

		_err = recv.client.enc.EncodeInt(i_index)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = recv.client.dec.DecodeString()

		return
	}
	resp_chan, err := recv.client.makeCall(7, enc, dec)
	if err != nil {
		return ret_val, errgo.NoteMask(err, "Could not make call to Buffer.GetLine")
	}
	resp := <-resp_chan
	if resp == nil {
		return ret_val, errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return ret_val, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	ret_val = resp.obj.(string)
	return ret_val, ret_err

}

func (recv *Buffer) SetLine(i_index int, i_line string) (ret_err error) {
	enc := func() (_err error) {
		_err = recv.client.enc.EncodeSliceLen(3)
		if _err != nil {
			return
		}

		_err = recv.client.encodeBuffer(*recv)
		if _err != nil {
			return
		}

		_err = recv.client.enc.EncodeInt(i_index)

		if _err != nil {
			return
		}

		_err = recv.client.enc.EncodeString(i_line)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_, _err = recv.client.dec.DecodeBytes()

		return
	}
	resp_chan, err := recv.client.makeCall(8, enc, dec)
	if err != nil {
		return errgo.NoteMask(err, "Could not make call to Buffer.SetLine")
	}
	resp := <-resp_chan
	if resp == nil {
		return errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	return ret_err

}

func (recv *Buffer) DelLine(i_index int) (ret_err error) {
	enc := func() (_err error) {
		_err = recv.client.enc.EncodeSliceLen(2)
		if _err != nil {
			return
		}

		_err = recv.client.encodeBuffer(*recv)
		if _err != nil {
			return
		}

		_err = recv.client.enc.EncodeInt(i_index)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_, _err = recv.client.dec.DecodeBytes()

		return
	}
	resp_chan, err := recv.client.makeCall(9, enc, dec)
	if err != nil {
		return errgo.NoteMask(err, "Could not make call to Buffer.DelLine")
	}
	resp := <-resp_chan
	if resp == nil {
		return errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	return ret_err

}

func (recv *Buffer) GetSlice(i_start int, i_end int, i_include_start bool, i_include_end bool) (ret_val []string, ret_err error) {
	enc := func() (_err error) {
		_err = recv.client.enc.EncodeSliceLen(5)
		if _err != nil {
			return
		}

		_err = recv.client.encodeBuffer(*recv)
		if _err != nil {
			return
		}

		_err = recv.client.enc.EncodeInt(i_start)

		if _err != nil {
			return
		}

		_err = recv.client.enc.EncodeInt(i_end)

		if _err != nil {
			return
		}

		_err = recv.client.enc.EncodeBool(i_include_start)

		if _err != nil {
			return
		}

		_err = recv.client.enc.EncodeBool(i_include_end)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = recv.client.decodeStringSlice()

		return
	}
	resp_chan, err := recv.client.makeCall(10, enc, dec)
	if err != nil {
		return ret_val, errgo.NoteMask(err, "Could not make call to Buffer.GetSlice")
	}
	resp := <-resp_chan
	if resp == nil {
		return ret_val, errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return ret_val, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	ret_val = resp.obj.([]string)
	return ret_val, ret_err

}

func (recv *Buffer) SetSlice(i_start int, i_end int, i_include_start bool, i_include_end bool, i_replacement []string) (ret_err error) {
	enc := func() (_err error) {
		_err = recv.client.enc.EncodeSliceLen(6)
		if _err != nil {
			return
		}

		_err = recv.client.encodeBuffer(*recv)
		if _err != nil {
			return
		}

		_err = recv.client.enc.EncodeInt(i_start)

		if _err != nil {
			return
		}

		_err = recv.client.enc.EncodeInt(i_end)

		if _err != nil {
			return
		}

		_err = recv.client.enc.EncodeBool(i_include_start)

		if _err != nil {
			return
		}

		_err = recv.client.enc.EncodeBool(i_include_end)

		if _err != nil {
			return
		}

		_err = recv.client.encodeStringSlice(i_replacement)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_, _err = recv.client.dec.DecodeBytes()

		return
	}
	resp_chan, err := recv.client.makeCall(11, enc, dec)
	if err != nil {
		return errgo.NoteMask(err, "Could not make call to Buffer.SetSlice")
	}
	resp := <-resp_chan
	if resp == nil {
		return errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	return ret_err

}

func (recv *Buffer) GetVar(i_name string) (ret_val interface{}, ret_err error) {
	enc := func() (_err error) {
		_err = recv.client.enc.EncodeSliceLen(2)
		if _err != nil {
			return
		}

		_err = recv.client.encodeBuffer(*recv)
		if _err != nil {
			return
		}

		_err = recv.client.enc.EncodeString(i_name)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = recv.client.dec.DecodeInterface()

		return
	}
	resp_chan, err := recv.client.makeCall(12, enc, dec)
	if err != nil {
		return ret_val, errgo.NoteMask(err, "Could not make call to Buffer.GetVar")
	}
	resp := <-resp_chan
	if resp == nil {
		return ret_val, errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return ret_val, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	ret_val = resp.obj.(interface{})
	return ret_val, ret_err

}

func (recv *Buffer) SetVar(i_name string, i_value interface{}) (ret_val interface{}, ret_err error) {
	enc := func() (_err error) {
		_err = recv.client.enc.EncodeSliceLen(3)
		if _err != nil {
			return
		}

		_err = recv.client.encodeBuffer(*recv)
		if _err != nil {
			return
		}

		_err = recv.client.enc.EncodeString(i_name)

		if _err != nil {
			return
		}

		_err = recv.client.enc.Encode(i_value)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = recv.client.dec.DecodeInterface()

		return
	}
	resp_chan, err := recv.client.makeCall(13, enc, dec)
	if err != nil {
		return ret_val, errgo.NoteMask(err, "Could not make call to Buffer.SetVar")
	}
	resp := <-resp_chan
	if resp == nil {
		return ret_val, errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return ret_val, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	ret_val = resp.obj.(interface{})
	return ret_val, ret_err

}

func (recv *Buffer) GetOption(i_name string) (ret_val interface{}, ret_err error) {
	enc := func() (_err error) {
		_err = recv.client.enc.EncodeSliceLen(2)
		if _err != nil {
			return
		}

		_err = recv.client.encodeBuffer(*recv)
		if _err != nil {
			return
		}

		_err = recv.client.enc.EncodeString(i_name)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = recv.client.dec.DecodeInterface()

		return
	}
	resp_chan, err := recv.client.makeCall(14, enc, dec)
	if err != nil {
		return ret_val, errgo.NoteMask(err, "Could not make call to Buffer.GetOption")
	}
	resp := <-resp_chan
	if resp == nil {
		return ret_val, errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return ret_val, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	ret_val = resp.obj.(interface{})
	return ret_val, ret_err

}

func (recv *Buffer) SetOption(i_name string, i_value interface{}) (ret_err error) {
	enc := func() (_err error) {
		_err = recv.client.enc.EncodeSliceLen(3)
		if _err != nil {
			return
		}

		_err = recv.client.encodeBuffer(*recv)
		if _err != nil {
			return
		}

		_err = recv.client.enc.EncodeString(i_name)

		if _err != nil {
			return
		}

		_err = recv.client.enc.Encode(i_value)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_, _err = recv.client.dec.DecodeBytes()

		return
	}
	resp_chan, err := recv.client.makeCall(15, enc, dec)
	if err != nil {
		return errgo.NoteMask(err, "Could not make call to Buffer.SetOption")
	}
	resp := <-resp_chan
	if resp == nil {
		return errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	return ret_err

}

func (recv *Buffer) GetNumber() (ret_val int, ret_err error) {
	enc := func() (_err error) {
		_err = recv.client.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = recv.client.encodeBuffer(*recv)
		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = recv.client.dec.DecodeInt()

		return
	}
	resp_chan, err := recv.client.makeCall(16, enc, dec)
	if err != nil {
		return ret_val, errgo.NoteMask(err, "Could not make call to Buffer.GetNumber")
	}
	resp := <-resp_chan
	if resp == nil {
		return ret_val, errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return ret_val, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	ret_val = resp.obj.(int)
	return ret_val, ret_err

}

func (recv *Buffer) GetName() (ret_val string, ret_err error) {
	enc := func() (_err error) {
		_err = recv.client.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = recv.client.encodeBuffer(*recv)
		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = recv.client.dec.DecodeString()

		return
	}
	resp_chan, err := recv.client.makeCall(17, enc, dec)
	if err != nil {
		return ret_val, errgo.NoteMask(err, "Could not make call to Buffer.GetName")
	}
	resp := <-resp_chan
	if resp == nil {
		return ret_val, errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return ret_val, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	ret_val = resp.obj.(string)
	return ret_val, ret_err

}

func (recv *Buffer) SetName(i_name string) (ret_err error) {
	enc := func() (_err error) {
		_err = recv.client.enc.EncodeSliceLen(2)
		if _err != nil {
			return
		}

		_err = recv.client.encodeBuffer(*recv)
		if _err != nil {
			return
		}

		_err = recv.client.enc.EncodeString(i_name)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_, _err = recv.client.dec.DecodeBytes()

		return
	}
	resp_chan, err := recv.client.makeCall(18, enc, dec)
	if err != nil {
		return errgo.NoteMask(err, "Could not make call to Buffer.SetName")
	}
	resp := <-resp_chan
	if resp == nil {
		return errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	return ret_err

}

func (recv *Buffer) IsValid() (ret_val bool, ret_err error) {
	enc := func() (_err error) {
		_err = recv.client.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = recv.client.encodeBuffer(*recv)
		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = recv.client.dec.DecodeBool()

		return
	}
	resp_chan, err := recv.client.makeCall(19, enc, dec)
	if err != nil {
		return ret_val, errgo.NoteMask(err, "Could not make call to Buffer.IsValid")
	}
	resp := <-resp_chan
	if resp == nil {
		return ret_val, errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return ret_val, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	ret_val = resp.obj.(bool)
	return ret_val, ret_err

}

func (recv *Buffer) Insert(i_lnum int, i_lines []string) (ret_err error) {
	enc := func() (_err error) {
		_err = recv.client.enc.EncodeSliceLen(3)
		if _err != nil {
			return
		}

		_err = recv.client.encodeBuffer(*recv)
		if _err != nil {
			return
		}

		_err = recv.client.enc.EncodeInt(i_lnum)

		if _err != nil {
			return
		}

		_err = recv.client.encodeStringSlice(i_lines)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_, _err = recv.client.dec.DecodeBytes()

		return
	}
	resp_chan, err := recv.client.makeCall(20, enc, dec)
	if err != nil {
		return errgo.NoteMask(err, "Could not make call to Buffer.Insert")
	}
	resp := <-resp_chan
	if resp == nil {
		return errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	return ret_err

}

func (recv *Buffer) GetMark(i_name string) (ret_val uint32, ret_err error) {
	enc := func() (_err error) {
		_err = recv.client.enc.EncodeSliceLen(2)
		if _err != nil {
			return
		}

		_err = recv.client.encodeBuffer(*recv)
		if _err != nil {
			return
		}

		_err = recv.client.enc.EncodeString(i_name)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = recv.client.dec.DecodeUint32()

		return
	}
	resp_chan, err := recv.client.makeCall(21, enc, dec)
	if err != nil {
		return ret_val, errgo.NoteMask(err, "Could not make call to Buffer.GetMark")
	}
	resp := <-resp_chan
	if resp == nil {
		return ret_val, errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return ret_val, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	ret_val = resp.obj.(uint32)
	return ret_val, ret_err

}

func (recv *Client) PushKeys(i_str string) (ret_err error) {
	enc := func() (_err error) {
		_err = recv.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = recv.enc.EncodeString(i_str)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_, _err = recv.dec.DecodeBytes()

		return
	}
	resp_chan, err := recv.makeCall(22, enc, dec)
	if err != nil {
		return errgo.NoteMask(err, "Could not make call to Client.PushKeys")
	}
	resp := <-resp_chan
	if resp == nil {
		return errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	return ret_err

}

func (recv *Client) Command(i_str string) (ret_err error) {
	enc := func() (_err error) {
		_err = recv.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = recv.enc.EncodeString(i_str)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_, _err = recv.dec.DecodeBytes()

		return
	}
	resp_chan, err := recv.makeCall(23, enc, dec)
	if err != nil {
		return errgo.NoteMask(err, "Could not make call to Client.Command")
	}
	resp := <-resp_chan
	if resp == nil {
		return errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	return ret_err

}

func (recv *Client) Feedkeys(i_keys string, i_mode string) (ret_err error) {
	enc := func() (_err error) {
		_err = recv.enc.EncodeSliceLen(2)
		if _err != nil {
			return
		}

		_err = recv.enc.EncodeString(i_keys)

		if _err != nil {
			return
		}

		_err = recv.enc.EncodeString(i_mode)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_, _err = recv.dec.DecodeBytes()

		return
	}
	resp_chan, err := recv.makeCall(24, enc, dec)
	if err != nil {
		return errgo.NoteMask(err, "Could not make call to Client.Feedkeys")
	}
	resp := <-resp_chan
	if resp == nil {
		return errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	return ret_err

}

func (recv *Client) ReplaceTermcodes(i_str string, i_from_part bool, i_do_lt bool, i_special bool) (ret_val string, ret_err error) {
	enc := func() (_err error) {
		_err = recv.enc.EncodeSliceLen(4)
		if _err != nil {
			return
		}

		_err = recv.enc.EncodeString(i_str)

		if _err != nil {
			return
		}

		_err = recv.enc.EncodeBool(i_from_part)

		if _err != nil {
			return
		}

		_err = recv.enc.EncodeBool(i_do_lt)

		if _err != nil {
			return
		}

		_err = recv.enc.EncodeBool(i_special)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = recv.dec.DecodeString()

		return
	}
	resp_chan, err := recv.makeCall(25, enc, dec)
	if err != nil {
		return ret_val, errgo.NoteMask(err, "Could not make call to Client.ReplaceTermcodes")
	}
	resp := <-resp_chan
	if resp == nil {
		return ret_val, errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return ret_val, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	ret_val = resp.obj.(string)
	return ret_val, ret_err

}

func (recv *Client) Eval(i_str string) (ret_val interface{}, ret_err error) {
	enc := func() (_err error) {
		_err = recv.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = recv.enc.EncodeString(i_str)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = recv.dec.DecodeInterface()

		return
	}
	resp_chan, err := recv.makeCall(26, enc, dec)
	if err != nil {
		return ret_val, errgo.NoteMask(err, "Could not make call to Client.Eval")
	}
	resp := <-resp_chan
	if resp == nil {
		return ret_val, errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return ret_val, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	ret_val = resp.obj.(interface{})
	return ret_val, ret_err

}

func (recv *Client) Strwidth(i_str string) (ret_val int, ret_err error) {
	enc := func() (_err error) {
		_err = recv.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = recv.enc.EncodeString(i_str)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = recv.dec.DecodeInt()

		return
	}
	resp_chan, err := recv.makeCall(27, enc, dec)
	if err != nil {
		return ret_val, errgo.NoteMask(err, "Could not make call to Client.Strwidth")
	}
	resp := <-resp_chan
	if resp == nil {
		return ret_val, errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return ret_val, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	ret_val = resp.obj.(int)
	return ret_val, ret_err

}

func (recv *Client) ListRuntimePaths() (ret_val []string, ret_err error) {
	enc := func() (_err error) {
		_err = recv.enc.EncodeSliceLen(0)
		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = recv.decodeStringSlice()

		return
	}
	resp_chan, err := recv.makeCall(28, enc, dec)
	if err != nil {
		return ret_val, errgo.NoteMask(err, "Could not make call to Client.ListRuntimePaths")
	}
	resp := <-resp_chan
	if resp == nil {
		return ret_val, errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return ret_val, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	ret_val = resp.obj.([]string)
	return ret_val, ret_err

}

func (recv *Client) ChangeDirectory(i_dir string) (ret_err error) {
	enc := func() (_err error) {
		_err = recv.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = recv.enc.EncodeString(i_dir)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_, _err = recv.dec.DecodeBytes()

		return
	}
	resp_chan, err := recv.makeCall(29, enc, dec)
	if err != nil {
		return errgo.NoteMask(err, "Could not make call to Client.ChangeDirectory")
	}
	resp := <-resp_chan
	if resp == nil {
		return errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	return ret_err

}

func (recv *Client) GetCurrentLine() (ret_val string, ret_err error) {
	enc := func() (_err error) {
		_err = recv.enc.EncodeSliceLen(0)
		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = recv.dec.DecodeString()

		return
	}
	resp_chan, err := recv.makeCall(30, enc, dec)
	if err != nil {
		return ret_val, errgo.NoteMask(err, "Could not make call to Client.GetCurrentLine")
	}
	resp := <-resp_chan
	if resp == nil {
		return ret_val, errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return ret_val, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	ret_val = resp.obj.(string)
	return ret_val, ret_err

}

func (recv *Client) SetCurrentLine(i_line string) (ret_err error) {
	enc := func() (_err error) {
		_err = recv.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = recv.enc.EncodeString(i_line)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_, _err = recv.dec.DecodeBytes()

		return
	}
	resp_chan, err := recv.makeCall(31, enc, dec)
	if err != nil {
		return errgo.NoteMask(err, "Could not make call to Client.SetCurrentLine")
	}
	resp := <-resp_chan
	if resp == nil {
		return errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	return ret_err

}

func (recv *Client) DelCurrentLine() (ret_err error) {
	enc := func() (_err error) {
		_err = recv.enc.EncodeSliceLen(0)
		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_, _err = recv.dec.DecodeBytes()

		return
	}
	resp_chan, err := recv.makeCall(32, enc, dec)
	if err != nil {
		return errgo.NoteMask(err, "Could not make call to Client.DelCurrentLine")
	}
	resp := <-resp_chan
	if resp == nil {
		return errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	return ret_err

}

func (recv *Client) GetVar(i_name string) (ret_val interface{}, ret_err error) {
	enc := func() (_err error) {
		_err = recv.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = recv.enc.EncodeString(i_name)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = recv.dec.DecodeInterface()

		return
	}
	resp_chan, err := recv.makeCall(33, enc, dec)
	if err != nil {
		return ret_val, errgo.NoteMask(err, "Could not make call to Client.GetVar")
	}
	resp := <-resp_chan
	if resp == nil {
		return ret_val, errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return ret_val, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	ret_val = resp.obj.(interface{})
	return ret_val, ret_err

}

func (recv *Client) SetVar(i_name string, i_value interface{}) (ret_val interface{}, ret_err error) {
	enc := func() (_err error) {
		_err = recv.enc.EncodeSliceLen(2)
		if _err != nil {
			return
		}

		_err = recv.enc.EncodeString(i_name)

		if _err != nil {
			return
		}

		_err = recv.enc.Encode(i_value)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = recv.dec.DecodeInterface()

		return
	}
	resp_chan, err := recv.makeCall(34, enc, dec)
	if err != nil {
		return ret_val, errgo.NoteMask(err, "Could not make call to Client.SetVar")
	}
	resp := <-resp_chan
	if resp == nil {
		return ret_val, errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return ret_val, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	ret_val = resp.obj.(interface{})
	return ret_val, ret_err

}

func (recv *Client) GetVvar(i_name string) (ret_val interface{}, ret_err error) {
	enc := func() (_err error) {
		_err = recv.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = recv.enc.EncodeString(i_name)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = recv.dec.DecodeInterface()

		return
	}
	resp_chan, err := recv.makeCall(35, enc, dec)
	if err != nil {
		return ret_val, errgo.NoteMask(err, "Could not make call to Client.GetVvar")
	}
	resp := <-resp_chan
	if resp == nil {
		return ret_val, errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return ret_val, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	ret_val = resp.obj.(interface{})
	return ret_val, ret_err

}

func (recv *Client) GetOption(i_name string) (ret_val interface{}, ret_err error) {
	enc := func() (_err error) {
		_err = recv.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = recv.enc.EncodeString(i_name)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = recv.dec.DecodeInterface()

		return
	}
	resp_chan, err := recv.makeCall(36, enc, dec)
	if err != nil {
		return ret_val, errgo.NoteMask(err, "Could not make call to Client.GetOption")
	}
	resp := <-resp_chan
	if resp == nil {
		return ret_val, errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return ret_val, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	ret_val = resp.obj.(interface{})
	return ret_val, ret_err

}

func (recv *Client) SetOption(i_name string, i_value interface{}) (ret_err error) {
	enc := func() (_err error) {
		_err = recv.enc.EncodeSliceLen(2)
		if _err != nil {
			return
		}

		_err = recv.enc.EncodeString(i_name)

		if _err != nil {
			return
		}

		_err = recv.enc.Encode(i_value)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_, _err = recv.dec.DecodeBytes()

		return
	}
	resp_chan, err := recv.makeCall(37, enc, dec)
	if err != nil {
		return errgo.NoteMask(err, "Could not make call to Client.SetOption")
	}
	resp := <-resp_chan
	if resp == nil {
		return errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	return ret_err

}

func (recv *Client) OutWrite(i_str string) (ret_err error) {
	enc := func() (_err error) {
		_err = recv.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = recv.enc.EncodeString(i_str)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_, _err = recv.dec.DecodeBytes()

		return
	}
	resp_chan, err := recv.makeCall(38, enc, dec)
	if err != nil {
		return errgo.NoteMask(err, "Could not make call to Client.OutWrite")
	}
	resp := <-resp_chan
	if resp == nil {
		return errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	return ret_err

}

func (recv *Client) ErrWrite(i_str string) (ret_err error) {
	enc := func() (_err error) {
		_err = recv.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = recv.enc.EncodeString(i_str)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_, _err = recv.dec.DecodeBytes()

		return
	}
	resp_chan, err := recv.makeCall(39, enc, dec)
	if err != nil {
		return errgo.NoteMask(err, "Could not make call to Client.ErrWrite")
	}
	resp := <-resp_chan
	if resp == nil {
		return errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	return ret_err

}

func (recv *Client) GetBuffers() (ret_val []Buffer, ret_err error) {
	enc := func() (_err error) {
		_err = recv.enc.EncodeSliceLen(0)
		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = recv.decodeBufferSlice()

		return
	}
	resp_chan, err := recv.makeCall(40, enc, dec)
	if err != nil {
		return ret_val, errgo.NoteMask(err, "Could not make call to Client.GetBuffers")
	}
	resp := <-resp_chan
	if resp == nil {
		return ret_val, errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return ret_val, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	ret_val = resp.obj.([]Buffer)
	return ret_val, ret_err

}

func (recv *Client) GetCurrentBuffer() (ret_val Buffer, ret_err error) {
	enc := func() (_err error) {
		_err = recv.enc.EncodeSliceLen(0)
		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = recv.decodeBuffer()

		return
	}
	resp_chan, err := recv.makeCall(41, enc, dec)
	if err != nil {
		return ret_val, errgo.NoteMask(err, "Could not make call to Client.GetCurrentBuffer")
	}
	resp := <-resp_chan
	if resp == nil {
		return ret_val, errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return ret_val, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	ret_val = resp.obj.(Buffer)
	return ret_val, ret_err

}

func (recv *Client) SetCurrentBuffer(i_buffer Buffer) (ret_err error) {
	enc := func() (_err error) {
		_err = recv.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = recv.encodeBuffer(i_buffer)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_, _err = recv.dec.DecodeBytes()

		return
	}
	resp_chan, err := recv.makeCall(42, enc, dec)
	if err != nil {
		return errgo.NoteMask(err, "Could not make call to Client.SetCurrentBuffer")
	}
	resp := <-resp_chan
	if resp == nil {
		return errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	return ret_err

}

func (recv *Client) GetWindows() (ret_val []Window, ret_err error) {
	enc := func() (_err error) {
		_err = recv.enc.EncodeSliceLen(0)
		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = recv.decodeWindowSlice()

		return
	}
	resp_chan, err := recv.makeCall(43, enc, dec)
	if err != nil {
		return ret_val, errgo.NoteMask(err, "Could not make call to Client.GetWindows")
	}
	resp := <-resp_chan
	if resp == nil {
		return ret_val, errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return ret_val, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	ret_val = resp.obj.([]Window)
	return ret_val, ret_err

}

func (recv *Client) GetCurrentWindow() (ret_val Window, ret_err error) {
	enc := func() (_err error) {
		_err = recv.enc.EncodeSliceLen(0)
		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = recv.decodeWindow()

		return
	}
	resp_chan, err := recv.makeCall(44, enc, dec)
	if err != nil {
		return ret_val, errgo.NoteMask(err, "Could not make call to Client.GetCurrentWindow")
	}
	resp := <-resp_chan
	if resp == nil {
		return ret_val, errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return ret_val, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	ret_val = resp.obj.(Window)
	return ret_val, ret_err

}

func (recv *Client) SetCurrentWindow(i_window Window) (ret_err error) {
	enc := func() (_err error) {
		_err = recv.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = recv.encodeWindow(i_window)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_, _err = recv.dec.DecodeBytes()

		return
	}
	resp_chan, err := recv.makeCall(45, enc, dec)
	if err != nil {
		return errgo.NoteMask(err, "Could not make call to Client.SetCurrentWindow")
	}
	resp := <-resp_chan
	if resp == nil {
		return errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	return ret_err

}

func (recv *Client) GetTabpages() (ret_val []Tabpage, ret_err error) {
	enc := func() (_err error) {
		_err = recv.enc.EncodeSliceLen(0)
		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = recv.decodeTabpageSlice()

		return
	}
	resp_chan, err := recv.makeCall(46, enc, dec)
	if err != nil {
		return ret_val, errgo.NoteMask(err, "Could not make call to Client.GetTabpages")
	}
	resp := <-resp_chan
	if resp == nil {
		return ret_val, errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return ret_val, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	ret_val = resp.obj.([]Tabpage)
	return ret_val, ret_err

}

func (recv *Client) GetCurrentTabpage() (ret_val Tabpage, ret_err error) {
	enc := func() (_err error) {
		_err = recv.enc.EncodeSliceLen(0)
		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = recv.decodeTabpage()

		return
	}
	resp_chan, err := recv.makeCall(47, enc, dec)
	if err != nil {
		return ret_val, errgo.NoteMask(err, "Could not make call to Client.GetCurrentTabpage")
	}
	resp := <-resp_chan
	if resp == nil {
		return ret_val, errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return ret_val, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	ret_val = resp.obj.(Tabpage)
	return ret_val, ret_err

}

func (recv *Client) SetCurrentTabpage(i_tabpage Tabpage) (ret_err error) {
	enc := func() (_err error) {
		_err = recv.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = recv.encodeTabpage(i_tabpage)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_, _err = recv.dec.DecodeBytes()

		return
	}
	resp_chan, err := recv.makeCall(48, enc, dec)
	if err != nil {
		return errgo.NoteMask(err, "Could not make call to Client.SetCurrentTabpage")
	}
	resp := <-resp_chan
	if resp == nil {
		return errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	return ret_err

}

func (recv *Client) Subscribe(i_event string) (ret_err error) {
	enc := func() (_err error) {
		_err = recv.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = recv.enc.EncodeString(i_event)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_, _err = recv.dec.DecodeBytes()

		return
	}
	resp_chan, err := recv.makeCall(49, enc, dec)
	if err != nil {
		return errgo.NoteMask(err, "Could not make call to Client.Subscribe")
	}
	resp := <-resp_chan
	if resp == nil {
		return errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	return ret_err

}

func (recv *Client) Unsubscribe(i_event string) (ret_err error) {
	enc := func() (_err error) {
		_err = recv.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = recv.enc.EncodeString(i_event)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_, _err = recv.dec.DecodeBytes()

		return
	}
	resp_chan, err := recv.makeCall(50, enc, dec)
	if err != nil {
		return errgo.NoteMask(err, "Could not make call to Client.Unsubscribe")
	}
	resp := <-resp_chan
	if resp == nil {
		return errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	return ret_err

}

func (recv *Client) RegisterProvider(i_method string) (ret_err error) {
	enc := func() (_err error) {
		_err = recv.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = recv.enc.EncodeString(i_method)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_, _err = recv.dec.DecodeBytes()

		return
	}
	resp_chan, err := recv.makeCall(51, enc, dec)
	if err != nil {
		return errgo.NoteMask(err, "Could not make call to Client.RegisterProvider")
	}
	resp := <-resp_chan
	if resp == nil {
		return errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	return ret_err

}

func (recv *Window) GetBuffer() (ret_val Buffer, ret_err error) {
	enc := func() (_err error) {
		_err = recv.client.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = recv.client.encodeWindow(*recv)
		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = recv.client.decodeBuffer()

		return
	}
	resp_chan, err := recv.client.makeCall(52, enc, dec)
	if err != nil {
		return ret_val, errgo.NoteMask(err, "Could not make call to Window.GetBuffer")
	}
	resp := <-resp_chan
	if resp == nil {
		return ret_val, errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return ret_val, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	ret_val = resp.obj.(Buffer)
	return ret_val, ret_err

}

func (recv *Window) GetCursor() (ret_val uint32, ret_err error) {
	enc := func() (_err error) {
		_err = recv.client.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = recv.client.encodeWindow(*recv)
		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = recv.client.dec.DecodeUint32()

		return
	}
	resp_chan, err := recv.client.makeCall(53, enc, dec)
	if err != nil {
		return ret_val, errgo.NoteMask(err, "Could not make call to Window.GetCursor")
	}
	resp := <-resp_chan
	if resp == nil {
		return ret_val, errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return ret_val, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	ret_val = resp.obj.(uint32)
	return ret_val, ret_err

}

func (recv *Window) SetCursor(i_pos uint32) (ret_err error) {
	enc := func() (_err error) {
		_err = recv.client.enc.EncodeSliceLen(2)
		if _err != nil {
			return
		}

		_err = recv.client.encodeWindow(*recv)
		if _err != nil {
			return
		}

		_err = recv.client.enc.EncodeUint32(i_pos)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_, _err = recv.client.dec.DecodeBytes()

		return
	}
	resp_chan, err := recv.client.makeCall(54, enc, dec)
	if err != nil {
		return errgo.NoteMask(err, "Could not make call to Window.SetCursor")
	}
	resp := <-resp_chan
	if resp == nil {
		return errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	return ret_err

}

func (recv *Window) GetHeight() (ret_val int, ret_err error) {
	enc := func() (_err error) {
		_err = recv.client.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = recv.client.encodeWindow(*recv)
		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = recv.client.dec.DecodeInt()

		return
	}
	resp_chan, err := recv.client.makeCall(55, enc, dec)
	if err != nil {
		return ret_val, errgo.NoteMask(err, "Could not make call to Window.GetHeight")
	}
	resp := <-resp_chan
	if resp == nil {
		return ret_val, errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return ret_val, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	ret_val = resp.obj.(int)
	return ret_val, ret_err

}

func (recv *Window) SetHeight(i_height int) (ret_err error) {
	enc := func() (_err error) {
		_err = recv.client.enc.EncodeSliceLen(2)
		if _err != nil {
			return
		}

		_err = recv.client.encodeWindow(*recv)
		if _err != nil {
			return
		}

		_err = recv.client.enc.EncodeInt(i_height)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_, _err = recv.client.dec.DecodeBytes()

		return
	}
	resp_chan, err := recv.client.makeCall(56, enc, dec)
	if err != nil {
		return errgo.NoteMask(err, "Could not make call to Window.SetHeight")
	}
	resp := <-resp_chan
	if resp == nil {
		return errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	return ret_err

}

func (recv *Window) GetWidth() (ret_val int, ret_err error) {
	enc := func() (_err error) {
		_err = recv.client.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = recv.client.encodeWindow(*recv)
		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = recv.client.dec.DecodeInt()

		return
	}
	resp_chan, err := recv.client.makeCall(57, enc, dec)
	if err != nil {
		return ret_val, errgo.NoteMask(err, "Could not make call to Window.GetWidth")
	}
	resp := <-resp_chan
	if resp == nil {
		return ret_val, errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return ret_val, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	ret_val = resp.obj.(int)
	return ret_val, ret_err

}

func (recv *Window) SetWidth(i_width int) (ret_err error) {
	enc := func() (_err error) {
		_err = recv.client.enc.EncodeSliceLen(2)
		if _err != nil {
			return
		}

		_err = recv.client.encodeWindow(*recv)
		if _err != nil {
			return
		}

		_err = recv.client.enc.EncodeInt(i_width)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_, _err = recv.client.dec.DecodeBytes()

		return
	}
	resp_chan, err := recv.client.makeCall(58, enc, dec)
	if err != nil {
		return errgo.NoteMask(err, "Could not make call to Window.SetWidth")
	}
	resp := <-resp_chan
	if resp == nil {
		return errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	return ret_err

}

func (recv *Window) GetVar(i_name string) (ret_val interface{}, ret_err error) {
	enc := func() (_err error) {
		_err = recv.client.enc.EncodeSliceLen(2)
		if _err != nil {
			return
		}

		_err = recv.client.encodeWindow(*recv)
		if _err != nil {
			return
		}

		_err = recv.client.enc.EncodeString(i_name)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = recv.client.dec.DecodeInterface()

		return
	}
	resp_chan, err := recv.client.makeCall(59, enc, dec)
	if err != nil {
		return ret_val, errgo.NoteMask(err, "Could not make call to Window.GetVar")
	}
	resp := <-resp_chan
	if resp == nil {
		return ret_val, errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return ret_val, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	ret_val = resp.obj.(interface{})
	return ret_val, ret_err

}

func (recv *Window) SetVar(i_name string, i_value interface{}) (ret_val interface{}, ret_err error) {
	enc := func() (_err error) {
		_err = recv.client.enc.EncodeSliceLen(3)
		if _err != nil {
			return
		}

		_err = recv.client.encodeWindow(*recv)
		if _err != nil {
			return
		}

		_err = recv.client.enc.EncodeString(i_name)

		if _err != nil {
			return
		}

		_err = recv.client.enc.Encode(i_value)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = recv.client.dec.DecodeInterface()

		return
	}
	resp_chan, err := recv.client.makeCall(60, enc, dec)
	if err != nil {
		return ret_val, errgo.NoteMask(err, "Could not make call to Window.SetVar")
	}
	resp := <-resp_chan
	if resp == nil {
		return ret_val, errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return ret_val, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	ret_val = resp.obj.(interface{})
	return ret_val, ret_err

}

func (recv *Window) GetOption(i_name string) (ret_val interface{}, ret_err error) {
	enc := func() (_err error) {
		_err = recv.client.enc.EncodeSliceLen(2)
		if _err != nil {
			return
		}

		_err = recv.client.encodeWindow(*recv)
		if _err != nil {
			return
		}

		_err = recv.client.enc.EncodeString(i_name)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = recv.client.dec.DecodeInterface()

		return
	}
	resp_chan, err := recv.client.makeCall(61, enc, dec)
	if err != nil {
		return ret_val, errgo.NoteMask(err, "Could not make call to Window.GetOption")
	}
	resp := <-resp_chan
	if resp == nil {
		return ret_val, errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return ret_val, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	ret_val = resp.obj.(interface{})
	return ret_val, ret_err

}

func (recv *Window) SetOption(i_name string, i_value interface{}) (ret_err error) {
	enc := func() (_err error) {
		_err = recv.client.enc.EncodeSliceLen(3)
		if _err != nil {
			return
		}

		_err = recv.client.encodeWindow(*recv)
		if _err != nil {
			return
		}

		_err = recv.client.enc.EncodeString(i_name)

		if _err != nil {
			return
		}

		_err = recv.client.enc.Encode(i_value)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_, _err = recv.client.dec.DecodeBytes()

		return
	}
	resp_chan, err := recv.client.makeCall(62, enc, dec)
	if err != nil {
		return errgo.NoteMask(err, "Could not make call to Window.SetOption")
	}
	resp := <-resp_chan
	if resp == nil {
		return errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	return ret_err

}

func (recv *Window) GetPosition() (ret_val uint32, ret_err error) {
	enc := func() (_err error) {
		_err = recv.client.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = recv.client.encodeWindow(*recv)
		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = recv.client.dec.DecodeUint32()

		return
	}
	resp_chan, err := recv.client.makeCall(63, enc, dec)
	if err != nil {
		return ret_val, errgo.NoteMask(err, "Could not make call to Window.GetPosition")
	}
	resp := <-resp_chan
	if resp == nil {
		return ret_val, errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return ret_val, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	ret_val = resp.obj.(uint32)
	return ret_val, ret_err

}

func (recv *Window) GetTabpage() (ret_val Tabpage, ret_err error) {
	enc := func() (_err error) {
		_err = recv.client.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = recv.client.encodeWindow(*recv)
		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = recv.client.decodeTabpage()

		return
	}
	resp_chan, err := recv.client.makeCall(64, enc, dec)
	if err != nil {
		return ret_val, errgo.NoteMask(err, "Could not make call to Window.GetTabpage")
	}
	resp := <-resp_chan
	if resp == nil {
		return ret_val, errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return ret_val, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	ret_val = resp.obj.(Tabpage)
	return ret_val, ret_err

}

func (recv *Window) IsValid() (ret_val bool, ret_err error) {
	enc := func() (_err error) {
		_err = recv.client.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = recv.client.encodeWindow(*recv)
		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = recv.client.dec.DecodeBool()

		return
	}
	resp_chan, err := recv.client.makeCall(65, enc, dec)
	if err != nil {
		return ret_val, errgo.NoteMask(err, "Could not make call to Window.IsValid")
	}
	resp := <-resp_chan
	if resp == nil {
		return ret_val, errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return ret_val, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	ret_val = resp.obj.(bool)
	return ret_val, ret_err

}

// helper functions for types

func (c *Client) encodeWindowSlice(s []Window) error {
	err := c.enc.EncodeSliceLen(len(s))
	if err != nil {
		return errgo.NoteMask(err, "Could not encode slice length")
	}

	for i := 0; i < len(s); i++ {

		err := c.encodeWindow(s[i])

		if err != nil {
			return errgo.Notef(err, "Could not encode Window at index %v", i)
		}
	}

	return nil
}

func (c *Client) decodeWindowSlice() ([]Window, error) {
	l, err := c.dec.DecodeSliceLen()
	if err != nil {
		return nil, errgo.NoteMask(err, "Could not decode slice length")
	}

	res := make([]Window, l)

	for i := 0; i < l; i++ {

		b, err := c.decodeWindow()

		if err != nil {
			return nil, errgo.Notef(err, "Could not decode Window at index %v", i)
		}
		res[i] = b
	}

	return res, nil
}

func (c *Client) encodeBufferSlice(s []Buffer) error {
	err := c.enc.EncodeSliceLen(len(s))
	if err != nil {
		return errgo.NoteMask(err, "Could not encode slice length")
	}

	for i := 0; i < len(s); i++ {

		err := c.encodeBuffer(s[i])

		if err != nil {
			return errgo.Notef(err, "Could not encode Buffer at index %v", i)
		}
	}

	return nil
}

func (c *Client) decodeBufferSlice() ([]Buffer, error) {
	l, err := c.dec.DecodeSliceLen()
	if err != nil {
		return nil, errgo.NoteMask(err, "Could not decode slice length")
	}

	res := make([]Buffer, l)

	for i := 0; i < l; i++ {

		b, err := c.decodeBuffer()

		if err != nil {
			return nil, errgo.Notef(err, "Could not decode Buffer at index %v", i)
		}
		res[i] = b
	}

	return res, nil
}

func (c *Client) encodeTabpageSlice(s []Tabpage) error {
	err := c.enc.EncodeSliceLen(len(s))
	if err != nil {
		return errgo.NoteMask(err, "Could not encode slice length")
	}

	for i := 0; i < len(s); i++ {

		err := c.encodeTabpage(s[i])

		if err != nil {
			return errgo.Notef(err, "Could not encode Tabpage at index %v", i)
		}
	}

	return nil
}

func (c *Client) decodeTabpageSlice() ([]Tabpage, error) {
	l, err := c.dec.DecodeSliceLen()
	if err != nil {
		return nil, errgo.NoteMask(err, "Could not decode slice length")
	}

	res := make([]Tabpage, l)

	for i := 0; i < l; i++ {

		b, err := c.decodeTabpage()

		if err != nil {
			return nil, errgo.Notef(err, "Could not decode Tabpage at index %v", i)
		}
		res[i] = b
	}

	return res, nil
}

func (c *Client) encodeStringSlice(s []string) error {
	err := c.enc.EncodeSliceLen(len(s))
	if err != nil {
		return errgo.NoteMask(err, "Could not encode slice length")
	}

	for i := 0; i < len(s); i++ {

		err := c.enc.EncodeString(s[i])

		if err != nil {
			return errgo.Notef(err, "Could not encode string at index %v", i)
		}
	}

	return nil
}

func (c *Client) decodeStringSlice() ([]string, error) {
	l, err := c.dec.DecodeSliceLen()
	if err != nil {
		return nil, errgo.NoteMask(err, "Could not decode slice length")
	}

	res := make([]string, l)

	for i := 0; i < l; i++ {

		b, err := c.dec.DecodeString()

		if err != nil {
			return nil, errgo.Notef(err, "Could not decode string at index %v", i)
		}
		res[i] = b
	}

	return res, nil
}
