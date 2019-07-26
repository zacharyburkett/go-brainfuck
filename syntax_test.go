package brainfuck_test

import (
	"testing"

	"github.com/neuronpool/go-brainfuck"
)

func TestValidateSyntax(t *testing.T) {
	tests := map[string]bool{
		`<>,.[[][[..++--]]]`:  true,
		`<>,.[[][[..++--]]]r`: false,
		`<>,.[[][[..++--]]][`: false,
	}
	for sample, answer := range tests {
		if err := (brainfuck.ValidateSyntax([]byte(sample)) == nil); err != answer {
			t.Errorf("%s returned %t, expected %t", sample, err, answer)
		}
	}
}
