package main

import (
	. "aoc"
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type Button struct {
	X float64
	Y float64
}
type Prize struct {
	X float64
	Y float64
}
type Machine struct {
	A     Button
	B     Button
	Prize Prize
}

func (m Machine) String() string {
	return fmt.Sprintf(`A: %v
B: %v
Prize: %v`, m.A, m.B, m.Prize)
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

		a_button := Button{X: float64(a_x), Y: float64(a_y)}
		b_button := Button{X: float64(b_x), Y: float64(b_y)}
		prize := Prize{X: float64(p_x), Y: float64(p_y)}
		machine := Machine{A: a_button, B: b_button, Prize: prize}
		ret = append(ret, machine)

	}

	return ret
}

func solve_equation(m Machine) (a_count, b_count float64) {
	a_x, a_y := m.A.X, m.A.Y
	b_x, b_y := m.B.X, m.B.Y
	p_x, p_y := m.Prize.X, m.Prize.Y

	ax_by := a_x * b_y
	px := p_x * b_y

	ay_bx := a_y * b_x
	py := p_y * b_x

	a_count = (px - py) / (ax_by - ay_bx)
	b_count = (p_y - a_y*a_count) / b_y

	return a_count, b_count
}

func part1(file_name string) {
	machines := read_input(file_name)

	tokens_required := 0
	for _, m := range machines {
		a_count, b_count := solve_equation(m)
		if a_count == math.Trunc(a_count) && b_count == math.Trunc(b_count) {
			tokens_required += 3*int(a_count) + int(b_count)
		}
	}

	fmt.Printf("P.1: %d", tokens_required)
}

func part2(file_name string) {
}
func main() {
	file_name := "input.txt"
	part1(file_name)
	part2(file_name)
}
