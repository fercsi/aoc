package aoc

// Coord represents a 2D coordinate point
type Coord struct {
	X, Y int
}

// Up returns the coordinate above the current one
func (coord Coord) Up() Coord {
	coord.Y -= 1
	return coord
}

// Down returns the coordinate below the current one
func (coord Coord) Down() Coord {
	coord.Y += 1
	return coord
}

// Left returns the coordinate left to the current one
func (coord Coord) Left() Coord {
	coord.X -= 1
	return coord
}

// Right returns the coordinate righz to the current one
func (coord Coord) Right() Coord {
	coord.X += 1
	return coord
}
