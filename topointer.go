package bavahelper

func PtrAny[T any](value T) *T {
	return &value
}
