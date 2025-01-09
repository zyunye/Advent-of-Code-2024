package main

import (
	. "aoc"
	"bufio"
	"fmt"
	"os"
)

func print_warehouse(w *[][]string) {
	for _, row := range *w {
		fmt.Println(row)
	}
}

func print_instruction(dir Position) {
	if dir == UP {
		fmt.Println("^")
	} else if dir == RIGHT {
		fmt.Println(">")
	} else if dir == DOWN {
		fmt.Println("v")
	} else if dir == LEFT {
		fmt.Println("<")
	}
}

func read_input_p1(file_name string) (warehouse [][]string, instructions []Position, robot_pos Position) {
	file, err := os.Open(file_name)
	CheckErr(err)

	scanner := bufio.NewScanner(file)

	warehouse = make([][]string, 0)
	instructions = make([]Position, 0)

	// Read in warehouse map
	r := 0
	for scanner.Scan() {

		row := scanner.Text()

		// If we hit the empty new line, exit so we can read instructions
		if len(row) == 0 {
			break
		}

		warehouse = append(warehouse, make([]string, 0))

		for c, v := range row {
			if v == '#' {
				warehouse[r] = append(warehouse[r], "#")
			} else if v == '@' {
				robot_pos = Position{R: r, C: c}
				warehouse[r] = append(warehouse[r], "@")
			} else if v == 'O' {
				warehouse[r] = append(warehouse[r], "O")
			} else if v == '.' {
				warehouse[r] = append(warehouse[r], ".")
			}
		}
		r++
	}

	// Read and parse robot instructions
	for scanner.Scan() {
		for _, v := range scanner.Text() {
			if v == '^' {
				instructions = append(instructions, UP)
			} else if v == '>' {
				instructions = append(instructions, RIGHT)
			} else if v == 'v' {
				instructions = append(instructions, DOWN)
			} else if v == '<' {
				instructions = append(instructions, LEFT)
			}
		}
	}

	return warehouse, instructions, robot_pos
}

func swap1(p1, p2 Position, warehouse *[][]string) {
	(*warehouse)[p1.R][p1.C], (*warehouse)[p2.R][p2.C] = (*warehouse)[p2.R][p2.C], (*warehouse)[p1.R][p1.C]
}

func move1(tile_pos Position, dir Position, warehouse *[][]string) (Position, bool) {
	tile := Get(tile_pos, warehouse)
	if tile == "#" {
		return tile_pos, false
	} else if tile == "." {
		return tile_pos, true
	} else if tile == "@" || tile == "O" {
		neighbor := tile_pos.Add(dir)
		if _, move_ok := move1(neighbor, dir, warehouse); move_ok {
			swap1(tile_pos, neighbor, warehouse)
			return neighbor, true
		} else {
			return tile_pos, false
		}
	} else {
		panic(fmt.Sprintf("Move was called on a weird tile: %s, %v", tile, tile_pos))
	}
}

func part1(file_name string) {
	warehouse, instructions, robot_pos := read_input_p1(file_name)

	for _, instruction := range instructions {
		robot_pos, _ = move1(robot_pos, instruction, &warehouse)
	}

	gps_sum := 0
	for r, row := range warehouse {
		for c, col := range row {
			if col == "O" {
				gps_sum += 100*r + c
			}
		}
	}

	fmt.Printf("P.1: %d\n", gps_sum)

}

func read_input_p2(file_name string) (warehouse [][]string, instructions []Position, robot_pos Position) {
	file, err := os.Open(file_name)
	CheckErr(err)

	scanner := bufio.NewScanner(file)

	warehouse = make([][]string, 0)
	instructions = make([]Position, 0)

	// Read in warehouse map
	r := 0
	for scanner.Scan() {

		row := scanner.Text()

		// If we hit the empty new line, exit so we can read instructions
		if len(row) == 0 {
			break
		}

		warehouse = append(warehouse, make([]string, 0))

		for _, v := range row {
			if v == '#' {
				warehouse[r] = append(warehouse[r], "#")
				warehouse[r] = append(warehouse[r], "#")
			} else if v == '@' {
				warehouse[r] = append(warehouse[r], "@")
				warehouse[r] = append(warehouse[r], ".")
			} else if v == 'O' {
				warehouse[r] = append(warehouse[r], "[")
				warehouse[r] = append(warehouse[r], "]")
			} else if v == '.' {
				warehouse[r] = append(warehouse[r], ".")
				warehouse[r] = append(warehouse[r], ".")
			}
		}
		r++
	}

	// Read and parse robot instructions
	for scanner.Scan() {
		for _, v := range scanner.Text() {
			if v == '^' {
				instructions = append(instructions, UP)
			} else if v == '>' {
				instructions = append(instructions, RIGHT)
			} else if v == 'v' {
				instructions = append(instructions, DOWN)
			} else if v == '<' {
				instructions = append(instructions, LEFT)
			}
		}
	}

	// Retrack robot position
	for r, row := range warehouse {
		for c, col := range row {
			if col == "@" {
				robot_pos = Position{R: r, C: c}
			}
		}
	}

	return warehouse, instructions, robot_pos
}

