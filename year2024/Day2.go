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

func (d Day2) SolvePart1(input string) (int64, error) {
	var reports [][]int
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		numbers := make([]int, len(parts))
		for i, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				return 0, err
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

	return int64(safe), nil
}

func (d Day2) SolvePart2(input string) (int64, error) {
	var reports [][]int
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		numbers := make([]int, len(parts))
		for i, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				return 0, err
			}
			numbers[i] = num
		}

		reports = append(reports, numbers)
	}

	safe := 0
	for _, report := range reports {
		if !common.IsMonotonic(report) {
			// Run problem dampener
			for i := 0; i < len(report); i++ {
				if common.IsMonotonic(common.RemoveAtIndex(report, i)) {
					break
				}
			}
		}
		if !isSafe(report) {
			for i := 0; i < len(report); i++ {
				if isSafe(common.RemoveAtIndex(report, i)) {
					safe++
					break
				}
			}
		} else {
			safe++
		}
	}

	return int64(safe), nil
}

func isSafe(report []int) bool {
	for i := 1; i < len(report); i++ {
		if math.Abs(float64(report[i-1]-report[i])) > 3.0 || report[i-1] == report[i] {
			return false
		}
	}
	return true
}
