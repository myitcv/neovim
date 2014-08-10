// **** THIS FILE IS GENERATED - DO NOT EDIT BY HAND
package neovim

import "github.com/juju/errgo"

func (recv *Buffer) GetLength() (ret_val int, ret_err error) {
	enc := func() (_err error) {
		_err = recv.client.enc.EncodeSliceLen(1)
		if _err != nil {
			return
		}

		_err = recv.encode()
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

		_err = recv.encode()
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

		_err = recv.encode()
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
