package day4

import (
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type PasswordCracker struct {
	min        int
	max        int
	conditions []func(int) bool
}

func sixDigits(n int) bool {
	var digits int

	for n > 0 {
		n /= 10
		digits++
	}

	return digits == 6
}

func twoAdjacentSame(n int) bool {
	prev := 11

	for n > 0 {
		nextDigit := n % 10
		if prev == nextDigit {
			return true
		}
		n /= 10
		prev = nextDigit
	}

	return false
}

func twoAdjacentOnly(n int) bool {
	prev := 10

	for n > 0 {
		current := n % 10
		next := (n % 100) / 10
		if prev == current && current == next {
			// we've hit a large group of sames and we need to skip
			for current == prev {
				n /= 10
				current = n % 10
			}
		}
		if prev == current {
			return true
		}
		n /= 10
		prev = current
	}

	return false
}

func noDecrease(n int) bool {
	var digits []int

	for n > 0 {
		nextDigit := n % 10
		digits = append(digits, nextDigit)
		n /= 10
	}

	slices.Reverse(digits)

	digitsCopy := make([]int, len(digits))
	copy(digitsCopy, digits)
	sortable := sort.IntSlice(digitsCopy)
	sortable.Sort()

	for i := 0; i < len(sortable); i++ {
		if sortable[i] != digits[i] {
			return false
		}
	}

	return true
}

func (pc PasswordCracker) Crack() []int {
	var valids []int

	for i := pc.min; i < pc.max+1; i++ {
		var conditionsMet int
		for _, cond := range pc.conditions {
			if cond(i) {
				conditionsMet++
			}
		}

		if conditionsMet == len(pc.conditions) {
			valids = append(valids, i)
		}
	}

	return valids
}

func partOne(minMax string) int {
	minMaxAry := strings.Split(minMax, "-")
	min, err := strconv.Atoi(minMaxAry[0])
	if err != nil {
		panic(err)
	}
	max, err := strconv.Atoi(minMaxAry[1])
	if err != nil {
		panic(err)
	}

	pc := PasswordCracker{
		min: min,
		max: max,
		conditions: []func(int) bool{
			sixDigits,
			twoAdjacentSame,
			noDecrease,
		},
	}

	return len(pc.Crack())
}

func partTwo(minMax string) int {
	minMaxAry := strings.Split(minMax, "-")
	min, err := strconv.Atoi(minMaxAry[0])
	if err != nil {
		panic(err)
	}
	max, err := strconv.Atoi(minMaxAry[1])
	if err != nil {
		panic(err)
	}

	pc := PasswordCracker{
		min: min,
		max: max,
		conditions: []func(int) bool{
			sixDigits,
			twoAdjacentOnly,
			noDecrease,
		},
	}

	return len(pc.Crack())
}

func Solution(input string) {
	partOneSolution := partOne(input)
	fmt.Printf("Part One: %d\n", partOneSolution)

	partTwoSolution := partTwo(input)
	fmt.Printf("Part Two: %d\n", partTwoSolution)
}
