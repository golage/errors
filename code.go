package errors

// Code implements error code number, you can extends this with define constants with this type
type Code uint

const (
	// CodeNil means no error exists
	CodeNil Code = iota

	// CodeUnimplemented means code does not implemented
	CodeUnimplemented

	// CodeUnknown means unknown error raised
	CodeUnknown

	// CodeInternal means internal module returns error
	CodeInternal

	// CodeUnauthenticated means user does not authenticated
	CodeUnauthenticated

	// CodePermissionDenied means user does not have access permissions
	CodePermissionDenied

	// CodeInvalidData means input parameters are invalid
	CodeInvalidData

	// CodeNotFound means there is no data found
	CodeNotFound

	// CodeAlreadyExists means there is same data is already exists
	CodeAlreadyExists
)
