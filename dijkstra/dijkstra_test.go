package dijkstra

import (
	"testing"
)

func TestDijkstra(t *testing.T) {
	// 0 --(1)--> 1
	// 0 --(4)--> 2
	// 1 --(2)--> 2
	// 1 --(6)--> 3
	// 2 --(3)--> 3
	graph := Graph{
		0: {{To: 1, Weight: 1}, {To: 2, Weight: 4}},
		1: {{To: 2, Weight: 2}, {To: 3, Weight: 6}},
		2: {{To: 3, Weight: 3}},
		3: {},
	}

	start := 0
	n := 4
	dist := Dijkstra(graph, start, n)

	expected := []int{0, 1, 3, 6}

	for i, d := range dist {
		if d != expected[i] {
			t.Errorf("unexpected distance to node %d: got %d, want %d", i, d, expected[i])
		}
	}
}
