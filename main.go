package main

import (
	"aoc-2019/common"
	"aoc-2019/day1"
	"aoc-2019/day2"
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

	switch day {
	case 1:
		input, err := common.ReadLinesFromInput(inputPath)
		if err != nil {
			log.Fatal(err)
		}
		day1.Solution(input)
	case 2:
		input, err := common.CommaSeparatedIntFromInput(inputPath)
		if err != nil {
			log.Fatal(err)
		}
		day2.Solution(input)
	}
}
