package main

import (
	"testing"

	. "gopkg.in/check.v1"
)

type PluginManagerTest struct {
}

func Test(t *testing.T) { TestingT(t) }

var _ = Suite(&PluginManagerTest{})

func (t *PluginManagerTest) SetUpTest(c *C) {
}

func (t *PluginManagerTest) TearDownTest(c *C) {
}

func (t *PluginManagerTest) TestStandard(c *C) {
	install("github.com/myitcv/neovim/example")
}
