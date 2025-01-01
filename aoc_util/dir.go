package aoc

type Position struct {
	C int
	R int
}

func (p1 Position) Add(p2 Position) Position {
	return Position{C: p1.C + p2.C, R: p1.R + p2.R}
}

func (p1 Position) Equal(p2 Position) bool {
	return p1.C == p2.C && p1.R == p2.R
}

var UP = Position{0, -1}
var DOWN = Position{0, 1}
var LEFT = Position{-1, 0}
var RIGHT = Position{1, 0}

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
	return Position{-1, -1}
}
