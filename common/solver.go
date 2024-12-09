package common

type Solver interface {
	SolvePart1(input string) (int64, error)
	SolvePart2(input string) (int64, error)
}
