package Helper

import "sort"

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for index, value := range a {
		if value != b[index] {
			return false
		}
	}
	return true
}

func GetInnerSquareAsLine(a [][]int, row int, column int) []int {
	innerSquare := make([]int, 9)
	for i := row*3; i < row*3+3; i++ {
		for j := column*3; j < column*3+3; j++ {
			innerSquare[(i%3)*3+j%3] = a[i][j]
		}
	}
	sort.Ints(innerSquare)
	return innerSquare
}