package aoc

import "fmt"

type Position struct {
	R int
	C int
}

func (p Position) String() string {
	return fmt.Sprintf(`{R:%d, C:%d}`, p.R, p.C)
}

func (p1 Position) Add(p2 Position) Position {
	return Position{R: p1.R + p2.R, C: p1.C + p2.C}
}
func (p Position) Mult(coeff int) Position {
	return Position{R: p.R * coeff, C: p.C * coeff}
}

func (p1 Position) Equal(p2 Position) bool {
	return p1.R == p2.R && p1.C == p2.C
}

var UP = Position{C: 0, R: -1}
var DOWN = Position{C: 0, R: 1}
var LEFT = Position{C: -1, R: 0}
var RIGHT = Position{C: 1, R: 0}

var UL = Position{C: -1, R: -1}
var UR = Position{C: 1, R: -1}
var DR = Position{C: 1, R: 1}
var DL = Position{C: -1, R: 1}

var TURN_ORDER = [4]Position{
	UP,
	RIGHT,
	DOWN,
	LEFT,
}

var TURN_MAP = map[Position]int{
	UP:    0,
	RIGHT: 1,
	DOWN:  2,
	LEFT:  3,
}

var ADJACENTS = [8]Position{
	UP,
	UR,
	RIGHT,
	DR,
	DOWN,
	DL,
	LEFT,
	UL,
}

func Turn(cur_dir Position, dir Position) Position {
	dir_ind := TURN_MAP[cur_dir]
	switch dir {
	case RIGHT:
		dir_ind = (dir_ind + 1) % 4
		return TURN_ORDER[dir_ind]
	case LEFT:
		dir_ind = (dir_ind - 1)
		if dir_ind < 0 {
			dir_ind = 3
		}
		return TURN_ORDER[dir_ind]
	case UP:
		return cur_dir
	case DOWN:
		dir_ind = (dir_ind + 2) % 4
		return TURN_ORDER[dir_ind]
	}
	// TODO: This should return an error instead of UL
	return Position{R: -1, C: -1}
}
