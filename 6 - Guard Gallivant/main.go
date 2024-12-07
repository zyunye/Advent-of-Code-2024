package main

import (
	. "aoc"
	"bufio"
	"fmt"
	"os"
	"sync"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

func read_input(file_name string) (Coordinate, [][]rune) {
	file, err := os.Open(file_name)
	check(err)
	defer file.Close()

	var start_pos Coordinate
	var lab [][]rune

	scanner := bufio.NewScanner(file)
	r := 0
	for scanner.Scan() {
		line := scanner.Text()
		lab = append(lab, make([]rune, len(line)))
		for c, v := range line {
			lab[r][c] = v
			if v == '^' {
				start_pos = Coordinate{R: r, C: c}
			}
		}
		r += 1
	}
	check(scanner.Err())
	return start_pos, lab
}

func step(pos Coordinate, dir Direction, lab [][]rune) (Coordinate, bool) {
	new_coord := Coordinate{R: pos.R + dir.Y, C: pos.C + dir.X}
	if new_coord.R < 0 || new_coord.R >= len(lab) {
		return Coordinate{R: -1, C: -1}, true
	} else if new_coord.C < 0 || new_coord.C >= len(lab[new_coord.R]) {
		return Coordinate{R: -1, C: -1}, true
	}
	return new_coord, false
}

func traverse(start Coordinate, lab [][]rune) <-chan Coordinate {
	ch := make(chan Coordinate)

	go func() {
		defer close(ch)
		cur_pos := start
		cur_dir := UP

		for {
			ch <- cur_pos
			next_pos, oob := step(cur_pos, cur_dir, lab)

			if oob {
				return
			}

			if lab[next_pos.R][next_pos.C] == '#' {
				cur_dir = Turn(cur_dir, RIGHT)
			} else {
				cur_pos = next_pos
			}
		}
	}()

	return ch
}

func part1(file_name string) {
	start_pos, lab := read_input(file_name)
	path_gen := traverse(start_pos, lab)

	seen := make(map[Coordinate]int)
	for coord := range path_gen {
		seen[coord] = 1
	}
	fmt.Printf("P.1: %d\n", len(seen))
}

func copy_arr(arr [][]rune) [][]rune {
	new_arr := make([][]rune, len(arr))
	for r, row := range arr {
		new_arr[r] = make([]rune, len(row))
		copy(new_arr[r], row)
	}
	return new_arr
}

func loop_checker(start_pos Coordinate, lab [][]rune, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	cur_path := traverse(start_pos, lab)
	iter_max := len(lab) * len(lab[0])
	cur_iter := 0

	seen := make(map[Coordinate]int)

	for coord := range cur_path {
		if cur_iter >= iter_max {
			results <- 1
			return
		}
		if _, ok := seen[coord]; ok {
			cur_iter += 1
		} else {
			seen[coord] = 1
		}
	}
}

func part2(file_name string) {
	start_pos, lab := read_input(file_name)

	path_gen := traverse(start_pos, lab)
	orig_path := make(map[Coordinate]int)
	for coord := range path_gen {
		orig_path[coord] = 1
	}

	found_loops := 0

	var wg sync.WaitGroup
	results := make(chan int)

	for seen_coord, _ := range orig_path {
		if !(seen_coord.R == start_pos.R && seen_coord.C == start_pos.C) {
			if lab[seen_coord.R][seen_coord.C] == '.' {
				new_lab := copy_arr(lab)
				new_lab[seen_coord.R][seen_coord.C] = '#'
				wg.Add(1)
				go loop_checker(start_pos, new_lab, results, &wg)
			}
		}
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		found_loops += result
	}

	fmt.Printf("P.2: %d\n", found_loops)
}

func main() {
	file_name := "input.txt"
	part1(file_name)
	part2(file_name)
}
