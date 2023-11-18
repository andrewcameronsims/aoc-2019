package day4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay4_twoAdjacentOnly(t *testing.T) {
	testCases := []struct {
		desc           string
		n              int
		expectedResult bool
	}{
		{
			desc:           "",
			n:              223450,
			expectedResult: true,
		},
		{
			desc:           "",
			n:              12223450,
			expectedResult: false,
		},
		{
			desc:           "",
			n:              213456,
			expectedResult: false,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := twoAdjacentOnly(tC.n)
			assert.Equal(t, tC.expectedResult, result)
		})
	}
}

func TestDay4_noDecrease(t *testing.T) {
	testCases := []struct {
		desc           string
		n              int
		expectedResult bool
	}{
		{
			desc:           "",
			n:              223450,
			expectedResult: false,
		},
		{
			desc:           "",
			n:              223456,
			expectedResult: true,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := noDecrease(tC.n)
			assert.Equal(t, tC.expectedResult, result)
		})
	}
}

func TestDay4_twoAdjacentSame(t *testing.T) {
	testCases := []struct {
		desc           string
		n              int
		expectedResult bool
	}{
		{
			desc:           "",
			n:              223450,
			expectedResult: true,
		},
		{
			desc:           "",
			n:              213456,
			expectedResult: false,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := twoAdjacentSame(tC.n)
			assert.Equal(t, tC.expectedResult, result)
		})
	}
}
