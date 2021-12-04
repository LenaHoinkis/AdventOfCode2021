package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "ex.input", "Relative file path to use as input.")

func main() {
	num, boards, err := ReadBingo(*inputFile)
	if err != nil {
		fmt.Print(err)
	}

	//copy
	var boardsCopy1, boardsCopy2 [][5][5]int
	for i := range boards {
		boardsCopy1 = append(boardsCopy1, boards[i])
		boardsCopy2 = append(boardsCopy2, boards[i])
	}

	//part1
	fmt.Println(calcWinningBoard(num, boardsCopy1))
	//part2
	fmt.Println(calcLeastWinningBoard(num, boardsCopy2))

}

//part1
func calcWinningBoard(num []int, boards [][5][5]int) int {
	for i := 0; i < len(num); i++ {
		for b, board := range boards {
			for x := 0; x < 5; x++ {
				for y := 0; y < 5; y++ {
					if board[x][y] == num[i] {
						boards[b][x][y] = -1
						board = boards[b]
						if isHit(boards[b], x, y) {
							return sumRemaining(board) * num[i]
						}
					}
				}
			}
		}
	}
	return 0
}

func isHit(board [5][5]int, x int, y int) bool {
	countx, county := 0, 0
	for z := 0; z < 5; z++ {
		if board[x][z] == -1 {
			countx++
		}
		if board[z][y] == -1 {
			county++
		}
		if countx == 5 || county == 5 {
			return true
		}
	}
	return false
}

func sumRemaining(b [5][5]int) int {
	sum := 0
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			if b[x][y] != -1 {
				sum += b[x][y]
			}
		}
	}
	return sum
}

//part2
func calcLeastWinningBoard(num []int, boards [][5][5]int) int {
	winner := 0
	for i := 0; i < len(num); i++ {
		for b, board := range boards {
			if boards[b][0][0] == -2 {
				continue
			}
			for x := 0; x < 5; x++ {
				for y := 0; y < 5; y++ {
					if boards[b][0][0] == -2 {
						break
					}
					if board[x][y] == num[i] {
						boards[b][x][y] = -1
						board = boards[b]
						if isHit(boards[b], x, y) {
							winner = sumRemaining(board) * num[i]
							boards[b][0][0] = -2
							board = boards[b]
						}
					}
				}
			}
		}
	}
	return winner
}

func ReadBingo(path string) ([]int, [][5][5]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var numbers []int
	scanner := bufio.NewScanner(file)

	//Read numbers from first line
	scanner.Scan()
	line := scanner.Text()
	splitted := strings.Split(line, ",")
	for _, v := range splitted {
		number, _ := strconv.Atoi(v)
		numbers = append(numbers, number)
	}
	scanner.Scan()

	//Read Boards
	//var lines []string
	var boards [][5][5]int
	var board [5][5]int
	x := 0
	for scanner.Scan() {
		line := scanner.Text()
		//lines = append(lines, line)
		if len(line) == 0 {
			boards = append(boards, board)
			x = 0
		} else {
			line = strings.TrimSpace(line)
			line = strings.ReplaceAll(line, "  ", " ")
			splitted := strings.Split(line, " ")
			for y, v := range splitted {
				board[x][y], _ = strconv.Atoi(v)
			}
			x++
		}
	}
	boards = append(boards, board)
	return numbers, boards, scanner.Err()
}
