package main

import (
	. "aoc"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func read_input(file_name string) ([]string, []string) {
	file, err := os.Open(file_name)
	CheckErr(err)

	scanner := bufio.NewScanner(file)

	available_patterns := make([]string, 0)
	desired_patterns := make([]string, 0)

	scanner.Scan()
	patterns := strings.Split(scanner.Text(), ",")
	for _, p := range patterns {
		available_patterns = append(available_patterns, strings.TrimSpace(p))
	}

	scanner.Scan()
	for scanner.Scan() {
		desired_patterns = append(desired_patterns, scanner.Text())
	}

	return available_patterns, desired_patterns
}

func check_pattern(pattern string, prefixes *[]string, cache *map[string]int) int {
	if _, found_pattern := (*cache)[pattern]; !found_pattern {
		if len(pattern) == 0 {
			return 1
		} else {
			num_solutions := 0
			for _, prefix := range *prefixes {
				if strings.HasPrefix(pattern, prefix) {
					num_solutions += check_pattern(pattern[len(prefix):], prefixes, cache)
				}
			}
			(*cache)[pattern] = num_solutions
		}
	}
	return (*cache)[pattern]
}

func part1(file_name string) {
	available, desired := read_input(file_name)
	cache := make(map[string]int)

	valid_patterns := 0
	for _, pattern := range desired {
		if check_pattern(pattern, &available, &cache) != 0 {
			valid_patterns++
		}
	}

	fmt.Printf("P.1: %d\n", valid_patterns)
}

func part2(file_name string) {
	available, desired := read_input(file_name)
	cache := make(map[string]int)

	total_ways := 0
	for _, pattern := range desired {
		total_ways += check_pattern(pattern, &available, &cache)
	}

	fmt.Printf("P.2: %d\n", total_ways)
}

func main() {
	file_name := "input.txt"

	part1(file_name)
	part2(file_name)
}
