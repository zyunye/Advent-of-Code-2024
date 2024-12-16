package aoc

import (
	"fmt"
	"encoding/json"
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
