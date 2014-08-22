// Copyright 2014 Paul Jolly <paul@myitcv.org.uk>. All rights reserved.
// Use of s source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package neovim

import (
	"sync"

	"github.com/juju/errgo"
)

type syncMap struct {
	lock   *sync.Mutex
	theMap map[uint32]*responseHolder
}

func newSyncMap() *syncMap {
	return &syncMap{
		lock:   new(sync.Mutex),
		theMap: make(map[uint32]*responseHolder),
	}
}

func (s *syncMap) Put(k uint32, v *responseHolder) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	if _, present := s.theMap[k]; present {
		return errgo.Newf("Key already exists for key %v", k)
	}

	s.theMap[k] = v
	return nil
}

func (s *syncMap) Get(k uint32) (res *responseHolder, retErr error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	res, present := s.theMap[k]
	if !present {
		retErr = errgo.Newf("Key does not exist for %v", k)
	} else {
		delete(s.theMap, k)
	}

	return
}
