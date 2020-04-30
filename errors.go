package errors

import (
	"fmt"
	"github.com/golage/errors/codes"
	"github.com/golage/errors/stacktrace"
)

// create new fundamental error
func New(code codes.Code, message string, args ...interface{}) Fundamental {
	return &fundamental{
		code:       code,
		message:    fmt.Sprintf(message, args...),
		stackTrace: stacktrace.Capture(1),
	}
}

// wrap existing error with fundamental error
func Wrap(cause error, code codes.Code, message string, args ...interface{}) Fundamental {
	return &fundamental{
		code:       code,
		message:    fmt.Sprintf("%v: %v", fmt.Sprintf(message, args...), Parse(cause).Message()),
		stackTrace: stacktrace.Capture(1),
	}
}

// parse any errors to fundamental error
func Parse(err error) Fundamental {
	switch err := err.(type) {
	case Fundamental:
		return err
	case stackTracer:
		return parseStackTracer(err)
	default:
		fnd := new(fundamental)
		fnd.Unmarshal(err.Error())
		return fnd
	}
}
