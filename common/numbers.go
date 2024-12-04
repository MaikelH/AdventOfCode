package common

type Number interface {
	int | float32 | float64
}

// IsMonotonic tests if the contents of the slice is monotonic
func IsMonotonic[T Number](slice []T) bool {
	isIncreasing := true
	isDecreasing := true

	for i := 1; i < len(slice); i++ {
		if slice[i] > slice[i-1] {
			isDecreasing = false
		}
		if slice[i] < slice[i-1] {
			isIncreasing = false
		}
	}

	return isIncreasing || isDecreasing
}
