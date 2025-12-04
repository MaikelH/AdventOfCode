package year2025

import (
	"bufio"
	"strconv"
	"strings"
)

type Day1 struct {
}

func (d Day1) SolvePart1(input string) (int64, error) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	dial := 50
	zeros := 0
	for scanner.Scan() {
		line := scanner.Text()
		steps, err := strconv.Atoi(line[1:])
		if err != nil {
			return 0, err
		}
		if strings.Contains(line, "R") {
			dial = dial + steps
			if dial > 100 {
				dial = dial % 100
			}
		} else {
			dial = dial - steps
			if dial < 0 {
				rem := dial % 100
				dial = 100 + rem
			}
		}
		if dial == 100 {
			dial = 0
		}
		if dial == 0 {
			zeros++
		}
	}

	return int64(zeros), nil
}

func (d Day1) SolvePart2(input string) (int64, error) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	dial := 50
	zeros := 0
	for scanner.Scan() {
		line := scanner.Text()
		steps, err := strconv.Atoi(line[1:])
		if err != nil {
			return 0, err
		}
		zeros += steps / 100
		steps = steps % 100
		if strings.Contains(line, "R") {
			dial = dial + steps
			if dial > 100 {
				dial = dial % 100
				zeros++
			}
		} else {
			dial = dial - steps
			if dial < 0 {
				if dial+steps > 0 {
					zeros++
				}
				rem := dial % 100
				dial = 100 + rem
			}
		}
		if dial == 100 {
			dial = 0
		}
		if dial == 0 {
			zeros++
		}
	}

	return int64(zeros), nil
}
