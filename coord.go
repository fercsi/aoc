package aoc

import (
	"iter"
	"math"
)

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

// Neighbours returns the adjacent coordinates of the given coordinate
//
// If an adjacent coordinate is outside the block defined by width and height,
// those coordinates are not returned
func (coord Coord) Neighbours(width, height int) []Coord {
	n4 := coord.AllNeighbours()
	neighbours := []Coord{}
	for _, c := range n4 {
		if c.X >= 0 && c.X < width && c.Y >= 0 && c.Y < height {
			neighbours = append(neighbours, c)
		}
	}
	return neighbours
}

// AllNeighbours returns all the adjacent coordinates of the given coordinate
//
// In contrast to simple Neighbours, it returns all adjacent coordinates
func (coord Coord) AllNeighbours() []Coord {
	return []Coord{
		coord.Up(),
		coord.Left(),
		coord.Right(),
		coord.Down(),
	}
}

// Neighbours8 returns the adjacent and diagonal coordinates of the given coordinate
//
// If a coordinate is outside the block defined by width and height, those
// coordinates are not returned
func (coord Coord) Neighbours8(width, height int) []Coord {
	n8 := coord.AllNeighbours8()
	neighbours := []Coord{}
	for _, c := range n8 {
		if c.X >= 0 && c.X < width && c.Y >= 0 && c.Y < height {
			neighbours = append(neighbours, c)
		}
	}
	return neighbours
}

// AllNeighbours8 returns all the adjacent and diagonal coordinates of the given coordinate
//
// In contrast to simple Neighbours8, it returns all coordinates
func (coord Coord) AllNeighbours8() []Coord {
	n4 := coord.AllNeighbours()
	return []Coord{
		n4[0].Left(), n4[0], n4[0].Right(),
		n4[1], n4[2],
		n4[3].Left(), n4[3], n4[3].Right(),
	}
}

// Distance returns the Manhattan-distance between two coordinates
func Distance(coord1, coord2 Coord) int {
	x1, y1 := coord1.X, coord1.Y
	x2, y2 := coord2.X, coord2.Y
	dist := 0
	if x1 > x2 {
		dist += x1 - x2
	} else {
		dist += x2 - x1
	}
	if y1 > y2 {
		dist += y1 - y2
	} else {
		dist += y2 - y1
	}
	return dist
}

// AtCoord returns the value found at coord of the given area
func AtCoord[E any](area [][]E, coord Coord) E {
	return area[coord.Y][coord.X]
}

// AtCoordUnlimited returns the value found at coordinate of the given area and
// default value if the area does not contain the coordinate
func AtCoordUnlimited[E any](area [][]E, coord Coord, defaultValue E) E {
	if coord.Y < 0 || coord.Y >= len(area) {
		return defaultValue
	}
	row := area[coord.Y]
	if coord.X < 0 || coord.X >= len(row) {
		return defaultValue
	}
	return row[coord.X]
}

// SetAtCoord sets the value at coord of a 2D slice of elements
func SetAtCoord[E any](area [][]E, coord Coord, val E) {
	area[coord.Y][coord.X] = val
}

// Boundaries returns the top left corner and the bottom right corner of the
// smallestboundary box containing each coordinates
func Boundaries(coords []Coord) (Coord, Coord) {
	minx := math.MaxInt
	maxx := math.MinInt
	miny := minx
	maxy := maxx
	for _, s := range coords {
		minx = min(minx, s.X)
		maxx = max(maxx, s.X)
		miny = min(miny, s.Y)
		maxy = max(maxy, s.Y)
	}
	return Coord{minx, miny}, Coord{maxx, maxy}
}

// GridCoords returns an iterator of Coord values for each point in a 2D rectangle
// defined by its top-left and bottom-right coordinates (inclusive).
func GridCoords(tl, br Coord) iter.Seq[Coord] {
	return func(yield func(Coord) bool) {
		for y := tl.Y; y <= br.Y; y++ {
			for x := tl.X; x <= br.X; x++ {
				if !yield(Coord{x, y}) {
					return
				}
			}
		}
	}
}
