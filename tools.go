package aoc

import "iter"

type addable interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64 |
    ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
    ~float32 | ~float64 |
    ~complex64 | ~complex128 |
    ~string
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
func SumFunc[E addable](s []E, f func(e E) E) E {
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
