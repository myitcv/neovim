package neovim

// **** THIS FILE IS GENERATED - DO NOT EDIT BY HAND

import (
	"github.com/juju/errgo"
	"github.com/myitcv/neovim/apidef"
)

// constants representing method ids

const (
	TypeBuffer uint8 = 0
	TypeTabpage
	TypeWindow
)

const (
	bufferDelLine           = "buffer_del_line"
	bufferLineCount         = "buffer_line_count"
	bufferGetLine           = "buffer_get_line"
	bufferGetMark           = "buffer_get_mark"
	bufferGetName           = "buffer_get_name"
	bufferGetNumber         = "buffer_get_number"
	bufferGetOption         = "buffer_get_option"
	bufferGetLineSlice      = "buffer_get_line_slice"
	bufferGetVar            = "buffer_get_var"
	bufferInsert            = "buffer_insert"
	bufferIsValid           = "buffer_is_valid"
	bufferSetLine           = "buffer_set_line"
	bufferSetName           = "buffer_set_name"
	bufferSetOption         = "buffer_set_option"
	bufferSetLineSlice      = "buffer_set_line_slice"
	bufferSetVar            = "buffer_set_var"
	tabpageGetVar           = "tabpage_get_var"
	tabpageGetWindow        = "tabpage_get_window"
	tabpageGetWindows       = "tabpage_get_windows"
	tabpageIsValid          = "tabpage_is_valid"
	tabpageSetVar           = "tabpage_set_var"
	clientChangeDirectory   = "vim_change_directory"
	clientCommand           = "vim_command"
	clientDelCurrentLine    = "vim_del_current_line"
	clientErrWrite          = "vim_err_write"
	clientEval              = "vim_eval"
	clientFeedkeys          = "vim_feedkeys"
	clientGetAPIInfo        = "vim_get_api_info"
	clientGetBuffers        = "vim_get_buffers"
	clientGetCurrentBuffer  = "vim_get_current_buffer"
	clientGetCurrentLine    = "vim_get_current_line"
	clientGetCurrentTabpage = "vim_get_current_tabpage"
	clientGetCurrentWindow  = "vim_get_current_window"
	clientGetOption         = "vim_get_option"
	clientGetTabpages       = "vim_get_tabpages"
	clientGetVar            = "vim_get_var"
	clientGetVvar           = "vim_get_vvar"
	clientGetWindows        = "vim_get_windows"
	clientListRuntimePaths  = "vim_list_runtime_paths"
	clientOutWrite          = "vim_out_write"
	clientPushKeys          = "vim_push_keys"
	clientregisterProvider  = "vim_register_provider"
	clientReplaceTermcodes  = "vim_replace_termcodes"
	clientReportError       = "vim_report_error"
	clientSetCurrentBuffer  = "vim_set_current_buffer"
	clientSetCurrentLine    = "vim_set_current_line"
	clientSetCurrentTabpage = "vim_set_current_tabpage"
	clientSetCurrentWindow  = "vim_set_current_window"
	clientSetOption         = "vim_set_option"
	clientSetVar            = "vim_set_var"
	clientStrwidth          = "vim_strwidth"
	clientsubscribe         = "vim_subscribe"
	clientunsubscribe       = "vim_unsubscribe"
	windowGetBuffer         = "window_get_buffer"
	windowGetCursor         = "window_get_cursor"
	windowGetHeight         = "window_get_height"
	windowGetOption         = "window_get_option"
	windowGetPosition       = "window_get_position"
	windowGetTabpage        = "window_get_tabpage"
	windowGetVar            = "window_get_var"
	windowGetWidth          = "window_get_width"
	windowIsValid           = "window_is_valid"
	windowSetCursor         = "window_set_cursor"
	windowSetHeight         = "window_set_height"
	windowSetOption         = "window_set_option"
	windowSetVar            = "window_set_var"
	windowSetWidth          = "window_set_width"
)

// methods on the API

