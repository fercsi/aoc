package aoc

import (
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
