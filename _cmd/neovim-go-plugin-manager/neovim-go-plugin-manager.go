package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"
	"time"

	_ "golang.org/x/tools/go/gcimporter"
	"golang.org/x/tools/go/types"
)

type Plugins struct {
	Plugins map[string][]string
}

var fInit = flag.Bool("init", false, "create the plugin host directory and corresponding config file")
var fPackage = flag.String("package", "", "the package in which the plugins are defined")

var homeDir = os.Getenv("HOME")
var pluginDir = fmt.Sprintf("%v/.nvim/plugins/go", homeDir)
var configFilePath = fmt.Sprintf("%v/plugins.json", pluginDir)
var pluginHostPath = fmt.Sprintf("%v/neovim-go-plugin-manager", pluginDir)

func main() {
	flag.Parse()
	if *fInit {
		log.Println("Init invoke; all other flags will be ignored")
		doInit()
		return
	}
	if *fPackage == "" {
		log.Fatalf("Need to specify a package")
	}
	newPluginHost := install(*fPackage)
	err := os.Rename(newPluginHost, pluginHostPath)
	if err != nil {
		log.Fatalf("Could not move new plugin host into place: %v\n", err)
	}
}

func doInit() {
	if homeDir == "" {
		log.Fatalf("$HOME environment variable not set")
	}
	log.Printf("HOME set: %v\n", homeDir)
	log.Println("Checking that plugin directory exists")
	pluginDir := fmt.Sprintf("%v/.vim/go", homeDir)
	err := os.MkdirAll(pluginDir, 0755)
	if err != nil {
		log.Fatalf("Error: %v\n", pluginDir, err)
	}
	log.Println("Checking that config file is in place")

	_, err = os.Stat(configFilePath)
	if err != nil {
		// file does not exist
		_, err = os.Create(configFilePath)
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

func install(pkg string) string {
	/*
		TODO

		1. Test this on windows

		Options

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
	err = os.Chdir(tmpDir)
	if err != nil {
		log.Fatalf("Could not change to temp directory: %v\n", err)
	}
	// defer func() {
	// 	log.Printf("Cleaning up temp directory %v\n", tmpDir)
	// 	os.RemoveAll(tmpDir)
	// }()

	log.Printf("Working directory is %v\n", tmpDir)

	envo := make(map[string]string)

	log.Printf("Just about to go get -d -t %v\n", pkg)
	// c := newCommand(envo, "go", "get", "-d", "-t", pkg)
	c := newCommand(envo, "go", "get", pkg)
	// c := newCommand(envo, "go", "version")
	// err = c.Run()
	output, err := c.CombinedOutput()
	fmt.Printf("We have output: %v\n", string(output))
	if err != nil {
		log.Fatalf("Could not go get plugin package: %v\n", string(output))
	}
	log.Printf("Now running go test %v\n", pkg)
	c = newCommand(envo, "go", "test", pkg)
	output, err = c.CombinedOutput()
	fmt.Printf("We have output: %v\n", string(output))
	if err != nil {
		log.Fatalf("Could not go test plugin: %v\n", err)
	}

	// Now we need to get the exported types that implement the neovim.Plugin
	// interface

	fmt.Printf("Now getting the types: %v\n", pkg)
	pluginTypes := getPluginImplementingTypes(pkg)

	realPluginTypes := make([]pluginType, 0)
	for _, t := range pluginTypes {
		elem := pluginType{
			Package: t.Pkg().Name(),
			Type:    t.Name(),
		}
		realPluginTypes = append(realPluginTypes, elem)
	}

	temp := template.New("plugin_host")
	template.Must(temp.Parse(pluginHostTemplate))
	pluginHostOutFile, err := os.OpenFile(tmpDir+"/plugin_host.go", os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("Could not create API out file: %v\n", err)
	}

	err = temp.Execute(pluginHostOutFile, realPluginTypes)
	if err != nil {
		log.Fatalf("Could not write plugin host: %v\n", err)
	}
	c = newCommand(envo, "go", "build", "-o", "neovim-go-plugin-manager")
	output, err = c.CombinedOutput()
	fmt.Printf("We have output: %v\n", string(output))
	if err != nil {
		log.Fatalf("Could not go build plugin host: %v\n", err)
	}

	return tmpDir + "/neovim-go-plugin-manager"

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

func getPluginImplementingTypes(pkg string) []types.Object {
	pp, err := build.Import(pkg, ".", 0)
	if err != nil {
		log.Fatalf("Could not import: %v\n", err)
	}
	fileSet := token.NewFileSet()
	astPkgs, err := parser.ParseDir(fileSet, pp.Dir, func(info os.FileInfo) bool {
		name := info.Name()
		return !info.IsDir() && !strings.HasPrefix(name, ".") && strings.HasSuffix(name, ".go") && !strings.HasSuffix(name, "_test.go")
	}, parser.AllErrors)
	if err != nil {
		log.Fatalf("Could not parse dir: %v\n", err)
	}

	exPkg := astPkgs["example"]
	files := make([]*ast.File, 0)
	for _, v := range exPkg.Files {
		files = append(files, v)
	}

	p, err := types.Check(pp.Dir, fileSet, files)
	if err != nil {
		log.Fatalf("Could not type check!: %v\n", err)
	}

	var nvim *types.Package
	for _, v := range p.Imports() {
		if v.Path() == "github.com/myitcv/neovim" {
			nvim = v
			break
		}
	}

	if nvim == nil {
		log.Fatalf("Could not get neovim package")
	}

	pi := getPluginInterface(nvim)
	if pi == nil {
		log.Fatalf("Could not get plugin interface!")
	}

	res := make([]types.Object, 0)

	scope := p.Scope()
	for _, n := range scope.Names() {
		obj := scope.Lookup(n)

		if types.Implements(types.NewPointer(obj.Type()), pi) && obj.Exported() {
			res = append(res, obj)
		}
	}

	return res
}

func getPluginInterface(p *types.Package) *types.Interface {
	scope := p.Scope()
	for _, n := range scope.Names() {
		obj := scope.Lookup(n)

		if typ, ok := obj.(*types.TypeName); ok {
			if inf, ok := typ.Type().Underlying().(*types.Interface); ok {
				if typ.Name() == "Plugin" {
					return inf
				}
			}
		}
	}
	return nil
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
	"github.com/myitcv/neovim/example"

	// list continues...
)

type pluginHost struct {
	client *neovim.Client
	log    neovim.Logger
}

func main() {
	transport := &neovim.StdWrapper{Stdin: os.Stdout, Stdout: os.Stdin}

	// pid := os.Getpid()
	// TODO enable unique file name
	// logFileName := fmt.Sprintf("/tmp/neovim_go_plugin_host_%v", pid)
	logFileName := fmt.Sprintf("/tmp/neovim_go_plugin_host")
	logFile, err := os.Create(logFileName)
	if err != nil {
		_log.Fatalf("Could not create log file %v: %v\n", logFileName, err)
	}
	log := _log.New(logFile, "", _log.Llongfile|_log.Ldate|_log.Ltime)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGQUIT)

	// to help with debugging
	// all this process to receive SIGQUIT and dump the stack
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

	host := &pluginHost{}

	client, err := neovim.NewClient(host.DoInit, transport, log)
	if err != nil {
		log.Fatalf("Could not connect to Neovim: %v\n", err)
	}
	host.client = client
	host.log = log
	client.Run()

	// TODO we could move the plugin init to an init-like method
	// that is called when we call plugin_load from nvim

	log.Println("Successfully connected to Neovim")

	<-client.KillChannel
	log.Printf("Got Kill Channel\n")
}

func (p *pluginHost) DoInit() error {
	log := p.log
	client := p.client
	var err error

	{{range $index, $element := .}}
	var p{{$index}} neovim.Plugin
	p{{$index}} = &{{$element.Package}}.{{$element.Type}}{} // see below
	tp{{$index}} := reflect.TypeOf(p{{$index}})
	log.Printf("Connecting %v\n", tp{{$index}})
	err = p{{$index}}.Init(client, &pluginLog{l: log, p: tp{{$index}}.String()})
	if err != nil {
		log.Fatalf("Could not Init %v: %v\n", tp{{$index}}, err)
	}
	log.Printf("Successfully called Init on %v\n", tp{{$index}})
	{{end}}

	// TODO we should return an error here
	return nil
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
