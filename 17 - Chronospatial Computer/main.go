package main

import (
	. "aoc"
	"bufio"
	"fmt"
	"os"
	"sort"
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
			computer.a = numerator >> combo

		} else if opcode == 1 { // bxl
			computer.b = computer.b ^ operand

		} else if opcode == 2 { // bst
			combo := combo_operand(computer, operand)
			computer.b = combo & 7

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
			computer.b = numerator >> combo

		} else if opcode == 7 { // cdv
			combo := combo_operand(computer, operand)
			numerator := computer.a
			computer.c = numerator >> combo
		}
		ptr += 2
	}

	return result
}

func part1(file_name string) {
	computer := read_input(file_name)
	fmt.Println(computer)
	ret := fsm(&computer)
	fmt.Printf("P.1: %v", ret)
}

func decompiled(a int) (a_next int, res int) {
	b := a & 7 // 2,4

	b = b ^ 3 // 1,3

	c := a >> b // 7,5

	a = a >> 3 // 0,3

	b = b ^ 4 // 1,4

	b = b ^ c // 4,_

	res = b & 7 // 5,5

	return a, res // 3, 0

}

func part2(file_name string) {
	computer := read_input(file_name)

	type ss struct {
		i int
		a int
	}

	search_stack := []ss{
		{i: len(computer.program), a: 0},
	}
	valid_a_vals := make([]int, 0)

	for len(search_stack) > 0 {
		st := Pop(&search_stack)
		cur_ind, cur_a := st.i, st.a

		if cur_ind == 0 {
			valid_a_vals = append(valid_a_vals, cur_a)
			continue
		}

		next_ind := cur_ind - 1
		actual_instr := computer.program[next_ind]

		for i := 0; i < 8; i++ {
			next_a := (cur_a << 3) | i
			a_ret, instr_ret := decompiled(next_a)

			if a_ret == cur_a && instr_ret == actual_instr {
				search_stack = append(search_stack, ss{i: next_ind, a: next_a})
			}
		}
	}

	sort.Ints(valid_a_vals)

	for _, v := range valid_a_vals {
		fmt.Println(v, fsm(&Computer{a: v, b: 0, c: 0, program: computer.program}))
	}
}

func main() {
	file_name := "input.txt"
	part1(file_name)
	part2(file_name)
}

// input program len == 16

// 64 to end of 1st 0
// 512 to end of 2nd 0
// 4096 to end of 3rd zero
// 0 enders cycle every 3rd power of 2

// This also means program length increases by 1 every cycle
// for a program length of 16, we need 16 cycles to pass
// this means 2^48 for our range?

//i need to search from [2^45 to 2^48)
// final number changes every 2^42 iterations

// every index of the program changes ever 2^(i*3) where i is its index iterations
// for example, the 1st index will swap every iteration i=0
// the 2nd index changes every 2^3 = 8 iterations i=1
// the 3rd index changes every 2^6 = 64 iterations i=2

// 0 = start:0	 		cycles:1024=2^10 	repeats:2^0
// 1 = start:2^(3*1)	cycles:1024=2^10	repeats:2^(3*1)
// 2 = start:2^(3^2)	cycles:1024=2^10	repeats:2^(3^2)
