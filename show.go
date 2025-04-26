package aoc

import (
    "fmt"
)

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
