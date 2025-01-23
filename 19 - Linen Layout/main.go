package main

import (
	. "aoc"
	"bufio"
	"fmt"
	"os"
	"sort"
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

type SortedMap struct {
	m           map[string]int
	inv_m       map[int][]string
	sorted_keys []string
}

func (sm *SortedMap) Sort(invert bool) {
	if invert {
		sort.Slice(sm.sorted_keys, func(i, j int) bool {
			return sm.m[sm.sorted_keys[i]] > sm.m[sm.sorted_keys[j]]
		})
	} else {
		sort.Slice(sm.sorted_keys, func(i, j int) bool {
			return sm.m[sm.sorted_keys[i]] < sm.m[sm.sorted_keys[j]]
		})
	}
}

func (sm *SortedMap) Add(k string, v int, invert_sort bool) {
	if _, ok := sm.m[k]; !ok {
		sm.m[k] = v

		if _, ok := sm.inv_m[v]; !ok {
			sm.inv_m[v] = make([]string, 0)
		}
		sm.inv_m[v] = append(sm.inv_m[v], k)

		sm.sorted_keys = append(sm.sorted_keys, k)
		sm.Sort(invert_sort)
	}
}

func check_pattern(pattern string, available *SortedMap, not_available *map[string]bool) bool {
	if len(pattern) == 0 {
		return true
	}
	if (*not_available)[pattern] {
		return false
	}

	_, ok := available.m[pattern]
	if ok {
		return true
	}

	for prefix_len := len(pattern) - 1; prefix_len > 0; prefix_len-- {

		// Check if the cache has a prefix of the length of the string we're currently looking for
		if _, ok := available.inv_m[prefix_len]; ok {

			// If we have cached strings of length `prefix_len`, check to see if our exact prefix exists
			if _, prefix_found := available.m[pattern[:prefix_len]]; prefix_found {

				// If our exact prefix exists, check to see if we can build the postfix
				postfix_valid := check_pattern(pattern[prefix_len:], available, not_available)

				if postfix_valid {
					// If our postfix also matches, cache this current string and return true
					available.Add(pattern, len(pattern), true)
					return true
				}
			}
		}
	}
	(*not_available)[pattern] = true
	return false
}

func part1(file_name string) {
	available, desired := read_input(file_name)
	not_available := make(map[string]bool)

	sorted_available := SortedMap{
		m:           make(map[string]int),
		inv_m:       make(map[int][]string),
		sorted_keys: make([]string, 0),
	}
	for k := range available {
		sorted_available.Add(k, len(k), true)
	}

	valid_designs := 0

	for _, cur_design := range desired {
		fmt.Println(cur_design)

		is_valid := check_pattern(cur_design, &sorted_available, &not_available)

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
