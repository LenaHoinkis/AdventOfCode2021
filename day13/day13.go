package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/lenahoinkis/AdventOfCode2021/utils"
)

type Vertex struct {
	X, Y int
}
type Fold struct {
	Axis  rune
	Point int
}

func main() {
	inputFile := flag.String("inputFile", "ex.input", "Relative file path to use as input.")
	flag.Parse()
	marks, folds, row, col, _ := PaperInstructions(*inputFile)

	row++
	col++
	board := make([]int, row*col)
	for i := range board {
		board[i] = 0
	}
	//board[i*row+j] = "abc" // like board[i][j] = "abc"
	for _, v := range marks {
		board[v.X+v.Y*row] = 1
	}

	for i, f := range folds {
		board, row, col = fold(row, col, board, f)
		//part 1 first fold
		if i == 0 {
			fmt.Println(sumBoard(0, board))
		}
	}

	readable := make([]string, row*col)
	for i, v := range board {
		if v != 0 {
			readable[i] = "X"
		} else {
			readable[i] = " "
		}
	}
	utils.PrintStringBoard(row, readable)
}

func sumBoard(compare int, board []int) int {
	sum := 0
	for _, v := range board {
		if v != compare {
			sum++
		}
	}
	return sum
}

func fold(row, col int, board []int, fold Fold) ([]int, int, int) {
	if fold.Axis == 'y' {
		return foldY(row, col, board, fold.Point)
	} else {
		return foldX(row, col, board, fold.Point)
	}
}

func foldY(row, col int, board []int, foldingPoint int) ([]int, int, int) {
	for x := 0; x < row; x++ {
		for y := 0; y < col-foldingPoint; y++ {
			board[x+y*row] += board[x+(col-y-1)*row]
		}
	}
	return board[:foldingPoint*row], row, col / 2
}

func foldX(row, col int, board []int, foldingPoint int) ([]int, int, int) {
	for x := 0; x < row-foldingPoint-1; x++ {
		for y := 0; y < col; y++ {
			board[x+y*row] += board[(row-x-1)+y*row]
		}
	}
	for x := 1; x < col+1; x++ {
		endOfLine := foldingPoint * x
		startOfNewLine := endOfLine + foldingPoint + 1
		board = append(board[:endOfLine], board[startOfNewLine:]...)
	}
	return board, row / 2, col
}

func PaperInstructions(path string) ([]Vertex, []Fold, int, int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, 0, 0, err
	}
	defer file.Close()

	var marks []Vertex
	var folds []Fold
	var w, h int
	isPaper := true
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			isPaper = false
			continue
		}
		if isPaper {
			nums := strings.Split(line, ",")
			x, _ := strconv.Atoi(nums[0])
			y, _ := strconv.Atoi(nums[1])
			if x > w {
				w = x
			}
			if y > h {
				h = y
			}
			point := Vertex{x, y}
			marks = append(marks, point)
		} else {
			var axis rune
			var point int
			fmt.Sscanf(line, "fold along %c=%d", &axis, &point)
			folds = append(folds, Fold{axis, point})
		}
	}
	return marks, folds, w, h, scanner.Err()
}
