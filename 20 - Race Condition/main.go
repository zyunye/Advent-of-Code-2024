package main

import (
	. "aoc"
	"bufio"
	"fmt"
	"os"
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
	fmt.Println(start)
	fmt.Println(end)

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

}

func main() {
	file_name := "input.txt"

	part1(file_name)
	part2(file_name)
}
