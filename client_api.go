package neovim

import (
	"fmt"

	"github.com/juju/errgo"
)

func (c *Client) GetBuffers() ([]Buffer, error) {
	resp_chan, err := c.makeCall("vim_get_buffers", nil, encodeNoArgs, decodeBufferArray)
	if err != nil {
		return nil, errgo.NoteMask(err, "Could not make call to GetBuffers")
	}
	resp := <-resp_chan
	if resp == nil {
		return nil, errgo.New("We got a nil response on resp_chan")
	}
	if resp.err != nil {
		return nil, errgo.NoteMask(err, "We got a non-nil error in our response")
	}
	ba := resp.obj.([]Buffer)
	fmt.Printf("We got: %v\n", ba)
	return ba, nil
}
