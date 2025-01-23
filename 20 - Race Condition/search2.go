package main

import (
	. "aoc"
	"container/heap"
	"math"
)

func heuristic2(cur_pos, goal Position) float64 {
	return math.Abs(float64(cur_pos.R)-float64(goal.R)) + math.Abs(float64(cur_pos.C)-float64(goal.C))
}

func get_adj_costs2(node Node, maze *[][]string) []Node {

	ret := make([]Node, 0)

	neighbors := GetOrthPositions(node.pos, maze)
	for _, neighbor := range neighbors {

		if (*maze)[neighbor.R][neighbor.C] == "#" {
			if node.ps_remaining > 0 {
				ret = append(ret, Node{pos: neighbor, ps_remaining: node.ps_remaining - 1})
			}
		} else {
			if node.ps_remaining < 20 {
				ret = append(ret, Node{pos: neighbor, ps_remaining: node.ps_remaining - 1})
			} else {
				ret = append(ret, Node{pos: neighbor, ps_remaining: 20})
			}
		}
	}

	return ret
}

func a_star2(start Position, end Position, maze *[][]string) (map[Node]Node, map[Node]float64) {
	frontier := make(PriorityQueue2, 0)
	frontier.Push(PQ2Tuple{Node: Node{pos: start, ps_remaining: 20}, priority: 0})

	came_from := make(map[Node]Node)
	cost_so_far := make(map[Node]float64)

	cost_so_far[Node{pos: start, ps_remaining: 20}] = 0

	for len(frontier) > 0 {
		current := frontier.Pop().(PQ2Tuple)
		cur_node := current.Node
		// cur_pos := current.pos
		// cur_ps_remaining := current.ps_remaining

		if current.pos.Equal(end) {
			continue
		}

		neighbors := get_adj_costs2(cur_node, maze)
		for _, adj_node := range neighbors {
			// cost := adj_node.ps_remaining
			new_cost := cost_so_far[cur_node] + 1
			csf, adj_seen := cost_so_far[adj_node]

			if !adj_seen || new_cost < csf {
				cost_so_far[adj_node] = new_cost
				priority := new_cost + heuristic(cur_node.pos, end) + float64(adj_node.ps_remaining)

				// PQ2Tuple{Node: Node{pos: start, ps_remaining: 20}, priority: 0}
				heap.Push(&frontier, PQ2Tuple{Node: adj_node, priority: priority})
				came_from[adj_node] = cur_node
			}
		}
	}
	return came_from, cost_so_far
}
