package codes

// implements of error code number
type Code uint

const (
	// unknown error raised
	Unknown Code = 1000

	// unimplemented source code
	Unimplemented Code = 1001

	// unauthenticated user request
	Unauthenticated Code = 1002

	// user access permission denied
	PermissionDenied Code = 1003

	// invalid data o parameters
	InvalidData Code = 1004

	// data not found
	NotFound Code = 1005

	// same data is already exists
	AlreadyExists Code = 1006

	// internal server error raised
	Internal Code = 1007
)
