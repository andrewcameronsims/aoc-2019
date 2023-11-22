package grid

type Facing int

const (
	Up Facing = iota
	Right
	Down
	Left
)

type Point struct {
	X, Y int
}
