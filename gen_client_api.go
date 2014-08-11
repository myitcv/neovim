package neovim

// **** THIS FILE IS GENERATED - DO NOT EDIT BY HAND

import "github.com/juju/errgo"

// constants representing method ids

const (
	Neovim_API               NeovimMethodId = 0
	Tabpage_GetWindows                      = 1
	Tabpage_GetVar                          = 2
	Tabpage_SetVar                          = 3
	Tabpage_GetWindow                       = 4
	Tabpage_IsValid                         = 5
	Buffer_GetLength                        = 6
	Buffer_GetLine                          = 7
	Buffer_SetLine                          = 8
	Buffer_DelLine                          = 9
	Buffer_GetSlice                         = 10
	Buffer_SetSlice                         = 11
	Buffer_GetVar                           = 12
	Buffer_SetVar                           = 13
	Buffer_GetOption                        = 14
	Buffer_SetOption                        = 15
	Buffer_GetNumber                        = 16
	Buffer_GetName                          = 17
	Buffer_SetName                          = 18
	Buffer_IsValid                          = 19
	Buffer_Insert                           = 20
	Buffer_GetMark                          = 21
	Client_PushKeys                         = 22
	Client_Command                          = 23
	Client_Feedkeys                         = 24
	Client_ReplaceTermcodes                 = 25
	Client_Eval                             = 26
	Client_Strwidth                         = 27
	Client_ListRuntimePaths                 = 28
	Client_ChangeDirectory                  = 29
	Client_GetCurrentLine                   = 30
	Client_SetCurrentLine                   = 31
	Client_DelCurrentLine                   = 32
	Client_GetVar                           = 33
	Client_SetVar                           = 34
	Client_GetVvar                          = 35
	Client_GetOption                        = 36
	Client_SetOption                        = 37
	Client_OutWrite                         = 38
	Client_ErrWrite                         = 39
	Client_GetBuffers                       = 40
	Client_GetCurrentBuffer                 = 41
	Client_SetCurrentBuffer                 = 42
	Client_GetWindows                       = 43
	Client_GetCurrentWindow                 = 44
	Client_SetCurrentWindow                 = 45
	Client_GetTabpages                      = 46
	Client_GetCurrentTabpage                = 47
	Client_SetCurrentTabpage                = 48
	Client_Subscribe                        = 49
	Client_Unsubscribe                      = 50
	Client_RegisterProvider                 = 51
	Window_GetBuffer                        = 52
	Window_GetCursor                        = 53
	Window_SetCursor                        = 54
	Window_GetHeight                        = 55
	Window_SetHeight                        = 56
	Window_GetWidth                         = 57
	Window_SetWidth                         = 58
	Window_GetVar                           = 59
	Window_SetVar                           = 60
	Window_GetOption                        = 61
	Window_SetOption                        = 62
	Window_GetPosition                      = 63
	Window_GetTabpage                       = 64
	Window_IsValid                          = 65
)

