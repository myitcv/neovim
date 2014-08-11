package neovim

import "github.com/juju/errgo"

func (c *Client) decodeBuffer() (ret_b Buffer, ret_err error) {
	b, err := c.dec.DecodeUint32()
	if err != nil {
		return ret_b, errgo.Notef(err, "Could not decode Buffer")
	}
	return Buffer{Id: b, client: c}, ret_err
}

func (c *Client) encodeBuffer(b *Buffer) error {
	err := c.enc.EncodeUint32(b.Id)
	if err != nil {
		return errgo.Notef(err, "Could not encode Buffer")
	}
	return nil
}
