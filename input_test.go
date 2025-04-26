package aoc

import (
	"os"
	"testing"
)

func TestGetRunType(t *testing.T) {
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	os.Args = []string{"cmd", "-t"}
	if rt := GetRunType(); rt != "-t" {
		t.Errorf("expected '-t', got %q", rt)
	}

	os.Args = []string{"cmd"}
	if rt := GetRunType(); rt != "" {
		t.Errorf("expected empty string, got %q", rt)
	}
}

func TestReadInput(t *testing.T) {
	err := os.WriteFile("input.txt", []byte("foo\nbar\n"), 0644)
	if err != nil {
		t.Fatalf("could not write test input: %v", err)
	}
	defer os.Remove("input.txt")

	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()
	os.Args = []string{"cmd"}

	lines := ReadInput()
	expected := []string{"foo", "bar"}
	if len(lines) != len(expected) {
		t.Fatalf("expected %d lines, got %d", len(expected), len(lines))
	}
	for i := range expected {
		if lines[i] != expected[i] {
			t.Errorf("line %d: expected %q, got %q", i, expected[i], lines[i])
		}
	}
}

func TestReadInput_WithError(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic, but no panic occurred")
		}
	}()

	os.Args = []string{"cmd", "-nope"}
	ReadInput()
}

func TestReadInputLine(t *testing.T) {
	err := os.WriteFile("input-t.txt", []byte("hello\nworld\n"), 0644)
	if err != nil {
		t.Fatalf("failed to write test input: %v", err)
	}
	defer os.Remove("input-t.txt")

	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()
	os.Args = []string{"cmd", "-t"}

	line := ReadInputLine()
	if line != "hello" {
		t.Errorf("expected 'hello', got %q", line)
	}
}

func TestReadInputBytes(t *testing.T) {
	err := os.WriteFile("input-x.txt", []byte("hi\nxy\n"), 0644)
	if err != nil {
		t.Fatalf("failed to write test input: %v", err)
	}
	defer os.Remove("input-x.txt")

	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()
	os.Args = []string{"cmd", "-x"}

	bytes := ReadInputBytes()
	expected := [][]byte{[]byte("hi"), []byte("xy")}
	if len(bytes) != len(expected) {
		t.Fatalf("expected %d rows, got %d", len(expected), len(bytes))
	}
	for i := range expected {
		if string(bytes[i]) != string(expected[i]) {
			t.Errorf("row %d: expected %q, got %q", i, expected[i], bytes[i])
		}
	}
}

func TestReadInputSplit(t *testing.T) {
	err := os.WriteFile("input-csv.txt", []byte("a,b\n1,2\n"), 0644)
	if err != nil {
		t.Fatalf("failed to write test input: %v", err)
	}
	defer os.Remove("input-csv.txt")

	os.Args = []string{"cmd", "-csv"}

	split := ReadInputSplit(",")
	want := [][]string{{"a", "b"}, {"1", "2"}}
	for i := range want {
		for j := range want[i] {
			if split[i][j] != want[i][j] {
				t.Errorf("split[%d][%d] = %q, want %q", i, j, split[i][j], want[i][j])
			}
		}
	}
}

func TestReadInputLineSplit(t *testing.T) {
	err := os.WriteFile("input-t4.txt", []byte("x:y:z\n"), 0644)
	if err != nil {
		t.Fatalf("failed to write test input: %v", err)
	}
	defer os.Remove("input-t4.txt")

	os.Args = []string{"cmd", "-t4"}

	split := ReadInputLineSplit(":")
	want := []string{"x", "y", "z"}
	for i := range want {
		if split[i] != want[i] {
			t.Errorf("split[%d] = %q, want %q", i, split[i], want[i])
		}
	}
}

func TestReadInputInt(t *testing.T) {
	err := os.WriteFile("input-int.txt", []byte("5\n6\n7\n"), 0644)
	if err != nil {
		t.Fatalf("failed to write test input: %v", err)
	}
	defer os.Remove("input-int.txt")

	os.Args = []string{"cmd", "-int"}

	ints := ReadInputInt()
	want := []int{5, 6, 7}
	for i := range want {
		if ints[i] != want[i] {
			t.Errorf("ints[%d] = %d, want %d", i, ints[i], want[i])
		}
	}
}

func TestParseLineIntList(t *testing.T) {
    line := "Numbers are: 1 -2 3,+4 and -05"
	want := []int{1, -2, 3, 4, -5}
	got := ParseLineIntList(line)
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("at index %d: got %d, want %d", i, got[i], want[i])
		}
	}
}

func TestReadInputLineIntList(t *testing.T) {
	err := os.WriteFile("input-ints.txt", []byte("10 -20 +30\n"), 0644)
	if err != nil {
		t.Fatalf("failed to write test input: %v", err)
	}
	defer os.Remove("input-ints.txt")

	os.Args = []string{"cmd", "-ints"}

	got := ReadInputLineIntList()
	want := []int{10, -20, 30}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("at index %d: got %d, want %d", i, got[i], want[i])
		}
	}
}

func TestReadInputIntList(t *testing.T) {
	// Készítsünk egy input fájlt, ami több sort tartalmaz számokkal
	err := os.WriteFile("input-matrix.txt", []byte("x: 1 2 / 3\ny: 4 5 / 6\nz: 7 8 / 9\n"), 0644)
	if err != nil {
		t.Fatalf("failed to write test input: %v", err)
	}
	defer os.Remove("input-matrix.txt")

	// Átállítjuk az os.Args-ot, hogy a megfelelő fájlt olvassa
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()
	os.Args = []string{"cmd", "-matrix"}

	got := ReadInputIntList()
	want := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	if len(got) != len(want) {
		t.Fatalf("expected %d rows, got %d", len(want), len(got))
	}
	for i := range want {
		for j := range want[i] {
			if got[i][j] != want[i][j] {
				t.Errorf("at row %d, col %d: got %d, want %d", i, j, got[i][j], want[i][j])
			}
		}
	}
}
