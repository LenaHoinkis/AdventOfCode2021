package utils

import "fmt"

func MakeMatrix(height, width int, def int) [][]int {
	m := make([][]int, width)
	for x := 0; x < width; x++ {
		m[x] = make([]int, height)
		for y := 0; y < height; y++ {
			m[x][y] = def
		}
	}
	return m
}

//PrintMatrix prints a 2d array of int
func PrintMatrix(is [][]int) {
	for x, i := range is {
		for y := range i {
			fmt.Print(is[x][y])
		}
		fmt.Println()
	}
}
