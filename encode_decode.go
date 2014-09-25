// Copyright 2014 Paul Jolly <paul@myitcv.org.uk>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package neovim

import "github.com/juju/errgo"

func (c *Client) decodeBuffer() (retVal Buffer, retErr error) {
	b, err := c.dec.R.ReadByte()
	if err != nil {
		return retVal, errgo.Notef(err, "Could not decode control byte")
	}

	if b != 0xd4 {
		return retVal, errgo.Newf("Expected code d4; got %v\n", b)
	}

	t, err := c.dec.DecodeUint8()
	if err != nil {
		return retVal, errgo.Notef(err, "Could not decode type")
	}

	if t != TypeBuffer {
		return retVal, errgo.Notef(err, "Expected TypeBuffer; got: %v\n", t)
	}

	bid, err := c.dec.DecodeUint8()
	if err != nil {
		return retVal, errgo.Notef(err, "Could not decode buffer ID")
	}
	return Buffer{ID: uint32(bid), client: c}, retErr
}

func (c *Client) encodeBuffer(b Buffer) error {
	err := c.enc.W.WriteByte(0xd4)
	if err != nil {
		return errgo.Notef(err, "Could not encode Buffer ext type")
	}
	err = c.enc.EncodeUint8(TypeBuffer)
	if err != nil {
		return errgo.Notef(err, "Could not encode Buffer type")
	}
	err = c.enc.EncodeUint8(uint8(b.ID))
	if err != nil {
		return errgo.Notef(err, "Could not encode Buffer")
	}
	return nil
}

func (c *Client) decodeWindow() (retVal Window, retErr error) {
	b, err := c.dec.DecodeUint32()
	if err != nil {
		return retVal, errgo.Notef(err, "Could not decode Window")
	}
	return Window{ID: b, client: c}, retErr
}

func (c *Client) encodeWindow(b Window) error {
	err := c.enc.EncodeUint32(b.ID)
	if err != nil {
		return errgo.Notef(err, "Could not encode Window")
	}
	return nil
}

func (c *Client) decodeTabpage() (retVal Tabpage, retErr error) {
	b, err := c.dec.DecodeUint32()
	if err != nil {
		return retVal, errgo.Notef(err, "Could not decode Tabpage")
	}
	return Tabpage{ID: b, client: c}, retErr
}

func (c *Client) encodeTabpage(b Tabpage) error {
	err := c.enc.EncodeUint32(b.ID)
	if err != nil {
		return errgo.Notef(err, "Could not encode Tabpage")
	}
	return nil
}
