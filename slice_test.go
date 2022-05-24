package bavahelper

import (
	"fmt"
	"testing"
)

func TestFind(t *testing.T) {
	type args[T any] struct {
		s  []T
		fn func(T) bool
	}
	tests := []struct {
		name string
		args args[string]
		want interface{}
	}{
		{
			name: "test",
			args: args[string]{
				s:  []string{"01", "02", "03"},
				fn: func(s string) bool { return s == "01" },
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Find(tt.args.s, tt.args.fn)

			fmt.Println(got)
		})
	}
}
