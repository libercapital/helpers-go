package bavahelper

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert/v2"
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

func TestFilter(t *testing.T) {
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
			want: []string{"01"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Filter(tt.args.s, tt.args.fn)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestReduce(t *testing.T) {
	type args[T any, T2 any] struct {
		s  []T
		fn func(*T2, T) T2
	}
	tests := []struct {
		name string
		args args[int, int]
		want interface{}
	}{
		{
			name: "test",
			args: args[int, int]{
				s: []int{1, 2, 3},
				fn: func(o *int, i int) int {
					if o != nil {
						return *o + i
					}

					return i
				},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Reduce(tt.args.s, tt.args.fn, nil)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestReverse(t *testing.T) {
	type args[T any] struct {
		s []T
	}
	tests := []struct {
		name string
		args args[string]
		want interface{}
	}{
		{
			name: "test",
			args: args[string]{
				s: []string{"01", "02", "03"},
			},
			want: []string{"03", "02", "01"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Reverse(tt.args.s)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSome(t *testing.T) {
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
			want: true,
		},
		{
			name: "test",
			args: args[string]{
				s:  []string{"01", "02", "03"},
				fn: func(s string) bool { return s == "04" },
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Some(tt.args.s, tt.args.fn)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestEvery(t *testing.T) {
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
			want: false,
		},
		{
			name: "test",
			args: args[string]{
				s:  []string{"01", "01", "01"},
				fn: func(s string) bool { return s == "01" },
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Every(tt.args.s, tt.args.fn)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestIncludes(t *testing.T) {
	type args[T any] struct {
		s []T
		a T
	}
	tests := []struct {
		name string
		args args[string]
		want interface{}
	}{
		{
			name: "test",
			args: args[string]{
				s: []string{"01", "02", "03"},
				a: "01",
			},
			want: true,
		},
		{
			name: "test",
			args: args[string]{
				s: []string{"01", "02", "03"},
				a: "04",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Includes(tt.args.s, tt.args.a)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSort(t *testing.T) {
	type args[T any] struct {
		s  []T
		fn func(T, T) bool
	}
	tests := []struct {
		name string
		args args[string]
		want interface{}
	}{
		{
			name: "test",
			args: args[string]{
				s:  []string{"04", "03", "02"},
				fn: func(a, b string) bool { return a < b },
			},
			want: []string{"02", "03", "04"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Sort(tt.args.s, tt.args.fn)
			assert.Equal(t, tt.want, got)
		})
	}
}
