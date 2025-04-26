package aoc

import (
	"bytes"
	"os"
	"testing"
)

func captureOutput(f func()) string {
	var buf bytes.Buffer
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	buf.ReadFrom(r)
	return buf.String()
}

func TestShowBytesCondensed(t *testing.T) {
	input := [][]byte{
		[]byte("abc"),
		[]byte("def"),
	}
	got := captureOutput(func() {
		ShowBytesCondensed(input)
	})
	want := "abc\ndef\n"
	if got != want {
		t.Errorf("unexpected output:\nGot:\n%q\nWant:\n%q", got, want)
	}
}

func TestShowIntsCondensed(t *testing.T) {
	input := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	got := captureOutput(func() {
		ShowIntsCondensed(input)
	})
	want := "123\n456\n"
	if got != want {
		t.Errorf("unexpected output:\nGot:\n%q\nWant:\n%q", got, want)
	}
}

func TestShowIntsCsv(t *testing.T) {
	input := [][]int{
		{10, 20, 30},
		{1, 2, 3},
	}
	got := captureOutput(func() {
		ShowIntsCsv(input)
	})
	want := "10, 20, 30\n1, 2, 3\n"
	if got != want {
		t.Errorf("unexpected output:\nGot:\n%q\nWant:\n%q", got, want)
	}
}
