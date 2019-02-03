package eval

import (
	"fmt"
	"testing"
)

func TestEval(t *testing.T) {
	code := `
package main

func Main() {
}
`
	err := Eval(code)
	if err != nil {
		fmt.Println(err.Error())
		t.Fail()
	}
}