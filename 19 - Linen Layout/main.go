package main

import (
	. "aoc"
	"bufio"
	"fmt"
	"os"
	"strings"
)

// type PrefixTree struct {
// 	node string
// 	next []string
// }

func read_input(file_name string) (map[string]bool, []string) {
	file, err := os.Open(file_name)
	CheckErr(err)

	scanner := bufio.NewScanner(file)

	available_patterns := make(map[string]bool, 0)
	desired_patterns := make([]string, 0)

	scanner.Scan()
	patterns := strings.Split(scanner.Text(), ",")
	for _, p := range patterns {
		available_patterns[strings.TrimSpace(p)] = true
	}

	scanner.Scan()
	for scanner.Scan() {
		desired_patterns = append(desired_patterns, scanner.Text())
	}

	return available_patterns, desired_patterns
}

func check_pattern(pattern string, available *map[string]bool) bool {
	if len(pattern) == 0 {
		return true
	} else {
		_, ok := (*available)[pattern]

		if ok {
			return true
		}

		found_path := false

		for i := 1; i < len(pattern); i++ {
			prefix := pattern[:i]
			postfix := pattern[i:]

			valid := check_pattern(postfix, available)
			if valid {
				(*available)[postfix] = true
				_, ok := (*available)[prefix]
				if ok {
					found_path = true
					(*available)[prefix] = true
				}

			}
		}
		if found_path {
			(*available)[pattern] = true
		}
		return found_path
	}
}

func part1(file_name string) {
	available, desired := read_input(file_name)

	valid_designs := 0

	for _, cur_design := range desired {
		fmt.Println(cur_design)

		// sub_pattern_len := 1

		// search_stack := make([]string, 0)
		// search_stack = append(search_stack, cur_design[:1])

		// for len(search_stack) > 0 {
		// 	cur_pattern := Pop(&search_stack)

		// }
		is_valid := check_pattern(cur_design, &available)

		if is_valid {
			valid_designs++
		}
	}

	fmt.Printf("P.1: %d\n", valid_designs)
}
func part2(file_name string) {}

func main() {
	file_name := "input.txt"

	part1(file_name)
	part2(file_name)
}
