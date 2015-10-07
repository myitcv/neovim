//go:generate msgp
package example

//msgp:tuple MyEvalResult
type MyEvalResult struct {
	S string
	I int
}
