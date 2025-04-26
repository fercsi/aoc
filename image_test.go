package aoc

import (
	"os"
	"testing"
)

func TestSaveBytesImage(t *testing.T) {
	image := [][]byte{
		{1, 2},
		{2, 1},
	}
	colors := map[byte]int{
		1: 0xFF0000, // red
		2: 0x00FF00, // green
	}

	tmpFile := "test_output.ppm"
	defer os.Remove(tmpFile)

	SaveBytesImage(tmpFile, image, colors)

	data, err := os.ReadFile(tmpFile)
	if err != nil {
		t.Fatalf("could not read output file: %v", err)
	}

	expectedHeader := "P6\n2 2\n255\n"
	expectedData := []byte{
		255, 0, 0,    // 1 = red
		0, 255, 0,    // 2 = green
		0, 255, 0,    // 2 = green
		255, 0, 0,    // 1 = red
	}

	// Check header
	if string(data[:len(expectedHeader)]) != expectedHeader {
		t.Errorf("unexpected header:\nGot:  %q\nWant: %q", data[:len(expectedHeader)], expectedHeader)
	}

	// Check pixels
	gotPixels := data[len(expectedHeader):]
	if len(gotPixels) != len(expectedData) {
		t.Fatalf("unexpected pixel data length: got %d, want %d", len(gotPixels), len(expectedData))
	}
	for i := range gotPixels {
		if gotPixels[i] != expectedData[i] {
			t.Errorf("pixel mismatch at byte %d: got %v, want %v", i, gotPixels[i], expectedData[i])
		}
	}
}
