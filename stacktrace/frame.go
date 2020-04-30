package stacktrace

import (
	"fmt"
	"runtime"
)

// implements of caller pointer
type Frame uintptr

// returns pointer number
func (f Frame) PC() uintptr {
	return uintptr(f) - 1
}

// returns true if frame pointer has a source function
func (f Frame) IsValid() bool {
	return runtime.FuncForPC(f.PC()) != nil
}

// returns function name
func (f Frame) Name() string {
	fn := runtime.FuncForPC(f.PC())
	if fn == nil {
		return ""
	}
	return fn.Name()
}

// return function source line
func (f Frame) Source() string {
	fn := runtime.FuncForPC(f.PC())
	if fn == nil {
		return ""
	}
	file, line := fn.FileLine(f.PC())
	return fmt.Sprintf("%v:%v", file, line)
}

// returns frame as a string
func (f Frame) String() string {
	fn := runtime.FuncForPC(f.PC())
	if fn == nil {
		return ""
	}
	return fmt.Sprintf("at %v in %v", f.Name(), f.Source())
}
