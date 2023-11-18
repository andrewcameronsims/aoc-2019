package cpu

import "log"

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
