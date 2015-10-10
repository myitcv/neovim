// Copyright 2014 Paul Jolly <paul@myitcv.org.uk>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package neovim

func (c *Client) decodeDictionary() (retVal map[string]interface{}, retErr error) {
	retVal = make(map[string]interface{})
	retErr = c.dec.ReadMapStrIntf(retVal)
	return retVal, retErr
}

func (c *Client) encodeDictionary(d map[string]interface{}) error {
	return c.enc.WriteMapStrIntf(d)
}
