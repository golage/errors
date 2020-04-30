package stacktrace

import (
	"github.com/stretchr/testify/assert"
	"runtime"
	"strings"
	"testing"
)

func getStacktrace() StackTrace {
	var st StackTrace
	var pcs [depth]uintptr
	n := runtime.Callers(0, pcs[:])
	for _, pc := range pcs[0:n] {
		st = append(st, Frame(pc))
	}
	return st
}

func TestStackTrace_String(t *testing.T) {
	t.Run("must returns each frame on one line", func(t *testing.T) {
		lines := len(strings.Split(getStacktrace().String(), "\n"))
		assert.Equal(t, lines, 5)
	})
}

func TestCapture(t *testing.T) {
	type args struct {
		skip int
	}
	tests := []struct {
		name string
		args args
		want StackTrace
	}{
		{
			name: "must returns frames of after 0 skip",
			args: args{
				skip: 0,
			},
			want: getStacktrace()[3:],
		},
		{
			name: "must returns frames of after 1 skip",
			args: args{
				skip: 1,
			},
			want: getStacktrace()[4:],
		},
		{
			name: "must returns frames of after 2 skips",
			args: args{
				skip: 2,
			},
			want: getStacktrace()[4:],
		},
		{
			name: "must returns no frame after skip all",
			args: args{
				skip: 3,
			},
		},
		{
			name: "must returns no frame after bigger than all skips",
			args: args{
				skip: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st := Capture(tt.args.skip)
			if len(st) > 1 {
				st = st[1:]
			}
			assert.Equal(t, st, tt.want)
		})
	}
}
