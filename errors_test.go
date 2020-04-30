package errors

import (
	"fmt"
	"github.com/golage/errors/codes"
	"github.com/golage/errors/stacktrace"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		code    codes.Code
		message string
		args    []interface{}
	}
	tests := []struct {
		name string
		args args
		want Fundamental
	}{
		{
			name: "must returns fundamental error with message",
			args: args{
				code:    codes.Internal,
				message: "text message",
				args:    nil,
			},
			want: &fundamental{
				code:       codes.Internal,
				message:    "text message",
				stackTrace: stacktrace.Capture(0),
			},
		},
		{
			name: "must returns fundamental error with template and args",
			args: args{
				code:    codes.NotFound,
				message: "%v message",
				args:    []interface{}{"text"},
			},
			want: &fundamental{
				code:       codes.NotFound,
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
	tests := []struct {
		name string
		args args
		want Fundamental
	}{
		{
			name: "must returns fundamental error from built-in error",
			args: args{
				err: fmt.Errorf("text message"),
			},
			want: &fundamental{
				code:    codes.Unknown,
				message: "text message",
			},
		},
		{
			name: "must returns fundamental error from built-in error with fundamental message",
			args: args{
				err: fmt.Errorf("error %d: %v", codes.AlreadyExists, "text message"),
			},
			want: &fundamental{
				code:    codes.AlreadyExists,
				message: "text message",
			},
		},
		{
			name: "must returns fundamental error from fundamental error",
			args: args{
				err: New(codes.Unauthenticated, "text message"),
			},
			want: &fundamental{
				code:       codes.Unauthenticated,
				message:    "text message",
				stackTrace: stacktrace.Capture(0),
			},
		},
		{
			name: "must returns fundamental error from pkg error",
			args: args{
				err: errors.New("text message"),
			},
			want: &fundamental{
				code:       codes.Unknown,
				message:    "text message",
				stackTrace: stacktrace.Capture(0),
			},
		},
		{
			name: "must returns fundamental error from pkg error with fundamental message",
			args: args{
				err: errors.Errorf("error %d: %v", codes.NotFound, "text message"),
			},
			want: &fundamental{
				code:       codes.NotFound,
				message:    "text message",
				stackTrace: stacktrace.Capture(0),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Parse(tt.args.err)
			assert.Equal(t, err.Error(), tt.want.Error())
			assert.Equal(t, err.Message(), tt.want.Message())
			assert.Equal(t, int(err.Code()), int(tt.want.Code()))
			if tt.want.StackTrace() != nil {
				assert.Equal(t, err.StackTrace()[1:], tt.want.StackTrace()[1:])
			}
		})
	}
}

func TestWrap(t *testing.T) {
	type args struct {
		cause   error
		code    codes.Code
		message string
		args    []interface{}
	}
	tests := []struct {
		name string
		args args
		want Fundamental
	}{
		{
			name: "must returns wrapped built-in error with message",
			args: args{
				cause:   fmt.Errorf("error"),
				code:    codes.Unimplemented,
				message: "text message",
			},
			want: &fundamental{
				code:       codes.Unimplemented,
				message:    "text message: error",
				stackTrace: stacktrace.Capture(0),
			},
		},
		{
			name: "must returns wrapped built-in error with template and args",
			args: args{
				cause:   fmt.Errorf("error"),
				code:    codes.Unauthenticated,
				message: "%v message",
				args:    []interface{}{"text"},
			},
			want: &fundamental{
				code:       codes.Unauthenticated,
				message:    "text message: error",
				stackTrace: stacktrace.Capture(0),
			},
		},
		{
			name: "must returns wrapped fundamental error with new code",
			args: args{
				cause:   New(codes.PermissionDenied, "error"),
				code:    codes.InvalidData,
				message: "text message",
			},
			want: &fundamental{
				code:       codes.InvalidData,
				message:    "text message: error",
				stackTrace: stacktrace.Capture(0),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Wrap(tt.args.cause, tt.args.code, tt.args.message, tt.args.args...)
			assert.Equal(t, err.Error(), tt.want.Error())
			assert.Equal(t, err.Message(), tt.want.Message())
			assert.Equal(t, int(err.Code()), int(tt.want.Code()))
			if tt.want.StackTrace() != nil {
				assert.Equal(t, err.StackTrace()[1:], tt.want.StackTrace()[1:])
			}
		})
	}
}
