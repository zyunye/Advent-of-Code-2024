package aoc

func Inbounds[T any](pos Position, arr *[][]T) bool {
	return !(pos.C < 0 || pos.R < 0 || pos.C >= len((*arr)[0]) || pos.R >= len(*arr))
}

func GetAdjPositions[T any](center Position, arr *[][]T) []Position {

	ret := make([]Position, 0)

	for _, adj_dir := range ADJACENTS {
		adj_pos := center.Add(adj_dir)
		if Inbounds(adj_pos, arr) {
			ret = append(ret, adj_pos)
		}
	}

	return ret
}

func GetAdjVals[T any](center Position, arr *[][]T) []T {
	adj_poses := GetAdjPositions(center, arr)
	ret := make([]T, 0)

	for _, pos := range adj_poses {
		v := (*arr)[pos.R][pos.C]
		ret = append(ret, v)
	}

	return ret
}

func GetOrthPositions[T any](center Position, arr *[][]T) []Position {

	ret := make([]Position, 0)

	for _, adj_dir := range TURN_ORDER {
		adj_pos := center.Add(adj_dir)
		if Inbounds(adj_pos, arr) {
			ret = append(ret, adj_pos)
		}
	}

	return ret
}

func GetOrthVals[T any](center Position, arr *[][]T) []T {
	adj_poses := GetOrthPositions(center, arr)
	ret := make([]T, 0)

	for _, pos := range adj_poses {
		v := (*arr)[pos.R][pos.C]
		ret = append(ret, v)
	}

	return ret
}

func Get[T any](pos Position, arr *[][]T) T {
	return (*arr)[pos.R][pos.C]
}

func (pt Position) CalcAdjPositions() []Position {
	ret := make([]Position, 0)

	for _, adj_dir := range ADJACENTS {
		adj_pos := pt.Add(adj_dir)
		ret = append(ret, adj_pos)
	}

	return ret
}

func (pt Position) CalcOrthPositions() []Position {
	ret := make([]Position, 0)

	for _, adj_dir := range TURN_ORDER {
		adj_pos := pt.Add(adj_dir)
		ret = append(ret, adj_pos)
	}

	return ret
}

func Pop[T any](arr *[]T) T {
	l := len(*arr)
	ret := (*arr)[l-1]
	*arr = (*arr)[:l-1]
	return ret
}

func Remove[T any](i int, arr *[]T) T {
	if i < 0 || i >= len((*arr)) {
		panic("Remove: Index out of bounds")
	}
	ret := (*arr)[i]
	*arr = append((*arr)[:i], (*arr)[i+1:]...)
	return ret
}

func RemoveNoRef[T any](i int, arr []T) ([]T, T) {
	if i < 0 || i >= len(arr) {
		panic("Remove: Index out of bounds")
	}
	ret := arr[i]
	arr = append(arr[:i], arr[i+1:]...)
	return arr, ret
}

func Insert[T any](i int, v T, arr *[]T) {
	if i < 0 {
		panic("Remove: Index out of bounds")
	} else if i == len((*arr)) {
		*arr = append((*arr), v)
	} else {
		*arr = append((*arr)[:i], append([]T{v}, (*arr)[i:]...)...)
	}
}

func Copy2DArr[T any](arr *[][]T) [][]T {
	maze_copy := make([][]T, len((*arr)))
	for i := range *arr {
		maze_copy[i] = append([]T(nil), (*arr)[i]...)
	}

	return maze_copy
}
