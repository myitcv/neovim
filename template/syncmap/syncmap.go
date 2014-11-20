// Template SyncMap type
package syncmap

import (
	"sync"

	"github.com/juju/errors"
)

// template type SyncMap(K, V)
type K string
type V string

type SyncMap struct {
	lock   *sync.Mutex
	theMap map[K]V
}

func newSyncMap() *SyncMap {
	return &SyncMap{
		lock:   new(sync.Mutex),
		theMap: make(map[K]V),
	}
}

func (s *SyncMap) Put(k K, v V) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	if _, present := s.theMap[k]; present {
		return errors.Errorf("Key already exists for key %v", k)
	}

	s.theMap[k] = v
	return nil
}

func (s *SyncMap) Get(k K) (res V, retErr error) {
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
