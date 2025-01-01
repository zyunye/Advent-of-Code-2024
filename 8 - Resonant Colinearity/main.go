package main

import (
	. "aoc"
	"bufio"
	"fmt"
	"os"
)

func read_input(file_name string) (map[string][]Position, int, int) {
	file, err := os.Open(file_name)
	CheckErr(err)

	nodes := make(map[string][]Position)

	scanner := bufio.NewScanner(file)

	r := 0
	max_c := 0
	for scanner.Scan() {
		for c, v := range scanner.Text() {
			max_c = max(max_c, c)

			v := string(v)
			if v == "." {
				continue
			}
			if _, ok := nodes[v]; !ok {
				nodes[v] = make([]Position, 0)
				nodes[v] = append(nodes[v], Position{R: r, C: c})
			} else {
				nodes[v] = append(nodes[v], Position{R: r, C: c})
			}
		}
		r += 1
	}

	CheckErr(scanner.Err())
	return nodes, r, max_c + 1
}

func deltas(c1 Position, c2 Position) (int, int) {
	return c2.R - c1.R, c2.C - c1.C
}

func calc_antinodes(c1 Position, c2 Position, dr int, dc int) (Position, Position) {
	anti1 := Position{R: c1.R - dr, C: c1.C - dc}
	anti2 := Position{R: c2.R + dr, C: c2.C + dc}

	return anti1, anti2
}

func is_inbounds(coord Position, r int, c int) bool {
	if coord.R < 0 || coord.R >= r {
		return false
	} else if coord.C < 0 || coord.C >= c {
		return false
	}
	return true
}

func part1(file_name string) {
	nodes, map_r, map_c := read_input(file_name)

	antinodes := make(map[Position]int)

	for key := range nodes {

		coords := nodes[key]

		for i := 0; i < len(coords)-1; i++ {
			for j := i + 1; j < len(coords); j++ {

				dr, dc := deltas(coords[i], coords[j])
				anti1, anti2 := calc_antinodes(coords[i], coords[j], dr, dc)

				if is_inbounds(anti1, map_r, map_c) {
					antinodes[anti1] = 1
				}

				if is_inbounds(anti2, map_r, map_c) {
					antinodes[anti2] = 1
				}
			}
		}
	}

	fmt.Printf("P.1: %d\n", len(antinodes))
}

func part2(file_name string) {
	nodes, map_r, map_c := read_input(file_name)

	antinodes := make(map[Position]int)

	for key := range nodes {

		coords := nodes[key]

		for i := 0; i < len(coords)-1; i++ {
			for j := i + 1; j < len(coords); j++ {
				antinodes[coords[i]] = 1
				antinodes[coords[j]] = 1

				dr, dc := deltas(coords[i], coords[j])
				anti1, anti2 := calc_antinodes(coords[i], coords[j], dr, dc)
				prev_anti1 := coords[i]
				prev_anti2 := coords[j]

				for is_inbounds(anti1, map_r, map_c) {
					antinodes[anti1] = 1
					tmp := anti1
					anti1, _ = calc_antinodes(anti1, prev_anti1, dr, dc)
					prev_anti1 = tmp
				}

				for is_inbounds(anti2, map_r, map_c) {
					antinodes[anti2] = 1
					tmp := anti2
					_, anti2 = calc_antinodes(anti2, prev_anti2, dr, dc)
					prev_anti2 = tmp
				}
			}
		}
	}
	
	fmt.Printf("P.2: %d\n", len(antinodes))
}

func main() {
	file_name := "input.txt"
	part1(file_name)
	part2(file_name)
}
