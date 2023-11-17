package day2

import (
	"fmt"
	"log"
)

type Computer struct {
	Memory  []int
	Pointer int
}

func (cpu *Computer) Run() {
	for {
		intcode := cpu.Memory[cpu.Pointer]
		switch intcode {
		case 1:
			cpu.add()
		case 2:
			cpu.mult()
		case 99:
			return
		default:
			log.Println("bad intcode, halting")
			return
		}
	}
}

func (cpu *Computer) add() {
	addrLeft := cpu.Memory[cpu.Pointer+1]
	addrRight := cpu.Memory[cpu.Pointer+2]
	addrResult := cpu.Memory[cpu.Pointer+3]

	cpu.Memory[addrResult] = cpu.Memory[addrLeft] + cpu.Memory[addrRight]

	cpu.Pointer += 4
}

func (cpu *Computer) mult() {
	addrLeft := cpu.Memory[cpu.Pointer+1]
	addrRight := cpu.Memory[cpu.Pointer+2]
	addrResult := cpu.Memory[cpu.Pointer+3]

	cpu.Memory[addrResult] = cpu.Memory[addrLeft] * cpu.Memory[addrRight]

	cpu.Pointer += 4
}

func Solution(input []int) {
	partOneSolution := partOne(input)
	fmt.Printf("Part One: %d\n", partOneSolution)

	partTwoSolution := partTwo(input)
	fmt.Printf("Part Two: %d\n", partTwoSolution)
}

func partOne(prog []int) int {
	progCopy := make([]int, len(prog))
	copy(progCopy, prog)

	cpu := &Computer{progCopy, 0}
	cpu.Run()

	return cpu.Memory[0]
}

func partTwo(prog []int) int {
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			progCopy := make([]int, len(prog))
			copy(progCopy, prog)

			cpu := &Computer{progCopy, 0}
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
