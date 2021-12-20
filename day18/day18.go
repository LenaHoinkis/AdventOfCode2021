package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"unicode"
)

// Solution with Binary Tree
func main() {
	inputFile := flag.String("inputFile", "ex1.input", "Relative file path to use as input.")
	flag.Parse()
	fmt.Println(inputFile)

	n := &Node{data: -1}
	n = n.insertLine("[[[[4,3],4],4],[7,[[8,4],9]]]")
	n = n.AddData("[1,1]")
	n.explodes(0)
	n.split()
	n.explodes(0)
	print(os.Stdout, n, 0, 'M')

	m := &Node{data: -1}
	m = m.insertLine("[[[[0,7],4],[[7,8],[6,0]]],[8,1]]")
	fmt.Println(m.magnitude())
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

func (n *Node) explodes(depth int) {
	//go deep starting left
	if depth != 5 {
		if n.left != nil {
			n.left.explodes(depth + 1)
		}
		if n.right != nil {
			n.right.explodes(depth + 1)
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
			nextRight.data += n.data

			//if we deleted the left already and the node is still -1
			if n.parent.data == -1 {
				n.parent.data = 0
			}

			n.parent.right = nil
			n = nil
		}
	}
}

func (n *Node) split() {
	//go deep starting left
	if n.left != nil {
		//safe the last left number in case of adding is required
		n.left.split()
	}
	if n.right != nil {
		n.right.split()
	}
	//split if required
	if n.data >= 10 {
		//rounds down automatically by default
		n.left = &Node{data: n.data / 2, left: nil, right: nil, parent: n}
		n.right = &Node{data: n.data/2 + 1, left: nil, right: nil, parent: n}
		n.data = -1
	}
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
	if n.parent.right != nil && n.parent.right != n {
		//check left
		n = n.parent.right
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
		//check right
		n = n.parent.left
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
