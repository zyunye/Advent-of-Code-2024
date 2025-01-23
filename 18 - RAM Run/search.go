package main

import (
	. "aoc"
	"container/heap"
	"math"
)

func heuristic(cur_pos, goal Position) float64 {
	return math.Abs(float64(cur_pos.R)-float64(goal.R)) + math.Abs(float64(cur_pos.C)-float64(goal.C))
}

func get_adj_costs(pos Position, maze *[][]string) []PriorityStep {

	ret := make([]PriorityStep, 0)

	neighbors := GetOrthPositions(pos, maze)
	for _, neighbor := range neighbors {
		if (*maze)[neighbor.R][neighbor.C] != "#" {
			ret = append(ret, PriorityStep{pos: neighbor, priority: 1})
		}
	}

	return ret
}

func a_star(start Position, end Position, maze *[][]string) (map[Position]Position, map[Position]float64) {
	frontier := make(PriorityQueue, 0)
	frontier.Push(PriorityStep{pos: start, priority: 0})
	came_from := make(map[Position]Position)
	cost_so_far := make(map[Position]float64)

	cost_so_far[start] = 0

	for len(frontier) > 0 {
		current := frontier.Pop().(PriorityStep)
		cur_pos := current.pos

		if current.pos.Equal(end) {
			continue
		}

		neighbors := get_adj_costs(cur_pos, maze)
		for _, adj := range neighbors {
			cost := adj.priority
			new_cost := cost_so_far[current.pos] + cost
			csf, adj_seen := cost_so_far[adj.pos]

			if !adj_seen || new_cost < csf {
				cost_so_far[adj.pos] = new_cost
				priority := new_cost + heuristic(cur_pos, end)

				heap.Push(&frontier, PriorityStep{pos: adj.pos, priority: priority})
				came_from[adj.pos] = cur_pos
			}
		}
	}
	return came_from, cost_so_far
}

func traceback(end Position, came_from *map[Position]Position, maze *[][]string) {

	stack := make([]Position, 0)
	stack = append(stack, end)

	for len(stack) > 0 {
		cur_pos := Pop(&stack)
		(*maze)[cur_pos.R][cur_pos.C] = "O"

		parent, ok := (*came_from)[cur_pos]

		if !ok {
			return
		}

		stack = append(stack, parent)
	}
}
