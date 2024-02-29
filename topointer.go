package helpers

func PtrAny[T any](value T) *T {
	return &value
}
