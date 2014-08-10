package neovim

import (
	"github.com/juju/errgo"
	"github.com/vmihailenco/msgpack"
)

func (c *Client) VimGetBuffers() (ret_b []Buffer, ret_err error) {
	dec := func(d *msgpack.Decoder) (ret_i interface{}, ret_err error) {
		ret_i, ret_err = decodeBufferSlice(d)
		return
	}
	resp_chan, err := c.makeCall(40, []interface{}{nil}, []Encoder{encodeNoArgs}, dec)
	if err != nil {
		return ret_b, errgo.NoteMask(err, "Could not make call to GetBuffers")
	}
	resp := <-resp_chan
	if resp == nil {
		return ret_b, errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return ret_b, errgo.NoteMask(err, "We got a non-nil error in our response")
	}
	ret_b = resp.obj.([]Buffer)
	return ret_b, ret_err
}

func (b *Buffer) GetLength() (ret_i int, ret_err error) {
	dec := func(d *msgpack.Decoder) (ret_i interface{}, ret_err error) {
		ret_i, ret_err = decodeInt(d)
		return
	}
	resp_chan, err := b.client.makeCall(6, []interface{}{b}, []Encoder{encodeBuffer}, dec)
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
	return ret, ret_err
}
