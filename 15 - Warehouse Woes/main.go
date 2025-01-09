package main

import (
	. "aoc"
	"bufio"
	"fmt"
	"os"
)

// type Tile struct {
// 	v string
// 	p Position
// }

// func (t *Tile) String() string {
// 	return t.v
// }
// func (t *Tile) Move(dir Position, warehouse *[][]Tile) bool {
// 	neighbor_pos := t.p.Add(dir)
// 	neighbor := (*warehouse)[neighbor_pos.R][neighbor_pos.C]

// 	if neighbor.Shove(dir, warehouse) {
// 		move_to := t.p.Add(dir)
// 		(*warehouse)[t.p.R][t.p.C] = Tile{p: Position{R: t.p.R, C: t.p.C}, v: "."}
// 		(*warehouse)[move_to.R][move_to.C] = *t
// 		t.p = move_to
// 		return true
// 	}
// 	return false
// }

// func (t *Tile) Shove(dir Position, warehouse *[][]Tile) bool {
// 	if t.v == "#" {
// 		return false
// 	} else if t.v == "." {
// 		return true
// 	} else if t.v == "O" || t.v == "@" {
// 		// neighbor_pos := t.p.Add(dir)
// 		// if (*warehouse)[neighbor_pos.R][neighbor_pos.C].Shove(dir, warehouse) {
// 		// move_to := t.p.Add(dir)
// 		// (*warehouse)[t.p.R][t.p.C] = Tile{p: Position{R: t.p.R, C: t.p.C}, v: "."}
// 		// (*warehouse)[move_to.R][move_to.C] = *t
// 		// t.p = move_to
// 		// return true
// 		// 	return t.Move(dir, warehouse)
// 		// }
// 		// return false
// 		return t.Move(dir, warehouse)
// 	} else {
// 		panic(fmt.Sprintf("Shove was called on a weird tile: %s, %v", t.v, t.p))
// 	}
// }

// func print_warehouse(w *[][]Tile) {
// 	var buffer bytes.Buffer

// 	for _, row := range *w {
// 		for _, col := range row {
// 			buffer.WriteString(col.String())
// 		}
// 		buffer.WriteString("\n")
// 	}
// 	fmt.Println(buffer.String())
// }

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

func swap(p1, p2 Position, warehouse *[][]string) {
	(*warehouse)[p1.R][p1.C], (*warehouse)[p2.R][p2.C] = (*warehouse)[p2.R][p2.C], (*warehouse)[p1.R][p1.C]
}

func move(tile_pos Position, dir Position, warehouse *[][]string) (Position, bool) {
	tile := Get(tile_pos, warehouse)
	if tile == "#" {
		return tile_pos, false
	} else if tile == "." {
		return tile_pos, true
	} else if tile == "@" || tile == "O" {
		neighbor := tile_pos.Add(dir)
		if _, move_ok := move(neighbor, dir, warehouse); move_ok {
			swap(tile_pos, neighbor, warehouse)
			return neighbor, true
		} else {
			return tile_pos, false
		}
	} else {
		panic(fmt.Sprintf("Move was called on a weird tile: %s, %v", tile, tile_pos))
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

func part1(file_name string) {
	warehouse, instructions, robot_pos := read_input_p1(file_name)

	for _, instruction := range instructions {
		robot_pos, _ = move(robot_pos, instruction, &warehouse)
	}

	gps_sum := 0
	for r, row := range warehouse {
		for c, col := range row {
			if col == "O" {
				gps_sum += 100*r + c
			}
		}
	}

	fmt.Printf("P.1: %d", gps_sum)

}
func part2(file_name string) {
}
func main() {
	file_name := "input.txt"
	part1(file_name)
	part2(file_name)
}
