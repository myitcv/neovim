//go:generate msgp
package neovim

//msgp:tuple Range
type Range struct {
	StartLine, EndLine int
}
