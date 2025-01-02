package main

import (
	. "aoc"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func read_input(file_name string) []string {
	file, err := os.Open(file_name)
	CheckErr(err)

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	return strings.Fields(line)
}

func split_stone(i int, pebbles *[]string) {
	val := (*pebbles)[i]
	val_len := len(val)

	left := val[:val_len/2]
	right := strings.TrimLeft(val[val_len/2:], "0")
	if right == "" {
		right = "0"
	}

	(*pebbles)[i] = left
	Insert(i+1, right, pebbles)

}

func part1(file_name string) {
	pebbles := read_input(file_name)

	for blink := 0; blink < 25; blink++ {
		fmt.Printf("Blink: %d, %d\n", blink, len(pebbles))
		for i := 0; i < len(pebbles); i++ {
			num, err := strconv.Atoi(pebbles[i])
			CheckErr(err)

			if num == 0 {
				pebbles[i] = "1"
			} else if len(pebbles[i])%2 == 0 {
				split_stone(i, &pebbles)
				i++
			} else {
				num *= 2024
				pebbles[i] = strconv.Itoa(num)
			}
		}
	}

	fmt.Printf("P.2: %d\n", len(pebbles))

}


func part2(file_name string) {
	pebbles := read_input(file_name)

	cache := make(map[string][]string)

	for blink := 0; blink < 75; blink++ {
		fmt.Printf("Blink: %d, %d\n", blink, len(pebbles))
		for i := 0; i < len(pebbles); i++ {
			num, err := strconv.Atoi(pebbles[i])
			CheckErr(err)

			if num == 0 {
				pebbles[i] = "1"
			} else if len(pebbles[i])%2 == 0 {
				split_stone(i, &pebbles)
				i++
			} else {
				num *= 2024
				pebbles[i] = strconv.Itoa(num)
			}
		}
	}

	fmt.Printf("P.2: %d\n", len(pebbles))
}

func main() {
	file_name := "input.txt"
	part1(file_name)
	part2(file_name)
}
