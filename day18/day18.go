package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"unicode"

	"github.com/lenahoinkis/AdventOfCode2021/utils"
)

// Solution with Binary Tree
func main() {
	inputFile := flag.String("inputFile", "ex2.input", "Relative file path to use as input.")
	flag.Parse()
	lines, _ := utils.ReadLines(*inputFile)

	t := &Node{data: -1}
	t = t.insertLine("[[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]],[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]]")
	t.explodes(0)
	print(os.Stdout, t, 0, 'M')

	x := &Node{data: -1}
	for i, line := range lines {
		if i == 0 {
			x = x.insertLine(line)
			continue
		}
		x = x.AddData(line)
		exploded, splited := true, true
		for exploded || splited {
			print(os.Stdout, x, 0, 'M')
			exploded = x.explodes(0)
			print(os.Stdout, x, 0, 'M')
			splited = x.split()
			fmt.Println()
		}
	}
	fmt.Println(x.magnitude())
}

type Node struct {
	left   *Node
	right  *Node
	parent *Node
	data   int
	// empty bool
}

func (n *Node) insertLine(line string) *Node {
	savedn := n
	for _, v := range line {
		switch v {
		case '[':
			n = n.insertEmptyNode() //create empty node on the left
		case ']':
			n = n.goOneUp()
		case ',':
			n = n.goUpAndInsertRight() //go up and create empty node on the right
		}
		if unicode.IsDigit(v) {
			number, _ := strconv.Atoi(string(v))
			n = n.insertNumber(number) // give current node the value
		}
	}
	return savedn

}

func (n *Node) insertEmptyNode() *Node {
	if n.left == nil {
		n.left = &Node{data: -1, left: nil, right: nil, parent: n}
	}
	return n.left
}

func (n *Node) goOneUp() *Node {
	return n.parent
}

func (n *Node) goUpAndInsertRight() *Node {
	n = n.parent //go up
	if n.right == nil {
		n.right = &Node{data: -1, left: nil, right: nil, parent: n}
	}
	return n.right
}

func (n *Node) insertNumber(number int) *Node {
	n.data = number
	return n
}

func (n *Node) AddData(line string) *Node {
	//add a new parent to insert the data
	n.parent = &Node{data: -1, left: n, right: nil, parent: nil}
	// add a right node here
	newRight := n.goUpAndInsertRight()
	newRight.insertLine(line)
	return n.parent
}

func (n *Node) explodes(depth int) bool {
	//go deep starting left
	explodeLeft, explodeRight := false, false
	if depth != 5 {
		if n.left != nil {
			explodeLeft = n.left.explodes(depth + 1)
		}
		if n.right != nil {
			explodeRight = n.right.explodes(depth + 1)
		}
	} else {
		//left case
		if n.parent.left == n {
			lastLeft := n.parent.searchLeft()
			//if we have a left field to add do it
			if lastLeft != nil {
				lastLeft.data += n.data
			}

			n.parent.left = nil
			n = nil

		} else {
			nextRight := n.parent.searchRight()
			if nextRight != nil {
				nextRight.data += n.data
			}

			//left is empty, we need to balance with 0
			if n.parent.left == nil {
				n.parent.data = 0
			}

			n.parent.right = nil
			n = nil
		}
		return true
	}
	return explodeLeft || explodeRight
}

func (n *Node) split() bool {
	splitLeft, splitRight := false, false
	//go deep starting left
	if n.left != nil {
		//safe the last left number in case of adding is required
		splitLeft = n.left.split()
	}
	if n.right != nil {
		splitRight = n.right.split()
	}
	//split if required
	if n.data >= 10 {
		//rounds down automatically by default
		n.left = &Node{data: n.data / 2, left: nil, right: nil, parent: n}
		n.right = &Node{data: n.data/2 + 1, left: nil, right: nil, parent: n}
		n.data = -1
		return true
	}
	return splitLeft || splitRight
}

func (n *Node) magnitude() int {
	//go deep starting left
	if n.left != nil {
		//safe the last left number in case of adding is required
		n.left.magnitude()
	}
	if n.right != nil {
		n.right.magnitude()
	}
	if n.left != nil && n.right != nil && n.left.data >= 0 && n.right.data >= 0 {
		n.data = 3*n.left.data + 2*n.right.data
	}
	return n.data
}

func (n *Node) searchRight() *Node {
	if n.parent == nil {
		return nil
	}
	//go up until we have a new route
	if n.parent.right != nil && n.parent.right != n {
		n = n.parent.right
		//as long as there is no left we go down right
		for n.left == nil {
			if n.right == nil {
				break
			}
			n = n.right
		}
		//go the left way down
		for n.left != nil {
			n = n.left
		}
		return n
	}
	return n.parent.searchRight()
}

func (n *Node) searchLeft() *Node {
	if n.parent == nil {
		return nil
	}
	if n.parent.left != nil && n.parent.left != n {
		n = n.parent.left
		//as long as there is no right we go down left
		for n.right == nil {
			if n.left == nil {
				break
			}
			n = n.left
		}
		//go the left way down
		for n.right != nil {
			n = n.right
		}
		return n
	}
	return n.parent.searchLeft()
}

func print(w io.Writer, node *Node, ns int, ch rune) {
	if node == nil {
		return
	}
	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.data)
	print(w, node.left, ns+2, 'L')
	print(w, node.right, ns+2, 'R')
}