func (n NeovimMethodId) String() string {
	switch n {
	case Neovim_API:
		return "API"
	case Tabpage_GetWindows:
		return "Tabpage_GetWindows"
	case Tabpage_GetVar:
		return "Tabpage_GetVar"
	case Tabpage_SetVar:
		return "Tabpage_SetVar"
	case Tabpage_GetWindow:
		return "Tabpage_GetWindow"
	case Tabpage_IsValid:
		return "Tabpage_IsValid"
	case Buffer_GetLength:
		return "Buffer_GetLength"
	case Buffer_GetLine:
		return "Buffer_GetLine"
	case Buffer_SetLine:
		return "Buffer_SetLine"
	case Buffer_DelLine:
		return "Buffer_DelLine"
	case Buffer_GetSlice:
		return "Buffer_GetSlice"
	case Buffer_SetSlice:
		return "Buffer_SetSlice"
	case Buffer_GetVar:
		return "Buffer_GetVar"
	case Buffer_SetVar:
		return "Buffer_SetVar"
	case Buffer_GetOption:
		return "Buffer_GetOption"
	case Buffer_SetOption:
		return "Buffer_SetOption"
	case Buffer_GetNumber:
		return "Buffer_GetNumber"
	case Buffer_GetName:
		return "Buffer_GetName"
	case Buffer_SetName:
		return "Buffer_SetName"
	case Buffer_IsValid:
		return "Buffer_IsValid"
	case Buffer_Insert:
		return "Buffer_Insert"
	case Buffer_GetMark:
		return "Buffer_GetMark"
	case Client_PushKeys:
		return "Client_PushKeys"
	case Client_Command:
		return "Client_Command"
	case Client_Feedkeys:
		return "Client_Feedkeys"
	case Client_ReplaceTermcodes:
		return "Client_ReplaceTermcodes"
	case Client_Eval:
		return "Client_Eval"
	case Client_Strwidth:
		return "Client_Strwidth"
	case Client_ListRuntimePaths:
		return "Client_ListRuntimePaths"
	case Client_ChangeDirectory:
		return "Client_ChangeDirectory"
	case Client_GetCurrentLine:
		return "Client_GetCurrentLine"
	case Client_SetCurrentLine:
		return "Client_SetCurrentLine"
	case Client_DelCurrentLine:
		return "Client_DelCurrentLine"
	case Client_GetVar:
		return "Client_GetVar"
	case Client_SetVar:
		return "Client_SetVar"
	case Client_GetVvar:
		return "Client_GetVvar"
	case Client_GetOption:
		return "Client_GetOption"
	case Client_SetOption:
		return "Client_SetOption"
	case Client_OutWrite:
		return "Client_OutWrite"
	case Client_ErrWrite:
		return "Client_ErrWrite"
	case Client_GetBuffers:
		return "Client_GetBuffers"
	case Client_GetCurrentBuffer:
		return "Client_GetCurrentBuffer"
	case Client_SetCurrentBuffer:
		return "Client_SetCurrentBuffer"
	case Client_GetWindows:
		return "Client_GetWindows"
	case Client_GetCurrentWindow:
		return "Client_GetCurrentWindow"
	case Client_SetCurrentWindow:
		return "Client_SetCurrentWindow"
	case Client_GetTabpages:
		return "Client_GetTabpages"
	case Client_GetCurrentTabpage:
		return "Client_GetCurrentTabpage"
	case Client_SetCurrentTabpage:
		return "Client_SetCurrentTabpage"
	case Client_Subscribe:
		return "Client_Subscribe"
	case Client_Unsubscribe:
		return "Client_Unsubscribe"
	case Client_RegisterProvider:
		return "Client_RegisterProvider"
	case Window_GetBuffer:
		return "Window_GetBuffer"
	case Window_GetCursor:
		return "Window_GetCursor"
	case Window_SetCursor:
		return "Window_SetCursor"
	case Window_GetHeight:
		return "Window_GetHeight"
	case Window_SetHeight:
		return "Window_SetHeight"
	case Window_GetWidth:
		return "Window_GetWidth"
	case Window_SetWidth:
		return "Window_SetWidth"
	case Window_GetVar:
		return "Window_GetVar"
	case Window_SetVar:
		return "Window_SetVar"
	case Window_GetOption:
		return "Window_GetOption"
	case Window_SetOption:
		return "Window_SetOption"
	case Window_GetPosition:
		return "Window_GetPosition"
	case Window_GetTabpage:
		return "Window_GetTabpage"
	case Window_IsValid:
		return "Window_IsValid"

	default:
		return ""
	}
}

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
	resp_chan, err := recv.client.makeCall(Tabpage_GetWindows, enc, dec)
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
	resp_chan, err := recv.client.makeCall(Tabpage_GetVar, enc, dec)
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
	resp_chan, err := recv.client.makeCall(Tabpage_SetVar, enc, dec)
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
	resp_chan, err := recv.client.makeCall(Tabpage_GetWindow, enc, dec)
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
	resp_chan, err := recv.client.makeCall(Tabpage_IsValid, enc, dec)
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
	resp_chan, err := recv.client.makeCall(Buffer_GetLength, enc, dec)
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
	resp_chan, err := recv.client.makeCall(Buffer_GetLine, enc, dec)
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
	resp_chan, err := recv.client.makeCall(Buffer_SetLine, enc, dec)
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
	resp_chan, err := recv.client.makeCall(Buffer_DelLine, enc, dec)
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
	resp_chan, err := recv.client.makeCall(Buffer_GetSlice, enc, dec)
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
	resp_chan, err := recv.client.makeCall(Buffer_SetSlice, enc, dec)
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
	resp_chan, err := recv.client.makeCall(Buffer_GetVar, enc, dec)
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
	resp_chan, err := recv.client.makeCall(Buffer_SetVar, enc, dec)
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
	resp_chan, err := recv.client.makeCall(Buffer_GetOption, enc, dec)
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
	resp_chan, err := recv.client.makeCall(Buffer_SetOption, enc, dec)
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
	resp_chan, err := recv.client.makeCall(Buffer_GetNumber, enc, dec)
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
	resp_chan, err := recv.client.makeCall(Buffer_GetName, enc, dec)
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
	resp_chan, err := recv.client.makeCall(Buffer_SetName, enc, dec)
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
	resp_chan, err := recv.client.makeCall(Buffer_IsValid, enc, dec)
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
	resp_chan, err := recv.client.makeCall(Buffer_Insert, enc, dec)
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
	resp_chan, err := recv.client.makeCall(Buffer_GetMark, enc, dec)
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
	resp_chan, err := recv.makeCall(Client_PushKeys, enc, dec)
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
	resp_chan, err := recv.makeCall(Client_Command, enc, dec)
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
	resp_chan, err := recv.makeCall(Client_Feedkeys, enc, dec)
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
	resp_chan, err := recv.makeCall(Client_ReplaceTermcodes, enc, dec)
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
	resp_chan, err := recv.makeCall(Client_Eval, enc, dec)
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
	resp_chan, err := recv.makeCall(Client_Strwidth, enc, dec)
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
	resp_chan, err := recv.makeCall(Client_ListRuntimePaths, enc, dec)
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
	resp_chan, err := recv.makeCall(Client_ChangeDirectory, enc, dec)
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
	resp_chan, err := recv.makeCall(Client_GetCurrentLine, enc, dec)
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
	resp_chan, err := recv.makeCall(Client_SetCurrentLine, enc, dec)
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
	resp_chan, err := recv.makeCall(Client_DelCurrentLine, enc, dec)
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
	resp_chan, err := recv.makeCall(Client_GetVar, enc, dec)
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
	resp_chan, err := recv.makeCall(Client_SetVar, enc, dec)
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
	resp_chan, err := recv.makeCall(Client_GetVvar, enc, dec)
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
	resp_chan, err := recv.makeCall(Client_GetOption, enc, dec)
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
	resp_chan, err := recv.makeCall(Client_SetOption, enc, dec)
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
	resp_chan, err := recv.makeCall(Client_OutWrite, enc, dec)
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
	resp_chan, err := recv.makeCall(Client_ErrWrite, enc, dec)
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
	resp_chan, err := recv.makeCall(Client_GetBuffers, enc, dec)
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
	resp_chan, err := recv.makeCall(Client_GetCurrentBuffer, enc, dec)
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
	resp_chan, err := recv.makeCall(Client_SetCurrentBuffer, enc, dec)
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
	resp_chan, err := recv.makeCall(Client_GetWindows, enc, dec)
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
	resp_chan, err := recv.makeCall(Client_GetCurrentWindow, enc, dec)
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
	resp_chan, err := recv.makeCall(Client_SetCurrentWindow, enc, dec)
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
	resp_chan, err := recv.makeCall(Client_GetTabpages, enc, dec)
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
	resp_chan, err := recv.makeCall(Client_GetCurrentTabpage, enc, dec)
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
	resp_chan, err := recv.makeCall(Client_SetCurrentTabpage, enc, dec)
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
	resp_chan, err := recv.makeCall(Client_Subscribe, enc, dec)
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
	resp_chan, err := recv.makeCall(Client_Unsubscribe, enc, dec)
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
	resp_chan, err := recv.makeCall(Client_RegisterProvider, enc, dec)
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
	resp_chan, err := recv.client.makeCall(Window_GetBuffer, enc, dec)
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
	resp_chan, err := recv.client.makeCall(Window_GetCursor, enc, dec)
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
	resp_chan, err := recv.client.makeCall(Window_SetCursor, enc, dec)
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
	resp_chan, err := recv.client.makeCall(Window_GetHeight, enc, dec)
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
	resp_chan, err := recv.client.makeCall(Window_SetHeight, enc, dec)
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
	resp_chan, err := recv.client.makeCall(Window_GetWidth, enc, dec)
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
	resp_chan, err := recv.client.makeCall(Window_SetWidth, enc, dec)
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
	resp_chan, err := recv.client.makeCall(Window_GetVar, enc, dec)
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
	resp_chan, err := recv.client.makeCall(Window_SetVar, enc, dec)
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
	resp_chan, err := recv.client.makeCall(Window_GetOption, enc, dec)
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
	resp_chan, err := recv.client.makeCall(Window_SetOption, enc, dec)
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
	resp_chan, err := recv.client.makeCall(Window_GetPosition, enc, dec)
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
	resp_chan, err := recv.client.makeCall(Window_GetTabpage, enc, dec)
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
	resp_chan, err := recv.client.makeCall(Window_IsValid, enc, dec)
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
