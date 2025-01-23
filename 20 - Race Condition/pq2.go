package main

import (
	. "aoc"
)

type Node struct {
	pos          Position
	ps_remaining int
}

type PQ2Tuple struct {
	Node
	priority float64
}

type PriorityQueue2 []PQ2Tuple

func (pq PriorityQueue2) Len() int {
	return len(pq)
}

func (pq PriorityQueue2) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue2) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue2) Push(item interface{}) {
	*pq = append(*pq, item.(PQ2Tuple))
}

func (pq *PriorityQueue2) Pop() interface{} {
	tmp := *pq
	n := len(tmp)
	item := tmp[n-1]
	*pq = tmp[0 : n-1]

	return item
}
