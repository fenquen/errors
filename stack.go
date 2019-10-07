package errors

import (
	"fmt"
	"runtime"
)

// stack represents a stack of program counters.
type stack []Frame

func (stack_ *stack) Format(fmtState fmt.State, verb rune) {
	switch verb {
	case 'v':
		switch {
		//case fmtState.Flag('+'):
		default:
			for _, frame := range *stack_ {
				_, _ = fmt.Fprintf(fmtState, "\n%+v", frame)
			}
		}
	}
}



func callers() *stack {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(3, pcs[:])

	result := pcs[0:n]
	var stack__ stack = make([]Frame, len(result))
	for a := range result {
		stack__[a] = Frame(result[a])
	}
	return &stack__
}
