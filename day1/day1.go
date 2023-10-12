package day1

import (
	"fmt"
	"strconv"
)

func Solution(input []string) {
	var modules []int
	for _, line := range input {
		mod, _ := strconv.Atoi(line)

		modules = append(modules, mod)
	}

	partOneSolution := partOne(modules)
	fmt.Printf("Part One: %d\n", partOneSolution)

	partTwoSolution := partTwo(modules)
	fmt.Printf("Part Two: %d\n", partTwoSolution)
}

func partOne(input []int) int {
	var answer int

	for _, module := range input {
		answer += module/3 - 2
	}

	return answer
}

func partTwo(input []int) int {
	var answer int

	var fuel int
	for _, module := range input {
		fuel = module/3 - 2
		answer += fuel

		for fuel > 0 {
			fuel = (fuel / 3) - 2
			if fuel > 0 {
				answer += fuel
			}
		}
	}

	return answer
}
