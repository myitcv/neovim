package neovim

// **** THIS FILE IS GENERATED - DO NOT EDIT BY HAND

import "github.com/juju/errors"

// constants representing method ids

const (
	typeBuffer  int8 = 0
	typeWindow  int8 = 1
	typeTabpage int8 = 2
)

const (
	bufferDelLine           = "buffer_del_line"
	bufferGetLine           = "buffer_get_line"
	bufferGetLineSlice      = "buffer_get_line_slice"
	bufferGetMark           = "buffer_get_mark"
	bufferGetName           = "buffer_get_name"
	bufferGetNumber         = "buffer_get_number"
	bufferGetOption         = "buffer_get_option"
	bufferGetVar            = "buffer_get_var"
	bufferInsert            = "buffer_insert"
	bufferIsValid           = "buffer_is_valid"
	bufferLineCount         = "buffer_line_count"
	bufferSetLine           = "buffer_set_line"
	bufferSetLineSlice      = "buffer_set_line_slice"
	bufferSetName           = "buffer_set_name"
	bufferSetOption         = "buffer_set_option"
	bufferSetVar            = "buffer_set_var"
	tabpageGetVar           = "tabpage_get_var"
	tabpageGetWindow        = "tabpage_get_window"
	tabpageGetWindows       = "tabpage_get_windows"
	tabpageIsValid          = "tabpage_is_valid"
	tabpageSetVar           = "tabpage_set_var"
	clientCallFunction      = "vim_call_function"
	clientChangeDirectory   = "vim_change_directory"
	clientCommand           = "vim_command"
	clientCommandOutput     = "vim_command_output"
	clientDelCurrentLine    = "vim_del_current_line"
	clientErrWrite          = "vim_err_write"
	clientEval              = "vim_eval"
	clientFeedkeys          = "vim_feedkeys"
	clientGetApiInfo        = "vim_get_api_info"
	clientGetBuffers        = "vim_get_buffers"
	clientGetColorMap       = "vim_get_color_map"
	clientGetCurrentBuffer  = "vim_get_current_buffer"
	clientGetCurrentLine    = "vim_get_current_line"
	clientGetCurrentTabpage = "vim_get_current_tabpage"
	clientGetCurrentWindow  = "vim_get_current_window"
	clientGetOption         = "vim_get_option"
	clientGetTabpages       = "vim_get_tabpages"
	clientGetVar            = "vim_get_var"
	clientGetVvar           = "vim_get_vvar"
	clientGetWindows        = "vim_get_windows"
	clientInput             = "vim_input"
	clientListRuntimePaths  = "vim_list_runtime_paths"
	clientNameToColor       = "vim_name_to_color"
	clientOutWrite          = "vim_out_write"
	clientReplaceTermcodes  = "vim_replace_termcodes"
	clientReportError       = "vim_report_error"
	clientSetCurrentBuffer  = "vim_set_current_buffer"
	clientSetCurrentLine    = "vim_set_current_line"
	clientSetCurrentTabpage = "vim_set_current_tabpage"
	clientSetCurrentWindow  = "vim_set_current_window"
	clientSetOption         = "vim_set_option"
	clientSetVar            = "vim_set_var"
	clientStrwidth          = "vim_strwidth"
	clientSubscribe         = "vim_subscribe"
	clientUnsubscribe       = "vim_unsubscribe"
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
		_err = b.client.enc.WriteArrayHeader(2)
		if _err != nil {
			return
		}

		_err = b.client.encodeBuffer(*b)
		if _err != nil {
			return
		}

		_err = b.client.enc.WriteInt(index)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_err = b.client.dec.ReadNil()

		return
	}
	respChan, err := b.client.makeCall(bufferDelLine, enc, dec)
	if err != nil {
		return b.client.panicOrReturn(errors.Annotate(err, "Could not make call to Buffer.DelLine"))
	}
	resp := <-respChan
	if resp == nil {
		return b.client.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return b.client.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	return nil

}

