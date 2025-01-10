package main

import (
	. "aoc"
)

type PriorityStep struct {
	pos      Position
	dir      Position
	priority float64
}

type PriorityQueue []*PriorityStep

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(item interface{}) {
	*pq = append(*pq, item.(*PriorityStep))
}

func (pq *PriorityQueue) Pop() interface{} {
	tmp := *pq
	n := len(tmp)
	item := tmp[n-1]
	*pq = tmp[0 : n-1]

	return item
}
