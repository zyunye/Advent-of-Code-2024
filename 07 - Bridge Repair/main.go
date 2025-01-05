package main

import (
	. "aoc"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Equation struct {
	total int
	nums  []int
}

func read_input(file_name string) []Equation {
	file, err := os.Open(file_name)
	CheckErr(err)

	var equations []Equation

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), ":")

		total, err := strconv.Atoi(tokens[0])
		CheckErr(err)

		nums_str := strings.Fields(tokens[1])
		nums := make([]int, len(nums_str))
		for i, n := range nums_str {
			nums[i], err = strconv.Atoi(n)
			CheckErr(err)
		}

		equations = append(equations, Equation{total: total, nums: nums})

	}

	CheckErr(scanner.Err())

	return equations
}

func permutations(length int, values []rune) <-chan string {
	generator := make(chan string)

	go func() {
		defer close(generator)

		if length == 0 {
			generator <- ""
			return
		}

		sub_perms := permutations(length-1, values)

		for perm := range sub_perms {
			for _, v := range values {
				generator <- perm + string(v)
			}
		}

	}()

	return generator
}

func part1(file_name string) {
	equations := read_input(file_name)
	var symbols = []rune{'+', '*'}

	matched_totals := 0

	for _, equation := range equations {
		total := equation.total
		nums := equation.nums
		op_combos := permutations(len(nums)-1, symbols)

		for combo := range op_combos {
			checked_total := nums[0]

			for i, op := range combo {
				n := nums[i+1]
				switch op {
				case '+':
					checked_total += n
				case '*':
					checked_total *= n
				}
			}

			if checked_total == total {
				matched_totals += total
				break
			}

		}
	}

	fmt.Printf("P.1: %d\n", matched_totals)
}

func part2(file_name string) {
	equations := read_input(file_name)
	var symbols = []rune{'+', '*', '|'}

	matched_totals := 0

	for _, equation := range equations {
		total := equation.total
		nums := equation.nums
		op_combos := permutations(len(nums)-1, symbols)

		for combo := range op_combos {
			checked_total := nums[0]

			for i, op := range combo {
				n := nums[i+1]
				switch op {
				case '+':
					checked_total += n
				case '*':
					checked_total *= n
				case '|':
					tmp := strconv.Itoa(checked_total) + strconv.Itoa(n)
					conv, err := strconv.Atoi(tmp)
					CheckErr(err)
					checked_total = conv
				}
			}

			if checked_total == total {
				matched_totals += total
				break
			}

		}

	}

	fmt.Printf("P.2: %d\n", matched_totals)
}

func main() {
	file_name := "input.txt"
	part1(file_name)
	part2(file_name)
}