// GetLine waiting for documentation from Neovim
func (b *Buffer) GetLine(index int) (string, error) {
	var retVal string
	enc := func() (_err error) {
		_err = b.client.enc.WriteArrayHeader(2)
		if _err != nil {
			return
		}

		_err = b.client.encodeBuffer(*b)
		if _err != nil {
			return
		}

		_err = b.client.enc.WriteInt(index)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = b.client.dec.ReadString()

		return
	}
	respChan, err := b.client.makeCall(bufferGetLine, enc, dec)
	if err != nil {
		return retVal, b.client.panicOrReturn(errors.Annotate(err, "Could not make call to Buffer.GetLine"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, b.client.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, b.client.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(string)
	return retVal, nil

}

// GetLineSlice waiting for documentation from Neovim
func (b *Buffer) GetLineSlice(start int, end int, includeStart bool, includeEnd bool) ([]string, error) {
	var retVal []string
	enc := func() (_err error) {
		_err = b.client.enc.WriteArrayHeader(5)
		if _err != nil {
			return
		}

		_err = b.client.encodeBuffer(*b)
		if _err != nil {
			return
		}

		_err = b.client.enc.WriteInt(start)

		if _err != nil {
			return
		}

		_err = b.client.enc.WriteInt(end)

		if _err != nil {
			return
		}

		_err = b.client.enc.WriteBool(includeStart)

		if _err != nil {
			return
		}

		_err = b.client.enc.WriteBool(includeEnd)

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
		return retVal, b.client.panicOrReturn(errors.Annotate(err, "Could not make call to Buffer.GetLineSlice"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, b.client.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, b.client.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.([]string)
	return retVal, nil

}

// GetMark waiting for documentation from Neovim
func (b *Buffer) GetMark(name string) ([]int, error) {
	var retVal []int
	enc := func() (_err error) {
		_err = b.client.enc.WriteArrayHeader(2)
		if _err != nil {
			return
		}

		_err = b.client.encodeBuffer(*b)
		if _err != nil {
			return
		}

		_err = b.client.enc.WriteString(name)

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
		return retVal, b.client.panicOrReturn(errors.Annotate(err, "Could not make call to Buffer.GetMark"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, b.client.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, b.client.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.([]int)
	return retVal, nil

}

// GetName waiting for documentation from Neovim
func (b *Buffer) GetName() (string, error) {
	var retVal string
	enc := func() (_err error) {
		_err = b.client.enc.WriteArrayHeader(1)
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

		_i, _err = b.client.dec.ReadString()

		return
	}
	respChan, err := b.client.makeCall(bufferGetName, enc, dec)
	if err != nil {
		return retVal, b.client.panicOrReturn(errors.Annotate(err, "Could not make call to Buffer.GetName"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, b.client.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, b.client.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(string)
	return retVal, nil

}

// GetNumber waiting for documentation from Neovim
func (b *Buffer) GetNumber() (int, error) {
	var retVal int
	enc := func() (_err error) {
		_err = b.client.enc.WriteArrayHeader(1)
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

		_i, _err = b.client.dec.ReadInt()

		return
	}
	respChan, err := b.client.makeCall(bufferGetNumber, enc, dec)
	if err != nil {
		return retVal, b.client.panicOrReturn(errors.Annotate(err, "Could not make call to Buffer.GetNumber"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, b.client.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, b.client.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(int)
	return retVal, nil

}

// GetOption waiting for documentation from Neovim
func (b *Buffer) GetOption(name string) (interface{}, error) {
	var retVal interface{}
	enc := func() (_err error) {
		_err = b.client.enc.WriteArrayHeader(2)
		if _err != nil {
			return
		}

		_err = b.client.encodeBuffer(*b)
		if _err != nil {
			return
		}

		_err = b.client.enc.WriteString(name)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = b.client.dec.ReadIntf()

		return
	}
	respChan, err := b.client.makeCall(bufferGetOption, enc, dec)
	if err != nil {
		return retVal, b.client.panicOrReturn(errors.Annotate(err, "Could not make call to Buffer.GetOption"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, b.client.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, b.client.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(interface{})
	return retVal, nil

}

// GetVar waiting for documentation from Neovim
func (b *Buffer) GetVar(name string) (interface{}, error) {
	var retVal interface{}
	enc := func() (_err error) {
		_err = b.client.enc.WriteArrayHeader(2)
		if _err != nil {
			return
		}

		_err = b.client.encodeBuffer(*b)
		if _err != nil {
			return
		}

		_err = b.client.enc.WriteString(name)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = b.client.dec.ReadIntf()

		return
	}
	respChan, err := b.client.makeCall(bufferGetVar, enc, dec)
	if err != nil {
		return retVal, b.client.panicOrReturn(errors.Annotate(err, "Could not make call to Buffer.GetVar"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, b.client.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, b.client.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(interface{})
	return retVal, nil

}

// Insert waiting for documentation from Neovim
func (b *Buffer) Insert(lnum int, lines []string) error {

	enc := func() (_err error) {
		_err = b.client.enc.WriteArrayHeader(3)
		if _err != nil {
			return
		}

		_err = b.client.encodeBuffer(*b)
		if _err != nil {
			return
		}

		_err = b.client.enc.WriteInt(lnum)

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

		_err = b.client.dec.ReadNil()

		return
	}
	respChan, err := b.client.makeCall(bufferInsert, enc, dec)
	if err != nil {
		return b.client.panicOrReturn(errors.Annotate(err, "Could not make call to Buffer.Insert"))
	}
	resp := <-respChan
	if resp == nil {
		return b.client.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return b.client.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	return nil

}

// IsValid waiting for documentation from Neovim
func (b *Buffer) IsValid() (bool, error) {
	var retVal bool
	enc := func() (_err error) {
		_err = b.client.enc.WriteArrayHeader(1)
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

		_i, _err = b.client.dec.ReadBool()

		return
	}
	respChan, err := b.client.makeCall(bufferIsValid, enc, dec)
	if err != nil {
		return retVal, b.client.panicOrReturn(errors.Annotate(err, "Could not make call to Buffer.IsValid"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, b.client.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, b.client.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(bool)
	return retVal, nil

}

// LineCount waiting for documentation from Neovim
func (b *Buffer) LineCount() (int, error) {
	var retVal int
	enc := func() (_err error) {
		_err = b.client.enc.WriteArrayHeader(1)
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

		_i, _err = b.client.dec.ReadInt()

		return
	}
	respChan, err := b.client.makeCall(bufferLineCount, enc, dec)
	if err != nil {
		return retVal, b.client.panicOrReturn(errors.Annotate(err, "Could not make call to Buffer.LineCount"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, b.client.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, b.client.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(int)
	return retVal, nil

}

// SetLine waiting for documentation from Neovim
func (b *Buffer) SetLine(index int, line string) error {

	enc := func() (_err error) {
		_err = b.client.enc.WriteArrayHeader(3)
		if _err != nil {
			return
		}

		_err = b.client.encodeBuffer(*b)
		if _err != nil {
			return
		}

		_err = b.client.enc.WriteInt(index)

		if _err != nil {
			return
		}

		_err = b.client.enc.WriteString(line)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_err = b.client.dec.ReadNil()

		return
	}
	respChan, err := b.client.makeCall(bufferSetLine, enc, dec)
	if err != nil {
		return b.client.panicOrReturn(errors.Annotate(err, "Could not make call to Buffer.SetLine"))
	}
	resp := <-respChan
	if resp == nil {
		return b.client.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return b.client.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	return nil

}

// SetLineSlice waiting for documentation from Neovim
func (b *Buffer) SetLineSlice(start int, end int, includeStart bool, includeEnd bool, replacement []string) error {

	enc := func() (_err error) {
		_err = b.client.enc.WriteArrayHeader(6)
		if _err != nil {
			return
		}

		_err = b.client.encodeBuffer(*b)
		if _err != nil {
			return
		}

		_err = b.client.enc.WriteInt(start)

		if _err != nil {
			return
		}

		_err = b.client.enc.WriteInt(end)

		if _err != nil {
			return
		}

		_err = b.client.enc.WriteBool(includeStart)

		if _err != nil {
			return
		}

		_err = b.client.enc.WriteBool(includeEnd)

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

		_err = b.client.dec.ReadNil()

		return
	}
	respChan, err := b.client.makeCall(bufferSetLineSlice, enc, dec)
	if err != nil {
		return b.client.panicOrReturn(errors.Annotate(err, "Could not make call to Buffer.SetLineSlice"))
	}
	resp := <-respChan
	if resp == nil {
		return b.client.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return b.client.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	return nil

}

// SetName waiting for documentation from Neovim
func (b *Buffer) SetName(name string) error {

	enc := func() (_err error) {
		_err = b.client.enc.WriteArrayHeader(2)
		if _err != nil {
			return
		}

		_err = b.client.encodeBuffer(*b)
		if _err != nil {
			return
		}

		_err = b.client.enc.WriteString(name)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_err = b.client.dec.ReadNil()

		return
	}
	respChan, err := b.client.makeCall(bufferSetName, enc, dec)
	if err != nil {
		return b.client.panicOrReturn(errors.Annotate(err, "Could not make call to Buffer.SetName"))
	}
	resp := <-respChan
	if resp == nil {
		return b.client.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return b.client.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	return nil

}

// SetOption waiting for documentation from Neovim
func (b *Buffer) SetOption(name string, value interface{}) error {

	enc := func() (_err error) {
		_err = b.client.enc.WriteArrayHeader(3)
		if _err != nil {
			return
		}

		_err = b.client.encodeBuffer(*b)
		if _err != nil {
			return
		}

		_err = b.client.enc.WriteString(name)

		if _err != nil {
			return
		}

		_err = b.client.enc.WriteIntf(value)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_err = b.client.dec.ReadNil()

		return
	}
	respChan, err := b.client.makeCall(bufferSetOption, enc, dec)
	if err != nil {
		return b.client.panicOrReturn(errors.Annotate(err, "Could not make call to Buffer.SetOption"))
	}
	resp := <-respChan
	if resp == nil {
		return b.client.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return b.client.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	return nil

}

// SetVar waiting for documentation from Neovim
func (b *Buffer) SetVar(name string, value interface{}) (interface{}, error) {
	var retVal interface{}
	enc := func() (_err error) {
		_err = b.client.enc.WriteArrayHeader(3)
		if _err != nil {
			return
		}

		_err = b.client.encodeBuffer(*b)
		if _err != nil {
			return
		}

		_err = b.client.enc.WriteString(name)

		if _err != nil {
			return
		}

		_err = b.client.enc.WriteIntf(value)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = b.client.dec.ReadIntf()

		return
	}
	respChan, err := b.client.makeCall(bufferSetVar, enc, dec)
	if err != nil {
		return retVal, b.client.panicOrReturn(errors.Annotate(err, "Could not make call to Buffer.SetVar"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, b.client.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, b.client.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(interface{})
	return retVal, nil

}

// GetVar waiting for documentation from Neovim
func (t *Tabpage) GetVar(name string) (interface{}, error) {
	var retVal interface{}
	enc := func() (_err error) {
		_err = t.client.enc.WriteArrayHeader(2)
		if _err != nil {
			return
		}

		_err = t.client.encodeTabpage(*t)
		if _err != nil {
			return
		}

		_err = t.client.enc.WriteString(name)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = t.client.dec.ReadIntf()

		return
	}
	respChan, err := t.client.makeCall(tabpageGetVar, enc, dec)
	if err != nil {
		return retVal, t.client.panicOrReturn(errors.Annotate(err, "Could not make call to Tabpage.GetVar"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, t.client.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, t.client.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(interface{})
	return retVal, nil

}

// GetWindow waiting for documentation from Neovim
func (t *Tabpage) GetWindow() (Window, error) {
	var retVal Window
	enc := func() (_err error) {
		_err = t.client.enc.WriteArrayHeader(1)
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
		return retVal, t.client.panicOrReturn(errors.Annotate(err, "Could not make call to Tabpage.GetWindow"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, t.client.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, t.client.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(Window)
	return retVal, nil

}

// GetWindows waiting for documentation from Neovim
func (t *Tabpage) GetWindows() ([]Window, error) {
	var retVal []Window
	enc := func() (_err error) {
		_err = t.client.enc.WriteArrayHeader(1)
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
		return retVal, t.client.panicOrReturn(errors.Annotate(err, "Could not make call to Tabpage.GetWindows"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, t.client.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, t.client.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.([]Window)
	return retVal, nil

}

// IsValid waiting for documentation from Neovim
func (t *Tabpage) IsValid() (bool, error) {
	var retVal bool
	enc := func() (_err error) {
		_err = t.client.enc.WriteArrayHeader(1)
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

		_i, _err = t.client.dec.ReadBool()

		return
	}
	respChan, err := t.client.makeCall(tabpageIsValid, enc, dec)
	if err != nil {
		return retVal, t.client.panicOrReturn(errors.Annotate(err, "Could not make call to Tabpage.IsValid"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, t.client.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, t.client.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(bool)
	return retVal, nil

}

// SetVar waiting for documentation from Neovim
func (t *Tabpage) SetVar(name string, value interface{}) (interface{}, error) {
	var retVal interface{}
	enc := func() (_err error) {
		_err = t.client.enc.WriteArrayHeader(3)
		if _err != nil {
			return
		}

		_err = t.client.encodeTabpage(*t)
		if _err != nil {
			return
		}

		_err = t.client.enc.WriteString(name)

		if _err != nil {
			return
		}

		_err = t.client.enc.WriteIntf(value)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = t.client.dec.ReadIntf()

		return
	}
	respChan, err := t.client.makeCall(tabpageSetVar, enc, dec)
	if err != nil {
		return retVal, t.client.panicOrReturn(errors.Annotate(err, "Could not make call to Tabpage.SetVar"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, t.client.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, t.client.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(interface{})
	return retVal, nil

}

// CallFunction waiting for documentation from Neovim
func (c *Client) CallFunction(fname string, args []interface{}) (interface{}, error) {
	var retVal interface{}
	enc := func() (_err error) {
		_err = c.enc.WriteArrayHeader(2)
		if _err != nil {
			return
		}

		_err = c.enc.WriteString(fname)

		if _err != nil {
			return
		}

		_err = c.enc.WriteIntf(args)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = c.dec.ReadIntf()

		return
	}
	respChan, err := c.makeCall(clientCallFunction, enc, dec)
	if err != nil {
		return retVal, c.panicOrReturn(errors.Annotate(err, "Could not make call to Client.CallFunction"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, c.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, c.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(interface{})
	return retVal, nil

}

// ChangeDirectory waiting for documentation from Neovim
func (c *Client) ChangeDirectory(dir string) error {

	enc := func() (_err error) {
		_err = c.enc.WriteArrayHeader(1)
		if _err != nil {
			return
		}

		_err = c.enc.WriteString(dir)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_err = c.dec.ReadNil()

		return
	}
	respChan, err := c.makeCall(clientChangeDirectory, enc, dec)
	if err != nil {
		return c.panicOrReturn(errors.Annotate(err, "Could not make call to Client.ChangeDirectory"))
	}
	resp := <-respChan
	if resp == nil {
		return c.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return c.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	return nil

}

// Command waiting for documentation from Neovim
func (c *Client) Command(str string) error {

	enc := func() (_err error) {
		_err = c.enc.WriteArrayHeader(1)
		if _err != nil {
			return
		}

		_err = c.enc.WriteString(str)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_err = c.dec.ReadNil()

		return
	}
	respChan, err := c.makeCall(clientCommand, enc, dec)
	if err != nil {
		return c.panicOrReturn(errors.Annotate(err, "Could not make call to Client.Command"))
	}
	resp := <-respChan
	if resp == nil {
		return c.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return c.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	return nil

}

// CommandOutput waiting for documentation from Neovim
func (c *Client) CommandOutput(str string) (string, error) {
	var retVal string
	enc := func() (_err error) {
		_err = c.enc.WriteArrayHeader(1)
		if _err != nil {
			return
		}

		_err = c.enc.WriteString(str)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = c.dec.ReadString()

		return
	}
	respChan, err := c.makeCall(clientCommandOutput, enc, dec)
	if err != nil {
		return retVal, c.panicOrReturn(errors.Annotate(err, "Could not make call to Client.CommandOutput"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, c.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, c.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(string)
	return retVal, nil

}

// DelCurrentLine waiting for documentation from Neovim
func (c *Client) DelCurrentLine() error {

	enc := func() (_err error) {
		_err = c.enc.WriteArrayHeader(0)
		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_err = c.dec.ReadNil()

		return
	}
	respChan, err := c.makeCall(clientDelCurrentLine, enc, dec)
	if err != nil {
		return c.panicOrReturn(errors.Annotate(err, "Could not make call to Client.DelCurrentLine"))
	}
	resp := <-respChan
	if resp == nil {
		return c.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return c.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	return nil

}

// ErrWrite waiting for documentation from Neovim
func (c *Client) ErrWrite(str string) error {

	enc := func() (_err error) {
		_err = c.enc.WriteArrayHeader(1)
		if _err != nil {
			return
		}

		_err = c.enc.WriteString(str)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_err = c.dec.ReadNil()

		return
	}
	respChan, err := c.makeCall(clientErrWrite, enc, dec)
	if err != nil {
		return c.panicOrReturn(errors.Annotate(err, "Could not make call to Client.ErrWrite"))
	}
	resp := <-respChan
	if resp == nil {
		return c.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return c.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	return nil

}

// Eval waiting for documentation from Neovim
func (c *Client) Eval(str string) (interface{}, error) {
	var retVal interface{}
	enc := func() (_err error) {
		_err = c.enc.WriteArrayHeader(1)
		if _err != nil {
			return
		}

		_err = c.enc.WriteString(str)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = c.dec.ReadIntf()

		return
	}
	respChan, err := c.makeCall(clientEval, enc, dec)
	if err != nil {
		return retVal, c.panicOrReturn(errors.Annotate(err, "Could not make call to Client.Eval"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, c.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, c.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(interface{})
	return retVal, nil

}

// Feedkeys waiting for documentation from Neovim
func (c *Client) Feedkeys(keys string, mode string, escapeCsi bool) error {

	enc := func() (_err error) {
		_err = c.enc.WriteArrayHeader(3)
		if _err != nil {
			return
		}

		_err = c.enc.WriteString(keys)

		if _err != nil {
			return
		}

		_err = c.enc.WriteString(mode)

		if _err != nil {
			return
		}

		_err = c.enc.WriteBool(escapeCsi)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_err = c.dec.ReadNil()

		return
	}
	respChan, err := c.makeCall(clientFeedkeys, enc, dec)
	if err != nil {
		return c.panicOrReturn(errors.Annotate(err, "Could not make call to Client.Feedkeys"))
	}
	resp := <-respChan
	if resp == nil {
		return c.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return c.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	return nil

}

// GetApiInfo waiting for documentation from Neovim
func (c *Client) GetApiInfo() ([]interface{}, error) {
	var retVal []interface{}
	enc := func() (_err error) {
		_err = c.enc.WriteArrayHeader(0)
		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = c.dec.ReadIntf()

		return
	}
	respChan, err := c.makeCall(clientGetApiInfo, enc, dec)
	if err != nil {
		return retVal, c.panicOrReturn(errors.Annotate(err, "Could not make call to Client.GetApiInfo"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, c.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, c.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.([]interface{})
	return retVal, nil

}

// GetBuffers waiting for documentation from Neovim
func (c *Client) GetBuffers() ([]Buffer, error) {
	var retVal []Buffer
	enc := func() (_err error) {
		_err = c.enc.WriteArrayHeader(0)
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
		return retVal, c.panicOrReturn(errors.Annotate(err, "Could not make call to Client.GetBuffers"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, c.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, c.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.([]Buffer)
	return retVal, nil

}

// GetColorMap waiting for documentation from Neovim
func (c *Client) GetColorMap() (map[string]interface{}, error) {
	var retVal map[string]interface{}
	enc := func() (_err error) {
		_err = c.enc.WriteArrayHeader(0)
		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = c.decodeDictionary()

		return
	}
	respChan, err := c.makeCall(clientGetColorMap, enc, dec)
	if err != nil {
		return retVal, c.panicOrReturn(errors.Annotate(err, "Could not make call to Client.GetColorMap"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, c.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, c.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(map[string]interface{})
	return retVal, nil

}

// GetCurrentBuffer waiting for documentation from Neovim
func (c *Client) GetCurrentBuffer() (Buffer, error) {
	var retVal Buffer
	enc := func() (_err error) {
		_err = c.enc.WriteArrayHeader(0)
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
		return retVal, c.panicOrReturn(errors.Annotate(err, "Could not make call to Client.GetCurrentBuffer"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, c.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, c.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(Buffer)
	return retVal, nil

}

// GetCurrentLine waiting for documentation from Neovim
func (c *Client) GetCurrentLine() (string, error) {
	var retVal string
	enc := func() (_err error) {
		_err = c.enc.WriteArrayHeader(0)
		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = c.dec.ReadString()

		return
	}
	respChan, err := c.makeCall(clientGetCurrentLine, enc, dec)
	if err != nil {
		return retVal, c.panicOrReturn(errors.Annotate(err, "Could not make call to Client.GetCurrentLine"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, c.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, c.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(string)
	return retVal, nil

}

// GetCurrentTabpage waiting for documentation from Neovim
func (c *Client) GetCurrentTabpage() (Tabpage, error) {
	var retVal Tabpage
	enc := func() (_err error) {
		_err = c.enc.WriteArrayHeader(0)
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
		return retVal, c.panicOrReturn(errors.Annotate(err, "Could not make call to Client.GetCurrentTabpage"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, c.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, c.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(Tabpage)
	return retVal, nil

}

// GetCurrentWindow waiting for documentation from Neovim
func (c *Client) GetCurrentWindow() (Window, error) {
	var retVal Window
	enc := func() (_err error) {
		_err = c.enc.WriteArrayHeader(0)
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
		return retVal, c.panicOrReturn(errors.Annotate(err, "Could not make call to Client.GetCurrentWindow"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, c.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, c.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(Window)
	return retVal, nil

}

// GetOption waiting for documentation from Neovim
func (c *Client) GetOption(name string) (interface{}, error) {
	var retVal interface{}
	enc := func() (_err error) {
		_err = c.enc.WriteArrayHeader(1)
		if _err != nil {
			return
		}

		_err = c.enc.WriteString(name)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = c.dec.ReadIntf()

		return
	}
	respChan, err := c.makeCall(clientGetOption, enc, dec)
	if err != nil {
		return retVal, c.panicOrReturn(errors.Annotate(err, "Could not make call to Client.GetOption"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, c.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, c.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(interface{})
	return retVal, nil

}

// GetTabpages waiting for documentation from Neovim
func (c *Client) GetTabpages() ([]Tabpage, error) {
	var retVal []Tabpage
	enc := func() (_err error) {
		_err = c.enc.WriteArrayHeader(0)
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
		return retVal, c.panicOrReturn(errors.Annotate(err, "Could not make call to Client.GetTabpages"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, c.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, c.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.([]Tabpage)
	return retVal, nil

}

// GetVar waiting for documentation from Neovim
func (c *Client) GetVar(name string) (interface{}, error) {
	var retVal interface{}
	enc := func() (_err error) {
		_err = c.enc.WriteArrayHeader(1)
		if _err != nil {
			return
		}

		_err = c.enc.WriteString(name)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = c.dec.ReadIntf()

		return
	}
	respChan, err := c.makeCall(clientGetVar, enc, dec)
	if err != nil {
		return retVal, c.panicOrReturn(errors.Annotate(err, "Could not make call to Client.GetVar"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, c.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, c.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(interface{})
	return retVal, nil

}

// GetVvar waiting for documentation from Neovim
func (c *Client) GetVvar(name string) (interface{}, error) {
	var retVal interface{}
	enc := func() (_err error) {
		_err = c.enc.WriteArrayHeader(1)
		if _err != nil {
			return
		}

		_err = c.enc.WriteString(name)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = c.dec.ReadIntf()

		return
	}
	respChan, err := c.makeCall(clientGetVvar, enc, dec)
	if err != nil {
		return retVal, c.panicOrReturn(errors.Annotate(err, "Could not make call to Client.GetVvar"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, c.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, c.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(interface{})
	return retVal, nil

}

// GetWindows waiting for documentation from Neovim
func (c *Client) GetWindows() ([]Window, error) {
	var retVal []Window
	enc := func() (_err error) {
		_err = c.enc.WriteArrayHeader(0)
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
		return retVal, c.panicOrReturn(errors.Annotate(err, "Could not make call to Client.GetWindows"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, c.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, c.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.([]Window)
	return retVal, nil

}

// Input waiting for documentation from Neovim
func (c *Client) Input(keys string) (int, error) {
	var retVal int
	enc := func() (_err error) {
		_err = c.enc.WriteArrayHeader(1)
		if _err != nil {
			return
		}

		_err = c.enc.WriteString(keys)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = c.dec.ReadInt()

		return
	}
	respChan, err := c.makeCall(clientInput, enc, dec)
	if err != nil {
		return retVal, c.panicOrReturn(errors.Annotate(err, "Could not make call to Client.Input"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, c.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, c.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(int)
	return retVal, nil

}

// ListRuntimePaths waiting for documentation from Neovim
func (c *Client) ListRuntimePaths() ([]string, error) {
	var retVal []string
	enc := func() (_err error) {
		_err = c.enc.WriteArrayHeader(0)
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
		return retVal, c.panicOrReturn(errors.Annotate(err, "Could not make call to Client.ListRuntimePaths"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, c.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, c.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.([]string)
	return retVal, nil

}

// NameToColor waiting for documentation from Neovim
func (c *Client) NameToColor(name string) (int, error) {
	var retVal int
	enc := func() (_err error) {
		_err = c.enc.WriteArrayHeader(1)
		if _err != nil {
			return
		}

		_err = c.enc.WriteString(name)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = c.dec.ReadInt()

		return
	}
	respChan, err := c.makeCall(clientNameToColor, enc, dec)
	if err != nil {
		return retVal, c.panicOrReturn(errors.Annotate(err, "Could not make call to Client.NameToColor"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, c.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, c.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(int)
	return retVal, nil

}

// OutWrite waiting for documentation from Neovim
func (c *Client) OutWrite(str string) error {

	enc := func() (_err error) {
		_err = c.enc.WriteArrayHeader(1)
		if _err != nil {
			return
		}

		_err = c.enc.WriteString(str)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_err = c.dec.ReadNil()

		return
	}
	respChan, err := c.makeCall(clientOutWrite, enc, dec)
	if err != nil {
		return c.panicOrReturn(errors.Annotate(err, "Could not make call to Client.OutWrite"))
	}
	resp := <-respChan
	if resp == nil {
		return c.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return c.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	return nil

}

// ReplaceTermcodes waiting for documentation from Neovim
func (c *Client) ReplaceTermcodes(str string, fromPart bool, doLt bool, special bool) (string, error) {
	var retVal string
	enc := func() (_err error) {
		_err = c.enc.WriteArrayHeader(4)
		if _err != nil {
			return
		}

		_err = c.enc.WriteString(str)

		if _err != nil {
			return
		}

		_err = c.enc.WriteBool(fromPart)

		if _err != nil {
			return
		}

		_err = c.enc.WriteBool(doLt)

		if _err != nil {
			return
		}

		_err = c.enc.WriteBool(special)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = c.dec.ReadString()

		return
	}
	respChan, err := c.makeCall(clientReplaceTermcodes, enc, dec)
	if err != nil {
		return retVal, c.panicOrReturn(errors.Annotate(err, "Could not make call to Client.ReplaceTermcodes"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, c.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, c.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(string)
	return retVal, nil

}

// ReportError waiting for documentation from Neovim
func (c *Client) ReportError(str string) error {

	enc := func() (_err error) {
		_err = c.enc.WriteArrayHeader(1)
		if _err != nil {
			return
		}

		_err = c.enc.WriteString(str)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_err = c.dec.ReadNil()

		return
	}
	respChan, err := c.makeCall(clientReportError, enc, dec)
	if err != nil {
		return c.panicOrReturn(errors.Annotate(err, "Could not make call to Client.ReportError"))
	}
	resp := <-respChan
	if resp == nil {
		return c.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return c.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	return nil

}

// SetCurrentBuffer waiting for documentation from Neovim
func (c *Client) SetCurrentBuffer(buffer Buffer) error {

	enc := func() (_err error) {
		_err = c.enc.WriteArrayHeader(1)
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

		_err = c.dec.ReadNil()

		return
	}
	respChan, err := c.makeCall(clientSetCurrentBuffer, enc, dec)
	if err != nil {
		return c.panicOrReturn(errors.Annotate(err, "Could not make call to Client.SetCurrentBuffer"))
	}
	resp := <-respChan
	if resp == nil {
		return c.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return c.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	return nil

}

// SetCurrentLine waiting for documentation from Neovim
func (c *Client) SetCurrentLine(line string) error {

	enc := func() (_err error) {
		_err = c.enc.WriteArrayHeader(1)
		if _err != nil {
			return
		}

		_err = c.enc.WriteString(line)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_err = c.dec.ReadNil()

		return
	}
	respChan, err := c.makeCall(clientSetCurrentLine, enc, dec)
	if err != nil {
		return c.panicOrReturn(errors.Annotate(err, "Could not make call to Client.SetCurrentLine"))
	}
	resp := <-respChan
	if resp == nil {
		return c.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return c.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	return nil

}

// SetCurrentTabpage waiting for documentation from Neovim
func (c *Client) SetCurrentTabpage(tabpage Tabpage) error {

	enc := func() (_err error) {
		_err = c.enc.WriteArrayHeader(1)
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

		_err = c.dec.ReadNil()

		return
	}
	respChan, err := c.makeCall(clientSetCurrentTabpage, enc, dec)
	if err != nil {
		return c.panicOrReturn(errors.Annotate(err, "Could not make call to Client.SetCurrentTabpage"))
	}
	resp := <-respChan
	if resp == nil {
		return c.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return c.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	return nil

}

// SetCurrentWindow waiting for documentation from Neovim
func (c *Client) SetCurrentWindow(window Window) error {

	enc := func() (_err error) {
		_err = c.enc.WriteArrayHeader(1)
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

		_err = c.dec.ReadNil()

		return
	}
	respChan, err := c.makeCall(clientSetCurrentWindow, enc, dec)
	if err != nil {
		return c.panicOrReturn(errors.Annotate(err, "Could not make call to Client.SetCurrentWindow"))
	}
	resp := <-respChan
	if resp == nil {
		return c.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return c.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	return nil

}

// SetOption waiting for documentation from Neovim
func (c *Client) SetOption(name string, value interface{}) error {

	enc := func() (_err error) {
		_err = c.enc.WriteArrayHeader(2)
		if _err != nil {
			return
		}

		_err = c.enc.WriteString(name)

		if _err != nil {
			return
		}

		_err = c.enc.WriteIntf(value)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_err = c.dec.ReadNil()

		return
	}
	respChan, err := c.makeCall(clientSetOption, enc, dec)
	if err != nil {
		return c.panicOrReturn(errors.Annotate(err, "Could not make call to Client.SetOption"))
	}
	resp := <-respChan
	if resp == nil {
		return c.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return c.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	return nil

}

// SetVar waiting for documentation from Neovim
func (c *Client) SetVar(name string, value interface{}) (interface{}, error) {
	var retVal interface{}
	enc := func() (_err error) {
		_err = c.enc.WriteArrayHeader(2)
		if _err != nil {
			return
		}

		_err = c.enc.WriteString(name)

		if _err != nil {
			return
		}

		_err = c.enc.WriteIntf(value)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = c.dec.ReadIntf()

		return
	}
	respChan, err := c.makeCall(clientSetVar, enc, dec)
	if err != nil {
		return retVal, c.panicOrReturn(errors.Annotate(err, "Could not make call to Client.SetVar"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, c.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, c.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(interface{})
	return retVal, nil

}

// Strwidth waiting for documentation from Neovim
func (c *Client) Strwidth(str string) (int, error) {
	var retVal int
	enc := func() (_err error) {
		_err = c.enc.WriteArrayHeader(1)
		if _err != nil {
			return
		}

		_err = c.enc.WriteString(str)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = c.dec.ReadInt()

		return
	}
	respChan, err := c.makeCall(clientStrwidth, enc, dec)
	if err != nil {
		return retVal, c.panicOrReturn(errors.Annotate(err, "Could not make call to Client.Strwidth"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, c.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, c.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(int)
	return retVal, nil

}

// Subscribe waiting for documentation from Neovim
func (c *Client) Subscribe(event string) error {

	enc := func() (_err error) {
		_err = c.enc.WriteArrayHeader(1)
		if _err != nil {
			return
		}

		_err = c.enc.WriteString(event)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_err = c.dec.ReadNil()

		return
	}
	respChan, err := c.makeCall(clientSubscribe, enc, dec)
	if err != nil {
		return c.panicOrReturn(errors.Annotate(err, "Could not make call to Client.Subscribe"))
	}
	resp := <-respChan
	if resp == nil {
		return c.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return c.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	return nil

}

// Unsubscribe waiting for documentation from Neovim
func (c *Client) Unsubscribe(event string) error {

	enc := func() (_err error) {
		_err = c.enc.WriteArrayHeader(1)
		if _err != nil {
			return
		}

		_err = c.enc.WriteString(event)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_err = c.dec.ReadNil()

		return
	}
	respChan, err := c.makeCall(clientUnsubscribe, enc, dec)
	if err != nil {
		return c.panicOrReturn(errors.Annotate(err, "Could not make call to Client.Unsubscribe"))
	}
	resp := <-respChan
	if resp == nil {
		return c.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return c.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	return nil

}

// GetBuffer waiting for documentation from Neovim
func (w *Window) GetBuffer() (Buffer, error) {
	var retVal Buffer
	enc := func() (_err error) {
		_err = w.client.enc.WriteArrayHeader(1)
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
		return retVal, w.client.panicOrReturn(errors.Annotate(err, "Could not make call to Window.GetBuffer"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, w.client.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, w.client.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(Buffer)
	return retVal, nil

}

// GetCursor waiting for documentation from Neovim
func (w *Window) GetCursor() ([]int, error) {
	var retVal []int
	enc := func() (_err error) {
		_err = w.client.enc.WriteArrayHeader(1)
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
		return retVal, w.client.panicOrReturn(errors.Annotate(err, "Could not make call to Window.GetCursor"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, w.client.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, w.client.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.([]int)
	return retVal, nil

}

// GetHeight waiting for documentation from Neovim
func (w *Window) GetHeight() (int, error) {
	var retVal int
	enc := func() (_err error) {
		_err = w.client.enc.WriteArrayHeader(1)
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

		_i, _err = w.client.dec.ReadInt()

		return
	}
	respChan, err := w.client.makeCall(windowGetHeight, enc, dec)
	if err != nil {
		return retVal, w.client.panicOrReturn(errors.Annotate(err, "Could not make call to Window.GetHeight"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, w.client.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, w.client.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(int)
	return retVal, nil

}

// GetOption waiting for documentation from Neovim
func (w *Window) GetOption(name string) (interface{}, error) {
	var retVal interface{}
	enc := func() (_err error) {
		_err = w.client.enc.WriteArrayHeader(2)
		if _err != nil {
			return
		}

		_err = w.client.encodeWindow(*w)
		if _err != nil {
			return
		}

		_err = w.client.enc.WriteString(name)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = w.client.dec.ReadIntf()

		return
	}
	respChan, err := w.client.makeCall(windowGetOption, enc, dec)
	if err != nil {
		return retVal, w.client.panicOrReturn(errors.Annotate(err, "Could not make call to Window.GetOption"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, w.client.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, w.client.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(interface{})
	return retVal, nil

}

// GetPosition waiting for documentation from Neovim
func (w *Window) GetPosition() ([]int, error) {
	var retVal []int
	enc := func() (_err error) {
		_err = w.client.enc.WriteArrayHeader(1)
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
		return retVal, w.client.panicOrReturn(errors.Annotate(err, "Could not make call to Window.GetPosition"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, w.client.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, w.client.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.([]int)
	return retVal, nil

}

// GetTabpage waiting for documentation from Neovim
func (w *Window) GetTabpage() (Tabpage, error) {
	var retVal Tabpage
	enc := func() (_err error) {
		_err = w.client.enc.WriteArrayHeader(1)
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
		return retVal, w.client.panicOrReturn(errors.Annotate(err, "Could not make call to Window.GetTabpage"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, w.client.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, w.client.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(Tabpage)
	return retVal, nil

}

// GetVar waiting for documentation from Neovim
func (w *Window) GetVar(name string) (interface{}, error) {
	var retVal interface{}
	enc := func() (_err error) {
		_err = w.client.enc.WriteArrayHeader(2)
		if _err != nil {
			return
		}

		_err = w.client.encodeWindow(*w)
		if _err != nil {
			return
		}

		_err = w.client.enc.WriteString(name)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = w.client.dec.ReadIntf()

		return
	}
	respChan, err := w.client.makeCall(windowGetVar, enc, dec)
	if err != nil {
		return retVal, w.client.panicOrReturn(errors.Annotate(err, "Could not make call to Window.GetVar"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, w.client.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, w.client.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(interface{})
	return retVal, nil

}

// GetWidth waiting for documentation from Neovim
func (w *Window) GetWidth() (int, error) {
	var retVal int
	enc := func() (_err error) {
		_err = w.client.enc.WriteArrayHeader(1)
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

		_i, _err = w.client.dec.ReadInt()

		return
	}
	respChan, err := w.client.makeCall(windowGetWidth, enc, dec)
	if err != nil {
		return retVal, w.client.panicOrReturn(errors.Annotate(err, "Could not make call to Window.GetWidth"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, w.client.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, w.client.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(int)
	return retVal, nil

}

// IsValid waiting for documentation from Neovim
func (w *Window) IsValid() (bool, error) {
	var retVal bool
	enc := func() (_err error) {
		_err = w.client.enc.WriteArrayHeader(1)
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

		_i, _err = w.client.dec.ReadBool()

		return
	}
	respChan, err := w.client.makeCall(windowIsValid, enc, dec)
	if err != nil {
		return retVal, w.client.panicOrReturn(errors.Annotate(err, "Could not make call to Window.IsValid"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, w.client.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, w.client.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(bool)
	return retVal, nil

}

// SetCursor waiting for documentation from Neovim
func (w *Window) SetCursor(pos []int) error {

	enc := func() (_err error) {
		_err = w.client.enc.WriteArrayHeader(2)
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

		_err = w.client.dec.ReadNil()

		return
	}
	respChan, err := w.client.makeCall(windowSetCursor, enc, dec)
	if err != nil {
		return w.client.panicOrReturn(errors.Annotate(err, "Could not make call to Window.SetCursor"))
	}
	resp := <-respChan
	if resp == nil {
		return w.client.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return w.client.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	return nil

}

// SetHeight waiting for documentation from Neovim
func (w *Window) SetHeight(height int) error {

	enc := func() (_err error) {
		_err = w.client.enc.WriteArrayHeader(2)
		if _err != nil {
			return
		}

		_err = w.client.encodeWindow(*w)
		if _err != nil {
			return
		}

		_err = w.client.enc.WriteInt(height)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_err = w.client.dec.ReadNil()

		return
	}
	respChan, err := w.client.makeCall(windowSetHeight, enc, dec)
	if err != nil {
		return w.client.panicOrReturn(errors.Annotate(err, "Could not make call to Window.SetHeight"))
	}
	resp := <-respChan
	if resp == nil {
		return w.client.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return w.client.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	return nil

}

// SetOption waiting for documentation from Neovim
func (w *Window) SetOption(name string, value interface{}) error {

	enc := func() (_err error) {
		_err = w.client.enc.WriteArrayHeader(3)
		if _err != nil {
			return
		}

		_err = w.client.encodeWindow(*w)
		if _err != nil {
			return
		}

		_err = w.client.enc.WriteString(name)

		if _err != nil {
			return
		}

		_err = w.client.enc.WriteIntf(value)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_err = w.client.dec.ReadNil()

		return
	}
	respChan, err := w.client.makeCall(windowSetOption, enc, dec)
	if err != nil {
		return w.client.panicOrReturn(errors.Annotate(err, "Could not make call to Window.SetOption"))
	}
	resp := <-respChan
	if resp == nil {
		return w.client.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return w.client.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	return nil

}

// SetVar waiting for documentation from Neovim
func (w *Window) SetVar(name string, value interface{}) (interface{}, error) {
	var retVal interface{}
	enc := func() (_err error) {
		_err = w.client.enc.WriteArrayHeader(3)
		if _err != nil {
			return
		}

		_err = w.client.encodeWindow(*w)
		if _err != nil {
			return
		}

		_err = w.client.enc.WriteString(name)

		if _err != nil {
			return
		}

		_err = w.client.enc.WriteIntf(value)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = w.client.dec.ReadIntf()

		return
	}
	respChan, err := w.client.makeCall(windowSetVar, enc, dec)
	if err != nil {
		return retVal, w.client.panicOrReturn(errors.Annotate(err, "Could not make call to Window.SetVar"))
	}
	resp := <-respChan
	if resp == nil {
		return retVal, w.client.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return retVal, w.client.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	retVal = resp.obj.(interface{})
	return retVal, nil

}

// SetWidth waiting for documentation from Neovim
func (w *Window) SetWidth(width int) error {

	enc := func() (_err error) {
		_err = w.client.enc.WriteArrayHeader(2)
		if _err != nil {
			return
		}

		_err = w.client.encodeWindow(*w)
		if _err != nil {
			return
		}

		_err = w.client.enc.WriteInt(width)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_err = w.client.dec.ReadNil()

		return
	}
	respChan, err := w.client.makeCall(windowSetWidth, enc, dec)
	if err != nil {
		return w.client.panicOrReturn(errors.Annotate(err, "Could not make call to Window.SetWidth"))
	}
	resp := <-respChan
	if resp == nil {
		return w.client.panicOrReturn(errors.New("We got a nil response on respChan"))
	}
	if resp.err != nil {
		return w.client.panicOrReturn(errors.Annotate(err, "We got a non-nil error in our response"))
	}

	return nil

}

func (c *Client) decodeBuffer() (retVal Buffer, retErr error) {
	retVal.client = c
	err := c.dec.ReadExtension(&retVal)
	if err != nil {
		return retVal, errors.Annotatef(err, "Could not decode extension Buffer")
	}

	return
}

func (c *Client) encodeBuffer(b Buffer) error {
	err := c.enc.WriteExtension(&b)
	if err != nil {
		return errors.Annotatef(err, "Could not encode Buffer")
	}
	return nil
}

func (c *Client) decodeWindow() (retVal Window, retErr error) {
	retVal.client = c
	err := c.dec.ReadExtension(&retVal)
	if err != nil {
		return retVal, errors.Annotatef(err, "Could not decode extension Window")
	}

	return
}

func (c *Client) encodeWindow(b Window) error {
	err := c.enc.WriteExtension(&b)
	if err != nil {
		return errors.Annotatef(err, "Could not encode Window")
	}
	return nil
}

func (c *Client) decodeTabpage() (retVal Tabpage, retErr error) {
	retVal.client = c
	err := c.dec.ReadExtension(&retVal)
	if err != nil {
		return retVal, errors.Annotatef(err, "Could not decode extension Tabpage")
	}

	return
}

func (c *Client) encodeTabpage(b Tabpage) error {
	err := c.enc.WriteExtension(&b)
	if err != nil {
		return errors.Annotatef(err, "Could not encode Tabpage")
	}
	return nil
}

// helper functions for types

func (c *Client) encodeBufferSlice(s []Buffer) error {
	err := c.enc.WriteArrayHeader(uint32(len(s)))
	if err != nil {
		return errors.Annotate(err, "Could not encode slice length")
	}

	for i := 0; i < len(s); i++ {

		err := c.encodeBuffer(s[i])

		if err != nil {
			return errors.Annotatef(err, "Could not encode Buffer at index %v", i)
		}
	}

	return nil
}

func (c *Client) decodeBufferSlice() ([]Buffer, error) {
	l, err := c.dec.ReadArrayHeader()
	if err != nil {
		return nil, errors.Annotate(err, "Could not decode slice length")
	}

	res := make([]Buffer, l)

	var i uint32
	for i = 0; i < l; i++ {

		b, err := c.decodeBuffer()

		if err != nil {
			return nil, errors.Annotatef(err, "Could not decode Buffer at index %v", i)
		}
		res[i] = b
	}

	return res, nil
}

func (c *Client) encodeTabpageSlice(s []Tabpage) error {
	err := c.enc.WriteArrayHeader(uint32(len(s)))
	if err != nil {
		return errors.Annotate(err, "Could not encode slice length")
	}

	for i := 0; i < len(s); i++ {

		err := c.encodeTabpage(s[i])

		if err != nil {
			return errors.Annotatef(err, "Could not encode Tabpage at index %v", i)
		}
	}

	return nil
}

func (c *Client) decodeTabpageSlice() ([]Tabpage, error) {
	l, err := c.dec.ReadArrayHeader()
	if err != nil {
		return nil, errors.Annotate(err, "Could not decode slice length")
	}

	res := make([]Tabpage, l)

	var i uint32
	for i = 0; i < l; i++ {

		b, err := c.decodeTabpage()

		if err != nil {
			return nil, errors.Annotatef(err, "Could not decode Tabpage at index %v", i)
		}
		res[i] = b
	}

	return res, nil
}

func (c *Client) encodeWindowSlice(s []Window) error {
	err := c.enc.WriteArrayHeader(uint32(len(s)))
	if err != nil {
		return errors.Annotate(err, "Could not encode slice length")
	}

	for i := 0; i < len(s); i++ {

		err := c.encodeWindow(s[i])

		if err != nil {
			return errors.Annotatef(err, "Could not encode Window at index %v", i)
		}
	}

	return nil
}

func (c *Client) decodeWindowSlice() ([]Window, error) {
	l, err := c.dec.ReadArrayHeader()
	if err != nil {
		return nil, errors.Annotate(err, "Could not decode slice length")
	}

	res := make([]Window, l)

	var i uint32
	for i = 0; i < l; i++ {

		b, err := c.decodeWindow()

		if err != nil {
			return nil, errors.Annotatef(err, "Could not decode Window at index %v", i)
		}
		res[i] = b
	}

	return res, nil
}

func (c *Client) encodeIntSlice(s []int) error {
	err := c.enc.WriteArrayHeader(uint32(len(s)))
	if err != nil {
		return errors.Annotate(err, "Could not encode slice length")
	}

	for i := 0; i < len(s); i++ {

		err := c.enc.WriteInt(s[i])

		if err != nil {
			return errors.Annotatef(err, "Could not encode int at index %v", i)
		}
	}

	return nil
}

func (c *Client) decodeIntSlice() ([]int, error) {
	l, err := c.dec.ReadArrayHeader()
	if err != nil {
		return nil, errors.Annotate(err, "Could not decode slice length")
	}

	res := make([]int, l)

	var i uint32
	for i = 0; i < l; i++ {

		b, err := c.dec.ReadInt()

		if err != nil {
			return nil, errors.Annotatef(err, "Could not decode int at index %v", i)
		}
		res[i] = b
	}

	return res, nil
}

func (c *Client) encodeStringSlice(s []string) error {
	err := c.enc.WriteArrayHeader(uint32(len(s)))
	if err != nil {
		return errors.Annotate(err, "Could not encode slice length")
	}

	for i := 0; i < len(s); i++ {

		err := c.enc.WriteString(s[i])

		if err != nil {
			return errors.Annotatef(err, "Could not encode string at index %v", i)
		}
	}

	return nil
}

func (c *Client) decodeStringSlice() ([]string, error) {
	l, err := c.dec.ReadArrayHeader()
	if err != nil {
		return nil, errors.Annotate(err, "Could not decode slice length")
	}

	res := make([]string, l)

	var i uint32
	for i = 0; i < l; i++ {

		b, err := c.dec.ReadString()

		if err != nil {
			return nil, errors.Annotatef(err, "Could not decode string at index %v", i)
		}
		res[i] = b
	}

	return res, nil
}
