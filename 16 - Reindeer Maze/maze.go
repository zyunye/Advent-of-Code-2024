package main

import (
	"fmt"
)

func print_maze(maze *[][]string) {
	for _, row := range *maze {
		fmt.Println(row)
	}
}
