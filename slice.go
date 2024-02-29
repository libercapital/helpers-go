package helpers

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

func Find[T any](s []T, fn func(T) bool) *T {
	for _, slice := range s {
		if fn(slice) {
			return &slice
		}
	}
	return nil
}

func Map[T any, T2 any](s []T, fn func(T) T2) []T2 {
	rt := make([]T2, len(s))
	for i, slice := range s {
		rt[i] = fn(slice)
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

func IncludesCustom[T any](s []T, el T, fn func(T, T) bool) bool {
	for _, item := range s {
		if fn(item, el) {
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

/*
Checks if a slice contains another slice from a group of slices, ignoring the order of its elements. Ex:

ContainsSubslice([]string{"b","a","c"}, [][]string{{"a","b"},{"c","d"}}) // returns true, because {"b","a","c"} has all the elements from {"a","b"} of the group.

ContainsSubslice([]string{"d","a","e"}, [][]string{{"a","b"},{"c","d"}}) // returns false, because {"d","a","e"} hasn't all the elements from some slices of the group {{"a","b"},{"c","d"}}.

ContainsSubslice([]int{1,3,1,4}, [][]int{{1,2,3,4,5},{1,1}}) // returns true, because {1,3,1,4} has all the elements from {1,1} of the group.

ContainsSubslice([]int{4,4,1}, [][]int{{1,2,3},{1,1,4}}) // returns false, because {4,4,1} hasn't all the elements from some slices of the group {{1,2,3},{1,1,4}}.
*/
func ContainsSubslice[T any](slice []T, bigslice [][]T) bool {
	for _, subslice := range bigslice {
		if IsSubslice(subslice, slice) {
			return true
		}
	}

	return false
}

/*
Checks if a slice contains another, ignoring the order of its elements. Ex:

IsSubslice([]string{"a","b"}, []string{"b","a","c"}) // returns true, because {"b","a","c"} has all the elements from {"a","b"}.

IsSubslice([]string{"c","d"}, []string{"b","a","c"}) // returns false, because {"b","a","c"} hasn't all the elements from {"c","d"}.

IsSubslice([]int{1,1}, []int{1,3,1,4}) // returns true, because {1,3,1,4} has all the elements from {1,1}.

IsSubslice([]int{1,1,4}, []int{4,4}) // returns false, because {4,4} hasn't all the elements from {1,1,4}.
*/
func IsSubslice[T any](subslice, slice []T) bool {
	if len(subslice) > len(slice) {
		return false
	}

	finalSubslice := make([]T, len(subslice))
	copy(finalSubslice, subslice)

	finalSlice := make([]T, len(slice))
	copy(finalSlice, slice)

	incr := 0
	for _, elem := range subslice {
		if index := contains(finalSlice, elem); index > -1 {
			finalSubslice = append(finalSubslice[:incr], finalSubslice[incr+1:]...)
			finalSlice = append(finalSlice[:index], finalSlice[index+1:]...)
		} else {
			incr++
		}
	}

	return len(finalSubslice) == 0
}

// Returns the index of the position of an element in a slice. When the element isn't in the slice, returns -1.
func contains[T any](slice []T, el T) int {
	for index, item := range slice {
		if reflect.DeepEqual(item, el) {
			return index
		}
	}
	return -1
}
