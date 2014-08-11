package neovim

import (
	"sync"

	"github.com/juju/errgo"
)

type sync_map struct {
	lock    *sync.Mutex
	the_map map[uint32]*response_holder
}

func newSyncMap() *sync_map {
	return &sync_map{
		lock:    new(sync.Mutex),
		the_map: make(map[uint32]*response_holder),
	}
}

func (this *sync_map) Put(k uint32, v *response_holder) error {
	this.lock.Lock()
	defer this.lock.Unlock()

	if _, present := this.the_map[k]; present {
		return errgo.Newf("Key already exists for key %v", k)
	}

	this.the_map[k] = v
	return nil
}

func (this *sync_map) Get(k uint32) (*response_holder, error) {
	this.lock.Lock()
	defer this.lock.Unlock()

	if res, present := this.the_map[k]; !present {
		return nil, errgo.Newf("Key does not exist for %v", k)
	} else {
		delete(this.the_map, k)
		return res, nil
	}
}
