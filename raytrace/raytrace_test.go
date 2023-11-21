package raytrace

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRayTrace_Trace(t *testing.T) {
	testCases := []struct {
		desc     string
		origin   Vector
		movement Vector
		world    [][]bool
		seen     *Vector
	}{
		{
			desc:     "",
			origin:   Vector{0, 0},
			movement: Vector{0, 1},
			world: [][]bool{
				{false, true, true},
				{true, false, true},
				{true, true, true},
			},
			seen: &Vector{0, 1},
		},
		{
			desc:     "",
			origin:   Vector{0, 0},
			movement: Vector{0, -1},
			world: [][]bool{
				{false, true, true},
				{true, false, true},
				{true, true, true},
			},
			seen: nil,
		},
		{
			desc:     "",
			origin:   Vector{0, 0},
			movement: Vector{1, 1},
			world: [][]bool{
				{false, true, true},
				{true, false, true},
				{true, true, true},
			},
			seen: &Vector{2, 2},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			rt := RayTracer{World: tC.world}
			seen := rt.Trace(tC.origin, tC.movement)

			assert.Equal(t, tC.seen, seen)
		})
	}
}
