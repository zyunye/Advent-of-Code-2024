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

		a_button := Button{X: a_x, Y: a_y}
		b_button := Button{X: b_x, Y: b_y}
		prize := Prize{X: p_x, Y: p_y}
		machine := Machine{A: a_button, B: b_button, Prize: prize}
		ret = append(ret, machine)

	}

	return ret
}

func extended_euclid(a, b int) (x, y, gcd int) {
	if b == 0 {
		return 1, 0, a
	}
	x1, y1, g := extended_euclid(b, a%b)
	x = y1
	y = x1 - y1*(a/b)
	gcd = g

	return x, y, gcd
}

func find_any_solution(a, b, c int) (x0, y0, gcd int) {
	x0, y0, gcd = extended_euclid(a, b)
	if c%gcd != 0 {
		return 0, 0, 0
	}

	x0 *= c / gcd
	y0 *= c / gcd

	if a < 0 {
		x0 = -x0
	}
	if b < 0 {
		y0 = -y0
	}

	return x0, y0, gcd
}

func shift_solution(x, y, a, b, cnt int) (x1, y1 int) {
	return x + cnt*b, y - cnt*a
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func find_all_solutions(a, b, c, minx, maxx, miny, maxy int) [][]int {
	x0, y0, gcd := find_any_solution(a, b, c)
	if gcd == 0 {
		return nil
	}

	a /= gcd
	b /= gcd

	sign_a := 1
	sign_b := 1

	x0, y0 = shift_solution(x0, y0, a, b, (minx-x0)/b)
	if x0 < minx {
		x0, y0 = shift_solution(x0, y0, a, b, sign_b)
	}
	if x0 > maxx {
		return nil
	}
	lx1 := x0

	x0, y0 = shift_solution(x0, y0, a, b, (maxx-x0)/b)
	if x0 > maxx {
		x0, y0 = shift_solution(x0, y0, a, b, -sign_b)
	}
	rx1 := x0

	x0, y0 = shift_solution(x0, y0, a, b, -(miny-y0)/a)
	if y0 < miny {
		x0, y0 = shift_solution(x0, y0, a, b, -sign_a)
	}
	if y0 > maxy {
		return nil
	}
	lx2 := x0

	x0, y0 = shift_solution(x0, y0, a, b, -(maxy-y0)/a)
	if y0 > maxy {
		x0, y0 = shift_solution(x0, y0, a, b, sign_a)
	}
	rx2 := x0

	if lx2 > rx2 {
		lx2, rx2 = rx2, lx2
	}
	lx := max(lx1, lx2)
	rx := min(rx1, rx2)

	ret := make([][]int, 0)

	k := 0
	for {
		x := lx + k*b/gcd
		y := (c - a*x) / b

		ret = append(ret, make([]int, 2))
		ret[k][0] = x
		ret[k][1] = y

		if x == rx || y < 0 {
			break
		}
		k++
	}

	return ret

}

func part1(file_name string) {
	machines := read_input(file_name)

	tokens_required := 0
	for _, m := range machines {
		x_solutions := find_all_solutions(m.A.X, m.B.X, m.Prize.X, 1, 100, 1, 100)
		y_solutions := find_all_solutions(m.A.Y, m.B.Y, m.Prize.Y, 1, 100, 1, 100)

		if len(x_solutions) == 0 || len(y_solutions) == 0 {
			continue
		}

		all_solutions := append(x_solutions, y_solutions...)
		seen_solutions := make(map[int]int)
		
		for _, solution_set := range all_solutions {
			count_A, count_B := solution_set[0], solution_set[1]
			if count_A > 100 || count_B > 100 {
				continue
			}
			if _, ok := seen_solutions[count_A]; ok {
				continue
			}

			test_X := m.A.X*count_A + m.B.X*count_B
			test_Y := m.A.Y*count_A + m.B.Y*count_B

			if test_X == m.Prize.X && test_Y == m.Prize.Y {
				tokens_required += count_A * 3 + count_B
				seen_solutions[count_A] = count_B
			}
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

// 28146 too low