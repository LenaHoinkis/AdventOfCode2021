package main

import (
	"flag"
	"fmt"
	"sort"

	"github.com/lenahoinkis/AdventOfCode2021/utils"
)

type Node struct {
	Index       int
	Predecessor *Node
	Risk        int
	Cost        int
	Done        bool
}

func main() {
	inputFile := flag.String("inputFile", "t.input", "Relative file path to use as input.")
	flag.Parse()
	board, rows, _ := utils.ReadIntsWithoutSeperator(*inputFile)
	utils.PrintIntBoard(rows, board)

	//init
	nodes := make([]Node, len(board))
	for i, v := range board {
		nodes[i].Risk = v
		nodes[i].Index = i
		nodes[i].Cost = -1
		nodes[i].Done = false
	}

	//part1
	fmt.Println(solveDijkstra(nodes, rows))

	//part2
	indexSmallBoard := 0
	moreRows := rows * 5
	biggerMap := make([]Node, len(board)*25)
	//big board 5x5
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			//small board
			for j := 0; j < rows; j++ {
				for k := 0; k < rows; k++ {
					// board[i*row+j] = "abc" // like board[i][j] = "abc"
					// (j*moreRows + k): normal iteration now everzthing is in the left corner
					// (y * rows): move always a board wide to the right (new y value)
					// (x*moreRows*rows)> move boardwise down (new x value)
					i := (j*moreRows + k) + (x * moreRows * rows) + (y * rows)
					risk := (nodes[indexSmallBoard].Risk+x+y-1)%9 + 1 // need -1 before mod because we do not start with 0. add it afterwards again
					biggerMap[i].Risk = risk
					biggerMap[i].Index = i
					biggerMap[i].Cost = -1
					biggerMap[i].Done = false
					//instad of calculating i made it easy with counting and starting over
					indexSmallBoard++
					if indexSmallBoard >= len(board) {
						indexSmallBoard = 0
					}
				}
			}
		}
	}
	fmt.Println(solveDijkstra(biggerMap, moreRows))

}

func solveDijkstra(board []Node, rows int) int {
	//start node
	board[0].Cost = 0

	//working queue with pointer of the board
	//The index in node objects are the same as the board index
	queue := make([]*Node, 1)
	queue[0] = &board[0]
	for len(queue) > 0 {
		node := smallestNodeInQueue(queue)
		neighbours := getNeighbours(*node, board, rows)
		for _, n := range neighbours {
			cost := node.Cost + n.Risk
			if cost < n.Cost || n.Cost == -1 {
				board[n.Index].Predecessor = node
				board[n.Index].Cost = cost
			}
			if !board[n.Index].Done && !isNodeAlreadyInQueue(&n, queue) {
				queue = append(queue, &board[n.Index])
			}
		}
		//printValues(rows, board)
		board[node.Index].Done = true
		queue = queue[1:]
	}
	return board[len(board)-1].Cost
}

func isNodeAlreadyInQueue(n *Node, q []*Node) bool {
	for _, v := range q {
		if v.Index == n.Index {
			return true
		}
	}
	return false
}

/*
was useful for testing
func printValues(row int, queue []Node) {
	for i, v := range queue {
		fmt.Printf("%d, ", v.Risk)
		if (i+1)%row == 0 {
			fmt.Println()
		}
	}
}*/

func getNeighbours(n Node, nodes []Node, rows int) []Node {
	neighbours := make([]Node, 0)
	//without this check we can go out of right side to the next row
	if n.Index%rows != rows {
		neighbours = addNeighbourIfPossible(n.Index+1, neighbours, nodes)
	}
	//without this check we can go out of left side to the last row
	if n.Index%rows != 0 {
		neighbours = addNeighbourIfPossible(n.Index-1, neighbours, nodes)
	}
	neighbours = addNeighbourIfPossible(n.Index+rows, neighbours, nodes)
	neighbours = addNeighbourIfPossible(n.Index-rows, neighbours, nodes)
	return neighbours
}

func addNeighbourIfPossible(i int, neighbours []Node, nodes []Node) []Node {
	if i > 0 && i < len(nodes) {
		neighbours = append(neighbours, nodes[i])
	}
	return neighbours
}

func smallestNodeInQueue(queue []*Node) *Node {
	if len(queue) >= 2 {
		sort.Slice(queue, func(i, j int) bool {
			return queue[i].Cost < queue[j].Cost
		})
	}
	return queue[0]
}
