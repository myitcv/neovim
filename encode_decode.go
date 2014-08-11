// Copyright 2014 Paul Jolly <paul@myitcv.org.uk>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package neovim

import "github.com/juju/errgo"

func (c *Client) decodeBuffer() (ret_b Buffer, ret_err error) {
	b, err := c.dec.DecodeUint32()
	if err != nil {
		return ret_b, errgo.Notef(err, "Could not decode Buffer")
	}
	return Buffer{Id: b, client: c}, ret_err
}

func (c *Client) encodeBuffer(b Buffer) error {
	err := c.enc.EncodeUint32(b.Id)
	if err != nil {
		return errgo.Notef(err, "Could not encode Buffer")
	}
	return nil
}

func (c *Client) decodeWindow() (ret_b Window, ret_err error) {
	b, err := c.dec.DecodeUint32()
	if err != nil {
		return ret_b, errgo.Notef(err, "Could not decode Window")
	}
	return Window{Id: b, client: c}, ret_err
}

func (c *Client) encodeWindow(b Window) error {
	err := c.enc.EncodeUint32(b.Id)
	if err != nil {
		return errgo.Notef(err, "Could not encode Window")
	}
	return nil
}

func (c *Client) decodeTabpage() (ret_b Tabpage, ret_err error) {
	b, err := c.dec.DecodeUint32()
	if err != nil {
		return ret_b, errgo.Notef(err, "Could not decode Tabpage")
	}
	return Tabpage{Id: b, client: c}, ret_err
}

func (c *Client) encodeTabpage(b Tabpage) error {
	err := c.enc.EncodeUint32(b.Id)
	if err != nil {
		return errgo.Notef(err, "Could not encode Tabpage")
	}
	return nil
}
