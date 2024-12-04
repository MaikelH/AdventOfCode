package common

func CountFrequency(numbers []int) map[int]int {
	freqMap := make(map[int]int)
	for _, num := range numbers {
		freqMap[num]++
	}
	return freqMap
}
