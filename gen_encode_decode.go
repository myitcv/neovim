package neovim

import "github.com/juju/errgo"

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
