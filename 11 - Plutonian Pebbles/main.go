package main

import (
	. "aoc"
	"bufio"
	"fmt"
	"math"
	"os"
	"runtime/pprof"
	"strconv"
	"strings"
	"time"
)

func read_input(file_name string) []int {
	file, err := os.Open(file_name)
	CheckErr(err)

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	str_nums := strings.Fields(line)

	ret := make([]int, len(str_nums))
	for i, v := range str_nums {
		ret[i], err = strconv.Atoi(v)
		CheckErr(err)
	}

	return ret
}

func get_num_len(num int) int {
	if num == 0 {
		return 1
	}

	digits := 0
	for num > 0 {
		num /= 10
		digits++
	}

	return digits
}

func split_stone(num int, num_len int) (int, int) {
	div := int(math.Pow10(num_len / 2))
	left := num / div
	right := num - (left * div)

	return left, right
}

func process_stone(num int, len_cache *map[int]int) (int, int) {
	num_len, ok := (*len_cache)[num]
	if !ok {
		num_len = get_num_len(num)
		(*len_cache)[num] = num_len
	}

	if num == 0 {
		return 1, -1
	} else if num_len%2 == 0 {
		return split_stone(num, num_len)
	} else {
		return num * 2024, -1
	}
}

func part_len_cache(file_name string) {
	pebbles := read_input(file_name)
	// pebbles := make([]int, 0, 1e10)
	// pebbles = append(pebbles, pebbles_orig...)

	len_cache := make(map[int]int)

	for blink := 0; blink < 25; blink++ {
		fmt.Printf("Blink: %d, %d\n", blink, len(pebbles))
		for i := 0; i < len(pebbles); i++ {
			num := pebbles[i]
			l, r := process_stone(num, &len_cache)

			pebbles[i] = l
			if r != -1 {
				Insert(i+1, r, &pebbles)
				i++
			}
		}
	}

	fmt.Printf("P.1_len: %d\n", len(pebbles))

}

func main() {
	file_name := "test.txt"
	var start time.Time

	f, err := os.Create("perf.prof")
	CheckErr(err)
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	// start := time.Now()
	// part1(file_name)
	// fmt.Printf("P1 time: %s", time.Since(start))

	start = time.Now()
	part_len_cache(file_name)
	fmt.Printf("P1 time: %s", time.Since(start))

}
