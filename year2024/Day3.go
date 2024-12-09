package year2024

import (
	"regexp"
	"strconv"
	"strings"
)

type Day3 struct {
}

func (d Day3) SolvePart1(input string) (string, error) {
	regex := regexp.MustCompile("mul\\(\\d+,\\d+\\)")
	matches := regex.FindAllString(input, -1)

	total := 0
	for _, match := range matches {
		match = match[4:]
		match = match[:len(match)-1]
		numbers := strings.Split(match, ",")
		x, err := strconv.Atoi(numbers[0])
		if err != nil {
			return "", err
		}
		y, err := strconv.Atoi(numbers[1])
		if err != nil {
			return "", err
		}

		total += x * y
	}

	return strconv.Itoa(total), nil
}

func (d Day3) SolvePart2(input string) (string, error) {
	regex := regexp.MustCompile("(mul\\(\\d+,\\d+\\))|(don't\\(\\))|(do\\(\\))")
	matches := regex.FindAllStringIndex(input, -1)

	total := 0
	stop := false
	for _, match := range matches {
		if strings.Contains(input[match[0]:match[1]], "don't()") {
			stop = true
			continue
		}
		if strings.Contains(input[match[0]:match[1]], "do()") {
			stop = false
			continue
		}

		if stop {
			continue
		}
		numbers := strings.Split(input[match[0]+4:match[1]-1], ",")
		x, err := strconv.Atoi(numbers[0])
		if err != nil {
			return "", err
		}
		y, err := strconv.Atoi(numbers[1])
		if err != nil {
			return "", err
		}

		total += x * y
	}

	return strconv.Itoa(total), nil
}
