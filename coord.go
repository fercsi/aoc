package aoc

type Coord struct {
    X, Y int
}

func (coord Coord) Up() Coord {
    coord.Y -= 1
    return coord
}

func (coord Coord) Down() Coord {
    coord.Y += 1
    return coord
}

func (coord Coord) Left() Coord {
    coord.X -= 1
    return coord
}

func (coord Coord) Right() Coord {
    coord.X += 1
    return coord
}
