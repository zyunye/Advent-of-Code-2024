package main

import (
	. "aoc"
	"bufio"
	"fmt"
	"math"
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

func (r *Robot) CalcEndPos(iters, max_r, max_c int) Position {
	end_pos := r.start.Add(r.velocity.Mult(iters))
	end_pos.C = end_pos.C % max_c
	end_pos.R = end_pos.R % max_r

	if end_pos.C < 0 {
		end_pos.C = max_c + end_pos.C
	}
	if end_pos.R < 0 {
		end_pos.R = max_r + end_pos.R
	}

	return end_pos
}

func (r *Robot) ModifyStartPos(iters, max_r, max_c int) {
	end_pos := r.CalcEndPos(iters, max_r, max_c)
	r.start = end_pos
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
		end_pos := robot.CalcEndPos(iters, bathroom_r, bathroom_c)

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

	fmt.Printf("P.1: %d\n", safety_factor)
}

func dump_bathroom(bathroom *[][]int, iter int) {
	file, err := os.OpenFile("out.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	CheckErr(err)
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString(fmt.Sprintf("Seconds: %d\n", iter))
	for _, row := range *bathroom {
		for _, col := range row {
			val := strconv.Itoa(col)

			if val == "0" {
				_, err := writer.WriteString(".")
				CheckErr(err)
			} else {
				_, err := writer.WriteString(val)
				CheckErr(err)
			}

		}
		writer.WriteString("\n")
	}
	writer.WriteString("\n")
	err = writer.Flush()
	CheckErr(err)
}

func part2(file_name string) {
	robots := read_input(file_name)

	bathroom_r := 103
	bathroom_c := 101

	mid_r := bathroom_r / 2
	mid_c := bathroom_c / 2

	safety_scores := make([]int, 0)

	// Search for safest bathroom to use as a starting point
	for i := 0; i < 100; i++ {

		quadrants := make(map[string]int)

		for _, robot := range robots {
			end_pos := robot.CalcEndPos(i, bathroom_r, bathroom_c)

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
		safety_scores = append(safety_scores, safety_factor)
	}

	smallest := math.MaxInt
	SEARCH_START := -1

	for i, score := range safety_scores {
		if score < smallest {
			smallest = score
			SEARCH_START = i
		}
	}

	// Modify robots so they start on this cycle
	for i, _ := range robots {
		robots[i].ModifyStartPos(SEARCH_START, bathroom_r, bathroom_c)
	}

	// Search multiples of bathroom column. (The starting image has vertical bands)
	for i := 0; i < 100; i++ {
		bathroom := make([][]int, bathroom_r)
		for br_c := range bathroom {
			bathroom[br_c] = make([]int, bathroom_c)
		}

		for _, robot := range robots {
			end_pos := robot.CalcEndPos(i*bathroom_c, bathroom_r, bathroom_c)
			bathroom[end_pos.R][end_pos.C] += 1
		}

		dump_bathroom(&bathroom, SEARCH_START + bathroom_c * i)
	}
}

func main() {
	file_name := "input.txt"
	part1(file_name)
	part2(file_name)
}
