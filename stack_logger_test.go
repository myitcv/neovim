package neovim_test

import (
	"runtime"

	"github.com/myitcv/neovim"
)

type stackLogger struct {
	_log neovim.Logger
}

func newStackLogger(underlying neovim.Logger) neovim.Logger {
	res := &stackLogger{}
	res._log = underlying
	return res
}

func (s *stackLogger) printStack() {
	buf := make([]byte, 1e6)
	i := runtime.Stack(buf, true)
	s._log.Printf("Got SIGQUIT, dumping stacks:\n%v", string(buf[0:i]))
}

func (s *stackLogger) Fatal(v ...interface{}) {
	s.printStack()
	s._log.Fatal(v...)
}
func (s *stackLogger) Fatalf(format string, v ...interface{}) {
	s.printStack()
	s._log.Fatalf(format, v...)
}
func (s *stackLogger) Fatalln(v ...interface{}) {
	s.printStack()
	s._log.Fatalln(v...)
}
func (s *stackLogger) Flags() int {
	return s._log.Flags()
}
func (s *stackLogger) Output(calldepth int, ss string) error {
	s.printStack()
	return s._log.Output(calldepth, ss)
}
func (s *stackLogger) Panic(v ...interface{}) {
	s._log.Panic(v...)
}
func (s *stackLogger) Panicf(format string, v ...interface{}) {
	s._log.Panicf(format, v...)
}
func (s *stackLogger) Panicln(v ...interface{}) {
	s._log.Panicln(v...)
}
func (s *stackLogger) Prefix() string {
	return s._log.Prefix()
}
func (s *stackLogger) Print(v ...interface{}) {
	s.printStack()
	s._log.Print(v...)
}
func (s *stackLogger) Printf(format string, v ...interface{}) {
	s.printStack()
	s._log.Printf(format, v...)
}
func (s *stackLogger) Println(v ...interface{}) {
	s.printStack()
	s._log.Println(v...)
}
func (s *stackLogger) SetFlags(flag int) {
	s._log.SetFlags(flag)
}
func (s *stackLogger) SetPrefix(prefix string) {
	s._log.SetPrefix(prefix)
}
