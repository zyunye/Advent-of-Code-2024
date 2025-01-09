package main

import (
	. "aoc"
	"bufio"
	"bytes"
	"fmt"
	"os"
)

type Tile struct {
	v string
	p Position
}

func (t *Tile) String() string {
	return t.v
}
func (t *Tile) Move(dir Position, warehouse *[][]Tile) {
	if t.Shove(dir, warehouse) {
		move_to := t.p.Add(dir)
		(*warehouse)[move_to.R][move_to.C] = *t
		t.p = move_to
	}
}

func (t *Tile) Shove(dir Position, warehouse *[][]Tile) bool {
	if t.v == "#" {
		return false
	} else if t.v == "." {
		return true
	} else if t.v == "O" || t.v == "@" {
		neighbor_pos := t.p.Add(dir)
		return (*warehouse)[neighbor_pos.R][neighbor_pos.C].Shove(dir, warehouse)
	} else {
		panic(fmt.Sprintf("Shove was called on a weird tile: %s, %v", t.v, t.p))
	}
}

func print_warehouse(w *[][]Tile) {
	var buffer bytes.Buffer

	for _, row := range *w {
		for _, col := range row {
			buffer.WriteString(col.String())
		}
		buffer.WriteString("\n")
	}
	fmt.Println(buffer.String())
}

func read_input(file_name string) (warehouse [][]Tile, instructions []Position, robot Tile) {
	file, err := os.Open(file_name)
	CheckErr(err)

	scanner := bufio.NewScanner(file)

	warehouse = make([][]Tile, 0)
	instructions = make([]Position, 0)

	// Read in warehouse map
	r := 0
	for scanner.Scan() {

		row := scanner.Text()

		// If we hit the empty new line, exit so we can read instructions
		if len(row) == 0 {
			break
		}

		warehouse = append(warehouse, make([]Tile, 0))

		for c, v := range row {
			if v == '#' {
				warehouse[r] = append(warehouse[r], Tile{p: Position{R: r, C: c}, v: "#"})
			} else if v == '@' {
				robot = Tile{p: Position{R: r, C: c}, v: "@"}
				warehouse[r] = append(warehouse[r], robot)
			} else if v == 'O' {
				warehouse[r] = append(warehouse[r], Tile{p: Position{R: r, C: c}, v: "O"})
			} else if v == '.' {
				warehouse[r] = append(warehouse[r], Tile{p: Position{R: r, C: c}, v: "."})
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

	return warehouse, instructions, robot
}

func part1(file_name string) {
	warehouse, instructions, robot := read_input(file_name)
	print_warehouse(&warehouse)
	fmt.Println(instructions)
	fmt.Println(robot)

	// for _, instruction := range instructions {

	// }
}
func part2(file_name string) {
}
func main() {
	file_name := "test.txt"
	part1(file_name)
	part2(file_name)
}
