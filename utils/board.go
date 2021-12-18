package utils

import "fmt"

func PrintIntBoard(row int, board []int) {
	for i, v := range board {
		fmt.Print(v)
		if (i+1)%row == 0 {
			fmt.Println()
		}
	}
}
func PrintStringBoard(row int, board []string) {
	for i, v := range board {
		fmt.Print(v)
		if (i+1)%row == 0 {
			fmt.Println()
		}
	}
}
