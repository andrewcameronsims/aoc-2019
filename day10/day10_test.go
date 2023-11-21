package day10

import (
	"aoc-2019/raytrace"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAsteroid(t *testing.T) {
	testCases := []struct {
		desc         string
		input        string
		expectedSeen int
	}{
		{
			desc: "",
			input: `......#.#.
					#..#.#....
					..#######.
					.#.#.###..
					.#..#.....
					..#....#.#
					#..#....#.
					.##.#..###
					##...#..#.
					.#....####`,
			expectedSeen: 33,
		},
		{
			desc: "",
			input: `.#..#
					.....
					#####
					....#
					...##`,
			expectedSeen: 8,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			world := makeWorld(tC.input)

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

					origin := raytrace.Vector{X: j, Y: i}
					seen := tracer.Trace360(origin)

					fmt.Printf("origin: %+v, seen: %d\n", origin, seen)

					if seen > maxSeen {
						maxSeen = seen
					}
				}
			}

			assert.Equal(t, tC.expectedSeen, maxSeen)
		})
	}
}
