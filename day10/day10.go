package day10

import (
	"aoc-2019/raytrace"
	"fmt"
	"strings"
)

func makeWorld(input string) [][]bool {
	lines := strings.Fields(input)
	runes := make([][]rune, len(lines))

	for i := 0; i < len(lines); i++ {
		runes[i] = []rune(lines[i])
	}

	world := make([][]bool, len(runes))
	for i := 0; i < len(runes); i++ {
		row := make([]bool, len(runes[i]))

		for j := 0; j < len(row); j++ {
			if runes[i][j] == '#' {
				row[j] = true
			}
		}

		world[i] = row
	}

	return world
}

func Solution(input [][]rune) {
	world := make([][]bool, len(input))

	for i := 0; i < len(input); i++ {
		row := make([]bool, len(input[i]))

		for j := 0; j < len(row); j++ {
			if input[i][j] == '#' {
				row[j] = true
			}
		}

		world[i] = row
	}

	tracer := raytrace.RayTracer{
		World: world,
	}

	var maxSeen int

	for i := 0; i < len(world); i++ {
		for j := 0; j < len(world[0]); j++ {
			if !world[i][j] {
				// must be another asteroid
				continue
			}

			origin := raytrace.Vector{X: i, Y: j}
			seen := tracer.Trace360(origin)
			if seen > maxSeen {
				maxSeen = seen
			}
		}
	}

	fmt.Printf("Day One: %d\n", maxSeen)
}
