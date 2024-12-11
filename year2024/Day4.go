package year2024

import (
	"AoC/common"
	"bufio"
	"regexp"
	"strings"
)

type Day4 struct {
}

func (d Day4) SolvePart1(input string) (int64, error) {
	regexF := regexp.MustCompile("XMAS")
	regexB := regexp.MustCompile("SAMX")

	chars := [][]byte{}
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		chars = append(chars, []byte(line))
	}

	matches := 0
	// Find row matches
	for j := 0; j < len(chars); j++ {
		forward := regexF.FindAllString(string(chars[j]), -1)
		backward := regexB.FindAllString(string(chars[j]), -1)
		matches += len(forward) + len(backward)
	}
	// Find column matches
	columns := common.GetColumns(chars)
	for j := 0; j < len(columns); j++ {
		forward := regexF.FindAllString(string(columns[j]), -1)
		backward := regexB.FindAllString(string(columns[j]), -1)
		matches += len(forward) + len(backward)
	}

	return int64(matches), nil
}

func (d Day4) SolvePart2(input string) (int64, error) {
	return 0, nil
}
