package aoc

import (
	"reflect"
	"testing"
)

func TestCoordDirections(t *testing.T) {
	start := Coord{X: 3, Y: 5}

	t.Run("Up", func(t *testing.T) {
		got := start.Up()
		want := Coord{X: 3, Y: 4}
		if got != want {
			t.Errorf("Up() = %v, want %v", got, want)
		}
	})

	t.Run("Down", func(t *testing.T) {
		got := start.Down()
		want := Coord{X: 3, Y: 6}
		if got != want {
			t.Errorf("Down() = %v, want %v", got, want)
		}
	})

	t.Run("Left", func(t *testing.T) {
		got := start.Left()
		want := Coord{X: 2, Y: 5}
		if got != want {
			t.Errorf("Left() = %v, want %v", got, want)
		}
	})

	t.Run("Right", func(t *testing.T) {
		got := start.Right()
		want := Coord{X: 4, Y: 5}
		if got != want {
			t.Errorf("Right() = %v, want %v", got, want)
		}
	})
}

func TestAllNeighbours(t *testing.T) {
	coord := Coord{X: 0, Y: 0}
	neighbours := coord.AllNeighbours()

	expectedNeighbours := []Coord{
		{X: 0, Y: -1}, // Up
		{X: -1, Y: 0}, // Left
		{X: 1, Y: 0},  // Right
		{X: 0, Y: 1},  // Down
	}

	if len(neighbours) != len(expectedNeighbours) {
		t.Errorf("Expected %d neighbours, got %d", len(expectedNeighbours), len(neighbours))
	}

	for i, n := range neighbours {
		if n != expectedNeighbours[i] {
			t.Errorf("Neighbour at index %d: expected %v, got %v", i, expectedNeighbours[i], n)
		}
	}
}

func TestNeighbours(t *testing.T) {
	coord := Coord{X: 1, Y: 1}
	width, height := 3, 3

	neighbours := coord.Neighbours(width, height)

	expectedNeighbours := []Coord{
		{X: 1, Y: 0}, // Up
		{X: 0, Y: 1}, // Left
		{X: 2, Y: 1}, // Right
		{X: 1, Y: 2}, // Down
	}

	if len(neighbours) != len(expectedNeighbours) {
		t.Errorf("Expected %d neighbours, got %d", len(expectedNeighbours), len(neighbours))
	}

	for i, n := range neighbours {
		if n != expectedNeighbours[i] {
			t.Errorf("Neighbour at index %d: expected %v, got %v", i, expectedNeighbours[i], n)
		}
	}

	// Test with a coordinate at bottom left corner
	coord = Coord{X: 0, Y: 2}
	neighbours = coord.Neighbours(width, height)

	expectedNeighbours = []Coord{
		{X: 0, Y: 1}, // Up
		{X: 1, Y: 2}, // Right
	}

	if len(neighbours) != len(expectedNeighbours) {
		t.Errorf("Expected %d neighbours for corner, got %d", len(expectedNeighbours), len(neighbours))
	}

	for i, n := range neighbours {
		if n != expectedNeighbours[i] {
			t.Errorf("Neighbour at index %d for corner: expected %v, got %v", i, expectedNeighbours[i], n)
		}
	}
}

func TestAllNeighbours8(t *testing.T) {
	coord := Coord{X: 1, Y: 1}
	neighbours := coord.AllNeighbours8()

	expected := []Coord{
		{X: 0, Y: 0}, // Up().Left()
		{X: 1, Y: 0}, // Up()
		{X: 2, Y: 0}, // Up().Right()
		{X: 0, Y: 1}, // Left()
		{X: 2, Y: 1}, // Right()
		{X: 0, Y: 2}, // Down().Left()
		{X: 1, Y: 2}, // Down()
		{X: 2, Y: 2}, // Down().Right()
	}

	if !reflect.DeepEqual(neighbours, expected) {
		t.Errorf("AllNeighbours8() = %v, expected %v", neighbours, expected)
	}
}

