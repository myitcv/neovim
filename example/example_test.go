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

func (t *ExampleTest) TestGetTwoNumbers(c *C) {
	t.client.RegisterSyncFunction("GetTwoNumbers", t.plug.newGetTwoNumbersResponder, true, true)
	topic := "GetTwoNumbers"
	commandDef := fmt.Sprintf(`call remote#define#FunctionOnChannel(1, "%v", 1, "%v", {'range': '', 'eval': '["42", 42]'})`, topic, topic)
	_ = t.client.Command(commandDef)
	res_i, _ := t.client.Eval(`GetTwoNumbers(5)`)
	exp := []interface{}{47, "42"}

	res, ok := res_i.([]interface{})
	c.Assert(ok, Equals, true)
	c.Assert(len(res), Equals, len(exp))
	c.Assert(int(res[0].(int64)), Equals, exp[0])
	c.Assert(string(res[1].([]byte)), Equals, exp[1])
}

func (t *ExampleTest) TestDoSomethingAsync(c *C) {
	t.client.RegisterSyncRequestHandler("DoSomethingAsync", t.plug.newGetTwoNumbersResponder, nil)
	topic := "DoSomethingAsync"
	commandDef := fmt.Sprintf(`call remote#define#FunctionOnChannel(1, "%v", 0, "%v", {})`, topic, topic)
	_ = t.client.Command(commandDef)
	ch := make(chan string)
	t.plug.AddDoSomethingAsyncChan(ch)
	t.client.Command(`call DoSomethingAsync("test")`)
	s := <-ch
	c.Assert(s, Equals, "test")
}
