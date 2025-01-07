package main

import (
	. "aoc"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Button struct {
	X int
	Y int
}
type Prize struct {
	X int
	Y int
}
type Machine struct {
	A     Button
	B     Button
	Prize Prize
}

func (m Machine) String() string {
	return fmt.Sprintf(`
A: %v
B: %v
Prize: %v
`, m.A, m.B, m.Prize)
}

func read_input(file_name string) []Machine {
	file, err := os.Open(file_name)
	CheckErr(err)

	scanner := bufio.NewScanner(file)

	ret := make([]Machine, 0)

	x_pattern := `X\+\d+`
	y_pattern := `Y\+\d+`
	xeq_pattern := `X=\d+`
	yeq_pattern := `Y=\d+`
	re_x, err := regexp.Compile(x_pattern)
	CheckErr(err)
	re_y, err := regexp.Compile(y_pattern)
	CheckErr(err)
	re_xeq, err := regexp.Compile(xeq_pattern)
	CheckErr(err)
	re_yeq, err := regexp.Compile(yeq_pattern)
	CheckErr(err)

	for scanner.Scan() {
		a := scanner.Text()
		scanner.Scan()
		b := scanner.Text()
		scanner.Scan()
		p := scanner.Text()
		scanner.Scan()
		scanner.Text()

		a_x, err := strconv.Atoi(re_x.FindString(a)[2:])
		CheckErr(err)
		a_y, err := strconv.Atoi(re_y.FindString(a)[2:])
		CheckErr(err)

		b_x, err := strconv.Atoi(re_x.FindString(b)[2:])
		CheckErr(err)
		b_y, err := strconv.Atoi(re_y.FindString(b)[2:])
		CheckErr(err)

		p_x, err := strconv.Atoi(re_xeq.FindString(p)[2:])
		CheckErr(err)
		p_y, err := strconv.Atoi(re_yeq.FindString(p)[2:])
		CheckErr(err)

		a_button := Button{X: a_x, Y: a_y}
		b_button := Button{X: b_x, Y: b_y}
		prize := Prize{X: p_x, Y: p_y}
		machine := Machine{A: a_button, B: b_button, Prize: prize}
		ret = append(ret, machine)

	}

	return ret
}

func part1(file_name string) {
	machines := read_input(file_name)
	fmt.Println(machines)
}
func part2(file_name string) {
}
func main() {
	file_name := "test.txt"
	part1(file_name)
	part2(file_name)
}
