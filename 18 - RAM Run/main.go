package main

import (
	. "aoc"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func read_input(file_name string) []Position {
	file, err := os.Open(file_name)
	CheckErr(err)

	scanner := bufio.NewScanner(file)

	ret := make([]Position, 0)

	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, ",")
		x, err := strconv.Atoi(tokens[0])
		CheckErr(err)
		y, err := strconv.Atoi(tokens[1])
		CheckErr(err)

		ret = append(ret, Position{R: y, C: x})
	}

	return ret
}

func init_grid(rows, cols int, bytes []Position, sim_length int) [][]string {

	grid := make([][]string, 0)

	for r := 0; r < rows; r++ {
		grid = append(grid, make([]string, cols))
		for c := 0; c < cols; c++ {
			grid[r][c] = "."
		}
	}

	for i := 0; i <= sim_length; i++ {
		corruption := bytes[i]
		grid[corruption.R][corruption.C] = "#"
	}

	return grid
}

func part1(file_name string, grid_r, grid_c, sim_length int) {
	bytes := read_input(file_name)

	grid := init_grid(grid_r, grid_c, bytes, sim_length)

	start := Position{R: 0, C: 0}
	end := Position{R: grid_r - 1, C: grid_c - 1}
	came_from, cost_so_far := a_star(start, end, &grid)

	_ = came_from
	// traceback(end, &came_from, &grid)
	// fmt.Println(came_from)
	// fmt.Println(cost_so_far)
	// PrintArrByRow(grid)

	fmt.Printf("P.1: %f\n", cost_so_far[end])
}

func part2(file_name string, grid_r, grid_c int) {
	bytes := read_input(file_name)

	start := Position{R: 0, C: 0}
	end := Position{R: grid_r - 1, C: grid_c - 1}

	for sim_length := 1024; sim_length < len(bytes); sim_length++ {
		grid := init_grid(grid_r, grid_c, bytes, sim_length)

		came_from, _ := a_star(start, end, &grid)

		fmt.Println(sim_length)

		_, ok := came_from[end]
		if !ok {
			fmt.Printf("P2. %v", bytes[sim_length])
			break
		}
	}

}

func main() {
	file_name := "input.txt"

	test_r, test_c := 7, 7
	test_sim_len := 12

	inp_r, inp_c := 71, 71
	inp_sim_len := 1024

	if file_name == "test.txt" {
		part1(file_name, test_r, test_c, test_sim_len)
		part2(file_name, test_r, test_c)
	} else if file_name == "input.txt" {
		part1(file_name, inp_r, inp_c, inp_sim_len)
		part2(file_name, inp_r, inp_c)
	}

}
