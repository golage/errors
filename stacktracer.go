package errors

import (
	"github.com/golage/errors/stacktrace"
	"github.com/pkg/errors"
)

type stackTracer interface {
	error
	StackTrace() errors.StackTrace
}

func parseStackTracer(err stackTracer) Fundamental {
	fnd := new(fundamental)
	fnd.Unmarshal(err.Error())
	for _, pc := range err.StackTrace() {
		frame := stacktrace.Frame(pc)
		if !frame.IsValid() {
			continue
		}
		fnd.stackTrace = append(fnd.stackTrace, frame)
	}
	return fnd
}
