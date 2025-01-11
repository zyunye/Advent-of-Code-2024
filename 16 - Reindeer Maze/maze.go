package main

import (
	"fmt"
	"os"
)

func print_maze(maze *[][]string) {
	for _, row := range *maze {
		fmt.Println(row)
	}
}

func print_maze_with_step(maze *[][]string, step Step, icon string) {
	for r, row := range *maze {
		for c, col := range row {
			if r == step.pos.R && c == step.pos.C {
				fmt.Print(icon)
			} else {
				fmt.Print(col)
			}
		}
		fmt.Println()
	}
}

func print_maze_file(maze *[][]string, fname string) {
	file, err := os.Create(fname)
	if err != nil {
		fmt.Println("Error creating file")
		return
	}
	defer file.Close()

	for _, row := range *maze {
		for _, col := range row {
			fmt.Fprintf(file, "%s ", col)
		}
		fmt.Fprintf(file, "\n")
	}
}
