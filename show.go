package aoc

import (
	"fmt"
)

// resultIndex keeps track of the current index across multiple ShowResult calls.
var resultIndex int = 0

// ShowResult prints each value in a numbered list format.
// The numbering continues across multiple calls.
func ShowResult(values ...any) {
	for _, value := range values {
		resultIndex++
		fmt.Printf("%d: %v\n", resultIndex, value)
	}
}

func ShowBytesCondensed(area [][]byte) {
	for _, row := range area {
		fmt.Println(string(row))
	}
}

func ShowIntsCondensed(area [][]int) {
	for _, row := range area {
		for _, item := range row {
			fmt.Print(item)
		}
		fmt.Println("")
	}
}

func ShowIntsCsv(area [][]int) {
	for _, row := range area {
		for i, item := range row {
			if i != 0 {
				fmt.Print(", ")
			}
			fmt.Print(item)
		}
		fmt.Println("")
	}
}
