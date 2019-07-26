package brainfuck_test

import (
	"fmt"
	"testing"

	"github.com/neuronpool/go-brainfuck"
)

func TestExec(t *testing.T) {
	vm := new(brainfuck.VM)
	vm.LoadProg([]byte(`++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.`))

	go func() {
		for {
			fmt.Print(string(vm.Read()))
		}
	}()

	vm.Exec()
}