func move2(tile_pos Position, dir Position, warehouse *[][]string) (Position, bool) {
	tile := Get(tile_pos, warehouse)
	if tile == "#" {
		return tile_pos, false

	} else if tile == "." {
		return tile_pos, true

	} else if tile == "@" {
		neighbor := tile_pos.Add(dir)
		if _, move_ok := move2(neighbor, dir, warehouse); move_ok {
			swap1(tile_pos, neighbor, warehouse)
			return neighbor, true
		} else {
			return tile_pos, false
		}

	} else if tile == "[" || tile == "]" {
		var left_pos Position
		var right_pos Position

		if tile == "[" {
			left_pos = tile_pos
			right_pos = tile_pos.Add(RIGHT)
		} else if tile == "]" {
			left_pos = tile_pos.Add(LEFT)
			right_pos = tile_pos
		}

		left_neighbor := left_pos.Add(dir)
		right_neighbor := right_pos.Add(dir)

		if dir == LEFT {
			if _, move_ok := move2(left_neighbor, dir, warehouse); move_ok {
				swap1(left_pos, left_neighbor, warehouse)
				swap1(right_pos, right_neighbor, warehouse)
				return left_neighbor, true
			} else {
				return tile_pos, false
			}
		} else if dir == RIGHT {
			if _, move_ok := move2(right_neighbor, dir, warehouse); move_ok {
				swap1(right_pos, right_neighbor, warehouse)
				swap1(left_pos, left_neighbor, warehouse)
				return right_neighbor, true
			} else {
				return tile_pos, false
			}
		} else {
			// Up Down movement
			_check_seen := func(pos Position, arr *[]Position) bool {
				for _, v := range *arr {
					if pos == v {
						return true
					}
				}
				return false
			}

			required_moves := make([]Position, 0)

			required_moves = append(required_moves, tile_pos)
			head := 0

			for head < len(required_moves) {
				cur_pos := required_moves[head]
				cur_val := Get(cur_pos, warehouse)

				if cur_val == "#" {
					return tile_pos, false
				} else if cur_val == "[" {
					if !_check_seen(cur_pos.Add(RIGHT), &required_moves) {
						required_moves = append(required_moves, cur_pos.Add(RIGHT))
					}
					if !_check_seen(cur_pos.Add(dir), &required_moves) && Get(cur_pos.Add(dir), warehouse) != "." {
						required_moves = append(required_moves, cur_pos.Add(dir))
					}
				} else if cur_val == "]" {
					if !_check_seen(cur_pos.Add(LEFT), &required_moves) {
						required_moves = append(required_moves, cur_pos.Add(LEFT))
					}
					if !_check_seen(cur_pos.Add(dir), &required_moves) && Get(cur_pos.Add(dir), warehouse) != "." {
						required_moves = append(required_moves, cur_pos.Add(dir))
					}
				}
				head += 1
			}

			for i := len(required_moves) - 1; i >= 0; i-- {
				pos := required_moves[i]
				swap1(pos, pos.Add(dir), warehouse)
			}

			return tile_pos, true
		}

	} else {
		panic(fmt.Sprintf("Move was called on a weird tile: %s, %v", tile, tile_pos))
	}
	panic("EOL on move2. Shouldn't be possible:")
}

func part2(file_name string) {
	warehouse, instructions, robot_pos := read_input_p2(file_name)
	// print_warehouse(&warehouse)
	// fmt.Println()

	for _, instruction := range instructions {
		// fmt.Printf("%d ", i)
		// print_instruction(instruction)
		robot_pos, _ = move2(robot_pos, instruction, &warehouse)
		// print_warehouse(&warehouse)
		// fmt.Println()
	}

	gps_sum := 0
	for r, row := range warehouse {
		for c, col := range row {
			if col == "[" {
				gps_sum += 100*r + c
			}
		}
	}

	fmt.Printf("P.2: %d\n", gps_sum)

}
func main() {
	file_name := "input.txt"
	part1(file_name)
	part2(file_name)
}
