package bavahelper

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

func Map[T any](s []T, fn func(T) T) []T {
	rt := make([]T, 0)
	for _, slice := range s {
		rt = append(rt, fn(slice))
	}
	return rt
}
