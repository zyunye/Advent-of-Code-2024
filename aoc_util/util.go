package aoc

import (
	"fmt"
)

func PrintRunes(arr [][]rune) {
	for _, row := range arr {
		fmt.Println(string(row))
	}
}
