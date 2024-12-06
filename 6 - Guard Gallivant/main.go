package main

import (
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

type Coordinate struct {
	r int
	c int
}

type Direction struct {
	X int
	Y int
}

var UP = Direction{0, -1}
var DOWN = Direction{0, 1}
var LEFT = Direction{-1, 0}
var RIGHT = Direction{1, 0}

func read_input(file_name string) (Coordinate, Direction, map[int][]int, map[int][]int, Coordinate, error) {
	file, err := os.Open(file_name)
	check(err)
	defer file.Close()

	var start_pos Coordinate
	var start_dir Direction
	var map_dim Coordinate
	row_obstructions := make(map[int][]int)
	col_obstructions := make(map[int][]int)

	scanner := bufio.NewScanner(file)
	r := 0
	for scanner.Scan() {
		for c, v := range scanner.Text() {
			if v == '^' {
				start_pos = Coordinate{r, c}
				start_dir = UP
			} else if v == '>' {
				start_pos = Coordinate{r, c}
				start_dir = RIGHT
			} else if v == 'v' {
				start_pos = Coordinate{r, c}
				start_dir = DOWN
			} else if v == '<' {
				start_pos = Coordinate{r, c}
				start_dir = LEFT
			} else if v == '#' {
				row_obstructions[r] = append(row_obstructions[r], c)
				col_obstructions[c] = append(col_obstructions[c], r)
			}
			map_dim.c = c + 1
		}

		r += 1
		map_dim.r = r
	}
	return start_pos, start_dir, row_obstructions, col_obstructions, map_dim, scanner.Err()
}

func search_ind(v int, arr []int) (int, int) {
	if v < arr[0] {
		return -1, arr[0]
	} else if v > arr[len(arr)-1] {
		return arr[len(arr)-1], -1
	} else {
		for i, n := range arr {
			if n < v && arr[i+1] > v {
				return n, arr[i+1]
			}
		}
	}
	return -1, -1
}

func find_nearest_obs(pos Coordinate, dir Direction, row_obs map[int][]int, col_obs map[int][]int) (Coordinate, bool) {
	if dir == UP {
		l, _ := search_ind(pos.r, col_obs[pos.c])
		if l == -1 {
			return Coordinate{}, true
		}
		return Coordinate{l, pos.c}, false
	} else if dir == DOWN {
		_, r := search_ind(pos.r, col_obs[pos.c])
		if r == -1 {
			return Coordinate{}, true
		}
		return Coordinate{r, pos.c}, false
	} else if dir == LEFT {
		l, _ := search_ind(pos.c, row_obs[pos.r])
		if l == -1 {
			return Coordinate{}, true
		}
		return Coordinate{pos.r, l}, false
	} else if dir == RIGHT {
		_, r := search_ind(pos.c, row_obs[pos.r])
		if r == -1 {
			return Coordinate{}, true
		}
		return Coordinate{pos.r, r}, false
	}
	return Coordinate{}, true
}

func part1(file_name string) map[Coordinate]int {
	start_pos, start_dir, row_obs, col_obs, map_dim, err := read_input(file_name)
	check(err)

	cur_pos := start_pos
	cur_dir := start_dir
	seen := make(map[Coordinate]int)

	for {
		obs, out_of_map := find_nearest_obs(cur_pos, cur_dir, row_obs, col_obs)
		if out_of_map {
			switch cur_dir {
			case UP:
				for r := cur_pos.r; r >= 0; r -= 1 {
					seen[Coordinate{r, cur_pos.c}] = 1
				}
			case RIGHT:
				for c := cur_pos.c; c < map_dim.c; c += 1 {
					seen[Coordinate{cur_pos.r, c}] = 1
				}
			case LEFT:
				for c := cur_pos.c; c >= 0; c -= 1 {
					seen[Coordinate{cur_pos.r, c}] = 1
				}
			case DOWN:
				for r := cur_pos.r; r < map_dim.r; r += 1 {
					seen[Coordinate{r, cur_pos.c}] = 1
				}
			}
			break
		}

		switch cur_dir {
		case UP:
			for r := cur_pos.r; r > obs.r; r -= 1 {
				seen[Coordinate{r, cur_pos.c}] = 1
			}
			cur_dir = RIGHT
			cur_pos = Coordinate{obs.r + 1, cur_pos.c}
		case RIGHT:
			for c := cur_pos.c; c < obs.c; c += 1 {
				seen[Coordinate{cur_pos.r, c}] = 1
			}
			cur_dir = DOWN
			cur_pos = Coordinate{cur_pos.r, obs.c - 1}
		case DOWN:
			for r := cur_pos.r; r < obs.r; r += 1 {
				seen[Coordinate{r, cur_pos.c}] = 1
			}
			cur_dir = LEFT
			cur_pos = Coordinate{obs.r - 1, cur_pos.c}
		case LEFT:
			for c := cur_pos.c; c > obs.c; c -= 1 {
				seen[Coordinate{cur_pos.r, c}] = 1
			}
			cur_dir = UP
			cur_pos = Coordinate{cur_pos.r, obs.c + 1}
		}
	}
	fmt.Printf("P.1: %d\n", len(seen))
	return seen
}

func part2(file_name string) {
	start_pos, start_dir, row_obs, col_obs, map_dim, err := read_input(file_name)
	check(err)

	cur_pos := start_pos
	cur_dir := start_dir
	type Key struct {
		coord Coordinate
		dir   Direction
	}
	seen := make(map[Key]int)
	var seen_in_order []Key

	for {
		obs, out_of_map := find_nearest_obs(cur_pos, cur_dir, row_obs, col_obs)
		if out_of_map {
			switch cur_dir {
			case UP:
				for r := cur_pos.r; r >= 0; r -= 1 {
					walked_pos := Coordinate{r, cur_pos.c}
					seen[Key{walked_pos, cur_dir}] = 1
					seen_in_order = append(seen_in_order, Key{walked_pos, cur_dir})
				}
			case RIGHT:
				for c := cur_pos.c; c < map_dim.c; c += 1 {
					walked_pos := Coordinate{cur_pos.r, c}
					seen[Key{walked_pos, cur_dir}] = 1
					seen_in_order = append(seen_in_order, Key{walked_pos, cur_dir})
				}
			case LEFT:
				for c := cur_pos.c; c >= 0; c -= 1 {
					walked_pos := Coordinate{cur_pos.r, c}
					seen[Key{walked_pos, cur_dir}] = 1
					seen_in_order = append(seen_in_order, Key{walked_pos, cur_dir})
				}
			case DOWN:
				for r := cur_pos.r; r < map_dim.r; r += 1 {
					walked_pos := Coordinate{r, cur_pos.c}
					seen[Key{walked_pos, cur_dir}] = 1
					seen_in_order = append(seen_in_order, Key{walked_pos, cur_dir})
				}
			}
			break
		}

		switch cur_dir {
		case UP:
			for r := cur_pos.r; r > obs.r; r -= 1 {
				walked_pos := Coordinate{r, cur_pos.c}
				seen[Key{walked_pos, cur_dir}] = 1
				seen_in_order = append(seen_in_order, Key{walked_pos, cur_dir})
			}
			cur_dir = RIGHT
			cur_pos = Coordinate{obs.r + 1, cur_pos.c}
		case RIGHT:
			for c := cur_pos.c; c < obs.c; c += 1 {
				walked_pos := Coordinate{cur_pos.r, c}
				seen[Key{walked_pos, cur_dir}] = 1
				seen_in_order = append(seen_in_order, Key{walked_pos, cur_dir})
			}
			cur_dir = DOWN
			cur_pos = Coordinate{cur_pos.r, obs.c - 1}
		case DOWN:
			for r := cur_pos.r; r < obs.r; r += 1 {
				walked_pos := Coordinate{r, cur_pos.c}
				seen[Key{walked_pos, cur_dir}] = 1
				seen_in_order = append(seen_in_order, Key{walked_pos, cur_dir})
			}
			cur_dir = LEFT
			cur_pos = Coordinate{obs.r - 1, cur_pos.c}
		case LEFT:
			for c := cur_pos.c; c > obs.c; c -= 1 {
				walked_pos := Coordinate{cur_pos.r, c}
				seen[Key{walked_pos, cur_dir}] = 1
				seen_in_order = append(seen_in_order, Key{walked_pos, cur_dir})
			}
			cur_dir = UP
			cur_pos = Coordinate{cur_pos.r, obs.c + 1}
		}
	}
	var potential_obs int
	for _, key := range seen_in_order {
		pos, dir := key.coord, key.dir
		switch dir {
		case UP:
			if _, ok := seen[Key{pos, RIGHT}]; ok {
				potential_obs += 1
			}
		case RIGHT:
		case DOWN:
		case LEFT:
			if _, ok := seen[Key{pos, UP}]; ok {
				potential_obs += 1
			}
		}
	}

	fmt.Printf("P.2: %d\n", potential_obs)
}

func main() {
	file_name := "test.txt"
	part1(file_name)
	part2(file_name)
}
