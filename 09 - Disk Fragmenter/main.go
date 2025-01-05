package main

import (
	. "aoc"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func read_input(file_name string) []int {
	file, err := os.Open(file_name)
	CheckErr(err)

	disk := make([]int, 0)

	scanner := bufio.NewScanner(file)
	cur_id := 0
	for scanner.Scan() {
		for i, v := range scanner.Text() {
			num_repeats, err := strconv.Atoi(string(v))
			CheckErr(err)

			if i%2 == 0 {
				for j := 0; j < num_repeats; j++ {
					disk = append(disk, cur_id)
				}
				cur_id += 1
			} else {
				for j := 0; j < num_repeats; j++ {
					disk = append(disk, -1)
				}
			}

		}
	}

	CheckErr(scanner.Err())
	return disk
}

func swap(disk *[]int, l int, r int) {
	(*disk)[l], (*disk)[r] = (*disk)[r], (*disk)[l]
}

func part1(file_name string) {
	disk := read_input(file_name)
	l, r := 0, len(disk)-1

	for l < r {
		l_val := disk[l]
		r_val := disk[r]

		if l_val != -1 {
			l += 1
		} else if r_val == -1 {
			r -= 1
		} else {
			swap(&disk, l, r)
			r -= 1
			l += 1
		}
	}

	checksum := 0

	for i := 0; i < len(disk); i++ {
		if disk[i] != -1 {
			checksum += disk[i] * i
		}
	}

	fmt.Printf("P.1: %d\n", checksum)
}

func get_segment_length(disk []int, i int, forward bool) (int, int, int) {
	val := disk[i]

	var start int
	var end int

	if !forward {
		start = -1
		end = i
		for i >= 0 && val == disk[i] {
			start = i
			i--
		}
	} else {
		start = i
		end = -1
		for i < len(disk) && val == disk[i] {
			end = i
			i++
		}
	}

	return start, end, end - start + 1
}

func part2(file_name string) {
	disk := read_input(file_name)
	r := len(disk) - 1

	for r >= 0 {
		disk_val := disk[r]

		if disk_val != -1 {
			part_start, _, part_len := get_segment_length(disk, r, false)

			for l := 0; l < r-part_len-1; l++ {
				search_val := disk[l]
				if search_val == -1 {
					space_start, _, space_len := get_segment_length(disk, l, true)

					if space_len >= part_len {
						for i := 0; i < part_len; i++ {
							swap(&disk, space_start+i, part_start+i)
						}
						break
					}
				}
			}
			r -= part_len
		} else {
			r--
		}
	}

	checksum := 0
	for i := 0; i < len(disk); i++ {
		if disk[i] != -1 {
			checksum += disk[i] * i
		}
	}

	fmt.Printf("P.2: %d\n", checksum)
}

func main() {
	file_name := "input.txt"
	part1(file_name)
	part2(file_name)
}
