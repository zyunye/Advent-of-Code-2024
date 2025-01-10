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

func heuristic(cur_pos, goal Position) float64 {
	return math.Abs(float64(cur_pos.R)-float64(goal.R)) + math.Abs(float64(cur_pos.C)-float64(goal.C))
}

func a_star(start Position, end Position, maze *[][]string) {

}

func part1(file_name string) {
	maze := read_input(file_name)

	start := Position{R: len(maze) - 2, C: 1}
	end := Position{R: 1, C: len(maze[0]) - 2}
	// print_maze(&maze)

	a_star(start, end, &maze)
}

func part2(file_name string) {
}

func main() {
	file_name := "input.txt"
	part1(file_name)
	part2(file_name)
}
