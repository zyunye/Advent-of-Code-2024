package main

import (
	. "aoc"
)

func print_manhattan_boundary(pos Position, radius int, maze *[][]string) {
	boundary := get_manhattan_boundary(pos, radius)

	maze_copy := Copy2DArr(maze)

	for _, coord := range boundary {
		if Inbounds(coord, &maze_copy) {
			maze_copy[coord.R][coord.C] = "X"
		}
	}

	PrintArrByRow(maze_copy)
}
