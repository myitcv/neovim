package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
	install("github.com/myitcv/neovim/example")
}

const (
	_PluginName string = "neovim-go-plugin-manager"
)

type pluginType struct {
	Package, Type string
}

func install(plugins ...string) {
	/*
		we need to collate a list plugins and their exported types
		things to report (and test)
		1. plugins that don't export any types
		2. plugins that don't compile (or pass tests)

		We should output for info the types we have found on a per package basis

		TODO

		1. Test this on windows

		Options

		1. Do you want to use isolated GOPATH - default false

		The steps required to achieve this:

		0. Create a temp directory and set GOPATH to that directory
		1. go get -d -t github.com/myitcv/neovim/example
		2. go test the plugin (will check it compiles)
		3. Parse files that are in the plugin package and extract exported types

		Exit codes should be checked for each
	*/

	t := time.Now()
	tmpDir := fmt.Sprintf("%v/%v_%v", os.TempDir(), _PluginName, t.UnixNano())
	err := os.Mkdir(tmpDir, 0700)
	if err != nil {
		log.Fatalf("Could not create a temp directory: %v\n", err)
	}
	// defer func() {
	// 	log.Printf("Cleaning up temp directory %v\n", tmpDir)
	// 	os.RemoveAll(tmpDir)
	// }()

	log.Printf("Working directory is %v\n", tmpDir)

	envo := make(map[string]string)
	envo["GOPATH"] = tmpDir

	for _, plugin := range plugins {
		log.Printf("Just about to go get -d -t %v\n", plugin)
		c := newCommand(envo, "go", "get", "-d", "-t", plugin)
		err = c.Run()
		if err != nil {
			log.Fatalf("Could not go get plugin: %v\n", err)
		}
		log.Printf("Now running go test %v\n", plugin)
		c = newCommand(envo, "go", "test", plugin)
		err = c.Run()
		if err != nil {
			log.Fatalf("Could not go test plugin: %v\n", err)
		}
	}

}

func newCommand(envo map[string]string, name string, args ...string) *exec.Cmd {
	res := exec.Command(name, args...)

	curEnv := os.Environ()
	newEnv := make([]string, len(envo))

	i := 0
	for k, v := range envo {
		newEnv[i] = fmt.Sprintf("%v=%v", k, v)
		i++
	}
	newEnv = append(newEnv, curEnv...)

	res.Env = newEnv
	return res
}
