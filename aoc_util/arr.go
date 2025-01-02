package aoc

func Inbounds[T any](pos Position, arr *[][]T) bool {
	rows := len(*arr)
	for r := 0; r < rows; r++ {
		len_c := len((*arr)[r])

		if len_c != rows {
			panic("Inbounds: Array not square")
		}
	}

	return !(pos.C < 0 || pos.R < 0 || pos.C >= len((*arr)[0]) || pos.R >= len(*arr))
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

func Insert[T any](i int, v T, arr *[]T) {
	if i < 0 {
		panic("Remove: Index out of bounds")
	} else if i == len((*arr)) {
		*arr = append((*arr), v)
	} else {
		*arr = append((*arr)[:i], append([]T{v}, (*arr)[i:]...)...)
	}
}
