package main

import (
	. "aoc"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Robot struct {
	start    Position
	velocity Position
}

func (r Robot) String() string {
	return fmt.Sprintf(`
Start: %v
Velocity: %v
`, r.start, r.velocity)
}

func read_input(file_name string) []Robot {
	file, err := os.Open(file_name)
	CheckErr(err)

	scanner := bufio.NewScanner(file)

	ret := make([]Robot, 0)

	for scanner.Scan() {
		tokens := strings.Fields(scanner.Text())
		pos_token := tokens[0][2:]
		vel_token := tokens[1][2:]

		pos_vals := strings.Split(pos_token, ",")
		vel_vals := strings.Split(vel_token, ",")

		px, err := strconv.Atoi(pos_vals[0])
		CheckErr(err)
		py, err := strconv.Atoi(pos_vals[1])
		CheckErr(err)
		vx, err := strconv.Atoi(vel_vals[0])
		CheckErr(err)
		vy, err := strconv.Atoi(vel_vals[1])
		CheckErr(err)

		ret = append(ret, Robot{
			start:    Position{R: py, C: px},
			velocity: Position{R: vy, C: vx},
		})
	}

	return ret

}

func part1(file_name string) {
	robots := read_input(file_name)

	bathroom_r := 103
	bathroom_c := 101
	iters := 100

	mid_r := bathroom_r / 2
	mid_c := bathroom_c / 2

	quadrants := make(map[string]int)

	for _, robot := range robots {
		end_pos := robot.start.Add(robot.velocity.Mult(iters))
		end_pos.C = end_pos.C % bathroom_c
		end_pos.R = end_pos.R % bathroom_r

		if end_pos.C < 0 {
			end_pos.C = bathroom_c + end_pos.C
		}
		if end_pos.R < 0 {
			end_pos.R = bathroom_r + end_pos.R
		}

		if end_pos.R < mid_r && end_pos.C > mid_c {
			quadrants["Q1"] += 1
		} else if end_pos.R < mid_r && end_pos.C < mid_c {
			quadrants["Q2"] += 1
		} else if end_pos.R > mid_r && end_pos.C < mid_c {
			quadrants["Q3"] += 1
		} else if end_pos.R > mid_r && end_pos.C > mid_c {
			quadrants["Q4"] += 1
		}

	}

	safety_factor := 1
	for _, v := range quadrants {
		safety_factor *= v
	}

	fmt.Printf("P.1: %d", safety_factor)

}
func part2(file_name string) {
	robots := read_input(file_name)

	bathroom_r := 103
	bathroom_c := 101
	iters := 100

	mid_r := bathroom_r / 2
	mid_c := bathroom_c / 2

}
func main() {
	file_name := "input.txt"
	part1(file_name)
	part2(file_name)
}
