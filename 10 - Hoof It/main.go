package main

import (
	. "aoc"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

func read_input(file_name string) [][]int {
	file, err := os.Open(file_name)
	CheckErr(err)

	trail_map := make([][]int, 0)

	scanner := bufio.NewScanner(file)

	r := 0
	for scanner.Scan() {
		trail_map = append(trail_map, make([]int, 0))
		for _, v := range scanner.Text() {
			height, err := strconv.Atoi(string(v))
			if err != nil {
				trail_map[r] = append(trail_map[r], -1)
			} else {
				trail_map[r] = append(trail_map[r], height)
			}
		}
		r++
	}

	return trail_map
}

func find_adjacent_inbounds(pos Position, trail_map *[][]int) []Position {
	climbable := make([]Position, 0)
	if Inbounds(pos.Add(UP), trail_map) {
		climbable = append(climbable, UP)
	}
	if Inbounds(pos.Add(RIGHT), trail_map) {
		climbable = append(climbable, RIGHT)
	}
	if Inbounds(pos.Add(DOWN), trail_map) {
		climbable = append(climbable, DOWN)
	}
	if Inbounds(pos.Add(LEFT), trail_map) {
		climbable = append(climbable, LEFT)
	}
	return climbable
}

type Step struct {
	prev Position
	cur  Position
}

func walk_trail(pos Position, trail_map *[][]int, results chan<- int, wg *sync.WaitGroup, all bool) {
	defer wg.Done()

	walk_stack := make([]Step, 0)
	walk_stack = append(walk_stack, Step{prev: Position{R: -2, C: -2}, cur: pos})

	seen_peaks := make(map[Position]bool)

	for len(walk_stack) != 0 {
		cur_step := Pop(&walk_stack)

		prev_pos := cur_step.prev
		cur_pos := cur_step.cur
		cur_height := (*trail_map)[cur_pos.R][cur_pos.C]

		if cur_height == 9 {
			if !all {
				seen_peaks[cur_pos] = true
			} else {
				results <- 1
			}
		}

		look_ahead := find_adjacent_inbounds(cur_pos, trail_map)
		for _, dir := range look_ahead {

			next_pos := cur_pos.Add(dir)

			if !next_pos.Equal(prev_pos) && cur_height+1 == (*trail_map)[next_pos.R][next_pos.C] {
				walk_stack = append(walk_stack, Step{prev: cur_pos, cur: next_pos})
			}
		}
	}
	if !all {
		results <- len(seen_peaks)
	}

}

func part1(file_name string) {
	trail_map := read_input(file_name)

	var wg sync.WaitGroup
	results := make(chan int)

	for r := 0; r < len(trail_map); r++ {
		for c := 0; c < len(trail_map); c++ {
			if trail_map[r][c] == 0 {
				wg.Add(1)
				go walk_trail(Position{R: r, C: c}, &trail_map, results, &wg, false)
			}
		}
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	trail_score := 0
	for result := range results {
		trail_score += result
	}

	fmt.Printf("P.1: %d\n", trail_score)
}

func part2(file_name string) {
	trail_map := read_input(file_name)

	var wg sync.WaitGroup
	results := make(chan int)

	for r := 0; r < len(trail_map); r++ {
		for c := 0; c < len(trail_map); c++ {
			if trail_map[r][c] == 0 {
				wg.Add(1)
				go walk_trail(Position{R: r, C: c}, &trail_map, results, &wg, true)
			}
		}
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	trail_score := 0
	for result := range results {
		trail_score += result
	}

	fmt.Printf("P.2: %d\n", trail_score)
}

func main() {
	file_name := "input.txt"
	part1(file_name)
	part2(file_name)
}
