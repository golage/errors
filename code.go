package errors

// implements of error code number,
// you can extends this with define constants as code type
type Code uint

const (
	// nil error object, can not create error with this code
	Nil Code = 0

	// unimplemented source code
	Unimplemented Code = 1

	// unknown error raised
	Unknown Code = 2

	// internal server error raised
	Internal Code = 3

	// unauthenticated user request
	Unauthenticated Code = 4

	// user access permission denied
	PermissionDenied Code = 5

	// invalid data o parameters
	InvalidData Code = 6

	// data not found
	NotFound Code = 7

	// same data is already exists
	AlreadyExists Code = 8
)
