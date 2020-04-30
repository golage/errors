package stacktrace

import (
	"github.com/stretchr/testify/assert"
	"runtime"
	"testing"
)

func getValidPC() uintptr {
	pc, _, _, _ := runtime.Caller(0)
	return pc
}

func TestFrame_IsValid(t *testing.T) {
	tests := []struct {
		name  string
		frame Frame
		want  bool
	}{
		{
			name:  "must returns true with valid pointer caller",
			frame: Frame(getValidPC()),
			want:  true,
		},
		{
			name:  "must returns false with invalid pointer caller",
			frame: Frame(1),
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.frame.IsValid(), tt.want)
		})
	}
}

func TestFrame_Name(t *testing.T) {
	tests := []struct {
		name  string
		frame Frame
		want  string
	}{
		{
			name:  "must returns name with valid pointer caller",
			frame: Frame(getValidPC()),
			want:  ".*stacktrace.getValidPC$",
		},
		{
			name:  "must returns empty string with invalid pointer caller",
			frame: Frame(1),
			want:  "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Regexp(t, tt.want, tt.frame.Name())
		})
	}
}

func TestFrame_PC(t *testing.T) {
	pc := getValidPC()
	tests := []struct {
		name  string
		frame Frame
		want  uintptr
	}{
		{
			name:  "must returns frame pointer caller",
			frame: Frame(pc),
			want:  pc - 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.frame.PC(), tt.want)
		})
	}
}

func TestFrame_Source(t *testing.T) {
	tests := []struct {
		name  string
		frame Frame
		want  string
	}{
		{
			name:  "must returns source with valid pointer caller",
			frame: Frame(getValidPC()),
			want:  "\\/stacktrace\\/frame_test.go:10*$",
		},
		{
			name:  "must returns empty string with invalid pointer caller",
			frame: Frame(1),
			want:  "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Regexp(t, tt.want, tt.frame.Source())
		})
	}
}

func TestFrame_String(t *testing.T) {
	tests := []struct {
		name  string
		frame Frame
		want  string
	}{
		{
			name:  "must returns marshal string with valid pointer caller",
			frame: Frame(getValidPC()),
			want:  "^at .*\\/stacktrace.getValidPC in .*\\/stacktrace\\/frame_test.go:10$",
		},
		{
			name:  "must returns empty string with invalid pointer caller",
			frame: Frame(1),
			want:  "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Regexp(t, tt.want, tt.frame.String())
		})
	}
}
