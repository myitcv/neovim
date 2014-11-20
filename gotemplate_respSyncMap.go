package neovim

import (
	"sync"

	"github.com/juju/errors"
)

// template type SyncMap(K, V)

type respSyncMap struct {
	lock   *sync.Mutex
	theMap map[uint32]*responseHolder
}

func newrespSyncMap() *respSyncMap {
	return &respSyncMap{
		lock:   new(sync.Mutex),
		theMap: make(map[uint32]*responseHolder),
	}
}

func (s *respSyncMap) Put(k uint32, v *responseHolder) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	if _, present := s.theMap[k]; present {
		return errors.Errorf("Key already exists for key %v", k)
	}

	s.theMap[k] = v
	return nil
}

func (s *respSyncMap) Get(k uint32) (res *responseHolder, retErr error) {
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
