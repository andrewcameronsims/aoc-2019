package main

import (
	"aoc-2019/common"
	"aoc-2019/day1"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	args := os.Args

	day, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatal(err)
	}

	inputPath := fmt.Sprintf("day%d/input", day)
	input, err := common.ReadInputFile(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	switch day {
	case 1:
		day1.Solution(input)
	}
}
