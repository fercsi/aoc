package aoc

import (
	"maps"
	"reflect"
	"strings"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Sum of ints", func(t *testing.T) {
		ints := []int{1, 2, 3, 4}
		want := 10
		got := Sum(ints)
		if got != want {
			t.Errorf("Sum(ints) = %d; want %d", got, want)
		}
	})

	t.Run("Sum of float64s", func(t *testing.T) {
		floats := []float64{1.5, 2.0, 3.25}
		want := 6.75
		got := Sum(floats)
		if got != want {
			t.Errorf("Sum(floats) = %f; want %f", got, want)
		}
	})

	t.Run("Sum of strings", func(t *testing.T) {
		strs := []string{"hello", " ", "world"}
		want := "hello world"
		got := Sum(strs)
		if got != want {
			t.Errorf(`Sum(strs) = "%s"; want "%s"`, got, want)
		}
	})
}

func TestSumFunc(t *testing.T) {
	t.Run("Sum of squares of ints", func(t *testing.T) {
		ints := []int{1, 2, 3}
		want := 14 // 1^2 + 2^2 + 3^2
		got := SumFunc(ints, func(x int) int {
			return x * x
		})
		if got != want {
			t.Errorf("SumFunc(squares) = %d; want %d", got, want)
		}
	})

	t.Run("Sum of uppercased strings", func(t *testing.T) {
		words := []string{"a", "b", "c"}
		want := "ABC"
		got := SumFunc(words, func(s string) string {
			return strings.ToUpper(s)
		})
		if got != want {
			t.Errorf("SumFunc(upper) = %s; want %s", got, want)
		}
	})

	t.Run("Convatenation of digits", func(t *testing.T) {
		digits := []int{1, 9, 7, 5}
		want := "1975"
		got := SumFunc(digits, func(d int) string {
			return string(byte(d) + '0')
		})
		if got != want {
			t.Errorf("SumFunc(digits) = %s; want %s", got, want)
		}
	})
}

func TestSumSeq(t *testing.T) {
	lut := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
	got := SumSeq(maps.Values(lut))
	want := 15

	if got != want {
		t.Errorf("SumiSeq = %d, want %d", got, want)
	}
}

func TestGrid2D(t *testing.T) {
	// Teljes bejárás teszt
	got := [][2]int{}
	for x, y := range Grid2D(0, 2, 3, 5) {
		got = append(got, [2]int{x, y})
	}

	want := [][2]int{
		{0, 3}, {1, 3},
		{0, 4}, {1, 4},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Grid2D iteration = %v, want %v", got, want)
	}

	// Megszakítás teszt
	count := 0
	Grid2D(0, 2, 3, 5)(func(x, y int) bool {
		count++
		return false // azonnal megszakítjuk
	})

	if count != 1 {
		t.Errorf("Expected early termination after 1 call, got %d", count)
	}
}
