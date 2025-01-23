package main

import (
	. "aoc"
	"container/heap"
	"math"
)

func heuristic(cur_pos, goal Position) float64 {
	return float64(manhattan_dist(cur_pos, goal))
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

func manhattan_dist(start Position, end Position) int {
	return int(math.Abs(float64(start.R)-float64(end.R))) + int(math.Abs(float64(start.C)-float64(end.C)))
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

func is_wall_too_thick(pos Position, dir Position, maze *[][]string) bool {
	check_pos := pos.Add(dir)

	if (*maze)[check_pos.R][check_pos.C] == "." {
		return true
	}

	check_pos = check_pos.Add(dir)

	if Inbounds(check_pos, maze) {
		return (*maze)[check_pos.R][check_pos.C] == "#"
	}
	return true
}

func traceback(end Position, came_from *map[Position]Position, maze *[][]string) ([][]string, map[Position]bool) {

	stack := make([]Position, 0)
	stack = append(stack, end)

	walls_to_check := make(map[Position]bool, 0)

	maze_copy := make([][]string, len((*maze)))
	for i := range *maze {
		maze_copy[i] = append([]string(nil), (*maze)[i]...)
	}

	for len(stack) > 0 {
		cur_pos := Pop(&stack)
		maze_copy[cur_pos.R][cur_pos.C] = "O"

		for _, dir := range TURN_ORDER {
			if !is_wall_too_thick(cur_pos, dir, maze) {
				walls_to_check[cur_pos.Add(dir)] = true
			}
		}

		parent, ok := (*came_from)[cur_pos]

		if !ok {
			break
		}

		stack = append(stack, parent)
	}
	return maze_copy, walls_to_check
}

func get_manhattan_boundary(pos Position, radius int) []Position {

	ret := make([]Position, 0)

	r := pos.R
	c := pos.C

	for dc := -radius; dc <= radius; dc++ {
		dr := radius - Abs(dc)

		ret = append(ret, Position{R: r + dr, C: c + dc})
		if dr != 0 {
			ret = append(ret, Position{R: r - dr, C: c + dc})
		}
	}

	return ret
}

func get_manhattan_boundary_filled(pos Position, radius int) []Position {

	ret := make([]Position, 0)

	r := pos.R
	c := pos.C

	for dc := -radius; dc <= radius; dc++ {
		dr_max := radius - Abs(dc)

		for dr := -dr_max; dr <= dr_max; dr++ {
			ret = append(ret, Position{R: r + dr, C: c + dc})
		}
	}

	return ret
}

func get_valid_points_within_boundary(pos Position, radius int, maze *[][]string) []Position {

	ret := make([]Position, 0)

	r := pos.R
	c := pos.C

	for dc := -radius; dc <= radius; dc++ {
		dr_max := radius - Abs(dc)

		for dr := -dr_max; dr <= dr_max; dr++ {

			check_pos := Position{R: r + dr, C: c + dc}
			if Inbounds(check_pos, maze) {
				ret = append(ret, Position{R: r + dr, C: c + dc})
				// if (*maze)[check_pos.R][check_pos.C] == "." {
				// 	ret = append(ret, Position{R: r + dr, C: c + dc})
				// } else {
				// 	pos_to_jump := manhattan_dist(pos, check_pos)
				// 	jump_to_end := manhattan_dist(check_pos, end)
				// 	if pos_to_jump+jump_to_end <= 20 {
				// 		ret = append(ret, Position{R: r + dr, C: c + dc})
				// 	}
				// }
			}

		}
	}

	return ret
}

func is_straight_line_unimpeded(start Position, end Position, maze *[][]string) bool {
	if start.R == end.R {
		min_c := min(start.C, end.C)
		max_c := max(start.C, end.C)
		for min_c < max_c {
			if (*maze)[start.R][min_c] == "#" {
				return false
			}
			min_c++
		}
		return true
	} else if start.C == end.C {
		min_r := min(start.R, end.R)
		max_r := max(start.R, end.R)
		for min_r < max_r {
			if (*maze)[min_r][start.C] == "#" {
				return false
			}
			min_r++
		}
		return true
	}
	return false
}
