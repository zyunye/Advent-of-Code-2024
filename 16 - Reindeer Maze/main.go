package main

import (
	. "aoc"
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func read_input(file_name string) [][]string {
	file, err := os.Open(file_name)
	CheckErr(err)

	scanner := bufio.NewScanner(file)

	maze := make([][]string, 0)

	r := 0
	for scanner.Scan() {
		maze = append(maze, make([]string, 0))
		for _, col := range scanner.Text() {
			maze[r] = append(maze[r], string(col))
		}
		r++
	}

	return maze
}

func heuristic(cur_pos, goal Position) float64 {
	return math.Abs(float64(cur_pos.R)-float64(goal.R)) + math.Abs(float64(cur_pos.C)-float64(goal.C))
}

func get_adj_costs(pos Position, facing Position, maze *[][]string) []CostStep {
	ret := make([]CostStep, 0)

	left_turn := Turn(facing, LEFT)
	ret = append(ret, CostStep{
		Step: Step{dir: left_turn, pos: pos},
		cost: 1000,
	})

	forward := pos.Add(facing)
	if Get(forward, maze) != "#" {
		ret = append(ret, CostStep{
			Step: Step{dir: facing, pos: forward},
			cost: 1,
		})
	}

	right_turn := Turn(facing, RIGHT)
	ret = append(ret, CostStep{
		Step: Step{dir: right_turn, pos: pos},
		cost: 1000,
	})

	return ret
}

func a_star(start Position, end Position, maze *[][]string) (map[Step][]CostStep, map[Step]float64) {
	frontier := make(PriorityQueue, 0)
	frontier.Push(PriorityStep{Step: Step{pos: start, dir: RIGHT}, priority: 0})

	came_from := make(map[Step][]CostStep)
	cost_so_far := make(map[Step]float64)

	for len(frontier) > 0 {
		current := heap.Pop(&frontier).(PriorityStep)
		cur_pos := current.pos
		cur_dir := current.dir

		// print_maze_with_step(maze, current.Step, "@")
		// fmt.Printf("At: %v, C: %f -- From: %v",
		// 	current.Step,
		// 	cost_so_far[current.Step],
		// 	came_from[current.Step])
		// fmt.Println()

		if cur_pos.Equal(end) {
			continue
		}

		adj_positions := get_adj_costs(cur_pos, cur_dir, maze)

		for _, adj := range adj_positions {
			cost := adj.cost
			new_cost := cost_so_far[current.Step] + cost
			csf, adj_seen := cost_so_far[adj.Step]

			// print_maze_with_step(maze, adj.Step, DirStr(adj.dir))
			// fmt.Printf("D: %s __ n: %f __ o: %f", DirStr(adj.dir), new_cost, csf)
			// fmt.Println()

			if !adj_seen || new_cost < csf {
				cost_so_far[adj.Step] = new_cost
				priority := new_cost + heuristic(adj.pos, end)
				heap.Push(&frontier, PriorityStep{Step: adj.Step, priority: priority})

				came_from[adj.Step] = make([]CostStep, 0)
				came_from[adj.Step] = append(came_from[adj.Step], CostStep{Step: current.Step, cost: cost_so_far[current.Step]})
			} else if new_cost <= csf {
				new_path := CostStep{Step: current.Step, cost: cost_so_far[current.Step]}

				// i := 0
				// for i < len(came_from[adj.Step]) {
				// 	if came_from[adj.Step][i].cost > new_path.cost {
				// 		came_from[adj.Step], _ = RemoveNoRef(i, came_from[adj.Step])
				// 	} else {
				// 		i++
				// 	}
				// }
				came_from[adj.Step] = append(came_from[adj.Step], new_path)
			}
		}

	}
	return came_from, cost_so_far
}

func trace_back(end Position, came_from *map[Step][]CostStep, cost_so_far *map[Step]float64, maze *[][]string) {

	end_paths := make([]CostStep, 0)
	for s, froms := range *came_from {
		if s.pos.Equal(end) {
			end_paths = append(end_paths, froms...)
		}
	}
	sort.Slice(end_paths, func(i, j int) bool {
		return end_paths[i].cost < end_paths[j].cost
	})

	parentage := make([]Position, 0)
	parentage = append(parentage, end_paths[0].pos)

	best_nodes := make(map[Position]bool)
	best_nodes[end] = true

	for len(parentage) > 0 {
		current_pos := Pop(&parentage)

		for step, froms := range *came_from {
			if step.pos.Equal(current_pos) {
				for _, cs := range froms {

					if _, ok := best_nodes[cs.pos]; !ok {
						parentage = append(parentage, cs.pos)
						best_nodes[cs.pos] = true
					}

					r, c := cs.pos.R, cs.pos.C
					times_treaded, err := strconv.Atoi((*maze)[r][c])
					if err != nil {
						(*maze)[r][c] = "1"
					} else {
						(*maze)[r][c] = strconv.Itoa(times_treaded + 1)
					}

					print_maze_file(maze, "out.txt")
					
				}
				fmt.Println()
			}
			
		}
	}

	fmt.Println(len(best_nodes))
}

func part1(file_name string) {
	maze := read_input(file_name)

	start := Position{R: len(maze) - 2, C: 1}
	end := Position{R: 1, C: len(maze[0]) - 2}
	print_maze(&maze)

	// came_from, cost_so_far := a_star(start, end, &maze)
	came_from, cost_so_far := a_star(start, end, &maze)
	// fmt.Printf("P.1: %f\n", cost_so_far[end])
	// print_came_from(&came_from, Position{R: 7, C: 5})
	// print_came_from(&came_from, Position{R: 7, C: 4})
	// print_came_from(&came_from, Position{R: 7, C: 3})

	// print_came_from(&came_from, Position{R: 8, C: 5})
	// print_came_from(&came_from, Position{R: 9, C: 5})

	trace_back(end, &came_from, &cost_so_far, &maze)

	for step, cost := range cost_so_far {
		// if step.pos.Equal(Position{R: 7, C: 5}) {
		if step.pos.Equal(end) {
			fmt.Println(step, cost)
			fmt.Println(came_from[step])
			fmt.Println()
		}
	}

}

func part2(file_name string) {
	// 819 too high
}

func main() {
	file_name := "input.txt"
	part1(file_name)
	part2(file_name)
}
