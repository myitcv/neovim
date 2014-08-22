package neovim

// **** THIS FILE IS GENERATED - DO NOT EDIT BY HAND

import "github.com/juju/errgo"

// constants representing method ids

const (
	neovimAPI               neovimMethodID = 0
	tabpageGetWindows                      = 1
	tabpageGetVar                          = 2
	tabpageSetVar                          = 3
	tabpageGetWindow                       = 4
	tabpageIsValid                         = 5
	bufferGetLength                        = 6
	bufferGetLine                          = 7
	bufferSetLine                          = 8
	bufferDelLine                          = 9
	bufferGetSlice                         = 10
	bufferSetSlice                         = 11
	bufferGetVar                           = 12
	bufferSetVar                           = 13
	bufferGetOption                        = 14
	bufferSetOption                        = 15
	bufferGetNumber                        = 16
	bufferGetName                          = 17
	bufferSetName                          = 18
	bufferIsValid                          = 19
	bufferInsert                           = 20
	bufferGetMark                          = 21
	clientPushKeys                         = 22
	clientCommand                          = 23
	clientFeedkeys                         = 24
	clientReplaceTermcodes                 = 25
	clientEval                             = 26
	clientStrwidth                         = 27
	clientListRuntimePaths                 = 28
	clientChangeDirectory                  = 29
	clientGetCurrentLine                   = 30
	clientSetCurrentLine                   = 31
	clientDelCurrentLine                   = 32
	clientGetVar                           = 33
	clientSetVar                           = 34
	clientGetVvar                          = 35
	clientGetOption                        = 36
	clientSetOption                        = 37
	clientOutWrite                         = 38
	clientErrWrite                         = 39
	clientGetBuffers                       = 40
	clientGetCurrentBuffer                 = 41
	clientSetCurrentBuffer                 = 42
	clientGetWindows                       = 43
	clientGetCurrentWindow                 = 44
	clientSetCurrentWindow                 = 45
	clientGetTabpages                      = 46
	clientGetCurrentTabpage                = 47
	clientSetCurrentTabpage                = 48
	clientSubscribe                        = 49
	clientUnsubscribe                      = 50
	clientRegisterProvider                 = 51
	windowGetBuffer                        = 52
	windowGetCursor                        = 53
	windowSetCursor                        = 54
	windowGetHeight                        = 55
	windowSetHeight                        = 56
	windowGetWidth                         = 57
	windowSetWidth                         = 58
	windowGetVar                           = 59
	windowSetVar                           = 60
	windowGetOption                        = 61
	windowSetOption                        = 62
	windowGetPosition                      = 63
	windowGetTabpage                       = 64
	windowIsValid                          = 65
)

func (n neovimMethodID) String() string {
	switch n {
	case neovimAPI:
		return "API"
	case tabpageGetWindows:
		return "TabpageGetWindows"
	case tabpageGetVar:
		return "TabpageGetVar"
	case tabpageSetVar:
		return "TabpageSetVar"
	case tabpageGetWindow:
		return "TabpageGetWindow"
	case tabpageIsValid:
		return "TabpageIsValid"
	case bufferGetLength:
		return "BufferGetLength"
	case bufferGetLine:
		return "BufferGetLine"
	case bufferSetLine:
		return "BufferSetLine"
	case bufferDelLine:
		return "BufferDelLine"
	case bufferGetSlice:
		return "BufferGetSlice"
	case bufferSetSlice:
		return "BufferSetSlice"
	case bufferGetVar:
		return "BufferGetVar"
	case bufferSetVar:
		return "BufferSetVar"
	case bufferGetOption:
		return "BufferGetOption"
	case bufferSetOption:
		return "BufferSetOption"
	case bufferGetNumber:
		return "BufferGetNumber"
	case bufferGetName:
		return "BufferGetName"
	case bufferSetName:
		return "BufferSetName"
	case bufferIsValid:
		return "BufferIsValid"
	case bufferInsert:
		return "BufferInsert"
	case bufferGetMark:
		return "BufferGetMark"
	case clientPushKeys:
		return "ClientPushKeys"
	case clientCommand:
		return "ClientCommand"
	case clientFeedkeys:
		return "ClientFeedkeys"
	case clientReplaceTermcodes:
		return "ClientReplaceTermcodes"
	case clientEval:
		return "ClientEval"
	case clientStrwidth:
		return "ClientStrwidth"
	case clientListRuntimePaths:
		return "ClientListRuntimePaths"
	case clientChangeDirectory:
		return "ClientChangeDirectory"
	case clientGetCurrentLine:
		return "ClientGetCurrentLine"
	case clientSetCurrentLine:
		return "ClientSetCurrentLine"
	case clientDelCurrentLine:
		return "ClientDelCurrentLine"
	case clientGetVar:
		return "ClientGetVar"
	case clientSetVar:
		return "ClientSetVar"
	case clientGetVvar:
		return "ClientGetVvar"
	case clientGetOption:
		return "ClientGetOption"
	case clientSetOption:
		return "ClientSetOption"
	case clientOutWrite:
		return "ClientOutWrite"
	case clientErrWrite:
		return "ClientErrWrite"
	case clientGetBuffers:
		return "ClientGetBuffers"
	case clientGetCurrentBuffer:
		return "ClientGetCurrentBuffer"
	case clientSetCurrentBuffer:
		return "ClientSetCurrentBuffer"
	case clientGetWindows:
		return "ClientGetWindows"
	case clientGetCurrentWindow:
		return "ClientGetCurrentWindow"
	case clientSetCurrentWindow:
		return "ClientSetCurrentWindow"
	case clientGetTabpages:
		return "ClientGetTabpages"
	case clientGetCurrentTabpage:
		return "ClientGetCurrentTabpage"
	case clientSetCurrentTabpage:
		return "ClientSetCurrentTabpage"
	case clientSubscribe:
		return "ClientSubscribe"
	case clientUnsubscribe:
		return "ClientUnsubscribe"
	case clientRegisterProvider:
		return "ClientRegisterProvider"
	case windowGetBuffer:
		return "WindowGetBuffer"
	case windowGetCursor:
		return "WindowGetCursor"
	case windowSetCursor:
		return "WindowSetCursor"
	case windowGetHeight:
		return "WindowGetHeight"
	case windowSetHeight:
		return "WindowSetHeight"
	case windowGetWidth:
		return "WindowGetWidth"
	case windowSetWidth:
		return "WindowSetWidth"
	case windowGetVar:
		return "WindowGetVar"
	case windowSetVar:
		return "WindowSetVar"
	case windowGetOption:
		return "WindowGetOption"
	case windowSetOption:
		return "WindowSetOption"
	case windowGetPosition:
		return "WindowGetPosition"
	case windowGetTabpage:
		return "WindowGetTabpage"
	case windowIsValid:
		return "WindowIsValid"

	default:
		return ""
	}
}

