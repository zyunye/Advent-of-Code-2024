package main

import (
	. "aoc"
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"sort"
)

type Farm struct {
	f [][]string
}

func (farm *Farm) Len() int {
	return len(farm.f)
}

func (farm *Farm) Get(pos Position) (string, error) {
	if !Inbounds(pos, &farm.f) {
		return "", errors.New("Farm Get OOB")
	}

	return farm.f[pos.R][pos.C], nil
}

type Plot struct {
	id        string
	points    []Position
	perimeter []Position

	core_area      int
	perimeter_area int
}

func (p Plot) String() string {
	return fmt.Sprintf(`
	id: %s
		points: %v
		perimeter: %v
		core_area: %d
		perimeter_area: %d
`, p.id, p.points, p.perimeter, p.core_area, p.perimeter_area)
}

func angle(pivot, p Position) float64 {
	return math.Atan2(float64(p.R)-float64(pivot.R), float64(p.C)-float64(pivot.C))
}

func (plot *Plot) sort_plot_perimeter() {
	// CCW sorting
	pts := plot.perimeter
	pivot := pts[0]

	for _, p := range pts {
		if p.R < pivot.R || (p.R == pivot.R && p.C < pivot.C) {
			pivot = p
		}
	}

	sort.SliceStable(pts, func(i, j int) bool {
		angle_i := angle(pivot, pts[i])
		angle_j := angle(pivot, pts[j])
		return angle_i > angle_j
	})
}

func read_input(file_name string) [][]string {
	file, err := os.Open(file_name)
	CheckErr(err)

	scanner := bufio.NewScanner(file)

	ret := make([][]string, 0)

	r := 0
	for scanner.Scan() {
		ret = append(ret, make([]string, 0))
		for _, v := range scanner.Text() {
			v := string(v)
			ret[r] = append(ret[r], v)
		}
		r++
	}

	return ret
}

func is_perimeter_coord(coord Position, farm *Farm) (bool, error) {
	v, err := farm.Get(coord)
	if err != nil {
		return false, err
	}

	for _, dir := range TURN_ORDER {
		adj_v, err := farm.Get(coord.Add(dir))

		if err != nil || adj_v != v {
			return true, nil
		}
	}

	return false, nil
}

func flood_fill(start Position, farm *Farm, seen *map[Position]bool) Plot {
	plot_id, err := farm.Get(start)
	if err != nil {
		panic("Start position for flood fill out of bounds")
	}

	stack := make([]Position, 0)
	stack = append(stack, start)

	new_plot := Plot{points: make([]Position, 0), perimeter: make([]Position, 0)}
	new_plot.id = plot_id

	for len(stack) > 0 {
		cur_pos := Pop(&stack)
		cur_val, err := farm.Get(cur_pos)

		// If current position is out of bounds, ignore
		if err != nil {
			continue
		}
		// If current position is not the same plot id as what we started with, ignore
		if cur_val != plot_id {
			continue
		}
		// Or, if we have seen it before during any flood fill process, ignore
		if _, ok := (*seen)[cur_pos]; ok {
			continue
		}

		// Check if current point is a perimeter point and treat respectively
		if is_perim, err := is_perimeter_coord(cur_pos, farm); err != nil || is_perim {
			new_plot.perimeter = append(new_plot.perimeter, cur_pos)
			new_plot.perimeter_area++
		} else {
			new_plot.points = append(new_plot.points, cur_pos)
			new_plot.core_area++
		}
		(*seen)[cur_pos] = true

		// Get the adjacent points and put it on the stack
		adj_pos := GetAdjPositions(cur_pos, &farm.f)
		stack = append(stack, adj_pos...)
	}

	return new_plot
}

func part1(file_name string) {
	f := read_input(file_name)

	farm := Farm{f: f}
	seen := make(map[Position]bool)
	plots := make(map[string][]Plot)

	for r, row := range farm.f {
		for c, id := range row {

			cur_pos := Position{R: r, C: c}

			if seen[cur_pos] {
				continue
			} else {
				new_plot := flood_fill(cur_pos, &farm, &seen)
				new_plot.sort_plot_perimeter()
				plots[id] = append(plots[id], new_plot)

			}
		}
	}

	for _, plot_list := range plots {
		for _, plot := range plot_list {
			fmt.Println(plot)
		}
	}

}

func part2(file_name string) {
}

func main() {
	file_name := "test.txt"
	part1(file_name)
	part2(file_name)
}
