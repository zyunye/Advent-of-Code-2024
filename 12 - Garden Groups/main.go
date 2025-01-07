package main

import (
	. "aoc"
	"bufio"
	"errors"
	"fmt"
	"os"
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
	tmp_pts := make([]Position, len(p.points))
	copy(tmp_pts, p.points)
	tmp_perim := make([]Position, len(p.perimeter))
	copy(tmp_perim, p.perimeter)

	for i, pt := range tmp_pts {
		tmp_pts[i] = Position{R: pt.R + 1, C: pt.C + 1}
	}
	for i, pt := range tmp_perim {
		tmp_perim[i] = Position{R: pt.R + 1, C: pt.C + 1}
	}

	return fmt.Sprintf(`
	id: %s
		points: %v
		perimeter: %v
		core_area: %d
		perimeter_area: %d
`, p.id, tmp_pts, tmp_perim, p.core_area, p.perimeter_area)
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

		// Get the orthogonal points and put it on the stack
		adj_pos := GetOrthPositions(cur_pos, &farm.f)
		stack = append(stack, adj_pos...)
	}

	return new_plot
}

func calc_perimeter(plot *Plot, farm *Farm) int {
	id := plot.id
	perim_points := plot.perimeter

	perimeter := 0

	for _, center_pt := range perim_points {
		orth_vals := GetOrthVals(center_pt, &farm.f)
		perimeter += 4
		for _, v := range orth_vals {
			if v == id {
				perimeter -= 1
			}
		}
	}

	return perimeter
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
				plots[id] = append(plots[id], new_plot)

			}
		}
	}

	// for _, plot_list := range plots {
	// 	for _, plot := range plot_list {
	// 		fmt.Println(plot)
	// 	}
	// }

	prices := make(map[string]int)
	for _, plot_list := range plots {
		for _, plot := range plot_list {
			id := plot.id
			core_area := plot.core_area
			perimeter_area := plot.perimeter_area
			perimeter := calc_perimeter(&plot, &farm)

			fmt.Println(plot)
			fmt.Println(perimeter)

			prices[id] += (core_area + perimeter_area) * perimeter
		}
	}

	total_price := 0
	for _, v := range prices {
		total_price += v
	}

	fmt.Printf("P.1: %d\n", total_price)

}

func part2(file_name string) {
}

func main() {
	file_name := "input.txt"
	part1(file_name)
	part2(file_name)
}