package aoc

import (
	"golang.org/x/exp/constraints"
	"iter"
)

// Sum returns the sum of all elements in the slice.
func Sum[E constraints.Ordered](s []E) E {
	var sum E
	for _, v := range s {
		sum += v
	}
	return sum
}

// SumFunc maps each element using f and returns the sum of the results.
func SumFunc[E constraints.Ordered](s []E, f func(e E) E) E {
	var sum E
	for _, v := range s {
		sum += f(v)
	}
	return sum
}

// Sum returns the sum of all elements defined by seq.
func SumSeq[E constraints.Ordered](seq iter.Seq[E]) E {
	var sum E
	for v := range seq {
		sum += v
	}
	return sum
}
