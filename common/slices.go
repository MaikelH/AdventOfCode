package common

import "slices"

func CountFrequency(numbers []int) map[int]int {
	freqMap := make(map[int]int)
	for _, num := range numbers {
		freqMap[num]++
	}
	return freqMap
}

func RemoveAtIndex[T any](slice []T, index int) []T {
	newSlice := make([]T, 0, len(slice)-1)
	newSlice = append(newSlice, slice[:index]...)
	newSlice = append(newSlice, slice[index+1:]...)

	return newSlice
}

func ContainsSlice[T comparable](mainSlice, sub []T) bool {
	if len(sub) == 0 {
		return true
	}
	if len(sub) > len(mainSlice) {
		return false
	}
	for i := 0; i <= len(mainSlice)-len(sub); i++ {
		if slices.Equal(mainSlice[i:i+len(sub)], sub) {
			return true
		}
	}
	return false
}

func GetColumns[T any](data [][]T) [][]T {
	columns := make([][]T, len(data[0]))
	for i := range columns {
		columns[i] = make([]T, len(data))
		for j := range columns[i] {
			columns[i][j] = data[j][i]
		}
	}
	return columns
}
