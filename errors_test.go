package errors

import (
	"fmt"
	"github.com/golage/errors/stacktrace"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		code    Code
		message string
		args    []interface{}
	}
	tests := []struct {
		name string
		args args
		want Fundamental
	}{
		{
			name: "must returns fundamental with message",
			args: args{
				code:    Internal,
				message: "text message",
				args:    nil,
			},
			want: &fundamental{
				code:       Internal,
				message:    "text message",
				stackTrace: stacktrace.Capture(0),
			},
		},
		{
			name: "must returns fundamental with template and args",
			args: args{
				code:    NotFound,
				message: "%v message",
				args:    []interface{}{"text"},
			},
			want: &fundamental{
				code:       NotFound,
				message:    "text message",
				stackTrace: stacktrace.Capture(0),
			},
		},
		{
			name: "must returns fundamental with unknown instead of nil code",
			args: args{
				code:    Nil,
				message: "text message",
				args:    nil,
			},
			want: &fundamental{
				code:       Unknown,
				message:    "text message",
				stackTrace: stacktrace.Capture(0),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := New(tt.args.code, tt.args.message, tt.args.args...)
			assert.Equal(t, err.Error(), tt.want.Error())
			assert.Equal(t, err.Message(), tt.want.Message())
			assert.Equal(t, int(err.Code()), int(tt.want.Code()))
			assert.Equal(t, err.StackTrace()[1:], tt.want.StackTrace()[1:])
		})
	}
}

func TestParse(t *testing.T) {
	type args struct {
		err error
	}
	type wants struct {
		err  Fundamental
		code Code
	}
	tests := []struct {
		name  string
		args  args
		wants wants
	}{
		{
			name: "must returns fundamental from built-in",
			args: args{
				err: fmt.Errorf("text message"),
			},
			wants: wants{
				err: &fundamental{
					code:    Unknown,
					message: "text message",
				},
				code: Unknown,
			},
		},
		{
			name: "must returns fundamental from built-in with fundamental marshal",
			args: args{
				err: fmt.Errorf("error %d: %v", AlreadyExists, "text message"),
			},
			wants: wants{
				err: &fundamental{
					code:    AlreadyExists,
					message: "text message",
				},
				code: AlreadyExists,
			},
		},
		{
			name: "must returns fundamental from fundamental",
			args: args{
				err: New(Unauthenticated, "text message"),
			},
			wants: wants{
				err: &fundamental{
					code:       Unauthenticated,
					message:    "text message",
					stackTrace: stacktrace.Capture(0),
				},
				code: Unauthenticated,
			},
		},
		{
			name: "must returns fundamental from pkg-error",
			args: args{
				err: errors.New("text message"),
			},
			wants: wants{
				err: &fundamental{
					code:       Unknown,
					message:    "text message",
					stackTrace: stacktrace.Capture(0),
				},
				code: Unknown,
			},
		},
		{
			name: "must returns fundamental from pkg-error with fundamental marshal",
			args: args{
				err: errors.Errorf("error %d: %v", NotFound, "text message"),
			},
			wants: wants{
				err: &fundamental{
					code:       NotFound,
					message:    "text message",
					stackTrace: stacktrace.Capture(0),
				},
				code: NotFound,
			},
		},
		{
			name: "must returns nil from nil input",
			args: args{
				err: nil,
			},
			wants: wants{
				err:  nil,
				code: Nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err, code := Parse(tt.args.err)
			assert.Equal(t, code, tt.wants.code)
			if tt.wants.err == nil {
				assert.Equal(t, err, nil)
			} else {
				assert.Equal(t, err.Error(), tt.wants.err.Error())
				assert.Equal(t, err.Message(), tt.wants.err.Message())
				assert.Equal(t, int(err.Code()), int(tt.wants.err.Code()))
				if tt.wants.err.StackTrace() != nil {
					assert.Equal(t, err.StackTrace()[1:], tt.wants.err.StackTrace()[1:])
				}
			}
		})
	}
}

func TestWrap(t *testing.T) {
	type args struct {
		cause   error
		code    Code
		message string
		args    []interface{}
	}
	tests := []struct {
		name string
		args args
		want Fundamental
	}{
		{
			name: "must returns nil with nil cause",
			args: args{
				cause:   nil,
				code:    PermissionDenied,
				message: "text message",
			},
			want: nil,
		},
		{
			name: "must returns wrapped built-in error with message",
			args: args{
				cause:   fmt.Errorf("error"),
				code:    Unimplemented,
				message: "text message",
			},
			want: &fundamental{
				code:       Unimplemented,
				message:    "text message: error",
				stackTrace: stacktrace.Capture(0),
			},
		},
		{
			name: "must returns wrapped built-in error with template and args",
			args: args{
				cause:   fmt.Errorf("error"),
				code:    Unauthenticated,
				message: "%v message",
				args:    []interface{}{"text"},
			},
			want: &fundamental{
				code:       Unauthenticated,
				message:    "text message: error",
				stackTrace: stacktrace.Capture(0),
			},
		},
		{
			name: "must returns wrapped fundamental with new code",
			args: args{
				cause:   New(PermissionDenied, "error"),
				code:    InvalidData,
				message: "text message",
			},
			want: &fundamental{
				code:       InvalidData,
				message:    "text message: error",
				stackTrace: stacktrace.Capture(0),
			},
		},
		{
			name: "must returns wrapped error with unknown instead of nil code",
			args: args{
				cause:   fmt.Errorf("error"),
				code:    Nil,
				message: "text message",
			},
			want: &fundamental{
				code:       Unknown,
				message:    "text message: error",
				stackTrace: stacktrace.Capture(0),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Wrap(tt.args.cause, tt.args.code, tt.args.message, tt.args.args...)
			if tt.want == nil {
				assert.Equal(t, err, nil)
			} else {
				assert.Equal(t, err.Error(), tt.want.Error())
				assert.Equal(t, err.Message(), tt.want.Message())
				assert.Equal(t, int(err.Code()), int(tt.want.Code()))
				if tt.want.StackTrace() != nil {
					assert.Equal(t, err.StackTrace()[1:], tt.want.StackTrace()[1:])
				}
			}
		})
	}
}
