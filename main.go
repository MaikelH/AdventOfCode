package main

import (
	"AoC/common"
	"AoC/year2024"
	"embed"
	"flag"
	"fmt"
	"log"
	"os"
)

//go:embed input/*
var inputFiles embed.FS

func main() {
	year := flag.Int("year", 2024, "Year of the Advent of Code challenge")
	day := flag.Int("day", 1, "Day of the Advent of Code challenge")
	test := flag.Int("test", 0, "Use test input files")
	session := os.Getenv("AOC_SESSION")
	flag.Parse()

	if session == "" {
		log.Fatal("Session token is required")
	}

	input, err := GetInputFile(*year, *day, session, *test)
	if err != nil {
		log.Fatalf("Failed to get input file: %v", err)
	}

	solutionFunc, err := getSolutionFunc(*year, *day)
	if err != nil {
		log.Fatalf("Failed to get solution function: %v", err)
	}

	fmt.Printf("Year %d Day %d\n", *year, *day)
	part1, err := solutionFunc.SolvePart1(input)
	if err != nil {
		log.Fatalf("Failed to solve part 1: %v", err)
	}
	fmt.Printf("Part 1: %s\n", part1)

	part2, err := solutionFunc.SolvePart2(input)
	if err != nil {
		log.Fatalf("Failed to solve part 2: %v", err)
	}
	fmt.Printf("Part 2: %s\n", part2)
}

func GetInputFile(year, day int, session string, test int) (string, error) {
	var file []byte
	var err error
	if test == 1 {
		file, err = inputFiles.ReadFile(fmt.Sprintf("input/%d/day-%d-test.txt", year, day))
	} else {
		file, err = inputFiles.ReadFile(fmt.Sprintf("input/%d/day-%d.txt", year, day))
	}
	if err == nil {
		return string(file), nil
	}

	file, err = common.DownloadInputFile(year, day, session)
	if err != nil {
		return "", err
	}

	err = os.WriteFile(fmt.Sprintf("input/%d/day-%d.txt", year, day), file, os.ModePerm)
	if err != nil {
		return "", err
	}

	return string(file), err
}

func getSolutionFunc(year, day int) (common.Solver, error) {
	solutions := map[int]map[int]common.Solver{
		2024: {
			1: year2024.Day1{},
			2: year2024.Day2{},
			3: year2024.Day3{},
			// Add other days here
		},
		// Add other years here
	}

	if yearSolutions, ok := solutions[year]; ok {
		if solutionFunc, ok := yearSolutions[day]; ok {
			return solutionFunc, nil
		}
	}

	return nil, fmt.Errorf("solution for year %d day %d not found", year, day)
}
