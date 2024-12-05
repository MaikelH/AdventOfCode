package common

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
