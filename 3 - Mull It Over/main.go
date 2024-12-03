package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func read_input(file_name string) string {
	b, err := os.ReadFile(file_name)
	check(err)
	return string(b)
}

func part1(file_path string) {
	file_txt := read_input(file_path)

	regex := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)
	all_muls := regex.FindAllString(file_txt, -1)

	sum := 0
	for _, mul := range all_muls {
		tokens := strings.Split(mul, ",")
		v1, err := strconv.Atoi(strings.TrimLeft(tokens[0], "mul("))
		check(err)
		v2, err := strconv.Atoi(strings.TrimRight(tokens[1], ")"))
		check(err)

		sum += v1 * v2

	}

	fmt.Printf("P.1: %d\n", sum)
}

func part2(file_path string) {
	file_txt := read_input(file_path)

	regex := regexp.MustCompile(`(mul\([0-9]+,[0-9]+\))|(do\(\)|don't\(\))`)
	all_tokens := regex.FindAllString(file_txt, -1)

	is_accumulating := true
	sum := 0
	for _, token := range all_tokens {
		if token == "do()" {
			is_accumulating = true
		} else if token == "don't()" {
			is_accumulating = false
		} else {
			if !is_accumulating {
				continue
			} else {
				tokens := strings.Split(token, ",")
				v1, err := strconv.Atoi(strings.TrimLeft(tokens[0], "mul("))
				check(err)
				v2, err := strconv.Atoi(strings.TrimRight(tokens[1], ")"))
				check(err)

				sum += v1 * v2
			}
		}
	}
	fmt.Printf("P.2: %d\n", sum)
}

func main() {
	file_path := "input.txt"
	part1(file_path)
	part2(file_path)
}
