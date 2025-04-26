package dijkstra

import (
	"testing"
)

func TestPriorityQueue_BasicOperations(t *testing.T) {
	pq := &PriorityQueue{}

	// Push Items
	item1 := &Item{node: 1, priority: 5}
	item2 := &Item{node: 2, priority: 3}
	item3 := &Item{node: 3, priority: 4}

	pq.Push(item1)
	pq.Push(item2)
	pq.Push(item3)

	if pq.Len() != 3 {
		t.Errorf("expected length 3, got %d", pq.Len())
	}

	// Check Less (priority order)
	if !pq.Less(1, 0) {
		t.Errorf("expected item2 (priority 3) to be less than item1 (priority 5)")
	}

	// Swap two elements
	pq.Swap(0, 1)
	if (*pq)[0].node != 2 || (*pq)[1].node != 1 {
		t.Errorf("swap failed: got nodes %d and %d", (*pq)[0].node, (*pq)[1].node)
	}

	// Pop elements
	popped := pq.Pop().(*Item)
	if popped.node != 3 {
		t.Errorf("expected node 3 to be popped, got %d", popped.node)
	}

	if pq.Len() != 2 {
		t.Errorf("expected length 2 after pop, got %d", pq.Len())
	}

	// Check item.index after pop
	if popped.index != -1 {
		t.Errorf("expected popped item index to be -1, got %d", popped.index)
	}
}

func TestPriorityQueue_PushAndPopOrder(t *testing.T) {
	pq := &PriorityQueue{}

	pq.Push(&Item{node: 10, priority: 1})
	pq.Push(&Item{node: 20, priority: 2})
	pq.Push(&Item{node: 30, priority: 3})

	// After pushing 3 items, length should be 3
	if pq.Len() != 3 {
		t.Errorf("expected length 3, got %d", pq.Len())
	}

	// Pop items and check the popped item
	pq.Pop()
	pq.Pop()
	popped3 := pq.Pop().(*Item)

	if popped3.priority != 1 {
		t.Errorf("expected last popped to have priority 1, got %d", popped3.priority)
	}
}
