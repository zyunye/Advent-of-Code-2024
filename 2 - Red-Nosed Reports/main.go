package main

import (
	"bufio"
	"fmt"
	_ "fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parse_int(s string) int64 {
	n, err := strconv.ParseInt(s, 10, 64)
	check(err)
	return n
}

func copy_concat[T any](arr1 []T, arr2 []T) []T {
	l1 := len(arr1)
	l2 := len(arr2)

	new_arr := make([]T, l1+l2)
	_ = copy(new_arr, arr1)
	_ = copy(new_arr[l1:], arr2)

	return new_arr
}

func read_input(file_name string) ([]string, error) {
	file, err := os.Open(file_name)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func line_to_ints(line []string) []int64 {
	var ret []int64
	for _, num := range line {
		ret = append(ret, parse_int(num))
	}
	return ret
}

func check_adjacent(n1 int64, n2 int64, increasing bool) bool {
	diff := n2 - n1
	abs_diff := math.Abs(float64(diff))
	if (increasing && diff <= 0) || (!increasing && diff >= 0) {
		return false
	} else if (abs_diff < 1) || (abs_diff > 3) {
		return false
	}
	return true
}

func check_ints(nums []int64) (bool, int) {
	var increasing bool
	if nums[1] == nums[0] {
		return false, 0
	} else if nums[1] > nums[0] {
		increasing = true
	} else {
		increasing = false
	}

	for i := 1; i < len(nums); i++ {
		n1 := nums[i-1]
		n2 := nums[i]

		is_safe := check_adjacent(n1, n2, increasing)
		if !is_safe {
			return false, i
		}
	}

	return true, -1
}

func part1(in_file string) {
	lines, err := read_input(in_file)
	check(err)

	var wg sync.WaitGroup
	rets_ch := make(chan int, 10)

	go func() {
		for _, line := range lines {
			wg.Add(1)
			go func(string) {
				defer wg.Done()
				tokens := strings.Fields(line)
				nums := line_to_ints(tokens)

				if is_safe, _ := check_ints(nums); is_safe {
					rets_ch <- 1
				}

			}(line)
		}
		wg.Wait()
		close(rets_ch)
	}()

	safe_sum := 0
	for ret := range rets_ch {
		safe_sum += ret
	}

	fmt.Printf("P.1: %d\n", safe_sum)

}

func part2(in_file string) {
	lines, err := read_input(in_file)
	check(err)

	var wg sync.WaitGroup
	rets_ch := make(chan int, 10)

	go func() {
		for _, line := range lines {
			wg.Add(1)
			go func(string) {
				defer wg.Done()
				tokens := strings.Fields(line)
				nums := line_to_ints(tokens)
				
				for i := 0; i < len(nums); i++ {
					excised := copy_concat(nums[:i], nums[i+1:])
					if is_safe, _ := check_ints(excised); is_safe {
						rets_ch <- 1
						return
					}
				}

			}(line)
		}
		wg.Wait()
		close(rets_ch)
	}()

	safe_sum := 0
	for ret := range rets_ch {
		safe_sum += ret
	}

	fmt.Printf("P.2: %d\n", safe_sum)
}

func main() {
	in_file := "input.txt"
	part1(in_file)
	part2(in_file)
}
