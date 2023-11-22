package day11

import (
	"aoc-2019/cpu"
	"aoc-2019/grid"
	"aoc-2019/set"
	"fmt"
	"strings"
)

type Robot struct {
	World      map[grid.Point]int
	Location   grid.Point
	Facing     grid.Facing
	Cycle      int
	PaintCount set.Set[grid.Point]
}

func (r *Robot) RenderWorld() string {
	var sb strings.Builder

	g := make([][]int, 10)
	for i := 0; i < len(g); i++ {
		row := make([]int, 50)

		for j := 0; j < len(row); j++ {
			point := grid.Point{X: j, Y: i}
			row[j] = r.World[point]
		}

		g[i] = row
	}

	for _, row := range g {
		for _, cell := range row {
			if cell == 1 {
				sb.WriteString("#")
			} else {
				sb.WriteString(".")
			}
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

func (r *Robot) In() int {
	paint := r.World[r.Location]

	return paint
}

func (r *Robot) Out(output int) {
	if r.Cycle%2 == 0 {
		// Paint Cycle
		switch output {
		case 0:
			r.World[r.Location] = 0
		case 1:
			r.World[r.Location] = 1
		default:
			panic("unknown command")
		}

		r.PaintCount.Add(r.Location)
	} else {
		// Move Cycle
		switch output {
		case 0:
			r.left()
		case 1:
			r.right()
		default:
			panic("unknown command")
		}

		r.move()
	}

	r.Cycle++
}

func (r *Robot) left() {
	new := r.Facing - 1
	if new < 0 {
		r.Facing = 3
	} else {
		r.Facing = (r.Facing - 1) % 4
	}
}

func (r *Robot) right() {
	r.Facing = (r.Facing + 1) % 4
}

func (r *Robot) move() {
	switch r.Facing {
	case grid.Up:
		r.Location.Y -= 1
	case grid.Left:
		r.Location.X -= 1
	case grid.Down:
		r.Location.Y += 1
	case grid.Right:
		r.Location.X += 1
	default:
		panic("unknown facing")
	}
}

func Solution(input []int) {
	fmt.Printf("Part One: %d\n", partOne(input))
	fmt.Println("Part Two:")
	world := partTwo(input)
	fmt.Println(world)
}

func partTwo(input []int) string {
	robot := Robot{
		World:      make(map[grid.Point]int),
		Location:   grid.Point{X: 0, Y: 0},
		Facing:     grid.Up,
		PaintCount: set.NewSet[grid.Point](),
	}

	robot.World[grid.Point{X: 0, Y: 0}] = 1

	computer := cpu.Computer{
		Memory: append(input, make([]int, 1000)...),
		Reader: &robot,
		Writer: &robot,
	}

	computer.Run()

	return robot.RenderWorld()
}

func partOne(input []int) int {
	robot := &Robot{
		World:      make(map[grid.Point]int),
		Location:   grid.Point{X: 0, Y: 0},
		Facing:     grid.Up,
		PaintCount: set.NewSet[grid.Point](),
	}

	computer := cpu.Computer{
		Memory: append(input, make([]int, 1000)...),
		Reader: robot,
		Writer: robot,
	}

	computer.Run()

	return len(robot.PaintCount)
}
