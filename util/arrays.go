package util

func Create2DArrayOfInt(rows, cols int) [][]int {
	tmp := make([][]int, rows)
	for i := range tmp {
		tmp[i] = make([]int, cols)
	}

	return tmp
}
