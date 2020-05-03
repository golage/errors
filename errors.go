package errors

import (
	"fmt"
	"github.com/golage/errors/stacktrace"
)

// New returns new instance fundamental error with args
func New(code Code, message string, args ...interface{}) Fundamental {
	if code == CodeNil {
		return nil
	}
	return &fundamental{
		code:       code,
		message:    fmt.Sprintf(message, args...),
		stackTrace: stacktrace.Capture(1),
	}
}

// Wrap returns new instance fundamental error with args from error cause
func Wrap(cause error, code Code, message string, args ...interface{}) Fundamental {
	parsed, _ := Parse(cause)
	if parsed == nil {
		return nil
	}
	if code == CodeNil {
		return nil
	}
	fnd := &fundamental{
		code:       code,
		message:    fmt.Sprintf("%v: %v", fmt.Sprintf(message, args...), parsed.Message()),
		stackTrace: parsed.StackTrace(),
	}
	if fnd.stackTrace == nil {
		fnd.stackTrace = stacktrace.Capture(1)
	}
	return fnd
}

// Parse returns fundamental error and code from all of error types
func Parse(err error) (Fundamental, Code) {
	switch err := err.(type) {
	case nil:
		return nil, CodeNil
	case Fundamental:
		return err, err.Code()
	case stackTracer:
		fnd := parseStackTracer(err)
		return fnd, fnd.Code()
	default:
		fnd := new(fundamental)
		fnd.Unmarshal(err.Error())
		return fnd, fnd.code
	}
}
