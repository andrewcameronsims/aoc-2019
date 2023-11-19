package day2

import (
	"aoc-2019/cpu"
	"fmt"
)

func Solution(input []int) {
	partOneSolution := partOne(input)
	fmt.Printf("Part One: %d\n", partOneSolution)

	partTwoSolution := partTwo(input)
	fmt.Printf("Part Two: %d\n", partTwoSolution)
}

func partOne(prog []int) int {
	progCopy := make([]int, len(prog))
	copy(progCopy, prog)

	cpu := cpu.NewComputer(progCopy)
	cpu.Run()

	return cpu.Memory[0]
}

func partTwo(prog []int) int {
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			progCopy := make([]int, len(prog))
			copy(progCopy, prog)

			cpu := cpu.NewComputer(progCopy)
			cpu.Memory[1] = i
			cpu.Memory[2] = j

			cpu.Run()

			result := cpu.Memory[0]

			if result == 19690720 {
				return 100*i + j
			}
		}
	}

	return 0
}