// DelLine waiting for documentation from Neovim
func (b *Buffer) DelLine(index int) error {

	enc := func() (_err error) {
		_err = b.client.enc.EncodeSliceLen(2)
		if _err != nil {
			return
		}

		_err = b.client.encodeBuffer(*b)
		if _err != nil {
			return
		}

		_err = b.client.enc.EncodeInt(index)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_, _err = b.client.dec.DecodeBytes()

		return
	}
	respChan, err := b.client.makeCall(bufferDelLine, enc, dec)
	if err != nil {
		return b.client.panicOrReturn(errgo.NoteMask(err, "Could not make call to Buffer.DelLine"))
	}
	resp := <-respChan
	if resp == nil {
		return b.client.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return b.client.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	return nil

}

// GetLineCount waiting for documentation from Neovim
func (b *Buffer) GetLineCount() (int, error) {
	var retVal int
	enc := func() (_err error) {
		_err = b.client.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = b.client.encodeBuffer(*b)
		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = b.client.dec.DecodeInt()

		return
	}
	respChan, err := b.client.makeCall(bufferLineCount, enc, dec)
	if err != nil {
		return retVal, b.client.panicOrReturn(errgo.NoteMask(err, "Could not make call to Buffer.LineCount"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, b.client.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, b.client.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(int)
	return retVal, nil

}

// GetLine waiting for documentation from Neovim
func (b *Buffer) GetLine(index int) (string, error) {
	res, err := b.GetLineRaw(index)
	return string(res), err
}

// GetLineRaw waiting for documentation from Neovim
func (b *Buffer) GetLineRaw(index int) ([]byte, error) {
	var retVal []byte
	enc := func() (_err error) {
		_err = b.client.enc.EncodeSliceLen(2)
		if _err != nil {
			return
		}

		_err = b.client.encodeBuffer(*b)
		if _err != nil {
			return
		}

		_err = b.client.enc.EncodeInt(index)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = b.client.dec.DecodeBytes()

		return
	}
	respChan, err := b.client.makeCall(bufferGetLine, enc, dec)
	if err != nil {
		return retVal, b.client.panicOrReturn(errgo.NoteMask(err, "Could not make call to Buffer.GetLine"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, b.client.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, b.client.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.([]byte)
	return retVal, nil

}

// GetMark waiting for documentation from Neovim
func (b *Buffer) GetMark(name string) ([]int, error) {
	return b.GetMarkRaw([]byte(name))
}

// GetMarkRaw waiting for documentation from Neovim
func (b *Buffer) GetMarkRaw(name []byte) ([]int, error) {
	var retVal []int
	enc := func() (_err error) {
		_err = b.client.enc.EncodeSliceLen(2)
		if _err != nil {
			return
		}

		_err = b.client.encodeBuffer(*b)
		if _err != nil {
			return
		}

		_err = b.client.enc.EncodeBytes(name)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = b.client.decodeIntSlice()

		return
	}
	respChan, err := b.client.makeCall(bufferGetMark, enc, dec)
	if err != nil {
		return retVal, b.client.panicOrReturn(errgo.NoteMask(err, "Could not make call to Buffer.GetMark"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, b.client.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, b.client.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.([]int)
	return retVal, nil

}

// GetName waiting for documentation from Neovim
func (b *Buffer) GetName() (string, error) {
	res, err := b.GetNameRaw()
	return string(res), err
}

// GetNameRaw waiting for documentation from Neovim
func (b *Buffer) GetNameRaw() ([]byte, error) {
	var retVal []byte
	enc := func() (_err error) {
		_err = b.client.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = b.client.encodeBuffer(*b)
		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = b.client.dec.DecodeBytes()

		return
	}
	respChan, err := b.client.makeCall(bufferGetName, enc, dec)
	if err != nil {
		return retVal, b.client.panicOrReturn(errgo.NoteMask(err, "Could not make call to Buffer.GetName"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, b.client.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, b.client.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.([]byte)
	return retVal, nil

}

// GetNumber waiting for documentation from Neovim
func (b *Buffer) GetNumber() (int, error) {
	var retVal int
	enc := func() (_err error) {
		_err = b.client.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = b.client.encodeBuffer(*b)
		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = b.client.dec.DecodeInt()

		return
	}
	respChan, err := b.client.makeCall(bufferGetNumber, enc, dec)
	if err != nil {
		return retVal, b.client.panicOrReturn(errgo.NoteMask(err, "Could not make call to Buffer.GetNumber"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, b.client.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, b.client.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(int)
	return retVal, nil

}

// GetOption waiting for documentation from Neovim
func (b *Buffer) GetOption(name string) (interface{}, error) {
	return b.GetOptionRaw([]byte(name))
}

// GetOptionRaw waiting for documentation from Neovim
func (b *Buffer) GetOptionRaw(name []byte) (interface{}, error) {
	var retVal interface{}
	enc := func() (_err error) {
		_err = b.client.enc.EncodeSliceLen(2)
		if _err != nil {
			return
		}

		_err = b.client.encodeBuffer(*b)
		if _err != nil {
			return
		}

		_err = b.client.enc.EncodeBytes(name)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = b.client.dec.DecodeInterface()

		return
	}
	respChan, err := b.client.makeCall(bufferGetOption, enc, dec)
	if err != nil {
		return retVal, b.client.panicOrReturn(errgo.NoteMask(err, "Could not make call to Buffer.GetOption"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, b.client.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, b.client.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(interface{})
	return retVal, nil

}

// GetLineSlice waiting for documentation from Neovim
func (b *Buffer) GetLineSlice(start int, end int, includeStart bool, includeEnd bool) ([]string, error) {
	res, err := b.GetLineSliceRaw(start, end, includeStart, includeEnd)

	if res != nil {
		res_s := make([]string, len(res))
		for i := range res {
			res_s[i] = string(res[i])
		}

		return res_s, err
	}

	return nil, err
}

// GetLineSliceRaw waiting for documentation from Neovim
func (b *Buffer) GetLineSliceRaw(start int, end int, includeStart bool, includeEnd bool) ([][]byte, error) {
	var retVal [][]byte
	enc := func() (_err error) {
		_err = b.client.enc.EncodeSliceLen(5)
		if _err != nil {
			return
		}

		_err = b.client.encodeBuffer(*b)
		if _err != nil {
			return
		}

		_err = b.client.enc.EncodeInt(start)

		if _err != nil {
			return
		}

		_err = b.client.enc.EncodeInt(end)

		if _err != nil {
			return
		}

		_err = b.client.enc.EncodeBool(includeStart)

		if _err != nil {
			return
		}

		_err = b.client.enc.EncodeBool(includeEnd)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = b.client.decodeStringSlice()

		return
	}
	respChan, err := b.client.makeCall(bufferGetLineSlice, enc, dec)
	if err != nil {
		return retVal, b.client.panicOrReturn(errgo.NoteMask(err, "Could not make call to Buffer.GetLineSlice"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, b.client.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, b.client.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.([][]byte)
	return retVal, nil

}

// GetVar waiting for documentation from Neovim
func (b *Buffer) GetVar(name string) (interface{}, error) {
	return b.GetVarRaw([]byte(name))
}

// GetVarRaw waiting for documentation from Neovim
func (b *Buffer) GetVarRaw(name []byte) (interface{}, error) {
	var retVal interface{}
	enc := func() (_err error) {
		_err = b.client.enc.EncodeSliceLen(2)
		if _err != nil {
			return
		}

		_err = b.client.encodeBuffer(*b)
		if _err != nil {
			return
		}

		_err = b.client.enc.EncodeBytes(name)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = b.client.dec.DecodeInterface()

		return
	}
	respChan, err := b.client.makeCall(bufferGetVar, enc, dec)
	if err != nil {
		return retVal, b.client.panicOrReturn(errgo.NoteMask(err, "Could not make call to Buffer.GetVar"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, b.client.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, b.client.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(interface{})
	return retVal, nil

}

// Insert waiting for documentation from Neovim
func (b *Buffer) Insert(lnum int, lines []string) error {
	var lines_b [][]byte
	if lines != nil {
		lines_b = make([][]byte, len(lines))
		for i := range lines {
			lines_b[i] = []byte(lines[i])
		}
	}
	return b.InsertRaw(lnum, lines_b)

}

// InsertRaw waiting for documentation from Neovim
func (b *Buffer) InsertRaw(lnum int, lines [][]byte) error {

	enc := func() (_err error) {
		_err = b.client.enc.EncodeSliceLen(3)
		if _err != nil {
			return
		}

		_err = b.client.encodeBuffer(*b)
		if _err != nil {
			return
		}

		_err = b.client.enc.EncodeInt(lnum)

		if _err != nil {
			return
		}

		_err = b.client.encodeStringSlice(lines)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_, _err = b.client.dec.DecodeBytes()

		return
	}
	respChan, err := b.client.makeCall(bufferInsert, enc, dec)
	if err != nil {
		return b.client.panicOrReturn(errgo.NoteMask(err, "Could not make call to Buffer.Insert"))
	}
	resp := <-respChan
	if resp == nil {
		return b.client.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return b.client.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	return nil

}

// IsValid waiting for documentation from Neovim
func (b *Buffer) IsValid() (bool, error) {
	var retVal bool
	enc := func() (_err error) {
		_err = b.client.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = b.client.encodeBuffer(*b)
		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = b.client.dec.DecodeBool()

		return
	}
	respChan, err := b.client.makeCall(bufferIsValid, enc, dec)
	if err != nil {
		return retVal, b.client.panicOrReturn(errgo.NoteMask(err, "Could not make call to Buffer.IsValid"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, b.client.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, b.client.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(bool)
	return retVal, nil

}

// SetLineRaw waiting for documentation from Neovim
func (b *Buffer) SetLine(index int, line string) error {
	return b.SetLineRaw(index, []byte(line))
}

// SetLineRaw waiting for documentation from Neovim
func (b *Buffer) SetLineRaw(index int, line []byte) error {

	enc := func() (_err error) {
		_err = b.client.enc.EncodeSliceLen(3)
		if _err != nil {
			return
		}

		_err = b.client.encodeBuffer(*b)
		if _err != nil {
			return
		}

		_err = b.client.enc.EncodeInt(index)

		if _err != nil {
			return
		}

		_err = b.client.enc.EncodeBytes(line)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_, _err = b.client.dec.DecodeBytes()

		return
	}
	respChan, err := b.client.makeCall(bufferSetLine, enc, dec)
	if err != nil {
		return b.client.panicOrReturn(errgo.NoteMask(err, "Could not make call to Buffer.SetLine"))
	}
	resp := <-respChan
	if resp == nil {
		return b.client.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return b.client.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	return nil

}

// SetName waiting for documentation from Neovim
func (b *Buffer) SetName(name string) error {
	return b.SetNameRaw([]byte(name))
}

// SetNameRaw waiting for documentation from Neovim
func (b *Buffer) SetNameRaw(name []byte) error {

	enc := func() (_err error) {
		_err = b.client.enc.EncodeSliceLen(2)
		if _err != nil {
			return
		}

		_err = b.client.encodeBuffer(*b)
		if _err != nil {
			return
		}

		_err = b.client.enc.EncodeBytes(name)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_, _err = b.client.dec.DecodeBytes()

		return
	}
	respChan, err := b.client.makeCall(bufferSetName, enc, dec)
	if err != nil {
		return b.client.panicOrReturn(errgo.NoteMask(err, "Could not make call to Buffer.SetName"))
	}
	resp := <-respChan
	if resp == nil {
		return b.client.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return b.client.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	return nil

}

// SetOption waiting for documentation from Neovim
func (b *Buffer) SetOption(name string, value interface{}) error {
	return b.SetOptionRaw([]byte(name), value)
}

// SetOptionRaw waiting for documentation from Neovim
func (b *Buffer) SetOptionRaw(name []byte, value interface{}) error {

	enc := func() (_err error) {
		_err = b.client.enc.EncodeSliceLen(3)
		if _err != nil {
			return
		}

		_err = b.client.encodeBuffer(*b)
		if _err != nil {
			return
		}

		_err = b.client.enc.EncodeBytes(name)

		if _err != nil {
			return
		}

		_err = b.client.enc.Encode(value)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_, _err = b.client.dec.DecodeBytes()

		return
	}
	respChan, err := b.client.makeCall(bufferSetOption, enc, dec)
	if err != nil {
		return b.client.panicOrReturn(errgo.NoteMask(err, "Could not make call to Buffer.SetOption"))
	}
	resp := <-respChan
	if resp == nil {
		return b.client.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return b.client.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	return nil

}

// SetLineSlice waiting for documentation from Neovim
func (b *Buffer) SetLineSlice(start int, end int, includeStart bool, includeEnd bool, replacement []string) error {
	var replacement_b [][]byte
	if replacement != nil {
		replacement_b = make([][]byte, len(replacement))
		for i := range replacement {
			replacement_b[i] = []byte(replacement[i])
		}
	}
	return b.SetLineSliceRaw(start, end, includeStart, includeEnd, replacement_b)
}

// SetLineSliceRaw waiting for documentation from Neovim
func (b *Buffer) SetLineSliceRaw(start int, end int, includeStart bool, includeEnd bool, replacement [][]byte) error {

	enc := func() (_err error) {
		_err = b.client.enc.EncodeSliceLen(6)
		if _err != nil {
			return
		}

		_err = b.client.encodeBuffer(*b)
		if _err != nil {
			return
		}

		_err = b.client.enc.EncodeInt(start)

		if _err != nil {
			return
		}

		_err = b.client.enc.EncodeInt(end)

		if _err != nil {
			return
		}

		_err = b.client.enc.EncodeBool(includeStart)

		if _err != nil {
			return
		}

		_err = b.client.enc.EncodeBool(includeEnd)

		if _err != nil {
			return
		}

		_err = b.client.encodeStringSlice(replacement)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_, _err = b.client.dec.DecodeBytes()

		return
	}
	respChan, err := b.client.makeCall(bufferSetLineSlice, enc, dec)
	if err != nil {
		return b.client.panicOrReturn(errgo.NoteMask(err, "Could not make call to Buffer.SetLineSlice"))
	}
	resp := <-respChan
	if resp == nil {
		return b.client.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return b.client.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	return nil

}

// SetVar waiting for documentation from Neovim
func (b *Buffer) SetVar(name string, value interface{}) (interface{}, error) {
	return b.SetVarRaw([]byte(name), value)
}

// SetVarRaw waiting for documentation from Neovim
func (b *Buffer) SetVarRaw(name []byte, value interface{}) (interface{}, error) {
	var retVal interface{}
	enc := func() (_err error) {
		_err = b.client.enc.EncodeSliceLen(3)
		if _err != nil {
			return
		}

		_err = b.client.encodeBuffer(*b)
		if _err != nil {
			return
		}

		_err = b.client.enc.EncodeBytes(name)

		if _err != nil {
			return
		}

		_err = b.client.enc.Encode(value)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = b.client.dec.DecodeInterface()

		return
	}
	respChan, err := b.client.makeCall(bufferSetVar, enc, dec)
	if err != nil {
		return retVal, b.client.panicOrReturn(errgo.NoteMask(err, "Could not make call to Buffer.SetVar"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, b.client.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, b.client.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(interface{})
	return retVal, nil

}

// GetVar waiting for documentation from Neovim
func (t *Tabpage) GetVar(name string) (interface{}, error) {
	return t.GetVarRaw([]byte(name))
}

// GetVarRaw waiting for documentation from Neovim
func (t *Tabpage) GetVarRaw(name []byte) (interface{}, error) {
	var retVal interface{}
	enc := func() (_err error) {
		_err = t.client.enc.EncodeSliceLen(2)
		if _err != nil {
			return
		}

		_err = t.client.encodeTabpage(*t)
		if _err != nil {
			return
		}

		_err = t.client.enc.EncodeBytes(name)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = t.client.dec.DecodeInterface()

		return
	}
	respChan, err := t.client.makeCall(tabpageGetVar, enc, dec)
	if err != nil {
		return retVal, t.client.panicOrReturn(errgo.NoteMask(err, "Could not make call to Tabpage.GetVar"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, t.client.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, t.client.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(interface{})
	return retVal, nil

}

// GetWindow waiting for documentation from Neovim
func (t *Tabpage) GetWindow() (Window, error) {
	var retVal Window
	enc := func() (_err error) {
		_err = t.client.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = t.client.encodeTabpage(*t)
		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = t.client.decodeWindow()

		return
	}
	respChan, err := t.client.makeCall(tabpageGetWindow, enc, dec)
	if err != nil {
		return retVal, t.client.panicOrReturn(errgo.NoteMask(err, "Could not make call to Tabpage.GetWindow"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, t.client.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, t.client.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(Window)
	return retVal, nil

}

// GetWindows waiting for documentation from Neovim
func (t *Tabpage) GetWindows() ([]Window, error) {
	var retVal []Window
	enc := func() (_err error) {
		_err = t.client.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = t.client.encodeTabpage(*t)
		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = t.client.decodeWindowSlice()

		return
	}
	respChan, err := t.client.makeCall(tabpageGetWindows, enc, dec)
	if err != nil {
		return retVal, t.client.panicOrReturn(errgo.NoteMask(err, "Could not make call to Tabpage.GetWindows"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, t.client.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, t.client.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.([]Window)
	return retVal, nil

}

// IsValid waiting for documentation from Neovim
func (t *Tabpage) IsValid() (bool, error) {
	var retVal bool
	enc := func() (_err error) {
		_err = t.client.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = t.client.encodeTabpage(*t)
		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = t.client.dec.DecodeBool()

		return
	}
	respChan, err := t.client.makeCall(tabpageIsValid, enc, dec)
	if err != nil {
		return retVal, t.client.panicOrReturn(errgo.NoteMask(err, "Could not make call to Tabpage.IsValid"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, t.client.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, t.client.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(bool)
	return retVal, nil

}

// SetVar waiting for documentation from Neovim
func (t *Tabpage) SetVar(name string, value interface{}) (interface{}, error) {
	return t.SetVarRaw([]byte(name), value)
}

// SetVarRaw waiting for documentation from Neovim
func (t *Tabpage) SetVarRaw(name []byte, value interface{}) (interface{}, error) {
	var retVal interface{}
	enc := func() (_err error) {
		_err = t.client.enc.EncodeSliceLen(3)
		if _err != nil {
			return
		}

		_err = t.client.encodeTabpage(*t)
		if _err != nil {
			return
		}

		_err = t.client.enc.EncodeBytes(name)

		if _err != nil {
			return
		}

		_err = t.client.enc.Encode(value)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = t.client.dec.DecodeInterface()

		return
	}
	respChan, err := t.client.makeCall(tabpageSetVar, enc, dec)
	if err != nil {
		return retVal, t.client.panicOrReturn(errgo.NoteMask(err, "Could not make call to Tabpage.SetVar"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, t.client.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, t.client.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(interface{})
	return retVal, nil

}

// ChangeDirectory waiting for documentation from Neovim
func (c *Client) ChangeDirectory(dir string) error {
	return c.ChangeDirectoryRaw([]byte(dir))
}

// ChangeDirectoryRaw waiting for documentation from Neovim
func (c *Client) ChangeDirectoryRaw(dir []byte) error {

	enc := func() (_err error) {
		_err = c.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = c.enc.EncodeBytes(dir)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_, _err = c.dec.DecodeBytes()

		return
	}
	respChan, err := c.makeCall(clientChangeDirectory, enc, dec)
	if err != nil {
		return c.panicOrReturn(errgo.NoteMask(err, "Could not make call to Client.ChangeDirectory"))
	}
	resp := <-respChan
	if resp == nil {
		return c.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return c.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	return nil

}

// Command waiting for documentation from Neovim
func (c *Client) Command(str string) error {
	return c.CommandRaw([]byte(str))
}

// CommandRaw waiting for documentation from Neovim
func (c *Client) CommandRaw(str []byte) error {

	enc := func() (_err error) {
		_err = c.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = c.enc.EncodeBytes(str)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_, _err = c.dec.DecodeBytes()

		return
	}
	respChan, err := c.makeCall(clientCommand, enc, dec)
	if err != nil {
		return c.panicOrReturn(errgo.NoteMask(err, "Could not make call to Client.Command"))
	}
	resp := <-respChan
	if resp == nil {
		return c.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return c.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	return nil

}

// DelCurrentLine waiting for documentation from Neovim
func (c *Client) DelCurrentLine() error {

	enc := func() (_err error) {
		_err = c.enc.EncodeSliceLen(0)
		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_, _err = c.dec.DecodeBytes()

		return
	}
	respChan, err := c.makeCall(clientDelCurrentLine, enc, dec)
	if err != nil {
		return c.panicOrReturn(errgo.NoteMask(err, "Could not make call to Client.DelCurrentLine"))
	}
	resp := <-respChan
	if resp == nil {
		return c.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return c.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	return nil

}

// ErrWrite waiting for documentation from Neovim
func (c *Client) ErrWrite(str string) error {
	return c.ErrWriteRaw([]byte(str))
}

// ErrWriteRaw waiting for documentation from Neovim
func (c *Client) ErrWriteRaw(str []byte) error {

	enc := func() (_err error) {
		_err = c.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = c.enc.EncodeBytes(str)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_, _err = c.dec.DecodeBytes()

		return
	}
	respChan, err := c.makeCall(clientErrWrite, enc, dec)
	if err != nil {
		return c.panicOrReturn(errgo.NoteMask(err, "Could not make call to Client.ErrWrite"))
	}
	resp := <-respChan
	if resp == nil {
		return c.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return c.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	return nil

}

// EvalRaw waiting for documentation from Neovim
func (c *Client) Eval(str string) (interface{}, error) {
	return c.EvalRaw([]byte(str))
}

// EvalRaw waiting for documentation from Neovim
func (c *Client) EvalRaw(str []byte) (interface{}, error) {
	var retVal interface{}
	enc := func() (_err error) {
		_err = c.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = c.enc.EncodeBytes(str)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = c.dec.DecodeInterface()

		return
	}
	respChan, err := c.makeCall(clientEval, enc, dec)
	if err != nil {
		return retVal, c.panicOrReturn(errgo.NoteMask(err, "Could not make call to Client.Eval"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, c.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, c.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(interface{})
	return retVal, nil

}

// Feedkeys waiting for documentation from Neovim
func (c *Client) Feedkeys(keys string, mode string) error {
	return c.FeedkeysRaw([]byte(keys), []byte(mode))
}

// FeedkeysRaw waiting for documentation from Neovim
func (c *Client) FeedkeysRaw(keys []byte, mode []byte) error {

	enc := func() (_err error) {
		_err = c.enc.EncodeSliceLen(2)
		if _err != nil {
			return
		}

		_err = c.enc.EncodeBytes(keys)

		if _err != nil {
			return
		}

		_err = c.enc.EncodeBytes(mode)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_, _err = c.dec.DecodeBytes()

		return
	}
	respChan, err := c.makeCall(clientFeedkeys, enc, dec)
	if err != nil {
		return c.panicOrReturn(errgo.NoteMask(err, "Could not make call to Client.Feedkeys"))
	}
	resp := <-respChan
	if resp == nil {
		return c.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return c.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	return nil

}

// GetBuffers waiting for documentation from Neovim
func (c *Client) GetAPIInfo() (uint8, *apidef.API, error) {
	var retChanID uint8
	var retAPI *apidef.API
	enc := func() (_err error) {
		_err = c.enc.EncodeSliceLen(0)
		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		l, _err := c.dec.DecodeSliceLen()
		if _err != nil {
			return
		}

		if l != 2 {
			return nil, errgo.Newf("Expected slice len to be 2; got %v", l)
		}

		chanID, _err := c.dec.DecodeUint8()
		if _err != nil {
			return
		}

		api, _err := apidef.GetAPI(c.dec)
		if _err != nil {
			return
		}

		_i = []interface{}{chanID, api}

		return
	}
	respChan, err := c.makeCall(clientGetAPIInfo, enc, dec)
	if err != nil {
		return retChanID, retAPI, c.panicOrReturn(errgo.NoteMask(err, "Could not make call to Client.GetBuffers"))
	}
	resp := <-respChan
	if resp == nil {
		return retChanID, retAPI, c.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retChanID, retAPI, c.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	retVal := resp.obj.([]interface{})
	retChanID = retVal[0].(uint8)
	retAPI = retVal[1].(*apidef.API)
	return retChanID, retAPI, nil

}

// GetBuffers waiting for documentation from Neovim
func (c *Client) GetBuffers() ([]Buffer, error) {
	var retVal []Buffer
	enc := func() (_err error) {
		_err = c.enc.EncodeSliceLen(0)
		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = c.decodeBufferSlice()

		return
	}
	respChan, err := c.makeCall(clientGetBuffers, enc, dec)
	if err != nil {
		return retVal, c.panicOrReturn(errgo.NoteMask(err, "Could not make call to Client.GetBuffers"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, c.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, c.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.([]Buffer)
	return retVal, nil

}

// GetCurrentBuffer waiting for documentation from Neovim
func (c *Client) GetCurrentBuffer() (Buffer, error) {
	var retVal Buffer
	enc := func() (_err error) {
		_err = c.enc.EncodeSliceLen(0)
		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = c.decodeBuffer()

		return
	}
	respChan, err := c.makeCall(clientGetCurrentBuffer, enc, dec)
	if err != nil {
		return retVal, c.panicOrReturn(errgo.NoteMask(err, "Could not make call to Client.GetCurrentBuffer"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, c.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, c.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(Buffer)
	return retVal, nil

}

// GetCurrentLine waiting for documentation from Neovim
func (c *Client) GetCurrentLine() (string, error) {
	res, err := c.GetCurrentLineRaw()
	return string(res), err
}

// GetCurrentLineRaw waiting for documentation from Neovim
func (c *Client) GetCurrentLineRaw() ([]byte, error) {
	var retVal []byte
	enc := func() (_err error) {
		_err = c.enc.EncodeSliceLen(0)
		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = c.dec.DecodeBytes()

		return
	}
	respChan, err := c.makeCall(clientGetCurrentLine, enc, dec)
	if err != nil {
		return retVal, c.panicOrReturn(errgo.NoteMask(err, "Could not make call to Client.GetCurrentLine"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, c.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, c.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.([]byte)
	return retVal, nil

}

// GetCurrentTabpage waiting for documentation from Neovim
func (c *Client) GetCurrentTabpage() (Tabpage, error) {
	var retVal Tabpage
	enc := func() (_err error) {
		_err = c.enc.EncodeSliceLen(0)
		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = c.decodeTabpage()

		return
	}
	respChan, err := c.makeCall(clientGetCurrentTabpage, enc, dec)
	if err != nil {
		return retVal, c.panicOrReturn(errgo.NoteMask(err, "Could not make call to Client.GetCurrentTabpage"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, c.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, c.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(Tabpage)
	return retVal, nil

}

// GetCurrentWindow waiting for documentation from Neovim
func (c *Client) GetCurrentWindow() (Window, error) {
	var retVal Window
	enc := func() (_err error) {
		_err = c.enc.EncodeSliceLen(0)
		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = c.decodeWindow()

		return
	}
	respChan, err := c.makeCall(clientGetCurrentWindow, enc, dec)
	if err != nil {
		return retVal, c.panicOrReturn(errgo.NoteMask(err, "Could not make call to Client.GetCurrentWindow"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, c.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, c.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(Window)
	return retVal, nil

}

// GetOption waiting for documentation from Neovim
func (c *Client) GetOption(name string) (interface{}, error) {
	return c.GetOptionRaw([]byte(name))
}

// GetOptionRaw waiting for documentation from Neovim
func (c *Client) GetOptionRaw(name []byte) (interface{}, error) {
	var retVal interface{}
	enc := func() (_err error) {
		_err = c.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = c.enc.EncodeBytes(name)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = c.dec.DecodeInterface()

		return
	}
	respChan, err := c.makeCall(clientGetOption, enc, dec)
	if err != nil {
		return retVal, c.panicOrReturn(errgo.NoteMask(err, "Could not make call to Client.GetOption"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, c.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, c.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(interface{})
	return retVal, nil

}

// GetTabpages waiting for documentation from Neovim
func (c *Client) GetTabpages() ([]Tabpage, error) {
	var retVal []Tabpage
	enc := func() (_err error) {
		_err = c.enc.EncodeSliceLen(0)
		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = c.decodeTabpageSlice()

		return
	}
	respChan, err := c.makeCall(clientGetTabpages, enc, dec)
	if err != nil {
		return retVal, c.panicOrReturn(errgo.NoteMask(err, "Could not make call to Client.GetTabpages"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, c.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, c.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.([]Tabpage)
	return retVal, nil

}

// GetVar waiting for documentation from Neovim
func (c *Client) GetVar(name string) (interface{}, error) {
	return c.GetVarRaw([]byte(name))
}

// GetVarRaw waiting for documentation from Neovim
func (c *Client) GetVarRaw(name []byte) (interface{}, error) {
	var retVal interface{}
	enc := func() (_err error) {
		_err = c.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = c.enc.EncodeBytes(name)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = c.dec.DecodeInterface()

		return
	}
	respChan, err := c.makeCall(clientGetVar, enc, dec)
	if err != nil {
		return retVal, c.panicOrReturn(errgo.NoteMask(err, "Could not make call to Client.GetVar"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, c.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, c.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(interface{})
	return retVal, nil

}

// GetVvar waiting for documentation from Neovim
func (c *Client) GetVvar(name string) (interface{}, error) {
	return c.GetVvarRaw([]byte(name))
}

// GetVvarRaw waiting for documentation from Neovim
func (c *Client) GetVvarRaw(name []byte) (interface{}, error) {
	var retVal interface{}
	enc := func() (_err error) {
		_err = c.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = c.enc.EncodeBytes(name)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = c.dec.DecodeInterface()

		return
	}
	respChan, err := c.makeCall(clientGetVvar, enc, dec)
	if err != nil {
		return retVal, c.panicOrReturn(errgo.NoteMask(err, "Could not make call to Client.GetVvar"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, c.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, c.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(interface{})
	return retVal, nil

}

// GetWindows waiting for documentation from Neovim
func (c *Client) GetWindows() ([]Window, error) {
	var retVal []Window
	enc := func() (_err error) {
		_err = c.enc.EncodeSliceLen(0)
		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = c.decodeWindowSlice()

		return
	}
	respChan, err := c.makeCall(clientGetWindows, enc, dec)
	if err != nil {
		return retVal, c.panicOrReturn(errgo.NoteMask(err, "Could not make call to Client.GetWindows"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, c.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, c.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.([]Window)
	return retVal, nil

}

// ListRuntimePaths waiting for documentation from Neovim
func (c *Client) ListRuntimePaths() ([]string, error) {
	var res_s []string
	res, err := c.ListRuntimePathsRaw()
	if res != nil {
		res_s := make([]string, len(res))
		for i := range res {
			res_s[i] = string(res[i])
		}
	}
	return res_s, err
}

// ListRuntimePathsRaw waiting for documentation from Neovim
func (c *Client) ListRuntimePathsRaw() ([][]byte, error) {
	var retVal [][]byte
	enc := func() (_err error) {
		_err = c.enc.EncodeSliceLen(0)
		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = c.decodeStringSlice()

		return
	}
	respChan, err := c.makeCall(clientListRuntimePaths, enc, dec)
	if err != nil {
		return retVal, c.panicOrReturn(errgo.NoteMask(err, "Could not make call to Client.ListRuntimePaths"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, c.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, c.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.([][]byte)
	return retVal, nil

}

// OutWrite waiting for documentation from Neovim
func (c *Client) OutWrite(str string) error {
	return c.OutWriteRaw([]byte(str))
}

// OutWriteRaw waiting for documentation from Neovim
func (c *Client) OutWriteRaw(str []byte) error {

	enc := func() (_err error) {
		_err = c.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = c.enc.EncodeBytes(str)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_, _err = c.dec.DecodeBytes()

		return
	}
	respChan, err := c.makeCall(clientOutWrite, enc, dec)
	if err != nil {
		return c.panicOrReturn(errgo.NoteMask(err, "Could not make call to Client.OutWrite"))
	}
	resp := <-respChan
	if resp == nil {
		return c.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return c.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	return nil

}

// PushKeys waiting for documentation from Neovim
func (c *Client) PushKeys(str string) error {
	return c.PushKeysRaw([]byte(str))
}

// PushKeysRaw waiting for documentation from Neovim
func (c *Client) PushKeysRaw(str []byte) error {

	enc := func() (_err error) {
		_err = c.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = c.enc.EncodeBytes(str)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_, _err = c.dec.DecodeBytes()

		return
	}
	respChan, err := c.makeCall(clientPushKeys, enc, dec)
	if err != nil {
		return c.panicOrReturn(errgo.NoteMask(err, "Could not make call to Client.PushKeys"))
	}
	resp := <-respChan
	if resp == nil {
		return c.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return c.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	return nil

}

// registerProvider waiting for documentation from Neovim
func (c *Client) registerProvider(feature string) error {
	return c.registerProviderRaw([]byte(feature))
}

// registerProviderRaw waiting for documentation from Neovim
func (c *Client) registerProviderRaw(feature []byte) error {

	enc := func() (_err error) {
		_err = c.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = c.enc.EncodeBytes(feature)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_, _err = c.dec.DecodeBytes()

		return
	}
	respChan, err := c.makeCall(clientregisterProvider, enc, dec)
	if err != nil {
		return c.panicOrReturn(errgo.NoteMask(err, "Could not make call to Client.registerProvider"))
	}
	resp := <-respChan
	if resp == nil {
		return c.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return c.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	return nil

}

// ReplaceTermcodes waiting for documentation from Neovim
func (c *Client) ReplaceTermcodes(str string, fromPart bool, doLt bool, special bool) (string, error) {
	res, err := c.ReplaceTermcodesRaw([]byte(str), fromPart, doLt, special)
	return string(res), err
}

// ReplaceTermcodesRaw waiting for documentation from Neovim
func (c *Client) ReplaceTermcodesRaw(str []byte, fromPart bool, doLt bool, special bool) ([]byte, error) {
	var retVal []byte
	enc := func() (_err error) {
		_err = c.enc.EncodeSliceLen(4)
		if _err != nil {
			return
		}

		_err = c.enc.EncodeBytes(str)

		if _err != nil {
			return
		}

		_err = c.enc.EncodeBool(fromPart)

		if _err != nil {
			return
		}

		_err = c.enc.EncodeBool(doLt)

		if _err != nil {
			return
		}

		_err = c.enc.EncodeBool(special)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = c.dec.DecodeBytes()

		return
	}
	respChan, err := c.makeCall(clientReplaceTermcodes, enc, dec)
	if err != nil {
		return retVal, c.panicOrReturn(errgo.NoteMask(err, "Could not make call to Client.ReplaceTermcodes"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, c.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, c.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.([]byte)
	return retVal, nil

}

// ReportError waiting for documentation from Neovim
func (c *Client) ReportError(str string) error {
	return c.ReportErrorRaw([]byte(str))
}

// ReportErrorRaw waiting for documentation from Neovim
func (c *Client) ReportErrorRaw(str []byte) error {

	enc := func() (_err error) {
		_err = c.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = c.enc.EncodeBytes(str)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_, _err = c.dec.DecodeBytes()

		return
	}
	respChan, err := c.makeCall(clientReportError, enc, dec)
	if err != nil {
		return c.panicOrReturn(errgo.NoteMask(err, "Could not make call to Client.ReportError"))
	}
	resp := <-respChan
	if resp == nil {
		return c.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return c.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	return nil

}

// SetCurrentBuffer waiting for documentation from Neovim
func (c *Client) SetCurrentBuffer(buffer Buffer) error {

	enc := func() (_err error) {
		_err = c.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = c.encodeBuffer(buffer)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_, _err = c.dec.DecodeBytes()

		return
	}
	respChan, err := c.makeCall(clientSetCurrentBuffer, enc, dec)
	if err != nil {
		return c.panicOrReturn(errgo.NoteMask(err, "Could not make call to Client.SetCurrentBuffer"))
	}
	resp := <-respChan
	if resp == nil {
		return c.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return c.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	return nil

}

// SetCurrentLine waiting for documentation from Neovim
func (c *Client) SetCurrentLine(line string) error {
	return c.SetCurrentLineRaw([]byte(line))
}

// SetCurrentLineRaw waiting for documentation from Neovim
func (c *Client) SetCurrentLineRaw(line []byte) error {

	enc := func() (_err error) {
		_err = c.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = c.enc.EncodeBytes(line)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_, _err = c.dec.DecodeBytes()

		return
	}
	respChan, err := c.makeCall(clientSetCurrentLine, enc, dec)
	if err != nil {
		return c.panicOrReturn(errgo.NoteMask(err, "Could not make call to Client.SetCurrentLine"))
	}
	resp := <-respChan
	if resp == nil {
		return c.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return c.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	return nil

}

// SetCurrentTabpage waiting for documentation from Neovim
func (c *Client) SetCurrentTabpage(tabpage Tabpage) error {

	enc := func() (_err error) {
		_err = c.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = c.encodeTabpage(tabpage)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_, _err = c.dec.DecodeBytes()

		return
	}
	respChan, err := c.makeCall(clientSetCurrentTabpage, enc, dec)
	if err != nil {
		return c.panicOrReturn(errgo.NoteMask(err, "Could not make call to Client.SetCurrentTabpage"))
	}
	resp := <-respChan
	if resp == nil {
		return c.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return c.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	return nil

}

// SetCurrentWindow waiting for documentation from Neovim
func (c *Client) SetCurrentWindow(window Window) error {

	enc := func() (_err error) {
		_err = c.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = c.encodeWindow(window)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_, _err = c.dec.DecodeBytes()

		return
	}
	respChan, err := c.makeCall(clientSetCurrentWindow, enc, dec)
	if err != nil {
		return c.panicOrReturn(errgo.NoteMask(err, "Could not make call to Client.SetCurrentWindow"))
	}
	resp := <-respChan
	if resp == nil {
		return c.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return c.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	return nil

}

// SetOption waiting for documentation from Neovim
func (c *Client) SetOption(name string, value interface{}) error {
	return c.SetOptionRaw([]byte(name), value)
}

// SetOptionRaw waiting for documentation from Neovim
func (c *Client) SetOptionRaw(name []byte, value interface{}) error {

	enc := func() (_err error) {
		_err = c.enc.EncodeSliceLen(2)
		if _err != nil {
			return
		}

		_err = c.enc.EncodeBytes(name)

		if _err != nil {
			return
		}

		_err = c.enc.Encode(value)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_, _err = c.dec.DecodeBytes()

		return
	}
	respChan, err := c.makeCall(clientSetOption, enc, dec)
	if err != nil {
		return c.panicOrReturn(errgo.NoteMask(err, "Could not make call to Client.SetOption"))
	}
	resp := <-respChan
	if resp == nil {
		return c.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return c.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	return nil

}

// SetVar waiting for documentation from Neovim
func (c *Client) SetVar(name string, value interface{}) (interface{}, error) {
	return c.SetVarRaw([]byte(name), value)
}

// SetVarRaw waiting for documentation from Neovim
func (c *Client) SetVarRaw(name []byte, value interface{}) (interface{}, error) {
	var retVal interface{}
	enc := func() (_err error) {
		_err = c.enc.EncodeSliceLen(2)
		if _err != nil {
			return
		}

		_err = c.enc.EncodeBytes(name)

		if _err != nil {
			return
		}

		_err = c.enc.Encode(value)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = c.dec.DecodeInterface()

		return
	}
	respChan, err := c.makeCall(clientSetVar, enc, dec)
	if err != nil {
		return retVal, c.panicOrReturn(errgo.NoteMask(err, "Could not make call to Client.SetVar"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, c.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, c.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(interface{})
	return retVal, nil

}

// Strwidth waiting for documentation from Neovim
func (c *Client) Strwidth(str string) (int, error) {
	return c.StrwidthRaw([]byte(str))
}

// StrwidthRaw waiting for documentation from Neovim
func (c *Client) StrwidthRaw(str []byte) (int, error) {
	var retVal int
	enc := func() (_err error) {
		_err = c.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = c.enc.EncodeBytes(str)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = c.dec.DecodeInt()

		return
	}
	respChan, err := c.makeCall(clientStrwidth, enc, dec)
	if err != nil {
		return retVal, c.panicOrReturn(errgo.NoteMask(err, "Could not make call to Client.Strwidth"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, c.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, c.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(int)
	return retVal, nil

}

// subscribe waiting for documentation from Neovim
func (c *Client) subscribe(event string) error {
	return c.subscribeRaw([]byte(event))
}

// subscribeRaw waiting for documentation from Neovim
func (c *Client) subscribeRaw(event []byte) error {

	enc := func() (_err error) {
		_err = c.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = c.enc.EncodeBytes(event)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_, _err = c.dec.DecodeBytes()

		return
	}
	respChan, err := c.makeCall(clientsubscribe, enc, dec)
	if err != nil {
		return c.panicOrReturn(errgo.NoteMask(err, "Could not make call to Client.subscribe"))
	}
	resp := <-respChan
	if resp == nil {
		return c.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return c.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	return nil

}

// unsubscribe waiting for documentation from Neovim
func (c *Client) unsubscribe(event string) error {
	return c.unsubscribeRaw([]byte(event))
}

// unsubscribeRaw waiting for documentation from Neovim
func (c *Client) unsubscribeRaw(event []byte) error {

	enc := func() (_err error) {
		_err = c.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = c.enc.EncodeBytes(event)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_, _err = c.dec.DecodeBytes()

		return
	}
	respChan, err := c.makeCall(clientunsubscribe, enc, dec)
	if err != nil {
		return c.panicOrReturn(errgo.NoteMask(err, "Could not make call to Client.unsubscribe"))
	}
	resp := <-respChan
	if resp == nil {
		return c.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return c.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	return nil

}

// GetBuffer waiting for documentation from Neovim
func (w *Window) GetBuffer() (Buffer, error) {
	var retVal Buffer
	enc := func() (_err error) {
		_err = w.client.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = w.client.encodeWindow(*w)
		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = w.client.decodeBuffer()

		return
	}
	respChan, err := w.client.makeCall(windowGetBuffer, enc, dec)
	if err != nil {
		return retVal, w.client.panicOrReturn(errgo.NoteMask(err, "Could not make call to Window.GetBuffer"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, w.client.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, w.client.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(Buffer)
	return retVal, nil

}

// GetCursor waiting for documentation from Neovim
func (w *Window) GetCursor() ([]int, error) {
	var retVal []int
	enc := func() (_err error) {
		_err = w.client.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = w.client.encodeWindow(*w)
		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = w.client.decodeIntSlice()

		return
	}
	respChan, err := w.client.makeCall(windowGetCursor, enc, dec)
	if err != nil {
		return retVal, w.client.panicOrReturn(errgo.NoteMask(err, "Could not make call to Window.GetCursor"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, w.client.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, w.client.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.([]int)
	return retVal, nil

}

// GetHeight waiting for documentation from Neovim
func (w *Window) GetHeight() (int, error) {
	var retVal int
	enc := func() (_err error) {
		_err = w.client.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = w.client.encodeWindow(*w)
		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = w.client.dec.DecodeInt()

		return
	}
	respChan, err := w.client.makeCall(windowGetHeight, enc, dec)
	if err != nil {
		return retVal, w.client.panicOrReturn(errgo.NoteMask(err, "Could not make call to Window.GetHeight"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, w.client.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, w.client.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(int)
	return retVal, nil

}

// GetOption waiting for documentation from Neovim
func (w *Window) GetOption(name string) (interface{}, error) {
	return w.GetOptionRaw([]byte(name))
}

// GetOptionRaw waiting for documentation from Neovim
func (w *Window) GetOptionRaw(name []byte) (interface{}, error) {
	var retVal interface{}
	enc := func() (_err error) {
		_err = w.client.enc.EncodeSliceLen(2)
		if _err != nil {
			return
		}

		_err = w.client.encodeWindow(*w)
		if _err != nil {
			return
		}

		_err = w.client.enc.EncodeBytes(name)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = w.client.dec.DecodeInterface()

		return
	}
	respChan, err := w.client.makeCall(windowGetOption, enc, dec)
	if err != nil {
		return retVal, w.client.panicOrReturn(errgo.NoteMask(err, "Could not make call to Window.GetOption"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, w.client.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, w.client.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(interface{})
	return retVal, nil

}

// GetPosition waiting for documentation from Neovim
func (w *Window) GetPosition() ([]int, error) {
	var retVal []int
	enc := func() (_err error) {
		_err = w.client.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = w.client.encodeWindow(*w)
		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = w.client.decodeIntSlice()

		return
	}
	respChan, err := w.client.makeCall(windowGetPosition, enc, dec)
	if err != nil {
		return retVal, w.client.panicOrReturn(errgo.NoteMask(err, "Could not make call to Window.GetPosition"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, w.client.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, w.client.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.([]int)
	return retVal, nil

}

// GetTabpage waiting for documentation from Neovim
func (w *Window) GetTabpage() (Tabpage, error) {
	var retVal Tabpage
	enc := func() (_err error) {
		_err = w.client.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = w.client.encodeWindow(*w)
		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = w.client.decodeTabpage()

		return
	}
	respChan, err := w.client.makeCall(windowGetTabpage, enc, dec)
	if err != nil {
		return retVal, w.client.panicOrReturn(errgo.NoteMask(err, "Could not make call to Window.GetTabpage"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, w.client.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, w.client.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(Tabpage)
	return retVal, nil

}

// GetVar waiting for documentation from Neovim
func (w *Window) GetVar(name string) (interface{}, error) {
	return w.GetVarRaw([]byte(name))
}

// GetVarRaw waiting for documentation from Neovim
func (w *Window) GetVarRaw(name []byte) (interface{}, error) {
	var retVal interface{}
	enc := func() (_err error) {
		_err = w.client.enc.EncodeSliceLen(2)
		if _err != nil {
			return
		}

		_err = w.client.encodeWindow(*w)
		if _err != nil {
			return
		}

		_err = w.client.enc.EncodeBytes(name)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = w.client.dec.DecodeInterface()

		return
	}
	respChan, err := w.client.makeCall(windowGetVar, enc, dec)
	if err != nil {
		return retVal, w.client.panicOrReturn(errgo.NoteMask(err, "Could not make call to Window.GetVar"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, w.client.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, w.client.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(interface{})
	return retVal, nil

}

// GetWidth waiting for documentation from Neovim
func (w *Window) GetWidth() (int, error) {
	var retVal int
	enc := func() (_err error) {
		_err = w.client.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = w.client.encodeWindow(*w)
		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = w.client.dec.DecodeInt()

		return
	}
	respChan, err := w.client.makeCall(windowGetWidth, enc, dec)
	if err != nil {
		return retVal, w.client.panicOrReturn(errgo.NoteMask(err, "Could not make call to Window.GetWidth"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, w.client.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, w.client.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(int)
	return retVal, nil

}

// IsValid waiting for documentation from Neovim
func (w *Window) IsValid() (bool, error) {
	var retVal bool
	enc := func() (_err error) {
		_err = w.client.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = w.client.encodeWindow(*w)
		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = w.client.dec.DecodeBool()

		return
	}
	respChan, err := w.client.makeCall(windowIsValid, enc, dec)
	if err != nil {
		return retVal, w.client.panicOrReturn(errgo.NoteMask(err, "Could not make call to Window.IsValid"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, w.client.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, w.client.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(bool)
	return retVal, nil

}

// SetCursor waiting for documentation from Neovim
func (w *Window) SetCursor(pos []int) error {

	enc := func() (_err error) {
		_err = w.client.enc.EncodeSliceLen(2)
		if _err != nil {
			return
		}

		_err = w.client.encodeWindow(*w)
		if _err != nil {
			return
		}

		_err = w.client.encodeIntSlice(pos)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_, _err = w.client.dec.DecodeBytes()

		return
	}
	respChan, err := w.client.makeCall(windowSetCursor, enc, dec)
	if err != nil {
		return w.client.panicOrReturn(errgo.NoteMask(err, "Could not make call to Window.SetCursor"))
	}
	resp := <-respChan
	if resp == nil {
		return w.client.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return w.client.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	return nil

}

// SetHeight waiting for documentation from Neovim
func (w *Window) SetHeight(height int) error {

	enc := func() (_err error) {
		_err = w.client.enc.EncodeSliceLen(2)
		if _err != nil {
			return
		}

		_err = w.client.encodeWindow(*w)
		if _err != nil {
			return
		}

		_err = w.client.enc.EncodeInt(height)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_, _err = w.client.dec.DecodeBytes()

		return
	}
	respChan, err := w.client.makeCall(windowSetHeight, enc, dec)
	if err != nil {
		return w.client.panicOrReturn(errgo.NoteMask(err, "Could not make call to Window.SetHeight"))
	}
	resp := <-respChan
	if resp == nil {
		return w.client.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return w.client.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	return nil

}

// SetOption waiting for documentation from Neovim
func (w *Window) SetOption(name string, value interface{}) error {
	return w.SetOptionRaw([]byte(name), value)
}

// SetOptionRaw waiting for documentation from Neovim
func (w *Window) SetOptionRaw(name []byte, value interface{}) error {

	enc := func() (_err error) {
		_err = w.client.enc.EncodeSliceLen(3)
		if _err != nil {
			return
		}

		_err = w.client.encodeWindow(*w)
		if _err != nil {
			return
		}

		_err = w.client.enc.EncodeBytes(name)

		if _err != nil {
			return
		}

		_err = w.client.enc.Encode(value)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_, _err = w.client.dec.DecodeBytes()

		return
	}
	respChan, err := w.client.makeCall(windowSetOption, enc, dec)
	if err != nil {
		return w.client.panicOrReturn(errgo.NoteMask(err, "Could not make call to Window.SetOption"))
	}
	resp := <-respChan
	if resp == nil {
		return w.client.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return w.client.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	return nil

}

// SetVar waiting for documentation from Neovim
func (w *Window) SetVar(name string, value interface{}) (interface{}, error) {
	return w.SetVarRaw([]byte(name), value)
}

// SetVarRaw waiting for documentation from Neovim
func (w *Window) SetVarRaw(name []byte, value interface{}) (interface{}, error) {
	var retVal interface{}
	enc := func() (_err error) {
		_err = w.client.enc.EncodeSliceLen(3)
		if _err != nil {
			return
		}

		_err = w.client.encodeWindow(*w)
		if _err != nil {
			return
		}

		_err = w.client.enc.EncodeBytes(name)

		if _err != nil {
			return
		}

		_err = w.client.enc.Encode(value)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = w.client.dec.DecodeInterface()

		return
	}
	respChan, err := w.client.makeCall(windowSetVar, enc, dec)
	if err != nil {
		return retVal, w.client.panicOrReturn(errgo.NoteMask(err, "Could not make call to Window.SetVar"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, w.client.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, w.client.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(interface{})
	return retVal, nil

}

// SetWidth waiting for documentation from Neovim
func (w *Window) SetWidth(width int) error {

	enc := func() (_err error) {
		_err = w.client.enc.EncodeSliceLen(2)
		if _err != nil {
			return
		}

		_err = w.client.encodeWindow(*w)
		if _err != nil {
			return
		}

		_err = w.client.enc.EncodeInt(width)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_, _err = w.client.dec.DecodeBytes()

		return
	}
	respChan, err := w.client.makeCall(windowSetWidth, enc, dec)
	if err != nil {
		return w.client.panicOrReturn(errgo.NoteMask(err, "Could not make call to Window.SetWidth"))
	}
	resp := <-respChan
	if resp == nil {
		return w.client.panicOrReturn(errgo.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return w.client.panicOrReturn(errgo.NoteMask(err, "We got a non-nil error in our response"))
	}

	return nil

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

func (c *Client) encodeStringSlice(s [][]byte) error {
	err := c.enc.EncodeSliceLen(len(s))
	if err != nil {
		return errgo.NoteMask(err, "Could not encode slice length")
	}

	for i := 0; i < len(s); i++ {

		err := c.enc.EncodeBytes(s[i])

		if err != nil {
			return errgo.Notef(err, "Could not encode []byte at index %v", i)
		}
	}

	return nil
}

func (c *Client) decodeStringSlice() ([][]byte, error) {
	l, err := c.dec.DecodeSliceLen()
	if err != nil {
		return nil, errgo.NoteMask(err, "Could not decode slice length")
	}

	res := make([][]byte, l)

	for i := 0; i < l; i++ {

		b, err := c.dec.DecodeBytes()

		if err != nil {
			return nil, errgo.Notef(err, "Could not decode []byte at index %v", i)
		}
		res[i] = b
	}

	return res, nil
}

func (c *Client) encodeIntSlice(s []int) error {
	err := c.enc.EncodeSliceLen(len(s))
	if err != nil {
		return errgo.NoteMask(err, "Could not encode slice length")
	}

	for i := 0; i < len(s); i++ {

		err := c.enc.EncodeInt(s[i])

		if err != nil {
			return errgo.Notef(err, "Could not encode int at index %v", i)
		}
	}

	return nil
}

func (c *Client) decodeIntSlice() ([]int, error) {
	l, err := c.dec.DecodeSliceLen()
	if err != nil {
		return nil, errgo.NoteMask(err, "Could not decode slice length")
	}

	res := make([]int, l)

	for i := 0; i < l; i++ {

		b, err := c.dec.DecodeInt()

		if err != nil {
			return nil, errgo.Notef(err, "Could not decode int at index %v", i)
		}
		res[i] = b
	}

	return res, nil
}