func TestNeighbours8(t *testing.T) {
	tests := []struct {
		name     string
		coord    Coord
		width    int
		height   int
		expected []Coord
	}{
		{
			name:   "Middle of 3x3 grid",
			coord:  Coord{X: 1, Y: 1},
			width:  3,
			height: 3,
			expected: []Coord{
				{X: 0, Y: 0},
				{X: 1, Y: 0},
				{X: 2, Y: 0},
				{X: 0, Y: 1},
				{X: 2, Y: 1},
				{X: 0, Y: 2},
				{X: 1, Y: 2},
				{X: 2, Y: 2},
			},
		},
		{
			name:   "Top-right corner of 3x3 grid",
			coord:  Coord{X: 2, Y: 0},
			width:  3,
			height: 3,
			expected: []Coord{
				{X: 1, Y: 0},
				{X: 1, Y: 1},
				{X: 2, Y: 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.coord.Neighbours8(tt.width, tt.height)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Neighbours8() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestDistance(t *testing.T) {
	tests := []struct {
		name     string
		coord1   Coord
		coord2   Coord
		expected int
	}{
		{
			name:     "Same point",
			coord1:   Coord{X: 0, Y: 0},
			coord2:   Coord{X: 0, Y: 0},
			expected: 0,
		},
		{
			name:     "Horizontal distance",
			coord1:   Coord{X: 0, Y: 0},
			coord2:   Coord{X: 3, Y: 0},
			expected: 3,
		},
		{
			name:     "Vertical distance",
			coord1:   Coord{X: 0, Y: 0},
			coord2:   Coord{X: 0, Y: 4},
			expected: 4,
		},
		{
			name:     "Diagonal (Manhattan) distance",
			coord1:   Coord{X: 1, Y: 2},
			coord2:   Coord{X: 4, Y: 6},
			expected: 7,
		},
		{
			name:     "Negative coordinates",
			coord1:   Coord{X: -2, Y: -3},
			coord2:   Coord{X: 1, Y: 1},
			expected: 7,
		},
		{
			name:     "Other direction",
			coord1:   Coord{X: 1, Y: 1},
			coord2:   Coord{X: -2, Y: -3},
			expected: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Distance(tt.coord1, tt.coord2)
			if got != tt.expected {
				t.Errorf("Distance(%v, %v) = %d; expected %d", tt.coord1, tt.coord2, got, tt.expected)
			}
		})
	}
}

func TestAtCoord(t *testing.T) {
	area := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	tests := []struct {
		name     string
		coord    Coord
		expected int
	}{
		{
			name:     "Top-left corner",
			coord:    Coord{X: 0, Y: 0},
			expected: 1,
		},
		{
			name:     "Center",
			coord:    Coord{X: 1, Y: 1},
			expected: 5,
		},
		{
			name:     "Bottom-right corner",
			coord:    Coord{X: 2, Y: 2},
			expected: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AtCoord(area, tt.coord)
			if got != tt.expected {
				t.Errorf("AtCoord(%v) = %d; expected %d", tt.coord, got, tt.expected)
			}
		})
	}
}

func TestAtCoordUnlimited(t *testing.T) {
	area := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	defaultValue := -1

	tests := []struct {
		name     string
		coord    Coord
		expected int
	}{
		{
			name:     "Inside bounds (top-left)",
			coord:    Coord{X: 0, Y: 0},
			expected: 1,
		},
		{
			name:     "Inside bounds (center)",
			coord:    Coord{X: 1, Y: 1},
			expected: 5,
		},
		{
			name:     "Inside bounds (bottom-right)",
			coord:    Coord{X: 2, Y: 2},
			expected: 9,
		},
		{
			name:     "Out of bounds (negative X)",
			coord:    Coord{X: -1, Y: 1},
			expected: defaultValue,
		},
		{
			name:     "Out of bounds (negative Y)",
			coord:    Coord{X: 1, Y: -1},
			expected: defaultValue,
		},
		{
			name:     "Out of bounds (too large X)",
			coord:    Coord{X: 3, Y: 1},
			expected: defaultValue,
		},
		{
			name:     "Out of bounds (too large Y)",
			coord:    Coord{X: 1, Y: 3},
			expected: defaultValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AtCoordUnlimited(area, tt.coord, defaultValue)
			if got != tt.expected {
				t.Errorf("AtCoordUnlimited(%v) = %d; expected %d", tt.coord, got, tt.expected)
			}
		})
	}
}

func TestBoundaries(t *testing.T) {
	coords := []Coord{
		{X: 1, Y: 5},
		{X: 7, Y: 4},
		{X: -1, Y: 0},
		{X: 3, Y: -5},
	}

	wantTl := Coord{-1, -5}
	wantBr := Coord{7, 5}

	tl, br := Boundaries(coords)

	if tl != wantTl {
		t.Errorf("Boundaries tl = %v, want %v", tl, wantTl)
	}
	if br != wantBr {
		t.Errorf("Boundaries br = %v, want %v", br, wantBr)
	}
}
