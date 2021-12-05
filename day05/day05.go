package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Vertex struct {
	X, Y int
}

func main() {
	inputFile := flag.String("inputFile", "ex.input", "Relative file path to use as input.")
	flag.Parse()

	coors, max, err := ReadPoints(*inputFile)
	if err != nil {
		fmt.Print(err)
	}

	m := make([][]int, max+1)
	for i := range m {
		m[i] = make([]int, max+1)
		for y := 0; y <= max; y++ {
			m[i][y] = 0
		}
	}

	//part1
	for _, coor := range coors {
		if coor[0].X == coor[1].X {
			sortedCoor := sortVertex(coor)
			for i := sortedCoor[0].Y; i <= sortedCoor[1].Y; i++ {
				m[i][sortedCoor[0].X]++
			}
		} else if coor[0].Y == coor[1].Y {
			sortedCoor := sortVertex(coor)
			for i := sortedCoor[0].X; i <= sortedCoor[1].X; i++ {
				m[sortedCoor[0].Y][i]++
			}
		}
	}

	//utils.PrintMatrix(m)
	fmt.Println(calcOverlap(m))

	//part2
	for _, coor := range coors {
		distance := math.Abs(float64(coor[1].Y)-float64(coor[0].Y)) + 1
		if coor[0].X < coor[1].X && coor[0].Y > coor[1].Y {
			for i := 0; i < int(distance); i++ {
				m[coor[0].Y-i][coor[0].X+i]++
			}
		} else if coor[0].X > coor[1].X && coor[0].Y < coor[1].Y {
			for i := 0; i < int(distance); i++ {
				m[coor[0].Y+i][coor[0].X-i]++
			}
		} else if coor[0].X < coor[1].X && coor[0].Y < coor[1].Y {
			for i := 0; i < int(distance); i++ {
				m[coor[0].Y+i][coor[0].X+i]++
			}
		} else if coor[0].X > coor[1].X && coor[0].Y > coor[1].Y {
			for i := 0; i < int(distance); i++ {
				m[coor[0].Y-i][coor[0].X-i]++
			}
		}
	}
	//utils.PrintMatrix(m)
	fmt.Println(calcOverlap(m))
}

func calcOverlap(m [][]int) int {
	sum := 0
	for _, i := range m {
		for _, v := range i {
			if v > 1 {
				sum++
			}
		}
	}
	return sum
}

func sortVertex(coor [2]Vertex) [2]Vertex {
	if coor[0].X > coor[1].X {
		temp := coor[0].X
		coor[0].X = coor[1].X
		coor[1].X = temp
	}
	if coor[0].Y > coor[1].Y {
		temp := coor[0].Y
		coor[0].Y = coor[1].Y
		coor[1].Y = temp
	}
	return coor
}

func ReadPoints(path string) ([][2]Vertex, int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, 0, err
	}
	defer file.Close()

	var points [][2]Vertex
	var point [2]Vertex

	scanner := bufio.NewScanner(file)

	x, max := 0, 0
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		splitted := strings.Split(line, " -> ")
		for i, v := range splitted {
			nums := strings.Split(v, ",")
			x, _ := strconv.Atoi(nums[0])
			y, _ := strconv.Atoi(nums[1])
			if x > max {
				max = x
			}
			if y > max {
				max = y
			}
			point[i] = Vertex{x, y}
		}
		points = append(points, point)
		x++
	}
	return points, max, err
}
