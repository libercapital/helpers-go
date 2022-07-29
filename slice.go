package bavahelper

import (
	"reflect"
	"sort"
)

func FindIndex[T1 any](s []T1, fn func(T1) bool) int {
	for index, slice := range s {
		if fn(slice) {
			return index
		}
	}

	return -1
}

func Find[T any](s []T, fn func(T) bool) T {
	for _, slice := range s {
		if fn(slice) {
			return slice
		}
	}
	return *new(T)
}

func Map[T any, T2 any](s []T, fn func(T) T2) []T2 {
	rt := make([]T2, 0)
	for _, slice := range s {
		rt = append(rt, fn(slice))
	}
	return rt
}

func Filter[T any](s []T, fn func(T) bool) []T {
	data := make([]T, 0)

	for _, item := range s {
		if fn(item) {
			data = append(data, item)
		}
	}

	return data
}

func Reduce[T any, T2 any](s []T, fn func(*T2, T) T2, initialElement *T2) T2 {
	out := initialElement

	for _, element := range s {
		val := fn(out, element)
		out = &val
	}

	return *out
}

func ForEach[T any](s []T, fn func(T)) {
	for _, item := range s {
		fn(item)
	}
}

func Reverse[T any](s []T) []T {
	out := make([]T, 0)

	for i := len(s) - 1; i >= 0; i-- {
		out = append(out, s[i])
	}

	return out
}

func Some[T any](s []T, fn func(T) bool) bool {
	for _, item := range s {
		if fn(item) {
			return true
		}
	}

	return false
}

func Every[T any](s []T, fn func(T) bool) bool {
	for _, item := range s {
		if !fn(item) {
			return false
		}
	}

	return true
}

func Includes[T any](s []T, el T) bool {
	for _, item := range s {
		if reflect.DeepEqual(item, el) {
			return true
		}
	}

	return false
}

func Sort[T any](s []T, fn func(T, T) bool) []T {
	sort.SliceStable(s, func(i, j int) bool {
		return fn(s[i], s[j])
	})

	return s
}
