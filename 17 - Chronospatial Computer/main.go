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

type Computer struct {
	a       int
	b       int
	c       int
	program []int
}

func (c Computer) String() string {
	return fmt.Sprintf(`Register A: %d
Register B: %d
Register C: %d
Program: %v`, c.a, c.b, c.c, c.program)
}

func read_input(file_name string) Computer {
	file, err := os.Open(file_name)
	CheckErr(err)

	scanner := bufio.NewScanner(file)

	computer := Computer{program: make([]int, 0)}

	scanner.Scan()
	fields := strings.Fields(scanner.Text())
	val, err := strconv.Atoi(fields[2])
	CheckErr(err)
	computer.a = val

	scanner.Scan()
	fields = strings.Fields(scanner.Text())
	val, err = strconv.Atoi(fields[2])
	CheckErr(err)
	computer.b = val

	scanner.Scan()
	fields = strings.Fields(scanner.Text())
	val, err = strconv.Atoi(fields[2])
	CheckErr(err)
	computer.c = val

	scanner.Scan()
	scanner.Scan()
	fields = strings.Fields(scanner.Text())
	fields = strings.Split(fields[1], ",")
	CheckErr(err)

	for _, v := range fields {
		val, err := strconv.Atoi(v)
		CheckErr(err)
		computer.program = append(computer.program, val)
	}

	return computer
}

func combo_operand(computer *Computer, operand int) int {
	if operand <= 3 {
		return operand
	} else if operand == 4 {
		return computer.a
	} else if operand == 5 {
		return computer.b
	} else if operand == 6 {
		return computer.c
	} else {
		panic("Operation received invalid operand")
	}
}

func fsm(computer *Computer) []int {
	ptr := 0

	result := make([]int, 0)

	for ptr < len(computer.program) {
		opcode := computer.program[ptr]
		operand := computer.program[ptr+1]

		if opcode == 0 { // adv
			combo := combo_operand(computer, operand)
			numerator := computer.a
			denominator := math.Pow(2, float64(combo))
			computer.a = int(numerator / int(denominator))

		} else if opcode == 1 { // bxl
			computer.b = computer.b ^ operand

		} else if opcode == 2 { // bst
			combo := combo_operand(computer, operand)
			computer.b = combo % 8

		} else if opcode == 3 { // jnz
			if computer.a != 0 {
				ptr = operand
				continue
			}

		} else if opcode == 4 { // bxc
			computer.b = computer.b ^ computer.c

		} else if opcode == 5 { // out
			combo := combo_operand(computer, operand)
			result = append(result, combo%8)

		} else if opcode == 6 { // bdv
			combo := combo_operand(computer, operand)
			numerator := computer.a
			denominator := math.Pow(2, float64(combo))
			computer.b = int(numerator / int(denominator))

		} else if opcode == 7 { // cdv
			combo := combo_operand(computer, operand)
			numerator := computer.a
			denominator := math.Pow(2, float64(combo))
			computer.c = int(numerator / int(denominator))
		}
		ptr += 2
	}

	return result
}

func part1(file_name string) {
	computer := read_input(file_name)
	fmt.Println(computer)
	res := fsm(&computer)
	computer_out(res)

}

func computer_out(results []int) {
	for _, v := range results {
		fmt.Printf("%d,", v)
	}
	fmt.Println()
}

func main() {
	file_name := "test.txt"
	part1(file_name)
}
