package neovim_test

import (
	"log"
	"sync"
	"testing"

	"github.com/myitcv/neovim"

	. "gopkg.in/check.v1"
	// "github.com/vmihailenco/msgpack"
)

type mp_response struct {
	t      int
	msg_id uint32
	err    interface{}
}

type NeovimTest struct {
	client *neovim.Client
}

func Test(t *testing.T) { TestingT(t) }

var _ = Suite(&NeovimTest{})

func (t *NeovimTest) SetUpTest(c *C) {
	client, err := neovim.NewUnixClient("/tmp/neovim", "unix")
	if err != nil {
		log.Fatalf("Could not setup client: %v", err)
	}
	t.client = client
}

func (t *NeovimTest) TestSimpleClient(c *C) {
	c.Assert(t.client, NotNil)
	ba, err := t.client.VimGetBuffers()
	c.Assert(err, IsNil)
	c.Assert(ba, NotNil)
}

func (t *NeovimTest) TestClient(c *C) {
	c.Assert(t.client, NotNil)
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ba, err := t.client.VimGetBuffers()
			c.Assert(err, IsNil)
			c.Assert(ba, NotNil)
		}()
	}
	wg.Wait()
}

func (t *NeovimTest) TestAPI(c *C) {
	api, err := t.client.API()
	c.Assert(err, IsNil)
	c.Assert(api, NotNil)
}
