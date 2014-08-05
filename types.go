package neovim

import (
	"net"
	"sync"
)

type Client struct {
	conn     net.Conn
	func_map map[string]uint32
	next_req uint32
	resp_map *sync_map
	lock     *sync.Mutex
}

type Buffer struct {
	id uint32
}
