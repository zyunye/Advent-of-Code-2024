package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

func read_input(file_name string) (map[int][]int, [][]int, error) {
	file, err := os.Open(file_name)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	orderings := make(map[int][]int)
	var updates [][]int
	scan_state := 0

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			scan_state = 1
			continue
		}

		if scan_state == 0 {
			tokens := strings.Split(line, "|")
			t1, err := strconv.Atoi(tokens[0])
			check(err)
			t2, err := strconv.Atoi(tokens[1])
			check(err)

			if _, ok := orderings[t1]; ok {
				orderings[t1] = append(orderings[t1], t2)
			} else {
				orderings[t1] = make([]int, 1)
				orderings[t1] = append(orderings[t1], t2)
			}

		} else {
			tokens := strings.Split(line, ",")
			var int_tokens []int
			for _, t := range tokens {
				t, err := strconv.Atoi(t)
				check(err)
				int_tokens = append(int_tokens, t)
			}

			updates = append(updates, int_tokens)
		}
	}

	return orderings, updates, scanner.Err()
}

func contains[T comparable](t T, slice []T) bool {
	for _, v := range slice {
		if v == t {
			return true
		}
	}
	return false
}

func traverse_pt1(update []int, orderings map[int][]int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i, num := range update {
		if i == 0 {
			continue
		}

		parent := update[i-1]
		if !contains(num, orderings[parent]) {
			return
		}
	}

	results <- update[len(update)/2]
}

func part1(file_name string) {
	orderings, updates, err := read_input(file_name)
	check(err)

	middle_sums := 0
	var wg sync.WaitGroup
	results := make(chan int)

	for _, update := range updates {
		wg.Add(1)
		go traverse_pt1(update, orderings, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		middle_sums += result
	}

	fmt.Printf("P.1: %d\n", middle_sums)
}

func insert_at[T any](slice []T, index int, v T) []T {
	slice = append(slice[:index], append([]T{v}, slice[index:]...)...)
	return slice
}

func traverse_pt2(update []int, orderings map[int][]int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	fixed := false
	for i, num := range update {
		if i == 0 {
			continue
		}
		parent := update[i-1]

		if !contains(num, orderings[parent]) {
			slice := append(update[:i], update[i+1:]...)
			for new_ind, n := range(slice) {
				if contains(n, orderings[num]) {
					slice = insert_at(slice, new_ind, num)
					break
				}
			}
			update = slice
			fixed = true
		}
	}

	if fixed {
		results <- update[len(update)/2]
	}

}

func part2(file_name string) {
	orderings, updates, err := read_input(file_name)
	check(err)

	middle_sums := 0
	var wg sync.WaitGroup
	results := make(chan int)

	for _, update := range updates {
		wg.Add(1)
		go traverse_pt2(update, orderings, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		middle_sums += result
	}

	fmt.Printf("P.2: %d\n", middle_sums)
}

func main() {
	file_name := "input.txt"
	part1(file_name)
	part2(file_name)
}
