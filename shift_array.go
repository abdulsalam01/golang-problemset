package main

import (
	"fmt"
)

func ShiftArray(array []int, i int) []int {
	// Write your code here.
	const dimen = 3
	var array2d = make([][]int, dimen)
	// convert to 2d array
	for i := 0; i < dimen; i++ {
		array2d[i] = make([]int, dimen)
		for j := 0; j < dimen; j++ {
			array2d[i][j] = array[i*dimen+j]
		}
	}

	// rotate matrix
	var _array [][]int
	for j := 0; j < i; j++ {
		_array = rotate(dimen, dimen, array2d)
	}

	var count int = 0
	// convert to 1d array
	for i := 0; i < dimen; i++ {
		for j := 0; j < dimen; j++ {
			array[count] = _array[i][j]
			count++
		}
	}

	return array
}

// helper
func rotate(dimenX int, dimenY int, arr [][]int) [][]int {
	row, col, prev, curr := 0, 0, 0, 0

	for row < dimenX && col < dimenY {

		if row+1 == dimenX || col+1 == dimenY {
			break
		}

		prev = arr[row+1][col]
		for i := col; i < dimenY; i++ {
			curr = arr[row][i]
			arr[row][i] = prev
			prev = curr
		}
		row++

		for i := row; i < dimenX; i++ {
			curr = arr[i][dimenX-1]
			arr[i][dimenX-1] = prev
			prev = curr
		}
		dimenY--

		if row < dimenX {
			for i := dimenY - 1; i >= col; i-- {
				curr = arr[dimenX-1][i]
				arr[dimenX-1][i] = prev
				prev = curr
			}
		}
		dimenX--

		if col < dimenY {
			for i := dimenX - 1; i >= row; i-- {
				curr = arr[i][col]
				arr[i][col] = prev
				prev = curr
			}
		}
		col++
	}

	return arr
}

func main() {
	fmt.Print(ShiftArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 1))
}
