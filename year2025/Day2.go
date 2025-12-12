package year2025

import (
	"AoC/common"
	"strconv"
	"strings"
)

type Day2 struct {
}

func (d Day2) SolvePart1(input string) (int64, error) {
	ranges := strings.Split(input, ",")
	ids := make([][]int, len(ranges))
	for i, r := range ranges {
		split := strings.Split(r, "-")
		start, _ := strconv.Atoi(split[0])
		end, _ := strconv.Atoi(split[1])
		ids[i] = []int{start, end}
	}

	sum := int64(0)
	for _, startEnd := range ids {
		for i := startEnd[0]; i <= startEnd[1]; i++ {
			idString := strconv.Itoa(i)
			if len(idString)%2 == 0 {
				middle := len(idString) / 2
				start := idString[0:middle]
				end := idString[middle:]
				if start == end {
					sum += int64(i)
				}
			}
		}
	}

	return sum, nil
}

func (d Day2) SolvePart2(input string) (int64, error) {
	ranges := strings.Split(input, ",")
	ids := make([][]int, len(ranges))
	for i, r := range ranges {
		split := strings.Split(r, "-")
		start, _ := strconv.Atoi(split[0])
		end, _ := strconv.Atoi(split[1])
		ids[i] = []int{start, end}
	}

	sum := int64(0)
	for _, startEnd := range ids {
		for i := startEnd[0]; i <= startEnd[1]; i++ {
			idString := strconv.Itoa(i)
			length := len(idString)
			divisors := common.Divisors(length)
			found := false

			for _, divisor := range divisors {
				if divisor == 1 || found {
					continue
				}
				parts := make(map[string]struct{}, divisor+1)
				cutLength := length / divisor
				for j := 0; j < len(idString); j = j + cutLength {
					parts[idString[j:j+cutLength]] = struct{}{}
				}

				if len(parts) == 1 {
					sum += int64(i)
					found = true
				}
			}
		}
	}

	return sum, nil
}
