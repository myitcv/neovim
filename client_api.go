package neovim

import "github.com/juju/errgo"

func (c *Client) GetBuffers() (ret_ba []Buffer, ret_err error) {
	dec := func() (_i interface{}, _err error) {
		_i, _err = c.decodeBufferSlice()
		return
	}
	resp_chan, err := c.makeCall(40, c.encodeArgs(), dec)
	if err != nil {
		return ret_ba, errgo.NoteMask(err, "Could not make call to Client.GetBuffers")
	}
	resp := <-resp_chan
	if resp == nil {
		return ret_ba, errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return ret_ba, errgo.NoteMask(err, "We got a non-nil error in our response")
	}
	ret_ba = resp.obj.([]Buffer)
	return ret_ba, ret_err
}

func (c *Client) GetCurrentBuffer() (ret_b Buffer, ret_err error) {
	dec := func() (_i interface{}, _err error) {
		_i, _err = c.decodeBuffer()
		return
	}
	resp_chan, err := c.makeCall(41, c.encodeArgs(), dec)
	if err != nil {
		return ret_b, errgo.NoteMask(err, "Could not make call to Client.GetCurrentBuffer")
	}
	resp := <-resp_chan
	if resp == nil {
		return ret_b, errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return ret_b, errgo.NoteMask(err, "We got a non-nil error in our response")
	}
	ret_b = resp.obj.(Buffer)
	return ret_b, ret_err
}

func (b *Buffer) GetLength() (ret_i int, ret_err error) {
	dec := func() (_i interface{}, _err error) {
		_i, _err = b.client.dec.DecodeInt()
		return
	}
	resp_chan, err := b.client.makeCall(6, b.client.encodeArgs(b.encode), dec)
	if err != nil {
		return ret_i, errgo.NoteMask(err, "Could not make call to Buffer.GetLength")
	}
	resp := <-resp_chan
	if resp == nil {
		return ret_i, errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return ret_i, errgo.NoteMask(err, "We got a non-nil error in our response")
	}
	ret_i = resp.obj.(int)
	return ret_i, ret_err
}
