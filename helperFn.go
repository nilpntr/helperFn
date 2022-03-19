package helperFn

func ToPointer[T any](value T) *T {
	return &value
}

func Ternary[T any](value *T, fallback T) T {
	if value != nil {
		return *value
	}
	return fallback
}

func TernaryPointer[T any](value *T, fallback *T) *T {
	if value != nil {
		return value
	}
	return fallback
}
