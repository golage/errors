package stacktrace

import (
	"runtime"
	"strings"
)

const (
	depth = 32
)

// Stacktrace implements is array of frames
type Stacktrace []Frame

// String returns stacktrace marshal string
func (st Stacktrace) String() string {
	var trace []string
	for _, frame := range st {
		trace = append(trace, frame.String())
	}
	return strings.Join(trace, "\n")
}

// Capture returns stacktrace of current line from skip
func Capture(skip int) Stacktrace {
	var pcs [depth]uintptr
	n := runtime.Callers(2+skip, pcs[:])

	var trace Stacktrace
	for _, pc := range pcs[0:n] {
		frame := Frame(pc)
		if !frame.IsValid() {
			continue
		}
		trace = append(trace, frame)
	}
	return trace
}
