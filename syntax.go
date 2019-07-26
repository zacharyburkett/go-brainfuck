package brainfuck

import (
	"errors"
	"fmt"
	"unicode"
)

// ============================================================
// Types and globals
// ============================================================

var tokens = map[byte]bool{
	'>': true,
	'<': true,
	'+': true,
	'-': true,
	'.': true,
	',': true,
	'[': true,
	']': true,
}

// ============================================================
// Functions
// ============================================================

// ValidateSyntax takes in a byte slice
// of brainfuck commands, and returns an error
// if the syntax is invalid. Invalid syntax can
// include 0 length programs, unknown characters,
// or unclosed loops.
func ValidateSyntax(prog []byte) error {
	if len(prog) == 0 {
		return errors.New("0 length program")
	}

	var stack []byte
	for i, b := range prog {
		// Skip whitespaces
		if unicode.IsSpace(rune(b)) {
			continue
		}

		if _, ok := tokens[b]; !ok {
			return fmt.Errorf("Char %d: Invalid token: %s", i, string(b))
		}

		// Check brackets with stack
		if b != '[' && b != ']' {
			continue
		}
		if len(stack) == 0 {
			stack = append(stack, b)
			continue
		}
		last := stack[len(stack)-1]
		if last == b {
			stack = append(stack, b)
			continue
		}
		if last == '[' && last != b {
			stack = stack[:len(stack)-1]
			continue
		}
	}

	if len(stack) != 0 {
		return errors.New("Brack mismatch")
	}
	return nil
}
