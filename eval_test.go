package eval

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEval(t *testing.T) {
	code := `
package main

func main() {
	fmt.Println("Hello world")
}
`
	err := Eval(code)
	assert.NoError(t, err)
}
