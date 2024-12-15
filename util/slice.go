package util

func FindInSlice[T comparable](slice []T, val T) bool {
	for i := range slice {
		if slice[i] == val {
			return true
		}
	}
	return false
}

func SliceGetMiddleValue[T any](slice *[]T) T {
	return (*slice)[len(*slice)/2]
}
