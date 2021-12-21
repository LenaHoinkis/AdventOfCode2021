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

// H Heuristik -> 2*row*5
func main() {
	inputFile := flag.String("inputFile", "data.input", "Relative file path to use as input.")
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
	//start node
	nodes[0].Cost = 0

	//working queue with pointer of all nodes
	//There index in node object are the same as the nodes index
	queue := make([]*Node, 1)
	queue[0] = &nodes[0]

	for len(queue) > 0 {
		node := smallestNodeInQueue(queue)
		neighbours := getNeighbours(*node, nodes, rows)
		for _, n := range neighbours {
			cost := node.Cost + n.Risk
			if cost < n.Cost || n.Cost == -1 {
				nodes[n.Index].Predecessor = node
				nodes[n.Index].Cost = cost
			}
			if !nodes[n.Index].Done && !isNodeAlreadyInQueue(&n, queue) {
				queue = append(queue, &nodes[n.Index])
			}
		}
		//printValues(rows, nodes)
		nodes[node.Index].Done = true
		queue = queue[1:]
	}
	fmt.Println(nodes[len(nodes)-1])
}

func isNodeAlreadyInQueue(n *Node, q []*Node) bool {
	for _, v := range q {
		if v.Index == n.Index {
			return true
		}
	}
	return false
}

func printValues(row int, queue []Node) {
	for i, v := range queue {
		fmt.Printf("%d, ", v.Cost)
		if (i+1)%row == 0 {
			fmt.Println()
		}
	}
}

func getNeighbours(n Node, nodes []Node, rows int) []Node {
	neighbours := make([]Node, 0)
	//without this check we can go out of right side to the next line
	if n.Index%rows != rows {
		neighbours = addNeighbourIfPossible(n.Index+1, neighbours, nodes)
	}
	//without this check we can go out of left side to the last line
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
