package aoc

import (
	"encoding/json"
	"fmt"
)

func PrintRunes(arr [][]rune) {
	for _, row := range arr {
		fmt.Println(string(row))
	}
}

func Dump(data interface{}) {
	b, _ := json.MarshalIndent(data, "", "  ")
	fmt.Print(string(b))
}

func PrintArrByRow[T any](arr [][]T) {
	for _, row := range arr {
		fmt.Println(row)
	}
}
