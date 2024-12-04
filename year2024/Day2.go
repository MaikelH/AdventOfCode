package year2024

import (
	"AoC/common"
	"bufio"
	"math"
	"strconv"
	"strings"
)

type Day2 struct {
}

func (d Day2) SolvePart1(input string) (string, error) {
	var reports [][]int
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		numbers := make([]int, len(parts))
		for i, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				return "", err
			}
			numbers[i] = num
		}

		reports = append(reports, numbers)
	}

	safe := 0
	for _, value := range reports {
		if !common.IsMonotonic(value) {
			continue
		}
		unsafe := false
		for i := 1; i < len(value); i++ {
			if math.Abs(float64(value[i-1]-value[i])) > 3.0 || value[i-1] == value[i] {
				unsafe = true
				break
			}
		}
		if !unsafe {
			safe++
		}
	}

	return strconv.Itoa(safe), nil
}

func (d Day2) SolvePart2(input string) (string, error) {
	return "", nil
}
