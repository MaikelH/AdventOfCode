package common

import (
	"math"
	"sort"
)

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

// Divisors returns a slice containing all divisors of a given integer n, sorted in ascending order.
// If n is 0, it returns an empty slice. Includes both positive and negative divisors if n is negative.
func Divisors(n int) []int {
	if n == 0 {
		return []int{}
	}
	m := n
	if m < 0 {
		m = -m
	}
	small := []int{}
	large := []int{}

	limit := int(math.Sqrt(float64(m)))
	for i := 1; i <= limit; i++ {
		if m%i == 0 {
			small = append(small, i)
			if i != m/i { // avoid duplicate when i*i == m
				large = append(large, m/i)
			}
		}
	}

	// Merge: small is ascending; large is descending, so reverse large
	for i, j := 0, len(large)-1; i < j; i, j = i+1, j-1 {
		large[i], large[j] = large[j], large[i]
	}
	res := append(small, large...)
	// Optional: sort to be safe if order matters
	sort.Ints(res)
	return res
}
