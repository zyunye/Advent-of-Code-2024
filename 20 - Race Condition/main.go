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

	fmt.Printf("P.1: %d\n", valid_cheats)
}

func part2(file_name string) {
	maze, start, end := read_input(file_name)
	orig_lineage, orig_costs := a_star(start, end, &maze)
	_ = orig_lineage
	_ = orig_costs

	cheats := make(map[int]int)

	for start_pos, start_cost := range orig_costs {
		for jump_pos, jump_cost := range orig_costs {
			if !start_pos.Equal(jump_pos) {
				start_to_jump_cost := manhattan_dist(start_pos, jump_pos)
				jump_to_end_cost := orig_costs[end] - jump_cost

				if start_to_jump_cost <= 20 {
					new_cost := start_cost + float64(start_to_jump_cost) + jump_to_end_cost

					if new_cost <= orig_costs[end] {
						cheats[int(orig_costs[end])-int(new_cost)] += 1
					}
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

	valid_cheats_count := 0
	for saved, count := range cheats {
		if saved >= 100 {
			valid_cheats_count += count
		}
	}

	fmt.Printf("P.2: %d\n", valid_cheats_count)

}

func main() {
	file_name := "input.txt"

	part1(file_name)
	part2(file_name)
}
