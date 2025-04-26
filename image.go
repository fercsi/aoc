package aoc

import (
	"fmt"
	"os"
)

func SaveBytesImage(filename string, image [][]byte, colors map[byte]int) {
	lut := map[byte][]byte{}
	for k, v := range colors {
		lut[k] = []byte{byte(v >> 16), byte(v >> 8), byte(v)}
	}

	width := len(image[0])
	height := len(image)
	header := fmt.Sprintf("P6\n%d %d\n255\n", width, height)
	content := []byte(header)

	for _, row := range image {
		for _, px := range row {
			content = append(content, lut[px]...)
		}
	}

	err := os.WriteFile(filename, content, 0666)
	check(err)
}
