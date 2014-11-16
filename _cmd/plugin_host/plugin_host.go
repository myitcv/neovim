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

	// list of types implementing neovim.Plugin

	// var p1 neovim.Plugin
	// p1 = &neogo.Neogo{}
	// err = p1.Init(client)
	// if err != nil {
	// 	log.Fatalf("Could not Init %v: %v\n", reflect.TypeOf(p1), err)
	// }
	// log.Printf("Successfully call Init on %v\n", reflect.TypeOf(p1))

	// list continues...
	<-client.KillChannel
	log.Printf("Got Kill Channel\n")
}

func (p *pluginHost) DoInit() error {
	log := p.log
	client := p.client

	var p2 neovim.Plugin
	p2 = &example.Example{} // see below
	tp2 := reflect.TypeOf(p2)
	log.Printf("Connecting %v\n", tp2)
	err := p2.Init(client, &pluginLog{l: log, p: tp2.String()})
	if err != nil {
		log.Fatalf("Could not Init %v: %v\n", tp2, err)
	}
	log.Printf("Successfully called Init on %v\n", tp2)

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
