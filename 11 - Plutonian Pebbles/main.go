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

func split_stone(i int, num_len int, pebbles *[]int) []int {
	val := (*pebbles)[i]
	div := int(math.Pow10(num_len / 2))
	left := val / div
	right := val - (left * div)

	(*pebbles)[i] = left
	Insert(i+1, right, pebbles)
	return []int{left, right}
}

func part_num_cache(file_name string) {
	pebbles := read_input(file_name)

	num_cache := make(map[int][]int)

	for blink := 0; blink < 25; blink++ {
		fmt.Printf("Blink: %d, %d\n", blink, len(pebbles))
		for i := 0; i < len(pebbles); i++ {

			num := pebbles[i]
			next, ok := num_cache[num]
			if !ok {
				num_len := get_num_len(num)
				if num == 0 {
					pebbles[i] = 1
				} else if num_len%2 == 0 {
					split_nums := split_stone(i, num_len, &pebbles)
					num_cache[num] = split_nums
					i++
				} else {
					pebbles[i] *= 2024
					num_cache[num] = []int{pebbles[i]}
				}
			} else {
				pebbles = append(pebbles[:i], append(next, pebbles[i+1:]...)...)
				if len(next) > 1 {
					i++
				}
			}
		}
	}

	fmt.Printf("P.1_num: %d\n", len(pebbles))

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
			num_len, ok := len_cache[num]
			if !ok {
				num_len = get_num_len(num)
				len_cache[num] = num_len
			}

			if num == 0 {
				pebbles[i] = 1
			} else if num_len%2 == 0 {
				split_stone(i, num_len, &pebbles)
				i++
			} else {
				pebbles[i] *= 2024
			}
		}
	}

	fmt.Printf("P.1_len: %d\n", len(pebbles))

}

func part2(file_name string) {
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

	part2(file_name)
}
