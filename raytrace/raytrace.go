package raytrace

import (
	"aoc-2019/set"
	"fmt"
)

type RayTracer struct {
	World [][]bool
}

type Vector struct {
	X, Y int
}

func (v Vector) Sum(other Vector) Vector {
	return Vector{
		X: v.X + other.X,
		Y: v.Y + other.Y,
	}
}

func (rt RayTracer) At(location Vector) (bool, error) {
	if location.Y >= len(rt.World) || location.Y < 0 || location.X >= len(rt.World[0]) || location.X < 0 {
		return false, fmt.Errorf("%+v out of bounds", location)
	}

	return rt.World[location.Y][location.X], nil
}

// Trace follows a ray from an origin point and returns the point at which its blocked
// This function will return nil if the ray reaches the edge of the world
func (rt RayTracer) Trace(origin, movement Vector) *Vector {
	location := origin.Sum(movement)

	objectAtLocation, err := rt.At(location)
	if err != nil {
		return nil
	}

	for !objectAtLocation && err == nil {
		location = location.Sum(movement)
		objectAtLocation, err = rt.At(location)
	}

	return &location
}

func (rt RayTracer) Trace360(origin Vector) int {
	allSeen := set.NewSet[Vector]()

	rays := set.NewSet[Vector]()
	// get a list of vectors corresponding to lines of sight/rays to trace
	// this current method is a hack but it should work well enough
	for y := 0; y < 1; y++ {
		for x := 0; x < 6; x++ {
			if x == 0 && y == 0 {
				continue
			}

			rays.AddAll(
				Vector{X: x, Y: y},
				Vector{X: -x, Y: y},
				Vector{X: x, Y: -y},
				Vector{X: -x, Y: -y},
			)
		}
	}

	for x := 0; x < 1; x++ {
		for y := 0; y < 6; y++ {
			if y == 0 && x == 0 {
				continue
			}

			rays.AddAll(
				Vector{X: x, Y: y},
				Vector{X: -x, Y: y},
				Vector{X: x, Y: -y},
				Vector{X: -x, Y: -y},
			)
		}
	}

	// call Trace on each of these and add the the
	for ray := range rays {
		seen := rt.Trace(origin, ray)
		if seen != nil {
			allSeen.Add(*seen)
		}
	}

	return len(allSeen)
}
