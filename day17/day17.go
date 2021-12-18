package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type Target struct {
	Xmin int
	Xmax int
	Ymin int
	Ymax int
}

// H Heuristik -> 2*row*5
func main() {
	inputFile := flag.String("inputFile", "data.input", "Relative file path to use as input.")
	flag.Parse()
	xmin, xmax, ymin, ymax, _ := readTargetArea(*inputFile)
	target := Target{xmin, xmax, ymin, ymax}

	//x,y position starts at 0,0
	//x velocity goes to 0 (not moving) forward
	//y velocity goes to -1 (falling down) up and down possible
	//need to find initial x,y velocity values
	xpos, ypos := 0, 0

	//start value velocity x y
	// solution for my data input: vxstart, vystart := 18, 99
	vxstart, vystart := 0, target.Ymin
	var missedDistance int
	var yBest int
	//find x with "Gerader Wurf"
	xd := 0
	for {
		_, missedDistance = throw(xpos, ypos, vxstart+xd, 0, target)
		if missedDistance <= 0 {
			break
		}
		xd++
	}
	//find y
	//-> we want as high as possible big y value
	//when it drops to 0 the velocity is the start velocity just minus
	//therfore i set the startpoint here for target.Ymin
	yd := 0
	for {
		yBest, missedDistance = throw(xpos, ypos, xd, -1*(vystart+yd), target)
		yd++
		if missedDistance <= 0 && yBest != 0 {
			break
		}
	}
	fmt.Println(yBest, missedDistance)
}

func throw(xpos, ypos, vxstart, vystart int, target Target) (int, int) {
	//do steps as long it is not out of range
	hit := false
	yBest := 0
	missedDistance := 0
	for {
		m, d := missedTarget(xpos, ypos, target)
		if m {
			missedDistance = d
			break
		}
		//change pos
		xpos += vxstart
		ypos += vystart

		//check highest y
		if ypos > yBest {
			yBest = ypos
		}

		//change velocity
		if vxstart != 0 {
			vxstart--
		}
		vystart--

		if hasHitTarget(xpos, ypos, target) {
			fmt.Println("hit!")
			hit = true
			break
		}
	}
	if hit {
		return yBest, 0
	}
	return 0, missedDistance
}

func missedTarget(xpos int, ypos int, target Target) (missedSide bool, distance int) {
	//left or right out of target
	if xpos < target.Xmin && ypos <= target.Ymax {
		return true, target.Xmin - xpos //missed for this value left -> throw stronger
	}
	if xpos > target.Xmax && ypos <= target.Ymax {
		return true, -(xpos - target.Xmax) //missed for this value right -> throw less strong
	}
	//under target
	if ypos < target.Ymin {
		return true, 0
	}
	return false, 0
}

func hasHitTarget(xpos int, ypos int, target Target) bool {
	if xpos >= target.Xmin && xpos <= target.Xmax && ypos <= target.Ymax && ypos >= target.Ymin {
		return true
	}
	return false
}

func readTargetArea(path string) (int, int, int, int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, 0, 0, 0, err
	}
	defer file.Close()

	var xmin, xmax, ymin, ymax int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Sscanf(line, "target area: x=%d..%d, y=%d..%d", &xmin, &xmax, &ymin, &ymax)
	}
	return xmin, xmax, ymin, ymax, scanner.Err()
}

// Abs returns the absolute value of x.
func Abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}
