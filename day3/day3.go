package day3

import (
	"aoc-2019/set"
	"fmt"
	"math"
	"strconv"
)

type Point struct {
	X, Y, Z int
}

func Solution(wires [][]string) {
	partOneSolution := partOne(wires)
	fmt.Printf("Part One: %d\n", partOneSolution)

	partTwoSolution := partTwo(wires)
	fmt.Printf("Part Two: %d\n", partTwoSolution)

}

func partTwo(wires [][]string) int {
	wireSets := make([][]Point, 2)

	for i, wire := range wires {
		var wireSet []Point
		p := Point{X: 0, Y: 0}
		wireSet = append(wireSet, p)

		for _, node := range wire {
			dir := node[0]
			mag := node[1:]

			steps, err := strconv.Atoi(mag)
			if err != nil {
				panic(err)
			}

			for i := 0; i < steps; i++ {
				switch dir {
				case 'R':
					p = Point{X: p.X + 1, Y: p.Y}
				case 'D':
					p = Point{X: p.X, Y: p.Y - 1}
				case 'U':
					p = Point{X: p.X, Y: p.Y + 1}
				case 'L':
					p = Point{X: p.X - 1, Y: p.Y}
				default:
					panic(dir)
				}
				wireSet = append(wireSet, p)
			}
		}
		wireSets[i] = wireSet
	}

	leftWire := wireSets[0]
	rightWire := wireSets[1]

	var crossed []Point

	for l, pl := range leftWire {
		for r, pr := range rightWire {
			if pl.X == pr.X && pl.Y == pr.Y {
				crossed = append(crossed, Point{X: pl.X, Y: pl.Y, Z: l + r})
			}
		}
	}

	shortest := math.MaxInt
	for _, x := range crossed {
		if x.Z == 0 {
			continue
		}

		if x.Z < shortest {
			shortest = x.Z
		}
	}

	return int(shortest)
}

func partOne(wires [][]string) int {
	wireSets := make([]set.Set[Point], 2)

	for i, wire := range wires {
		p := Point{X: 0, Y: 0}
		wireSet := set.NewSet[Point]()
		wireSet.Add(p)

		for _, node := range wire {
			wireSet.Add(p)
			dir := node[0]
			mag := node[1:]

			steps, err := strconv.Atoi(mag)
			if err != nil {
				panic(err)
			}

			for i := 0; i < steps; i++ {
				switch dir {
				case 'R':
					p = Point{X: p.X + 1, Y: p.Y}
				case 'D':
					p = Point{X: p.X, Y: p.Y - 1}
				case 'U':
					p = Point{X: p.X, Y: p.Y + 1}
				case 'L':
					p = Point{X: p.X - 1, Y: p.Y}
				default:
					panic(dir)
				}
				wireSet.Add(p)
			}
		}
		wireSets[i] = wireSet
	}

	leftWire := wireSets[0]
	rightWire := wireSets[1]

	crossed := leftWire.Intersect(rightWire)

	closest := math.MaxFloat64
	for x := range crossed {
		if x.X == 0 && x.Y == 0 {
			continue
		}

		distance := math.Abs(float64(x.X)) + math.Abs(float64(x.Y))
		if distance < closest {
			closest = distance
		}

	}

	return int(closest)
}
