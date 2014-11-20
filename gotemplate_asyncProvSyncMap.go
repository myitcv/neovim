package neovim

import (
	"sync"

	"github.com/juju/errors"
)

// template type SyncMap(K, V)

type asyncProvSyncMap struct {
	lock   *sync.Mutex
	theMap map[string]AsyncDecoder
}

func newasyncProvSyncMap() *asyncProvSyncMap {
	return &asyncProvSyncMap{
		lock:   new(sync.Mutex),
		theMap: make(map[string]AsyncDecoder),
	}
}

func (s *asyncProvSyncMap) Put(k string, v AsyncDecoder) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	if _, present := s.theMap[k]; present {
		return errors.Errorf("Key already exists for key %v", k)
	}

	s.theMap[k] = v
	return nil
}

func (s *asyncProvSyncMap) Get(k string) (res AsyncDecoder, retErr error) {
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
