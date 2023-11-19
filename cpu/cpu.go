package cpu

import (
	"fmt"
	"log"
)

type mode int

const (
	Position mode = iota
	Immediate
)

type IO interface {
	In() int
	Out(output int)
}

type Interactive struct{}

func (io Interactive) In() int {
	var input int
	fmt.Print("input> ")
	_, err := fmt.Scanf("%d", &input)
	if err != nil {
		panic(err)
	}

	return input
}

func (io Interactive) Out(output int) {
	fmt.Println(output)
}

type Computer struct {
	Memory  []int
	Pointer int
	cmd     int
	modes   []mode
	io      IO
}

func NewComputer(prog []int) *Computer {
	return &Computer{
		Memory: prog,
		io:     Interactive{},
	}
}

func (cpu *Computer) Run() {
	for {
		intcode := cpu.Memory[cpu.Pointer]
		cpu.parseIntcode(intcode)

		switch cpu.cmd {
		case 1:
			cpu.add()
		case 2:
			cpu.mult()
		case 3:
			cpu.in()
		case 4:
			cpu.out()
		case 5:
			cpu.jt()
		case 6:
			cpu.jf()
		case 7:
			cpu.lt()
		case 8:
			cpu.eq()
		case 99:
			return
		default:
			log.Printf("bad intcode %d, halting", intcode)
			return
		}
	}
}

func (cpu *Computer) parseIntcode(intcode int) {
	var modes []mode

	cmd := intcode % 100
	intcode /= 100

	for i := 0; i < 3; i++ {
		modes = append(modes, mode(intcode%10))
		intcode /= 10
	}

	cpu.cmd = cmd
	cpu.modes = modes
}

func (cpu *Computer) jt() {
	cond := cpu.read(cpu.modes[0], cpu.Pointer+1)
	if cond != 0 {
		newPointer := cpu.read(cpu.modes[1], cpu.Pointer+2)
		cpu.Pointer = newPointer
		return
	}

	cpu.Pointer += 3
}

func (cpu *Computer) jf() {
	cond := cpu.read(cpu.modes[0], cpu.Pointer+1)
	if cond == 0 {
		newPointer := cpu.read(cpu.modes[1], cpu.Pointer+2)
		cpu.Pointer = newPointer
		return
	}

	cpu.Pointer += 3
}

func (cpu *Computer) lt() {
	left := cpu.read(cpu.modes[0], cpu.Pointer+1)
	right := cpu.read(cpu.modes[1], cpu.Pointer+2)

	var result int
	if left < right {
		result = 1
	}

	addr := cpu.Memory[cpu.Pointer+3]
	cpu.Memory[addr] = result
	cpu.Pointer += 4
}

func (cpu *Computer) eq() {
	left := cpu.read(cpu.modes[0], cpu.Pointer+1)
	right := cpu.read(cpu.modes[1], cpu.Pointer+2)

	var result int
	if left == right {
		result = 1
	}

	addr := cpu.Memory[cpu.Pointer+3]
	cpu.Memory[addr] = result
	cpu.Pointer += 4
}

func (cpu *Computer) in() {
	addr := cpu.Memory[cpu.Pointer+1]
	cpu.Memory[addr] = cpu.io.In()

	cpu.Pointer += 2
}

func (cpu *Computer) out() {
	output := cpu.read(cpu.modes[0], cpu.Pointer+1)
	cpu.io.Out(output)

	cpu.Pointer += 2
}

func (cpu *Computer) add() {
	left := cpu.read(cpu.modes[0], cpu.Pointer+1)
	right := cpu.read(cpu.modes[1], cpu.Pointer+2)
	addrResult := cpu.Memory[cpu.Pointer+3]

	cpu.Memory[addrResult] = left + right

	cpu.Pointer += 4
}

func (cpu *Computer) mult() {
	left := cpu.read(cpu.modes[0], cpu.Pointer+1)
	right := cpu.read(cpu.modes[1], cpu.Pointer+2)
	addrResult := cpu.Memory[cpu.Pointer+3]

	cpu.Memory[addrResult] = left * right

	cpu.Pointer += 4
}

func (cpu *Computer) read(mode mode, addr int) int {
	val := cpu.Memory[addr]
	switch mode {
	case Immediate:
		return val
	case Position:
		return cpu.Memory[val]
	default:
		panic("unknown mode")
	}
}
