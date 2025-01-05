package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
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

type Direction struct {
	X int
	Y int
}

var UP = Direction{0, -1}
var DOWN = Direction{0, 1}
var LEFT = Direction{-1, 0}
var RIGHT = Direction{1, 0}
var UPLEFT = Direction{-1, -1}
var UPRIGHT = Direction{1, -1}
var DOWNLEFT = Direction{-1, 1}
var DOWNRIGHT = Direction{1, 1}

func walk(x int, y int, dir Direction) (int, int) {
	return x + dir.X, y + dir.Y
}

func lines_to_arr(lines []string) [][]rune {
	var char_map [][]rune

	for r, line := range lines {
		char_map = append(char_map, make([]rune, len(line)))

		for c, char := range line {

			char_map[r][c] = char
		}
	}

	return char_map
}

func print_runes(char_map [][]rune) {
	for _, row := range char_map {
		for _, char := range row {
			fmt.Print(string(char))
		}
		fmt.Println()
	}
}

func check_inbound(x int, y int, char_map [][]rune) bool {
	if y < 0 || y >= len(char_map) {
		return false
	} else if x < 0 || x >= len(char_map[y]) {
		return false
	}
	return true
}

func walker_p1(x int, y int, dir Direction, char_map [][]rune, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	state := 'X'
	curX := x
	curY := y
	for i := 0; i < 4; i++ {
		if !check_inbound(curX, curY, char_map) {
			return
		}

		if i == 0 && char_map[curY][curX] == state {
			curX, curY = walk(curX, curY, dir)
			state = 'M'
		} else if i == 1 && char_map[curY][curX] == state {
			curX, curY = walk(curX, curY, dir)
			state = 'A'
		} else if i == 2 && char_map[curY][curX] == state {
			curX, curY = walk(curX, curY, dir)
			state = 'S'
		} else if i == 3 && char_map[curY][curX] == state {
			results <- 1
			return
		} else {
			return
		}
	}
}

func part1(file_name string) {
	lines, err := read_input(file_name)
	check(err)
	char_map := lines_to_arr(lines)

	xmases_found := 0
	var wg sync.WaitGroup
	results := make(chan int)

	for r, _ := range char_map {
		for c, _ := range char_map[r] {
			wg.Add(8)
			go walker_p1(c, r, UPLEFT, char_map, results, &wg)
			go walker_p1(c, r, UP, char_map, results, &wg)
			go walker_p1(c, r, UPRIGHT, char_map, results, &wg)
			go walker_p1(c, r, RIGHT, char_map, results, &wg)
			go walker_p1(c, r, DOWNRIGHT, char_map, results, &wg)
			go walker_p1(c, r, DOWN, char_map, results, &wg)
			go walker_p1(c, r, DOWNLEFT, char_map, results, &wg)
			go walker_p1(c, r, LEFT, char_map, results, &wg)
		}
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		xmases_found += result
	}

	fmt.Printf("P.1: %d\n", xmases_found)

}

func walker_p2(x int, y int, char_map [][]rune, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 4; i++ {
		if char_map[y][x] == 'A' {
			ulx, uly := walk(x, y, UPLEFT)
			if !check_inbound(ulx, uly, char_map) {
				return
			}
			urx, ury := walk(x, y, UPRIGHT)
			if !check_inbound(urx, ury, char_map) {
				return
			}
			dlx, dly := walk(x, y, DOWNLEFT)
			if !check_inbound(dlx, dly, char_map) {
				return
			}
			drx, dry := walk(x, y, DOWNRIGHT)
			if !check_inbound(drx, dry, char_map) {
				return
			}

			if (char_map[uly][ulx] == 'M' && char_map[dry][drx] == 'S') || (char_map[uly][ulx] == 'S' && char_map[dry][drx] == 'M') {
				if (char_map[ury][urx] == 'M' && char_map[dly][dlx] == 'S') || (char_map[ury][urx] == 'S' && char_map[dly][dlx] == 'M') {
					results <- 1
					return
				}
			}
		}
	}
}

func part2(file_name string) {
	lines, err := read_input(file_name)
	check(err)
	char_map := lines_to_arr(lines)

	x_mases_found := 0
	var wg sync.WaitGroup
	results := make(chan int)

	for r, _ := range char_map {
		for c, _ := range char_map[r] {
			wg.Add(1)
			go walker_p2(c, r, char_map, results, &wg)
		}
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		x_mases_found += result
	}

	fmt.Printf("P.2: %d\n", x_mases_found)
}

func main() {
	file_name := "input.txt"
	part1(file_name)
	part2(file_name)
}
