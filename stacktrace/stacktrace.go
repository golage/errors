package stacktrace

import (
	"runtime"
	"strings"
)

const (
	depth = 32
)

// implement stacktrace as array of frames
type StackTrace []Frame

// marshal stacktrace as a string
func (st StackTrace) String() string {
	var trace []string
	for _, frame := range st {
		trace = append(trace, frame.String())
	}
	return strings.Join(trace, "\n")
}

// capture stacktrace of current line
func Capture(skip int) StackTrace {
	var pcs [depth]uintptr
	n := runtime.Callers(2+skip, pcs[:])

	var trace StackTrace
	for _, pc := range pcs[0:n] {
		frame := Frame(pc)
		if !frame.IsValid() {
			continue
		}
		trace = append(trace, frame)
	}
	return trace
}
