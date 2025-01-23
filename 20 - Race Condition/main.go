package main

import (
	. "aoc"
	"bufio"
	"fmt"
	"os"
	"sort"
)

func read_input(file_name string) ([][]string, Position, Position) {
	file, err := os.Open(file_name)
	CheckErr(err)

	scanner := bufio.NewScanner(file)

	maze := make([][]string, 0)

	row := 0
	start_pos := Position{}
	end_pos := Position{}

	for scanner.Scan() {
		maze = append(maze, make([]string, 0))
		for col, v := range scanner.Text() {
			v := string(v)
			if v == "S" {
				start_pos.R = row
				start_pos.C = col
			} else if v == "E" {
				end_pos.R = row
				end_pos.C = col
			}

			maze[row] = append(maze[row], v)
		}
		row++
	}
	return maze, start_pos, end_pos
}

func part1(file_name string) {
	maze, start, end := read_input(file_name)

	came_from, cost_so_far := a_star(start, end, &maze)
	_, walls_to_check := traceback(end, &came_from, &maze)

	original_cost := cost_so_far[end]
	cheats := make(map[int]int)

	for wall := range walls_to_check {
		maze_copy := make([][]string, len(maze))
		for i := range maze {
			maze_copy[i] = append([]string(nil), maze[i]...)
		}

		maze_copy[wall.R][wall.C] = "."

		_, cost_so_far := a_star(start, end, &maze_copy)

		cheats[int(original_cost)-int(cost_so_far[end])] += 1
	}

	valid_cheats := 0
	for k, v := range cheats {
		if k >= 100 {
			valid_cheats += v
		}
	}

	fmt.Printf("P.1: %d", valid_cheats)
}

func part2(file_name string) {
	maze, start, end := read_input(file_name)
	orig_lineage, orig_costs := a_star(start, end, &maze)
	_ = orig_lineage
	_ = orig_costs

	search_stack := make([]Position, 0)
	search_stack = append(search_stack, start)
	seen := make(map[Position]bool)

	radius := 20

	cheats := make(map[int]int)

	for len(search_stack) > 0 {
		cur_pos := Pop(&search_stack)
		seen[cur_pos] = true

		jumps_in_range := get_valid_points_within_boundary(cur_pos, radius, &maze)

		for _, jump := range jumps_in_range {
			cur_cost := orig_costs[cur_pos]
			end_cost := orig_costs[end]
			jump_cost := orig_costs[jump]

			new_cost := cur_cost + float64(end_cost-jump_cost)

			if new_cost < end_cost-cur_cost {
				saved_cost := end_cost - new_cost

				cheats[int(saved_cost)] += 1
			}

			// if jump.Equal(end) {

			// 	cur_cost := orig_costs[cur_pos]
			// 	new_cost := cur_cost + float64(manhattan_dist(cur_pos, end))

			// 	saved_cost := orig_costs[end] - new_cost
			// 	cheats[int(saved_cost)] += 1
			// } else {
			// cost_remaining_after_jump := orig_costs[end] - orig_costs[jump]
			// if cost_remaining_after_jump < orig_costs[cur_pos] {
			// 	saved_cost := orig_costs[jump] - orig_costs[cur_pos]
			// 	cheats[int(saved_cost)] += 1
			// }
			// }
		}

		// dist_to_end := manhattan_dist(cur_pos, end)

		// if dist_to_end <= radius {
		// 	cur_cost := orig_costs[cur_pos]
		// 	new_cost := cur_cost + float64(dist_to_end)

		// 	saved_cost := orig_costs[end] - new_cost
		// 	cheats[int(saved_cost)] += 1
		// }

		neighbors := GetOrthPositions(cur_pos, &maze)
		for _, neighbor := range neighbors {
			if maze[neighbor.R][neighbor.C] == "." {
				if _, ok := seen[neighbor]; !ok {
					search_stack = append(search_stack, neighbor)
				}

			}
		}
	}

	func(m map[int]int) {
		type kv struct {
			Key   int
			Value int
		}

		var kvPairs []kv
		for key, value := range m {
			kvPairs = append(kvPairs, kv{Key: key, Value: value})
		}

		sort.Slice(kvPairs, func(i, j int) bool {
			return kvPairs[i].Key < kvPairs[j].Key
		})

		total_cheats := 0
		for _, pair := range kvPairs {
			fmt.Printf("%d cheats saved %d ps\n", pair.Value, pair.Key)
			total_cheats += pair.Value
		}
		fmt.Println(total_cheats)

	}(cheats)

}

func main() {
	file_name := "test.txt"

	// part1(file_name)
	part2(file_name)
}
