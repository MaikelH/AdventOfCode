package year2024

import (
	"AoC/common"
	"bufio"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Day1 struct {
}

func (d Day1) SolvePart1(input string) (int64, error) {
	var leftList []int
	var rightList []int
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		left, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatalf("Failed to convert string to int: %v", err)
		}

		right, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatalf("Failed to convert string to int: %v", err)
		}

		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	difference := 0
	for i := 0; i < len(leftList); i++ {
		difference += int(math.Abs(float64(leftList[i] - rightList[i])))
	}

	return int64(difference), nil
}

func (d Day1) SolvePart2(input string) (int64, error) {
	var leftList []int
	var rightList []int
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		left, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatalf("Failed to convert string to int: %v", err)
		}

		right, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatalf("Failed to convert string to int: %v", err)
		}

		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}

	frequencies := common.CountFrequency(rightList)
	score := 0
	for _, value := range leftList {
		score += frequencies[value] * value
	}

	return int64(score), nil
}
