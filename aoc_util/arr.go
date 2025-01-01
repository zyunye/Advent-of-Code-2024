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
