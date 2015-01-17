// Copyright 2014 Paul Jolly <paul@myitcv.org.uk>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package neovim

import "github.com/juju/errors"

func (c *Client) decodeString() (retVal string, retErr error) {
	b, err := c.dec.ReadBytes(nil)
	if err != nil {
		return retVal, errors.Annotatef(err, "Could not decode string raw bytes")
	}
	return string(b), retErr
}

func (c *Client) encodeString(s string) error {
	err := c.enc.WriteBytes([]byte(s))
	if err != nil {
		return errors.Annotatef(err, "Could not encode string raw bytes")
	}
	return nil
}
