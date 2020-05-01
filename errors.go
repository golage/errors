package errors

import (
	"fmt"
	"github.com/golage/errors/stacktrace"
)

// create new fundamental error
func New(code Code, message string, args ...interface{}) Fundamental {
	if code == Nil {
		code = Unknown
	}
	return &fundamental{
		code:       code,
		message:    fmt.Sprintf(message, args...),
		stackTrace: stacktrace.Capture(1),
	}
}

// wrap existing error with fundamental error
func Wrap(cause error, code Code, message string, args ...interface{}) Fundamental {
	parsed, _ := Parse(cause)
	if parsed == nil {
		return nil
	}
	if code == Nil {
		code = Unknown
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

// parse any errors to fundamental error
func Parse(err error) (Fundamental, Code) {
	switch err := err.(type) {
	case nil:
		return nil, Nil
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
