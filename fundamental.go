package errors

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/golage/errors/stacktrace"
)

var (
	regex = regexp.MustCompile("error ([0-9]*): (.*)")
)

// interface of fundamental errors
type Fundamental interface {
	error

	// returns error code
	Code() Code

	// return error message
	Message() string

	// return error stacktrace
	StackTrace() stacktrace.StackTrace
}

type fundamental struct {
	code       Code
	message    string
	stackTrace stacktrace.StackTrace
}

func (err fundamental) Marshal() string {
	return fmt.Sprintf("error %d: %v", err.code, err.message)
}

func (err *fundamental) Unmarshal(message string) {
	err.code = Unknown
	err.message = message
	matches := regex.FindStringSubmatch(message)
	if len(matches) == 3 {
		if code, e := strconv.Atoi(matches[1]); e == nil {
			err.code = Code(code)
			err.message = matches[2]
		}
	}
}

func (err fundamental) Error() string {
	return err.Marshal()
}

func (err fundamental) Code() Code {
	return err.code
}

func (err fundamental) Message() string {
	return err.message
}

func (err fundamental) StackTrace() stacktrace.StackTrace {
	return err.stackTrace
}
