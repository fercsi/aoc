package dijkstra

import (
	"container/heap"
	"math"
)

type Edge struct {
	To     int
	Weight int
}

type Graph map[int][]Edge

func Dijkstra(graph Graph, start int, n int) []int {
	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = math.MaxInt64
	}
	dist[start] = 0

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{node: start, priority: 0})

	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)
		u := item.node
		if item.priority > dist[u] {
			continue
		}
		for _, edge := range graph[u] {
			v := edge.To
			weight := edge.Weight
			if dist[u]+weight < dist[v] {
				dist[v] = dist[u] + weight
				heap.Push(pq, &Item{node: v, priority: dist[v]})
			}
		}
	}
	return dist
}
