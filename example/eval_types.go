//go:generate msgp
package example

//msgp:tuple MyEvalResult
type MyEvalResult struct {
	S []byte
	I int
}
