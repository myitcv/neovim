// Copyright 2014 Paul Jolly <paul@myitcv.org.uk>. All rights reserved.
// Use of s source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package neovim

import (
	"sync"

	"github.com/juju/errors"
)

type syncRespMap struct {
	lock   *sync.Mutex
	theMap map[uint32]*responseHolder
}

func newSyncRespMap() *syncRespMap {
	return &syncRespMap{
		lock:   new(sync.Mutex),
		theMap: make(map[uint32]*responseHolder),
	}
}

func (s *syncRespMap) Put(k uint32, v *responseHolder) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	if _, present := s.theMap[k]; present {
		return errors.Errorf("Key already exists for key %v", k)
	}

	s.theMap[k] = v
	return nil
}

func (s *syncRespMap) Get(k uint32) (res *responseHolder, retErr error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	res, present := s.theMap[k]
	if !present {
		retErr = errors.Errorf("Key does not exist for %v", k)
	} else {
		delete(s.theMap, k)
	}

	return
}

type syncProviderMap struct {
	lock   *sync.Mutex
	theMap map[string]SyncDecoder
}

func newSyncProviderMap() *syncProviderMap {
	return &syncProviderMap{
		lock:   new(sync.Mutex),
		theMap: make(map[string]SyncDecoder),
	}
}

func (s *syncProviderMap) Put(k string, v SyncDecoder) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	if _, present := s.theMap[k]; present {
		return errors.Errorf("Key already exists for key %v", k)
	}

	s.theMap[k] = v
	return nil
}

func (s *syncProviderMap) Get(k string) (res SyncDecoder, retErr error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	res, present := s.theMap[k]
	if !present {
		retErr = errors.Errorf("Key does not exist for %v", k)
	}

	return
}

type asyncProviderMap struct {
	lock   *sync.Mutex
	theMap map[string]AsyncDecoder
}

func newAsyncProviderMap() *asyncProviderMap {
	return &asyncProviderMap{
		lock:   new(sync.Mutex),
		theMap: make(map[string]AsyncDecoder),
	}
}

func (s *asyncProviderMap) Put(k string, v AsyncDecoder) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	if _, present := s.theMap[k]; present {
		return errors.Errorf("Key already exists for key %v", k)
	}

	s.theMap[k] = v
	return nil
}

func (s *asyncProviderMap) Get(k string) (res AsyncDecoder, retErr error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	res, present := s.theMap[k]
	if !present {
		retErr = errors.Errorf("Key does not exist for %v", k)
	}

	return
}
