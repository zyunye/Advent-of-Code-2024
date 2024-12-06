package main

import (
	. "aoc"
	"bufio"
	"fmt"
	"os"
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

func walk(pos Coordinate, dir Direction, lab [][]rune) (Coordinate, bool) {
	new_coord := Coordinate{R: pos.R + dir.Y, C: pos.C + dir.X}
	if new_coord.R < 0 || new_coord.R >= len(lab) {
		return Coordinate{R: -1, C: -1}, true
	} else if new_coord.C < 0 || new_coord.C >= len(lab[new_coord.R]) {
		return Coordinate{R: -1, C: -1}, true
	}
	return new_coord, false
}

func part1(file_name string) {
	start_pos, lab := read_input(file_name)

	cur_pos := start_pos
	cur_dir := UP
	seen := make(map[Coordinate]int)
	for {
		seen[cur_pos] = 1
		next_pos, oob := walk(cur_pos, cur_dir, lab)

		if oob {
			break
		}

		if lab[next_pos.R][next_pos.C] == '#' {
			cur_dir = Turn(cur_dir, RIGHT)
		} else {
			cur_pos = next_pos
		}
	}
	fmt.Printf("P.1: %d\n", len(seen))
}

func part2(file_name string) {

}

func main() {
	file_name := "input.txt"
	part1(file_name)
	part2(file_name)
}
