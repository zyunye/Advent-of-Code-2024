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

	} else if tile == "[" {
		my_neighbor_pos := tile_pos.Add(dir)
		partner_pos := tile_pos.Add(RIGHT)

		if dir == LEFT {
			if _, my_move_ok := move2(my_neighbor_pos, dir, warehouse); my_move_ok {
				swap1(my_neighbor_pos, tile_pos, warehouse)
				swap1(tile_pos, partner_pos, warehouse)
				return my_neighbor_pos, true
			}

		} else if dir == RIGHT {
			return move2(partner_pos, dir, warehouse)

		} else {
			my_lookahead := Get(tile_pos.Add(dir), warehouse)
			partner_lookahead := Get(partner_pos.Add(dir), warehouse)

			if my_lookahead == "." && partner_lookahead == "." {
				me_towards := tile_pos.Add(dir)
				partner_towards := partner_pos.Add(dir)
				swap1(tile_pos, me_towards, warehouse)
				swap1(partner_pos, partner_towards, warehouse)
				return me_towards, true

			} else if my_lookahead == "#" || partner_lookahead == "#" {
				return tile_pos, false

			} else {
				_, my_move_ok := move2(my_neighbor_pos, dir, warehouse)
				if !my_move_ok {
					return tile_pos, false
				}
				_, partner_move_ok := move2(partner_pos, dir, warehouse)
				if !partner_move_ok {
					return tile_pos, false
				}

				me_towards := tile_pos.Add(dir)
				partner_towards := partner_pos.Add(dir)
				swap1(tile_pos, me_towards, warehouse)
				swap1(partner_pos, partner_towards, warehouse)
				return me_towards, true
			}
		}

	} else if tile == "]" {
		my_neighbor_pos := tile_pos.Add(dir)
		partner_pos := tile_pos.Add(LEFT)

		if dir == RIGHT {
			if _, my_move_ok := move2(my_neighbor_pos, dir, warehouse); my_move_ok {
				swap1(my_neighbor_pos, tile_pos, warehouse)
				swap1(tile_pos, partner_pos, warehouse)
				return my_neighbor_pos, true
			}

		} else if dir == LEFT {
			return move2(partner_pos, dir, warehouse)

		} else {
			my_lookahead := Get(tile_pos.Add(dir), warehouse)
			partner_lookahead := Get(partner_pos.Add(dir), warehouse)

			if my_lookahead == "." && partner_lookahead == "." {
				me_towards := tile_pos.Add(dir)
				partner_towards := partner_pos.Add(dir)
				swap1(tile_pos, me_towards, warehouse)
				swap1(partner_pos, partner_towards, warehouse)
				return me_towards, true

			} else if my_lookahead == "#" || partner_lookahead == "#" {
				return tile_pos, false

			} else {

				_, my_move_ok := move2(my_neighbor_pos, dir, warehouse)
				if !my_move_ok {
					return tile_pos, false
				}
				_, partner_move_ok := move2(partner_pos, dir, warehouse)
				if !partner_move_ok {
					return tile_pos, false
				}

				me_towards := tile_pos.Add(dir)
				partner_towards := partner_pos.Add(dir)
				swap1(tile_pos, me_towards, warehouse)
				swap1(partner_pos, partner_towards, warehouse)
				return me_towards, true
			}
		}
	} else {
		panic(fmt.Sprintf("Move was called on a weird tile: %s, %v", tile, tile_pos))
	}
	panic("EOL on move2. Shouldn't be possible:")
}

func part2(file_name string) {
	warehouse, instructions, robot_pos := read_input_p2(file_name)
	print_warehouse(&warehouse)
	fmt.Println()

	for _, instruction := range instructions {
		print_instruction(instruction)
		robot_pos, _ = move2(robot_pos, instruction, &warehouse)
		print_warehouse(&warehouse)
		fmt.Println()
	}

}
func main() {
	file_name := "test2.txt"
	part1(file_name)
	part2(file_name)
}
