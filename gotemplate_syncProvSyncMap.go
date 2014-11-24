package neovim

import (
	"sync"

	"github.com/juju/errors"
)

// template type SyncMap(K, V)

type syncProvSyncMap struct {
	lock   *sync.Mutex
	theMap map[string]NewSyncDecoder
}

func newsyncProvSyncMap() *syncProvSyncMap {
	return &syncProvSyncMap{
		lock:   new(sync.Mutex),
		theMap: make(map[string]NewSyncDecoder),
	}
}

func (s *syncProvSyncMap) Put(k string, v NewSyncDecoder) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	if _, present := s.theMap[k]; present {
		return errors.Errorf("Key already exists for key %v", k)
	}

	s.theMap[k] = v
	return nil
}

func (s *syncProvSyncMap) Get(k string) (res NewSyncDecoder, retErr error) {
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
