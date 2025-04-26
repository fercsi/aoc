package aoc

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func GetRunType() string {
	if len(os.Args) > 1 {
		return os.Args[1]
	}
	return ""
}

func ReadInput() []string {
	filename := "input.txt"
	if len(os.Args) > 1 {
		filename = "input" + os.Args[1] + ".txt"
	}
	input, err := os.ReadFile(filename)
	check(err)
	lines := strings.Split(string(input), "\n")
	return lines[:len(lines)-1]
}

func ReadInputLine() string {
	return ReadInput()[0]
}

func ReadInputBytes() [][]byte {
	input := ReadInput()
	inputBytes := make([][]byte, len(input))
	for i, line := range input {
		inputBytes[i] = []byte(line)
	}
	return inputBytes
}

func ReadInputSplit(delimiter string) [][]string {
	input := ReadInput()
	var result [][]string
	for _, line := range input {
		result = append(result, strings.Split(line, delimiter))
	}
	return result
}

func ReadInputLineSplit(delimiter string) []string {
	input := ReadInputLine()
	return strings.Split(input, delimiter)
}

func ReadInputInt() []int {
	input := ReadInput()
	var result []int
	for _, line := range input {
		num, err := strconv.Atoi(line)
		check(err)
		result = append(result, num)
	}
	return result
}

func ReadInputIntList() [][]int {
	input := ReadInput()
	var result [][]int
	for _, line := range input {
		result = append(result, ParseLineIntList(line))
	}
	return result
}

func ReadInputLineIntList() []int {
	input := ReadInputLine()
	return ParseLineIntList(input)
}

func ParseLineIntList(line string) []int {
	rx := regexp.MustCompile(`[+-]?\d+`)
	matches := rx.FindAllString(line, -1)
	var row []int
	for _, str := range matches {
		num, err := strconv.Atoi(str)
		check(err)
		row = append(row, num)
	}
	return row
}
