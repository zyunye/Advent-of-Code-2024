package aoc

type Direction struct {
	X int
	Y int
}

var UP = Direction{0, -1}
var DOWN = Direction{0, 1}
var LEFT = Direction{-1, 0}
var RIGHT = Direction{1, 0}

var TURN_ORDER = [4]Direction{
	UP,
	RIGHT,
	DOWN,
	LEFT,
}

var TURN_MAP = map[Direction]int{
	UP:    0,
	RIGHT: 1,
	DOWN:  2,
	LEFT:  3,
}

func Turn(cur_dir Direction, dir Direction) Direction {
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
	return Direction{-1, -1}
}
