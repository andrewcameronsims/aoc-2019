package cpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestIO struct {
	outputs []int
	inputs  []int
}

func (io *TestIO) In() int {
	input := io.inputs[0]
	io.inputs = io.inputs[1:]

	return input
}

func (io *TestIO) Out(output int) {
	io.outputs = append(io.outputs, output)
}

func TestComputer_Run(t *testing.T) {
	testCases := []struct {
		desc           string
		program        []int
		input          int
		expectedOutput int
	}{
		{
			desc: "",
			program: []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
				1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
				999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			input:          7,
			expectedOutput: 999,
		},
		{
			desc: "",
			program: []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
				1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
				999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			input:          8,
			expectedOutput: 1000,
		},
		{
			desc: "",
			program: []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
				1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
				999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			input:          9,
			expectedOutput: 1001,
		},
		{
			desc:           "",
			program:        []int{1102, 34915192, 34915192, 7, 4, 7, 99, 0},
			expectedOutput: 1219070632396864,
		},
		{
			desc:           "",
			program:        []int{104, 1125899906842624, 99},
			expectedOutput: 1125899906842624,
		},
		{
			desc:           "",
			program:        []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99},
			expectedOutput: 109,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			testIo := TestIO{
				inputs: []int{tC.input},
			}
			cpu := &Computer{Reader: &testIo, Writer: &testIo, Memory: append(tC.program, make([]int, 100)...)}
			cpu.Run()

			output := testIo.outputs[0]
			assert.Equal(t, tC.expectedOutput, output)
		})
	}
}

func TestComputer_parseIntcode(t *testing.T) {
	testCases := []struct {
		desc          string
		intcode       int
		expectedCmd   int
		expectedModes []mode
	}{
		{
			desc:          "",
			intcode:       1002,
			expectedCmd:   2,
			expectedModes: []mode{Position, Immediate, Position},
		},
		{
			desc:          "",
			intcode:       2,
			expectedCmd:   2,
			expectedModes: []mode{Position, Position, Position},
		},
		{
			desc:          "",
			intcode:       11002,
			expectedCmd:   2,
			expectedModes: []mode{Position, Immediate, Immediate},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			assertThat := assert.New(t)
			cpu := Computer{}
			cpu.parseIntcode(tC.intcode)

			assertThat.Equal(tC.expectedCmd, cpu.cmd)
			assertThat.Equal(tC.expectedModes, cpu.modes)
		})
	}
}
