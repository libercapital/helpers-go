package bavahelper

import (
	"fmt"
	"reflect"
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

func isEquals(element, target string) bool {
	return reflect.DeepEqual(element, target)
}

func TestIncludesCustom(t *testing.T) {
	type args[T any] struct {
		s  []T
		el T
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
				s:  []string{"01", "02", "03"},
				el: "01",
				fn: isEquals,
			},
			want: true,
		},
		{
			name: "test",
			args: args[string]{
				s:  []string{"01", "02", "03"},
				el: "04",
				fn: isEquals,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IncludesCustom(tt.args.s, tt.args.el, tt.args.fn)
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

func TestHasSubslice(t *testing.T) {
	type args[T any] struct {
		slice    []T
		bigslice [][]T
	}
	tests := []struct {
		name string
		args args[int8]
		want bool
	}{
		{
			name: "ReturnsTrueCaseA",
			args: args[int8]{
				slice:    []int8{2, 1, 3},
				bigslice: [][]int8{{1, 2}, {3, 4}},
			},
			want: true,
		},
		{
			name: "ReturnsFalseCaseB",
			args: args[int8]{
				slice:    []int8{4, 1, 5},
				bigslice: [][]int8{{1, 2}, {3, 4}},
			},
			want: false,
		},
		{
			name: "ReturnsTrueCaseC",
			args: args[int8]{
				slice:    []int8{1, 3, 1, 4},
				bigslice: [][]int8{{1, 2, 3, 4, 5}, {1, 1}},
			},
			want: true,
		},
		{
			name: "ReturnsFalseCaseD",
			args: args[int8]{
				slice:    []int8{4, 4, 1},
				bigslice: [][]int8{{1, 2, 3}, {1, 1, 4}},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ContainsSubslice(tt.args.slice, tt.args.bigslice)
			assert.Equal(t, tt.want, got)
		})
	}
}
