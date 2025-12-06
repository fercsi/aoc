package aoc

import "iter"

type addable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 |
		~complex64 | ~complex128 |
		~string
}

type mulable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 |
		~complex64 | ~complex128
}

// Sum returns the sum of all elements in the slice.
func Sum[E addable](s []E) E {
	var sum E
	for _, v := range s {
		sum += v
	}
	return sum
}

// SumFunc maps each element using f and returns the sum of the results.
func SumFunc[S any, E addable](s []S, f func(e S) E) E {
	var sum E
	for _, v := range s {
		sum += f(v)
	}
	return sum
}

// Sum returns the sum of all elements defined by seq.
func SumSeq[E addable](seq iter.Seq[E]) E {
	var sum E
	for v := range seq {
		sum += v
	}
	return sum
}

// Prod returns the product of all elements in the slice.
func Prod[E mulable](s []E) E {
	prod := E(1)
	for _, v := range s {
		prod *= v
	}
	return prod
}

// ProdFunc maps each element using f and returns the product of the results.
func ProdFunc[S any, E mulable](s []S, f func(e S) E) E {
	prod := E(1)
	for _, v := range s {
		prod *= f(v)
	}
	return prod
}

// Prod returns the product of all elements defined by seq.
func ProdSeq[E mulable](seq iter.Seq[E]) E {
	prod := E(1)
	for v := range seq {
		prod *= v
	}
	return prod
}

// CountIf returns the number of elements in s which pass the f test funvtion
func CountIf[S any](s []S, f func(S) bool) int {
	count := 0
	for _, v := range s {
		if f(v) {
			count++
		}
	}
	return count
}

// Grid2D returns an iterator of x, y values for each point in a 2D rectangle
// defined by startx, endx, starty, endy values (ends not inclusive).
func Grid2D(startx, endx, starty, endy int) iter.Seq2[int, int] {
	return func(yield func(int, int) bool) {
		for y := starty; y < endy; y++ {
			for x := startx; x < endx; x++ {
				if !yield(x, y) {
					return
				}
			}
		}
	}
}

// Transpose returns the column-wise transpose of the given lines. Missing
// characters are filled with spaces, and trailing spaces are trimmed while
// leading spaces are preserved.
func Transpose(lines []string) []string {
	maxLen := 0
	for _, s := range lines {
		if len(s) > maxLen {
			maxLen = len(s)
		}
	}

	out := make([]string, 0, maxLen)

	for col := 0; col < maxLen; col++ {
		buf := make([]byte, len(lines))
		for row := range lines {
			if col < len(lines[row]) {
				buf[row] = lines[row][col]
			} else {
				buf[row] = ' '
			}
		}

		end := len(buf)
		for end > 0 && buf[end-1] == ' ' {
			end--
		}

		out = append(out, string(buf[:end]))
	}

	return out
}
