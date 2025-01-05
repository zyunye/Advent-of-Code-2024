package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
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

func parse_and_split(lines []string) ([]int64, []int64) {
	var arr1 []int64
	var arr2 []int64
	for _, line := range lines {
		nums := strings.Fields(line)
		n1, err := strconv.ParseInt(nums[0], 10, 64)
		check(err)
		n2, err := strconv.ParseInt(nums[1], 10, 64)
		check(err)

		arr1 = append(arr1, n1)
		arr2 = append(arr2, n2)
		
	}

	return arr1, arr2
}

func part1() {
	lines, err := read_input("input.txt")
	check(err)
	arr1, arr2 := parse_and_split(lines)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {defer wg.Done(); slices.Sort(arr1)}()
	go func() {defer wg.Done(); slices.Sort(arr2)}()
	wg.Wait()

	var sum int64
	for i := range(arr1) {
		n1 := arr1[i]
		n2 := arr2[i]
		diff := int64(math.Abs(float64(n1 - n2)))
		sum += diff
	}

	fmt.Printf("P.1: %d\n", sum)
}

func part2() {
	lines, err := read_input("input.txt")
	check(err)
	arr1, arr2 := parse_and_split(lines)

	num_map := make(map[int64]int64)
	for _, n2 := range(arr2) {
		num_map[n2] += 1
	}

	var sum int64
	for _, n1 := range(arr1) {
		sum += n1 * num_map[n1]
	}

	fmt.Printf("P.2: %d\n", sum)
}

func main() {
	part1()
	part2()
}
