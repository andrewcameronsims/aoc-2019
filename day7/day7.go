package day7

import (
	"aoc-2019/cpu"
	"aoc-2019/permute"
	"fmt"
	"math"
	"sync"
)

type AmplifierLoop struct {
	Channels   []*cpu.ChannelBus
	Amplifiers []*Amplifier
}

func (al *AmplifierLoop) Run() int {
	var wg sync.WaitGroup

	for _, amp := range al.Amplifiers {
		wg.Add(1)

		go func(amp *Amplifier) {
			amp.cpu.Run()
			wg.Done()
		}(amp)

	}
	wg.Wait()

	lastOutput := al.Channels[0].Log[len(al.Channels[0].Log)-1]

	return lastOutput
}

func copyProgram(p []int) []int {
	pCopy := make([]int, len(p))
	copy(pCopy, p)

	return pCopy
}

func NewAmplifierLoop(phases []int, prog []int) *AmplifierLoop {
	bus1 := &cpu.ChannelBus{
		Q: make(chan int, 100),
	}
	bus1.Q <- phases[0]
	bus1.Q <- 0

	bus2 := &cpu.ChannelBus{
		Q: make(chan int, 100),
	}
	bus2.Q <- phases[1]

	bus3 := &cpu.ChannelBus{
		Q: make(chan int, 100),
	}
	bus3.Q <- phases[2]

	bus4 := &cpu.ChannelBus{
		Q: make(chan int, 100),
	}
	bus4.Q <- phases[3]

	bus5 := &cpu.ChannelBus{
		Q: make(chan int, 100),
	}
	bus5.Q <- phases[4]

	amp1 := &Amplifier{
		cpu: &cpu.Computer{
			Memory: copyProgram(prog),
			Reader: bus1,
			Writer: bus2,
		},
	}
	amp2 := &Amplifier{
		cpu: &cpu.Computer{
			Memory: copyProgram(prog),
			Reader: bus2,
			Writer: bus3,
		},
	}
	amp3 := &Amplifier{
		cpu: &cpu.Computer{
			Memory: copyProgram(prog),
			Reader: bus3,
			Writer: bus4,
		},
	}
	amp4 := &Amplifier{
		cpu: &cpu.Computer{
			Memory: copyProgram(prog),
			Reader: bus4,
			Writer: bus5,
		},
	}
	amp5 := &Amplifier{
		cpu: &cpu.Computer{
			Memory: copyProgram(prog),
			Reader: bus5,
			Writer: bus1,
		},
	}

	return &AmplifierLoop{
		Channels: []*cpu.ChannelBus{
			bus1,
			bus2,
			bus3,
			bus4,
			bus5,
		},
		Amplifiers: []*Amplifier{
			amp1,
			amp2,
			amp3,
			amp4,
			amp5,
		},
	}
}

type Amplifier struct {
	cpu *cpu.Computer
	io  *cpu.StandardIO
}

func NewAmplifier(program []int, phase, input int) *Amplifier {
	programCopy := make([]int, len(program))
	copy(programCopy, program)

	io := cpu.StandardIO{
		Inputs: []int{phase, input},
	}
	cpu := cpu.Computer{
		Memory: program,
		Reader: &io,
		Writer: &io,
	}

	return &Amplifier{
		cpu: &cpu,
		io:  &io,
	}
}

func (a *Amplifier) Out() int {
	a.cpu.Run()
	return a.io.Outputs[0]
}

func Solution(program []int) {
	fmt.Printf("Day One: %d\n", partOne(program))
	fmt.Printf("Day Two: %v\n", partTwo(program))
}

func partTwo(program []int) int {
	maxThrust := math.MinInt

	phaseSettings := []int{5, 6, 7, 8, 9}
	perms := permute.Permutations(phaseSettings)

	for _, perm := range perms {
		loop := NewAmplifierLoop(perm, program)
		thrust := loop.Run()

		if thrust > maxThrust {
			maxThrust = thrust
		}
	}

	return maxThrust
}

func partOne(program []int) int {
	maxThrust := math.MinInt

	ampInputs := []int{0, 1, 2, 3, 4}
	perms := permute.Permutations(ampInputs)

	for _, perm := range perms {
		amp1 := NewAmplifier(program, perm[0], 0)
		amp2 := NewAmplifier(program, perm[1], amp1.Out())
		amp3 := NewAmplifier(program, perm[2], amp2.Out())
		amp4 := NewAmplifier(program, perm[3], amp3.Out())
		amp5 := NewAmplifier(program, perm[4], amp4.Out())

		thrust := amp5.Out()

		if thrust > maxThrust {
			maxThrust = thrust
		}
	}

	return maxThrust
}
