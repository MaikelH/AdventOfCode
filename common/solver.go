package common

type Solver interface {
	SolvePart1(input string) (string, error)
	SolvePart2(input string) (string, error)
}
