package example

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"testing"

	"github.com/juju/errors"
	"github.com/myitcv/neovim"
	. "gopkg.in/check.v1"
)

type ExampleTest struct {
	client *neovim.Client
	nvim   *exec.Cmd
	plug   *Example
}

func Test(t *testing.T) { TestingT(t) }

var _ = Suite(&ExampleTest{})

func (t *ExampleTest) SetUpTest(c *C) {
	t.nvim = exec.Command("nvim", "-u", "/dev/null")
	t.nvim.Dir = "/tmp"
	client, err := neovim.NewCmdClient(neovim.NullInitMethod, t.nvim, nil)
	if err != nil {
		log.Fatalf("Could not setup client: %v", errors.Details(err))
	}
	client.PanicOnError = true
	t.client = client
	client.Run()

	plug := &Example{}
	err = plug.Init(t.client, log.New(os.Stderr, "", log.LstdFlags))
	if err != nil {
		log.Fatalf("Could not Init plugin: %v\n", err)
	}
	t.plug = plug
}

func (t *ExampleTest) TearDownTest(c *C) {
	err := t.plug.Shutdown()
	if err != nil {
		log.Fatalf("Could not Shutdown plugin: %v\n", err)
	}
	done := make(chan struct{})
	go func() {
		state, err := t.nvim.Process.Wait()
		if err != nil {
			log.Fatalf("Process did not exit cleanly: %v, %v\n", err, state)
		}
		done <- struct{}{}
	}()
	err = t.client.Close()
	if err != nil {
		log.Fatalf("Could not close client: %v\n", err)
	}
	<-done
}

func (t *ExampleTest) TestGetANumber(c *C) {
	t.client.RegisterSyncRequestHandler("GetANumber", t.plug.newGetANumberResponder)
	topic := "GetANumber"
	commandDef := fmt.Sprintf(`call remote#define#FunctionOnChannel(1, "%v", 1, "%v", {})`, topic, topic)
	_ = t.client.Command(commandDef)
	res, _ := t.client.Eval(`GetANumber()`)
	c.Assert(res, Equals, int64(42))
}