// methods on the API

// GetWindows waiting for documentation from Neovim
func (t *Tabpage) GetWindows() ([]Window, error) {
	var retVal []Window
	var retErr error
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
		return retVal, errgo.NoteMask(err, "Could not make call to Tabpage.GetWindows")
	}
	resp := <-respChan
	if resp == nil {
		return retVal, errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return retVal, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	retVal = resp.obj.([]Window)
	return retVal, retErr

}

// GetVar waiting for documentation from Neovim
func (t *Tabpage) GetVar(name string) (interface{}, error) {
	var retVal interface{}
	var retErr error
	enc := func() (_err error) {
		_err = t.client.enc.EncodeSliceLen(2)
		if _err != nil {
			return
		}

		_err = t.client.encodeTabpage(*t)
		if _err != nil {
			return
		}

		_err = t.client.enc.EncodeString(name)

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
		return retVal, errgo.NoteMask(err, "Could not make call to Tabpage.GetVar")
	}
	resp := <-respChan
	if resp == nil {
		return retVal, errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return retVal, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	retVal = resp.obj.(interface{})
	return retVal, retErr

}

// SetVar waiting for documentation from Neovim
func (t *Tabpage) SetVar(name string, value interface{}) (interface{}, error) {
	var retVal interface{}
	var retErr error
	enc := func() (_err error) {
		_err = t.client.enc.EncodeSliceLen(3)
		if _err != nil {
			return
		}

		_err = t.client.encodeTabpage(*t)
		if _err != nil {
			return
		}

		_err = t.client.enc.EncodeString(name)

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
		return retVal, errgo.NoteMask(err, "Could not make call to Tabpage.SetVar")
	}
	resp := <-respChan
	if resp == nil {
		return retVal, errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return retVal, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	retVal = resp.obj.(interface{})
	return retVal, retErr

}

// GetWindow waiting for documentation from Neovim
func (t *Tabpage) GetWindow() (Window, error) {
	var retVal Window
	var retErr error
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
		return retVal, errgo.NoteMask(err, "Could not make call to Tabpage.GetWindow")
	}
	resp := <-respChan
	if resp == nil {
		return retVal, errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return retVal, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	retVal = resp.obj.(Window)
	return retVal, retErr

}

// IsValid waiting for documentation from Neovim
func (t *Tabpage) IsValid() (bool, error) {
	var retVal bool
	var retErr error
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
		return retVal, errgo.NoteMask(err, "Could not make call to Tabpage.IsValid")
	}
	resp := <-respChan
	if resp == nil {
		return retVal, errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return retVal, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	retVal = resp.obj.(bool)
	return retVal, retErr

}

// GetLength waiting for documentation from Neovim
func (b *Buffer) GetLength() (int, error) {
	var retVal int
	var retErr error
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
	respChan, err := b.client.makeCall(bufferGetLength, enc, dec)
	if err != nil {
		return retVal, errgo.NoteMask(err, "Could not make call to Buffer.GetLength")
	}
	resp := <-respChan
	if resp == nil {
		return retVal, errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return retVal, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	retVal = resp.obj.(int)
	return retVal, retErr

}

// GetLine waiting for documentation from Neovim
func (b *Buffer) GetLine(index int) (string, error) {
	var retVal string
	var retErr error
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

		_i, _err = b.client.dec.DecodeString()

		return
	}
	respChan, err := b.client.makeCall(bufferGetLine, enc, dec)
	if err != nil {
		return retVal, errgo.NoteMask(err, "Could not make call to Buffer.GetLine")
	}
	resp := <-respChan
	if resp == nil {
		return retVal, errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return retVal, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	retVal = resp.obj.(string)
	return retVal, retErr

}

// SetLine waiting for documentation from Neovim
func (b *Buffer) SetLine(index int, line string) error {

	var retErr error
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

		_err = b.client.enc.EncodeString(line)

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
		return errgo.NoteMask(err, "Could not make call to Buffer.SetLine")
	}
	resp := <-respChan
	if resp == nil {
		return errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	return retErr

}

// DelLine waiting for documentation from Neovim
func (b *Buffer) DelLine(index int) error {

	var retErr error
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
		return errgo.NoteMask(err, "Could not make call to Buffer.DelLine")
	}
	resp := <-respChan
	if resp == nil {
		return errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	return retErr

}

// GetSlice waiting for documentation from Neovim
func (b *Buffer) GetSlice(start int, end int, includeStart bool, includeEnd bool) ([]string, error) {
	var retVal []string
	var retErr error
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
	respChan, err := b.client.makeCall(bufferGetSlice, enc, dec)
	if err != nil {
		return retVal, errgo.NoteMask(err, "Could not make call to Buffer.GetSlice")
	}
	resp := <-respChan
	if resp == nil {
		return retVal, errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return retVal, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	retVal = resp.obj.([]string)
	return retVal, retErr

}

// SetSlice waiting for documentation from Neovim
func (b *Buffer) SetSlice(start int, end int, includeStart bool, includeEnd bool, replacement []string) error {

	var retErr error
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
	respChan, err := b.client.makeCall(bufferSetSlice, enc, dec)
	if err != nil {
		return errgo.NoteMask(err, "Could not make call to Buffer.SetSlice")
	}
	resp := <-respChan
	if resp == nil {
		return errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	return retErr

}

// GetVar waiting for documentation from Neovim
func (b *Buffer) GetVar(name string) (interface{}, error) {
	var retVal interface{}
	var retErr error
	enc := func() (_err error) {
		_err = b.client.enc.EncodeSliceLen(2)
		if _err != nil {
			return
		}

		_err = b.client.encodeBuffer(*b)
		if _err != nil {
			return
		}

		_err = b.client.enc.EncodeString(name)

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
		return retVal, errgo.NoteMask(err, "Could not make call to Buffer.GetVar")
	}
	resp := <-respChan
	if resp == nil {
		return retVal, errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return retVal, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	retVal = resp.obj.(interface{})
	return retVal, retErr

}

// SetVar waiting for documentation from Neovim
func (b *Buffer) SetVar(name string, value interface{}) (interface{}, error) {
	var retVal interface{}
	var retErr error
	enc := func() (_err error) {
		_err = b.client.enc.EncodeSliceLen(3)
		if _err != nil {
			return
		}

		_err = b.client.encodeBuffer(*b)
		if _err != nil {
			return
		}

		_err = b.client.enc.EncodeString(name)

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
		return retVal, errgo.NoteMask(err, "Could not make call to Buffer.SetVar")
	}
	resp := <-respChan
	if resp == nil {
		return retVal, errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return retVal, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	retVal = resp.obj.(interface{})
	return retVal, retErr

}

// GetOption waiting for documentation from Neovim
func (b *Buffer) GetOption(name string) (interface{}, error) {
	var retVal interface{}
	var retErr error
	enc := func() (_err error) {
		_err = b.client.enc.EncodeSliceLen(2)
		if _err != nil {
			return
		}

		_err = b.client.encodeBuffer(*b)
		if _err != nil {
			return
		}

		_err = b.client.enc.EncodeString(name)

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
		return retVal, errgo.NoteMask(err, "Could not make call to Buffer.GetOption")
	}
	resp := <-respChan
	if resp == nil {
		return retVal, errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return retVal, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	retVal = resp.obj.(interface{})
	return retVal, retErr

}

// SetOption waiting for documentation from Neovim
func (b *Buffer) SetOption(name string, value interface{}) error {

	var retErr error
	enc := func() (_err error) {
		_err = b.client.enc.EncodeSliceLen(3)
		if _err != nil {
			return
		}

		_err = b.client.encodeBuffer(*b)
		if _err != nil {
			return
		}

		_err = b.client.enc.EncodeString(name)

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
		return errgo.NoteMask(err, "Could not make call to Buffer.SetOption")
	}
	resp := <-respChan
	if resp == nil {
		return errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	return retErr

}

// GetNumber waiting for documentation from Neovim
func (b *Buffer) GetNumber() (int, error) {
	var retVal int
	var retErr error
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
		return retVal, errgo.NoteMask(err, "Could not make call to Buffer.GetNumber")
	}
	resp := <-respChan
	if resp == nil {
		return retVal, errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return retVal, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	retVal = resp.obj.(int)
	return retVal, retErr

}

// GetName waiting for documentation from Neovim
func (b *Buffer) GetName() (string, error) {
	var retVal string
	var retErr error
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

		_i, _err = b.client.dec.DecodeString()

		return
	}
	respChan, err := b.client.makeCall(bufferGetName, enc, dec)
	if err != nil {
		return retVal, errgo.NoteMask(err, "Could not make call to Buffer.GetName")
	}
	resp := <-respChan
	if resp == nil {
		return retVal, errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return retVal, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	retVal = resp.obj.(string)
	return retVal, retErr

}

// SetName waiting for documentation from Neovim
func (b *Buffer) SetName(name string) error {

	var retErr error
	enc := func() (_err error) {
		_err = b.client.enc.EncodeSliceLen(2)
		if _err != nil {
			return
		}

		_err = b.client.encodeBuffer(*b)
		if _err != nil {
			return
		}

		_err = b.client.enc.EncodeString(name)

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
		return errgo.NoteMask(err, "Could not make call to Buffer.SetName")
	}
	resp := <-respChan
	if resp == nil {
		return errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	return retErr

}

// IsValid waiting for documentation from Neovim
func (b *Buffer) IsValid() (bool, error) {
	var retVal bool
	var retErr error
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
		return retVal, errgo.NoteMask(err, "Could not make call to Buffer.IsValid")
	}
	resp := <-respChan
	if resp == nil {
		return retVal, errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return retVal, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	retVal = resp.obj.(bool)
	return retVal, retErr

}

// Insert waiting for documentation from Neovim
func (b *Buffer) Insert(lnum int, lines []string) error {

	var retErr error
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
		return errgo.NoteMask(err, "Could not make call to Buffer.Insert")
	}
	resp := <-respChan
	if resp == nil {
		return errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	return retErr

}

// GetMark waiting for documentation from Neovim
func (b *Buffer) GetMark(name string) (uint32, error) {
	var retVal uint32
	var retErr error
	enc := func() (_err error) {
		_err = b.client.enc.EncodeSliceLen(2)
		if _err != nil {
			return
		}

		_err = b.client.encodeBuffer(*b)
		if _err != nil {
			return
		}

		_err = b.client.enc.EncodeString(name)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = b.client.dec.DecodeUint32()

		return
	}
	respChan, err := b.client.makeCall(bufferGetMark, enc, dec)
	if err != nil {
		return retVal, errgo.NoteMask(err, "Could not make call to Buffer.GetMark")
	}
	resp := <-respChan
	if resp == nil {
		return retVal, errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return retVal, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	retVal = resp.obj.(uint32)
	return retVal, retErr

}

// PushKeys waiting for documentation from Neovim
func (c *Client) PushKeys(str string) error {

	var retErr error
	enc := func() (_err error) {
		_err = c.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = c.enc.EncodeString(str)

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
		return errgo.NoteMask(err, "Could not make call to Client.PushKeys")
	}
	resp := <-respChan
	if resp == nil {
		return errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	return retErr

}

// Command waiting for documentation from Neovim
func (c *Client) Command(str string) error {

	var retErr error
	enc := func() (_err error) {
		_err = c.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = c.enc.EncodeString(str)

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
		return errgo.NoteMask(err, "Could not make call to Client.Command")
	}
	resp := <-respChan
	if resp == nil {
		return errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	return retErr

}

// Feedkeys waiting for documentation from Neovim
func (c *Client) Feedkeys(keys string, mode string) error {

	var retErr error
	enc := func() (_err error) {
		_err = c.enc.EncodeSliceLen(2)
		if _err != nil {
			return
		}

		_err = c.enc.EncodeString(keys)

		if _err != nil {
			return
		}

		_err = c.enc.EncodeString(mode)

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
		return errgo.NoteMask(err, "Could not make call to Client.Feedkeys")
	}
	resp := <-respChan
	if resp == nil {
		return errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	return retErr

}

// ReplaceTermcodes waiting for documentation from Neovim
func (c *Client) ReplaceTermcodes(str string, fromPart bool, doLt bool, special bool) (string, error) {
	var retVal string
	var retErr error
	enc := func() (_err error) {
		_err = c.enc.EncodeSliceLen(4)
		if _err != nil {
			return
		}

		_err = c.enc.EncodeString(str)

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

		_i, _err = c.dec.DecodeString()

		return
	}
	respChan, err := c.makeCall(clientReplaceTermcodes, enc, dec)
	if err != nil {
		return retVal, errgo.NoteMask(err, "Could not make call to Client.ReplaceTermcodes")
	}
	resp := <-respChan
	if resp == nil {
		return retVal, errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return retVal, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	retVal = resp.obj.(string)
	return retVal, retErr

}

// Eval waiting for documentation from Neovim
func (c *Client) Eval(str string) (interface{}, error) {
	var retVal interface{}
	var retErr error
	enc := func() (_err error) {
		_err = c.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = c.enc.EncodeString(str)

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
		return retVal, errgo.NoteMask(err, "Could not make call to Client.Eval")
	}
	resp := <-respChan
	if resp == nil {
		return retVal, errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return retVal, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	retVal = resp.obj.(interface{})
	return retVal, retErr

}

// Strwidth waiting for documentation from Neovim
func (c *Client) Strwidth(str string) (int, error) {
	var retVal int
	var retErr error
	enc := func() (_err error) {
		_err = c.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = c.enc.EncodeString(str)

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
		return retVal, errgo.NoteMask(err, "Could not make call to Client.Strwidth")
	}
	resp := <-respChan
	if resp == nil {
		return retVal, errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return retVal, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	retVal = resp.obj.(int)
	return retVal, retErr

}

// ListRuntimePaths waiting for documentation from Neovim
func (c *Client) ListRuntimePaths() ([]string, error) {
	var retVal []string
	var retErr error
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
		return retVal, errgo.NoteMask(err, "Could not make call to Client.ListRuntimePaths")
	}
	resp := <-respChan
	if resp == nil {
		return retVal, errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return retVal, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	retVal = resp.obj.([]string)
	return retVal, retErr

}

// ChangeDirectory waiting for documentation from Neovim
func (c *Client) ChangeDirectory(dir string) error {

	var retErr error
	enc := func() (_err error) {
		_err = c.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = c.enc.EncodeString(dir)

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
		return errgo.NoteMask(err, "Could not make call to Client.ChangeDirectory")
	}
	resp := <-respChan
	if resp == nil {
		return errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	return retErr

}

// GetCurrentLine waiting for documentation from Neovim
func (c *Client) GetCurrentLine() (string, error) {
	var retVal string
	var retErr error
	enc := func() (_err error) {
		_err = c.enc.EncodeSliceLen(0)
		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_i, _err = c.dec.DecodeString()

		return
	}
	respChan, err := c.makeCall(clientGetCurrentLine, enc, dec)
	if err != nil {
		return retVal, errgo.NoteMask(err, "Could not make call to Client.GetCurrentLine")
	}
	resp := <-respChan
	if resp == nil {
		return retVal, errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return retVal, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	retVal = resp.obj.(string)
	return retVal, retErr

}

// SetCurrentLine waiting for documentation from Neovim
func (c *Client) SetCurrentLine(line string) error {

	var retErr error
	enc := func() (_err error) {
		_err = c.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = c.enc.EncodeString(line)

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
		return errgo.NoteMask(err, "Could not make call to Client.SetCurrentLine")
	}
	resp := <-respChan
	if resp == nil {
		return errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	return retErr

}

// DelCurrentLine waiting for documentation from Neovim
func (c *Client) DelCurrentLine() error {

	var retErr error
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
		return errgo.NoteMask(err, "Could not make call to Client.DelCurrentLine")
	}
	resp := <-respChan
	if resp == nil {
		return errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	return retErr

}

// GetVar waiting for documentation from Neovim
func (c *Client) GetVar(name string) (interface{}, error) {
	var retVal interface{}
	var retErr error
	enc := func() (_err error) {
		_err = c.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = c.enc.EncodeString(name)

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
		return retVal, errgo.NoteMask(err, "Could not make call to Client.GetVar")
	}
	resp := <-respChan
	if resp == nil {
		return retVal, errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return retVal, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	retVal = resp.obj.(interface{})
	return retVal, retErr

}

// SetVar waiting for documentation from Neovim
func (c *Client) SetVar(name string, value interface{}) (interface{}, error) {
	var retVal interface{}
	var retErr error
	enc := func() (_err error) {
		_err = c.enc.EncodeSliceLen(2)
		if _err != nil {
			return
		}

		_err = c.enc.EncodeString(name)

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
		return retVal, errgo.NoteMask(err, "Could not make call to Client.SetVar")
	}
	resp := <-respChan
	if resp == nil {
		return retVal, errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return retVal, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	retVal = resp.obj.(interface{})
	return retVal, retErr

}

// GetVvar waiting for documentation from Neovim
func (c *Client) GetVvar(name string) (interface{}, error) {
	var retVal interface{}
	var retErr error
	enc := func() (_err error) {
		_err = c.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = c.enc.EncodeString(name)

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
		return retVal, errgo.NoteMask(err, "Could not make call to Client.GetVvar")
	}
	resp := <-respChan
	if resp == nil {
		return retVal, errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return retVal, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	retVal = resp.obj.(interface{})
	return retVal, retErr

}

// GetOption waiting for documentation from Neovim
func (c *Client) GetOption(name string) (interface{}, error) {
	var retVal interface{}
	var retErr error
	enc := func() (_err error) {
		_err = c.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = c.enc.EncodeString(name)

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
		return retVal, errgo.NoteMask(err, "Could not make call to Client.GetOption")
	}
	resp := <-respChan
	if resp == nil {
		return retVal, errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return retVal, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	retVal = resp.obj.(interface{})
	return retVal, retErr

}

// SetOption waiting for documentation from Neovim
func (c *Client) SetOption(name string, value interface{}) error {

	var retErr error
	enc := func() (_err error) {
		_err = c.enc.EncodeSliceLen(2)
		if _err != nil {
			return
		}

		_err = c.enc.EncodeString(name)

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
		return errgo.NoteMask(err, "Could not make call to Client.SetOption")
	}
	resp := <-respChan
	if resp == nil {
		return errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	return retErr

}

// OutWrite waiting for documentation from Neovim
func (c *Client) OutWrite(str string) error {

	var retErr error
	enc := func() (_err error) {
		_err = c.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = c.enc.EncodeString(str)

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
		return errgo.NoteMask(err, "Could not make call to Client.OutWrite")
	}
	resp := <-respChan
	if resp == nil {
		return errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	return retErr

}

// ErrWrite waiting for documentation from Neovim
func (c *Client) ErrWrite(str string) error {

	var retErr error
	enc := func() (_err error) {
		_err = c.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = c.enc.EncodeString(str)

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
		return errgo.NoteMask(err, "Could not make call to Client.ErrWrite")
	}
	resp := <-respChan
	if resp == nil {
		return errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	return retErr

}

// GetBuffers waiting for documentation from Neovim
func (c *Client) GetBuffers() ([]Buffer, error) {
	var retVal []Buffer
	var retErr error
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
		return retVal, errgo.NoteMask(err, "Could not make call to Client.GetBuffers")
	}
	resp := <-respChan
	if resp == nil {
		return retVal, errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return retVal, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	retVal = resp.obj.([]Buffer)
	return retVal, retErr

}

// GetCurrentBuffer waiting for documentation from Neovim
func (c *Client) GetCurrentBuffer() (Buffer, error) {
	var retVal Buffer
	var retErr error
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
		return retVal, errgo.NoteMask(err, "Could not make call to Client.GetCurrentBuffer")
	}
	resp := <-respChan
	if resp == nil {
		return retVal, errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return retVal, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	retVal = resp.obj.(Buffer)
	return retVal, retErr

}

// SetCurrentBuffer waiting for documentation from Neovim
func (c *Client) SetCurrentBuffer(buffer Buffer) error {

	var retErr error
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
		return errgo.NoteMask(err, "Could not make call to Client.SetCurrentBuffer")
	}
	resp := <-respChan
	if resp == nil {
		return errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	return retErr

}

// GetWindows waiting for documentation from Neovim
func (c *Client) GetWindows() ([]Window, error) {
	var retVal []Window
	var retErr error
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
		return retVal, errgo.NoteMask(err, "Could not make call to Client.GetWindows")
	}
	resp := <-respChan
	if resp == nil {
		return retVal, errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return retVal, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	retVal = resp.obj.([]Window)
	return retVal, retErr

}

// GetCurrentWindow waiting for documentation from Neovim
func (c *Client) GetCurrentWindow() (Window, error) {
	var retVal Window
	var retErr error
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
		return retVal, errgo.NoteMask(err, "Could not make call to Client.GetCurrentWindow")
	}
	resp := <-respChan
	if resp == nil {
		return retVal, errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return retVal, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	retVal = resp.obj.(Window)
	return retVal, retErr

}

// SetCurrentWindow waiting for documentation from Neovim
func (c *Client) SetCurrentWindow(window Window) error {

	var retErr error
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
		return errgo.NoteMask(err, "Could not make call to Client.SetCurrentWindow")
	}
	resp := <-respChan
	if resp == nil {
		return errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	return retErr

}

// GetTabpages waiting for documentation from Neovim
func (c *Client) GetTabpages() ([]Tabpage, error) {
	var retVal []Tabpage
	var retErr error
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
		return retVal, errgo.NoteMask(err, "Could not make call to Client.GetTabpages")
	}
	resp := <-respChan
	if resp == nil {
		return retVal, errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return retVal, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	retVal = resp.obj.([]Tabpage)
	return retVal, retErr

}

// GetCurrentTabpage waiting for documentation from Neovim
func (c *Client) GetCurrentTabpage() (Tabpage, error) {
	var retVal Tabpage
	var retErr error
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
		return retVal, errgo.NoteMask(err, "Could not make call to Client.GetCurrentTabpage")
	}
	resp := <-respChan
	if resp == nil {
		return retVal, errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return retVal, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	retVal = resp.obj.(Tabpage)
	return retVal, retErr

}

// SetCurrentTabpage waiting for documentation from Neovim
func (c *Client) SetCurrentTabpage(tabpage Tabpage) error {

	var retErr error
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
		return errgo.NoteMask(err, "Could not make call to Client.SetCurrentTabpage")
	}
	resp := <-respChan
	if resp == nil {
		return errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	return retErr

}

// Subscribe waiting for documentation from Neovim
func (c *Client) Subscribe(event string) error {

	var retErr error
	enc := func() (_err error) {
		_err = c.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = c.enc.EncodeString(event)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_, _err = c.dec.DecodeBytes()

		return
	}
	respChan, err := c.makeCall(clientSubscribe, enc, dec)
	if err != nil {
		return errgo.NoteMask(err, "Could not make call to Client.Subscribe")
	}
	resp := <-respChan
	if resp == nil {
		return errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	return retErr

}

// Unsubscribe waiting for documentation from Neovim
func (c *Client) Unsubscribe(event string) error {

	var retErr error
	enc := func() (_err error) {
		_err = c.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = c.enc.EncodeString(event)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_, _err = c.dec.DecodeBytes()

		return
	}
	respChan, err := c.makeCall(clientUnsubscribe, enc, dec)
	if err != nil {
		return errgo.NoteMask(err, "Could not make call to Client.Unsubscribe")
	}
	resp := <-respChan
	if resp == nil {
		return errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	return retErr

}

// RegisterProvider waiting for documentation from Neovim
func (c *Client) RegisterProvider(method string) error {

	var retErr error
	enc := func() (_err error) {
		_err = c.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = c.enc.EncodeString(method)

		if _err != nil {
			return
		}

		return
	}
	dec := func() (_i interface{}, _err error) {

		_, _err = c.dec.DecodeBytes()

		return
	}
	respChan, err := c.makeCall(clientRegisterProvider, enc, dec)
	if err != nil {
		return errgo.NoteMask(err, "Could not make call to Client.RegisterProvider")
	}
	resp := <-respChan
	if resp == nil {
		return errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	return retErr

}

// GetBuffer waiting for documentation from Neovim
func (w *Window) GetBuffer() (Buffer, error) {
	var retVal Buffer
	var retErr error
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
		return retVal, errgo.NoteMask(err, "Could not make call to Window.GetBuffer")
	}
	resp := <-respChan
	if resp == nil {
		return retVal, errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return retVal, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	retVal = resp.obj.(Buffer)
	return retVal, retErr

}

// GetCursor waiting for documentation from Neovim
func (w *Window) GetCursor() (uint32, error) {
	var retVal uint32
	var retErr error
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

		_i, _err = w.client.dec.DecodeUint32()

		return
	}
	respChan, err := w.client.makeCall(windowGetCursor, enc, dec)
	if err != nil {
		return retVal, errgo.NoteMask(err, "Could not make call to Window.GetCursor")
	}
	resp := <-respChan
	if resp == nil {
		return retVal, errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return retVal, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	retVal = resp.obj.(uint32)
	return retVal, retErr

}

// SetCursor waiting for documentation from Neovim
func (w *Window) SetCursor(pos uint32) error {

	var retErr error
	enc := func() (_err error) {
		_err = w.client.enc.EncodeSliceLen(2)
		if _err != nil {
			return
		}

		_err = w.client.encodeWindow(*w)
		if _err != nil {
			return
		}

		_err = w.client.enc.EncodeUint32(pos)

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
		return errgo.NoteMask(err, "Could not make call to Window.SetCursor")
	}
	resp := <-respChan
	if resp == nil {
		return errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	return retErr

}

// GetHeight waiting for documentation from Neovim
func (w *Window) GetHeight() (int, error) {
	var retVal int
	var retErr error
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
		return retVal, errgo.NoteMask(err, "Could not make call to Window.GetHeight")
	}
	resp := <-respChan
	if resp == nil {
		return retVal, errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return retVal, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	retVal = resp.obj.(int)
	return retVal, retErr

}

// SetHeight waiting for documentation from Neovim
func (w *Window) SetHeight(height int) error {

	var retErr error
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
		return errgo.NoteMask(err, "Could not make call to Window.SetHeight")
	}
	resp := <-respChan
	if resp == nil {
		return errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	return retErr

}

// GetWidth waiting for documentation from Neovim
func (w *Window) GetWidth() (int, error) {
	var retVal int
	var retErr error
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
		return retVal, errgo.NoteMask(err, "Could not make call to Window.GetWidth")
	}
	resp := <-respChan
	if resp == nil {
		return retVal, errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return retVal, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	retVal = resp.obj.(int)
	return retVal, retErr

}

// SetWidth waiting for documentation from Neovim
func (w *Window) SetWidth(width int) error {

	var retErr error
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
		return errgo.NoteMask(err, "Could not make call to Window.SetWidth")
	}
	resp := <-respChan
	if resp == nil {
		return errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	return retErr

}

// GetVar waiting for documentation from Neovim
func (w *Window) GetVar(name string) (interface{}, error) {
	var retVal interface{}
	var retErr error
	enc := func() (_err error) {
		_err = w.client.enc.EncodeSliceLen(2)
		if _err != nil {
			return
		}

		_err = w.client.encodeWindow(*w)
		if _err != nil {
			return
		}

		_err = w.client.enc.EncodeString(name)

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
		return retVal, errgo.NoteMask(err, "Could not make call to Window.GetVar")
	}
	resp := <-respChan
	if resp == nil {
		return retVal, errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return retVal, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	retVal = resp.obj.(interface{})
	return retVal, retErr

}

// SetVar waiting for documentation from Neovim
func (w *Window) SetVar(name string, value interface{}) (interface{}, error) {
	var retVal interface{}
	var retErr error
	enc := func() (_err error) {
		_err = w.client.enc.EncodeSliceLen(3)
		if _err != nil {
			return
		}

		_err = w.client.encodeWindow(*w)
		if _err != nil {
			return
		}

		_err = w.client.enc.EncodeString(name)

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
		return retVal, errgo.NoteMask(err, "Could not make call to Window.SetVar")
	}
	resp := <-respChan
	if resp == nil {
		return retVal, errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return retVal, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	retVal = resp.obj.(interface{})
	return retVal, retErr

}

// GetOption waiting for documentation from Neovim
func (w *Window) GetOption(name string) (interface{}, error) {
	var retVal interface{}
	var retErr error
	enc := func() (_err error) {
		_err = w.client.enc.EncodeSliceLen(2)
		if _err != nil {
			return
		}

		_err = w.client.encodeWindow(*w)
		if _err != nil {
			return
		}

		_err = w.client.enc.EncodeString(name)

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
		return retVal, errgo.NoteMask(err, "Could not make call to Window.GetOption")
	}
	resp := <-respChan
	if resp == nil {
		return retVal, errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return retVal, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	retVal = resp.obj.(interface{})
	return retVal, retErr

}

// SetOption waiting for documentation from Neovim
func (w *Window) SetOption(name string, value interface{}) error {

	var retErr error
	enc := func() (_err error) {
		_err = w.client.enc.EncodeSliceLen(3)
		if _err != nil {
			return
		}

		_err = w.client.encodeWindow(*w)
		if _err != nil {
			return
		}

		_err = w.client.enc.EncodeString(name)

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
		return errgo.NoteMask(err, "Could not make call to Window.SetOption")
	}
	resp := <-respChan
	if resp == nil {
		return errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	return retErr

}

// GetPosition waiting for documentation from Neovim
func (w *Window) GetPosition() (uint32, error) {
	var retVal uint32
	var retErr error
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

		_i, _err = w.client.dec.DecodeUint32()

		return
	}
	respChan, err := w.client.makeCall(windowGetPosition, enc, dec)
	if err != nil {
		return retVal, errgo.NoteMask(err, "Could not make call to Window.GetPosition")
	}
	resp := <-respChan
	if resp == nil {
		return retVal, errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return retVal, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	retVal = resp.obj.(uint32)
	return retVal, retErr

}

// GetTabpage waiting for documentation from Neovim
func (w *Window) GetTabpage() (Tabpage, error) {
	var retVal Tabpage
	var retErr error
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
		return retVal, errgo.NoteMask(err, "Could not make call to Window.GetTabpage")
	}
	resp := <-respChan
	if resp == nil {
		return retVal, errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return retVal, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	retVal = resp.obj.(Tabpage)
	return retVal, retErr

}

// IsValid waiting for documentation from Neovim
func (w *Window) IsValid() (bool, error) {
	var retVal bool
	var retErr error
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
		return retVal, errgo.NoteMask(err, "Could not make call to Window.IsValid")
	}
	resp := <-respChan
	if resp == nil {
		return retVal, errgo.New("We got a nil response on respChan")
	}
	if resp.err != nil {
		return retVal, errgo.NoteMask(err, "We got a non-nil error in our response")
	}

	retVal = resp.obj.(bool)
	return retVal, retErr

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
