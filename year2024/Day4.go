package year2024

import (
	"bufio"
	"regexp"
	"strings"
)

type Day4 struct {
}

func (d Day4) SolvePart1(input string) (int64, error) {
	regexF := regexp.MustCompile("XMAS")
	regexB := regexp.MustCompile("SAMX")

	chars := make([][]byte, 10)
	scanner := bufio.NewScanner(strings.NewReader(input))
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		chars[i] = []byte(line)
		i++
	}

	matches := 0
	// Find horizontal matches
	for j := 0; j < len(chars); j++ {
		forward := regexF.FindAllString(string(chars[j]), -1)
		backward := regexB.FindAllString(string(chars[j]), -1)
		matches += len(forward) + len(backward)
	}
	// Find vertical matches

	return int64(matches), nil
}

func (d Day4) SolvePart2(input string) (int64, error) {
	return 0, nil
}
