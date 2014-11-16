package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

type Plugins struct {
	Plugins map[string][]string
}

var fInit = flag.Bool("init", false, "create the plugin host directory and corresponding config file")
var fPackage = flag.String("package", "", "the package in which the plugins are defined")

var homeDir = os.Getenv("HOME")
var pluginDir = fmt.Sprintf("%v/.vim/go", homeDir)
var configFilePath = fmt.Sprintf("%v/plugins.json", pluginDir)

func main() {
	flag.Parse()
	if *fInit {
		log.Println("Init invoke; all other flags will be ignored")
		doInit()
		return
	}
	// if *fPackage == "" {
	// 	log.Fatalf("Need to specify a package")
	// }
	// if len(flag.Args()) == 0 {
	// 	log.Fatalf("Need to specify at least one plugin type to install")
	// }
	// install(*fPackage, flag.Args())
}

func doInit() {
	if homeDir == "" {
		log.Fatalf("$HOME environment variable not set")
	}
	log.Printf("HOME set: %v\n", homeDir)
	log.Println("Checking that plugin directory exists")
	pluginDir := fmt.Sprintf("%v/.vim/go", homeDir)
	pd, err := os.MkdirAll(pluginDir, 0755)
	if err != nil {
		log.Fatalf("Error: %v\n", pluginDir, err)
	}
	log.Println("Checking that config file is in place")

	statConfigFile, err := os.Stat(configFilePath)
	if err != nil {
		// file does not exist
		configFile, err := os.Create(configFilePath)
		if err != nil {
			log.Fatalf("Could not create config file at %v: %v\n", configFilePath, err)
		}
	}

}

const (
	_PluginName string = "neovim-go-plugin-manager"
)

type pluginType struct {
	Package, Type string
}

func install(pkg string, plugins ...string) {
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

	log.Printf("Just about to go get -d -t %v\n", pkg)
	c := newCommand(envo, "go", "get", "-d", "-t", pkg)
	err = c.Run()
	if err != nil {
		log.Fatalf("Could not go get plugin package: %v\n", err)
	}
	log.Printf("Now running go test %v\n", pkg)
	c = newCommand(envo, "go", "test", pkg)
	err = c.Run()
	if err != nil {
		log.Fatalf("Could not go test plugin: %v\n", err)
	}

	// TODO at this point it would be good to verify that the types provided
	// are in fact part of this package and furthermore that they implement
	// the neovim.Plugin interface

	// check that the plugin directory exists
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

var pluginHostTemplate = `
package main

import (
	"fmt"
	_log "log"
	"os"
	"os/signal"
	"reflect"
	"runtime"
	"syscall"

	"github.com/myitcv/neovim"
	{{range .Packages }}
	{{.}}
	{{end}}
)

func main() {
	transport := &neovim.StdWrapper{Stdin: os.Stdout, Stdout: os.Stdin}

	pid := os.Getpid()
	logFileName := fmt.Sprintf("/tmp/neovim_go_plugin_host_%v", pid)
	logFile, err := os.Create(logFileName)
	if err != nil {
		_log.Fatalf("Could not create log file %v: %v\n", logFileName, err)
	}
	log := _log.New(logFile, "", _log.Llongfile|_log.Ldate|_log.Ltime)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGQUIT)

	go func() {
		for {
			select {
			case <-sig:
				buf := make([]byte, 1e6)
				i := runtime.Stack(buf, true)
				log.Printf("Got SIGQUIT, dumping stacks:\n%v", string(buf[0:i]))
			}
		}
	}()

	client, err := neovim.NewClient(transport, log)
	if err != nil {
		log.Fatalf("Could not connect to Neovim: %v\n", err)
	}

	log.Println("Successfully connected to Neovim")

	// list of types implementing neovim.Plugin

	{{range $index, $element := .Types}}
	var p{{$index}} neovim.Plugin
	p{{$index}} = &example.Example{} // see below
	tp{{$index}} := reflect.TypeOf(p{{$index}})
	log.Printf("Connecting %v\n", tp{{$index}})
	p{{$index}}.Init(client, &pluginLog{l: log, p: tp{{$index}}.String()})
	if err != nil {
		log.Fatalf("Could not Init %v: %v\n", tp{{$index}}, err)
	}
	log.Printf("Successfully called Init on %v\n", tp{{$index}})
	{{end}}

	// list continues...
	<-client.KillChannel
	log.Printf("Got Kill Channel\n")
}

type pluginLog struct {
	l neovim.Logger
	p string
}

func newPluginLog(l neovim.Logger, p string) neovim.Logger {
	res := &pluginLog{l: l, p: "[" + p + "] "}
	return res
}

func genArgs(i interface{}, v ...interface{}) []interface{} {
	args := []interface{}{i}
	args = append(args, v...)
	return args
}

func (p *pluginLog) Fatal(v ...interface{}) {
	p.l.Fatal(genArgs(p.p, v))
}

func (p *pluginLog) Fatalf(format string, v ...interface{}) {
	p.l.Fatalf(p.p+format, v...)
}

func (p *pluginLog) Fatalln(v ...interface{}) {
	p.l.Fatalln(genArgs(p.p, v))
}

func (p *pluginLog) Flags() int {
	return p.l.Flags()
}

func (p *pluginLog) Output(calldepth int, s string) error {
	return p.l.Output(calldepth, p.p+s)
}

func (p *pluginLog) Panic(v ...interface{}) {
	p.l.Panic(genArgs(p.p, v))
}

func (p *pluginLog) Panicf(format string, v ...interface{}) {
	p.l.Panicf(p.p+format, v...)
}

func (p *pluginLog) Panicln(v ...interface{}) {
	p.l.Panicln(genArgs(p.p, v))
}

func (p *pluginLog) Prefix() string {
	// TODO this might want to include the p.p
	panic("Not supported")
	return p.l.Prefix()
}

func (p *pluginLog) Print(v ...interface{}) {
	p.l.Print(genArgs(p.p, v))
}

func (p *pluginLog) Printf(format string, v ...interface{}) {
	p.l.Printf(p.p+format, v...)
}

func (p *pluginLog) Println(v ...interface{}) {
	p.l.Println(genArgs(p.p, v))
}

func (p *pluginLog) SetFlags(flag int) {
	panic("Not supported")
}

func (p *pluginLog) SetPrefix(prefix string) {
	panic("Not supported")
}
`
