package main

import (
	. "aoc"
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
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

func print_maze(maze *[][]string) {
	for _, row := range *maze {
		fmt.Println(row)
	}
}

func print_maze_with_iter(pos Position, maze *[][]string) {
	for r, row := range *maze {
		for c, col := range row {
			if pos.R == r && pos.C == c {
				fmt.Print("S")
			} else {
				fmt.Print(col)
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func print_full_path(end Position, path *map[Step]Step, maze *[][]string) {
	var cur_step Step
	for s, _ := range *path {
		if s.pos.Equal(end) {
			cur_step = s
			break
		}
	}

	for {
		from, ok := (*path)[cur_step]
		if !ok {
			break
		}
		pos := from.pos
		facing := from.dir

		if facing.Equal(UP) {
			(*maze)[pos.R][pos.C] = "^"
		} else if facing.Equal(RIGHT) {
			(*maze)[pos.R][pos.C] = ">"
		} else if facing.Equal(DOWN) {
			(*maze)[pos.R][pos.C] = "v"
		} else if facing.Equal(LEFT) {
			(*maze)[pos.R][pos.C] = "<"
		}

		cur_step = from
	}

	print_maze(maze)
}

func cost(cur_pos Position, next_pos Position, cur_facing Position) int {
	move := next_pos.Subtract(cur_pos)

	if move == cur_facing {
		return 1
	} else if move == Turn(cur_facing, LEFT) || move == Turn(cur_facing, RIGHT) {
		return 1001
	} else {
		panic(fmt.Sprintf("Cost did not receive adjacent positions, or tried to turn 180 -- cur:%v, next:%v, facing:%v", cur_pos, next_pos, cur_facing))
	}
}

func heuristic(cur_pos, goal Position) float64 {
	// return math.Abs(float64(cur_pos.R)-float64(goal.R)) + math.Abs(float64(cur_pos.C)-float64(goal.C))
	return math.Sqrt(math.Pow(float64(cur_pos.R)-float64(goal.R), 2) + math.Pow(float64(cur_pos.C)-float64(goal.C), 2))
}

func a_star(start Position, end Position, maze *[][]string) (map[Step]Step, map[Position]float64) {

	start_step := Step{pos: start, dir: RIGHT, priority: 0}

	frontier := &PriorityQueue{
		&start_step,
	}
	came_from := make(map[Step]Step)
	cost_so_far := make(map[Position]float64)

	cost_so_far[start] = 0

	for len(*frontier) != 0 {
		current := heap.Pop(frontier).(*Step)
		cur_pos := current.pos
		cur_facing := current.dir

		// print_maze_with_iter(cur_pos, maze)

		// if cur_pos.Equal(end) {
		// 	break
		// }

		valid_moves := map[Position]Position{
			Turn(cur_facing, LEFT):  cur_pos.Add(Turn(cur_facing, LEFT)),  //left
			cur_facing:              cur_pos.Add(cur_facing),              //forward
			Turn(cur_facing, RIGHT): cur_pos.Add(Turn(cur_facing, RIGHT)), //right
		}

		for facing, move := range valid_moves {
			if !Inbounds(move, maze) {
				continue
			}
			graph_val := Get(move, maze)
			if graph_val == "#" {
				continue
			}

			new_cost := cost_so_far[cur_pos] + float64(cost(cur_pos, move, cur_facing))

			prev_cost, seen := cost_so_far[move]
			if !seen || new_cost < prev_cost {
				cost_so_far[move] = new_cost
				priority := heuristic(move, end)

				next := Step{pos: move, dir: facing, priority: priority}

				heap.Push(frontier, &next)
				came_from[next] = *current
			}
		}
	}
	return came_from, cost_so_far
}

func part1(file_name string) {
	maze := read_input(file_name)

	start := Position{R: len(maze) - 2, C: 1}
	end := Position{R: 1, C: len(maze[0]) - 2}
	print_maze(&maze)

	path, costs := a_star(start, end, &maze)

	fmt.Println(len(path))

	fmt.Printf("P.1: %f\n", costs[end])
	//83468 too low
	print_full_path(end, &path, &maze)

}

func part2(file_name string) {
}

func main() {
	file_name := "test.txt"
	part1(file_name)
	part2(file_name)
}
