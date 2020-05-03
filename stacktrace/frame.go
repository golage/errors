package stacktrace

import (
	"fmt"
	"runtime"
)

// Frame implements of caller pointer
type Frame uintptr

// PC returns frame caller pointer
func (f Frame) PC() uintptr {
	return uintptr(f) - 1
}

// IsValid returns true if frame has source function
func (f Frame) IsValid() bool {
	return runtime.FuncForPC(f.PC()) != nil
}

// Name returns frame function name
func (f Frame) Name() string {
	fn := runtime.FuncForPC(f.PC())
	if fn == nil {
		return ""
	}
	return fn.Name()
}

// Source returns function source line
func (f Frame) Source() string {
	fn := runtime.FuncForPC(f.PC())
	if fn == nil {
		return ""
	}
	file, line := fn.FileLine(f.PC())
	return fmt.Sprintf("%v:%v", file, line)
}

// String returns frame string marshal
func (f Frame) String() string {
	fn := runtime.FuncForPC(f.PC())
	if fn == nil {
		return ""
	}
	return fmt.Sprintf("at %v in %v", f.Name(), f.Source())
}
