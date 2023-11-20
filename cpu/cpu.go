package cpu

import (
	"fmt"
	"log"
)

type mode int

const (
	Position mode = iota
	Immediate
	Relative
)

type Reader interface {
	In() int
}

type Writer interface {
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

type StandardIO struct {
	Outputs []int
	Inputs  []int
}

func (io *StandardIO) In() int {
	input := io.Inputs[0]
	io.Inputs = io.Inputs[1:]

	return input
}

func (io *StandardIO) Out(output int) {
	io.Outputs = append(io.Outputs, output)
}

type Computer struct {
	Memory   []int
	Pointer  int
	Reader   Reader
	Writer   Writer
	cmd      int
	modes    []mode
	relative int
}

func NewComputer(prog []int) *Computer {
	memory := make([]int, 1000)

	return &Computer{
		Memory: append(prog, memory...),
		Reader: Interactive{},
		Writer: Interactive{},
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
		case 9:
			cpu.rel()
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

func (cpu *Computer) rel() {
	offset := cpu.read(cpu.modes[0], cpu.Pointer+1)
	cpu.relative += offset

	cpu.Pointer += 2
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

	cpu.write(cpu.modes[2], cpu.Pointer+3, result)

	cpu.Pointer += 4
}

func (cpu *Computer) eq() {
	left := cpu.read(cpu.modes[0], cpu.Pointer+1)
	right := cpu.read(cpu.modes[1], cpu.Pointer+2)

	var result int
	if left == right {
		result = 1
	}

	cpu.write(cpu.modes[2], cpu.Pointer+3, result)

	cpu.Pointer += 4
}

func (cpu *Computer) in() {
	input := cpu.Reader.In()
	cpu.write(cpu.modes[0], cpu.Pointer+1, input)

	cpu.Pointer += 2
}

func (cpu *Computer) out() {
	output := cpu.read(cpu.modes[0], cpu.Pointer+1)
	cpu.Writer.Out(output)

	cpu.Pointer += 2
}

func (cpu *Computer) add() {
	left := cpu.read(cpu.modes[0], cpu.Pointer+1)
	right := cpu.read(cpu.modes[1], cpu.Pointer+2)

	cpu.write(cpu.modes[2], cpu.Pointer+3, left+right)

	cpu.Pointer += 4
}

func (cpu *Computer) mult() {
	left := cpu.read(cpu.modes[0], cpu.Pointer+1)
	right := cpu.read(cpu.modes[1], cpu.Pointer+2)

	cpu.write(cpu.modes[2], cpu.Pointer+3, left*right)

	cpu.Pointer += 4
}

func (cpu *Computer) write(mode mode, addr, value int) {
	val := cpu.Memory[addr]

	switch mode {
	case Position:
		cpu.Memory[val] = value
	case Relative:
		cpu.Memory[val+cpu.relative] = value
	default:
		panic("unknown mode")
	}
}

func (cpu *Computer) read(mode mode, addr int) int {
	val := cpu.Memory[addr]

	switch mode {
	case Immediate:
		return val
	case Position:
		return cpu.Memory[val]
	case Relative:
		return cpu.Memory[val+cpu.relative]
	default:
		panic("unknown mode")
	}
}
