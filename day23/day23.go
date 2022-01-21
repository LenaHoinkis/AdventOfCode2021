package main

import (
	"flag"
	"fmt"
	"math"
	"sort"

	"github.com/lenahoinkis/AdventOfCode2021/utils"
)

/*
00 01 02 05 06 09 10 13 14 17 18
	  03	07	  11	15
	  04	08	  12	16
*/

type amphipod struct {
	current_field int
	color         string
	moved         bool
	OnPlace       bool
	cost          int
}

func NewAmphipod(i int, color string) *amphipod {
	a := new(amphipod)
	a.current_field = i
	a.color = color
	a.OnPlace = false
	return a
}

var inputFile = flag.String("inputFile", "ex1.input", "Relative file path to use as input.")

func main() {
	flag.Parse()
	lines, _ := utils.ReadLines(*inputFile)

	var amphipods []*amphipod
	i, lvl := 3, 0
	for _, l := range lines {
		for _, c := range l {
			if c >= 'A' && c <= 'D' {
				amphipods = append(amphipods, NewAmphipod(i+lvl, string(c)))
				i += 4
				if i == 19 {
					i = 3
					lvl++
				}
			}
		}
	}
	fmt.Println(solve(amphipods, 99999))
}

func solve(as []*amphipod, solution int) int {
	if getCosts(as) > solution {
		return solution
	}

	for i, a := range as {
		as[i] = a.setOnPlace(as)
	}

	// for every amphipod
	if allSolved(as) {
		return getCosts(as)
	}
	for i, a := range as {
		if !a.OnPlace {
			// get the valid places
			validFields := a.getValidFields(as)
			for _, field := range validFields {
				var newAs []*amphipod
				for x, v := range as {
					if x == i {
						newA := NewAmphipod(a.current_field, a.color)
						newA.cost = a.cost
						newA = newA.move(field)
						newAs = append(newAs, newA)
					} else {
						newAs = append(newAs, v)
					}
				}
				costs := solve(newAs, solution)
				if costs < solution {
					solution = costs
					fmt.Println(solution)
				}
			}
		}
	}
	return solution
}

func allSolved(as []*amphipod) bool {
	for _, a := range as {
		if !a.OnPlace {
			return false
		}
	}
	return true
}

func getCosts(as []*amphipod) (sum int) {
	for _, a := range as {
		sum += a.cost
	}
	return sum
}

func (a *amphipod) addCost(moves int) *amphipod {
	switch a.color {
	case "A":
		a.cost += 1 * moves
	case "B":
		a.cost += 10 * moves
	case "C":
		a.cost += 100 * moves
	case "D":
		a.cost += 1000 * moves
	}
	return a
}

func (a *amphipod) move(nextField int) *amphipod {
	a = a.calcCost(nextField)
	a.current_field = nextField
	a.moved = true
	return a
}

func (a *amphipod) calcCost(nextField int) *amphipod {
	moves := 0
	startField, endField := a.current_field, nextField
	//moving in
	if isInRoom(a.current_field, 0) {
		moves++
		startField--
	}
	if isInRoom(a.current_field, 1) {
		moves += 2
		startField -= 2
	}

	//moving out
	if isInRoom(nextField, 0) {
		moves++
		endField--
	}
	if isInRoom(nextField, 1) {
		moves += 2
		endField -= 2
	}

	//moving on floor
	x, y := 0, 0
	if startField < endField {
		x = startField
		y = endField
	} else {
		x = endField
		y = startField
	}

	for i := x; i < y; i++ {
		moves++
		if isInRoom(i, 1) {
			moves -= 2
		}
	}

	a = a.addCost(int(math.Abs(float64(moves))))
	return a
}

func (a *amphipod) getValidFields(others []*amphipod) []int {
	if a.OnPlace {
		return []int{}
	}

	var validFields []int

	//floor is only allowed on the first
	if !a.moved {
		validFields = []int{0, 1, 5, 9, 13, 17, 18}
	}

	//add fitting fields of color
	f1, f2 := fieldForColor(a.color)
	validFields = append(validFields, f1, f2)
	sort.Ints(validFields)

	//remove current field
	validFields = remove(validFields, a.current_field)

	//remove occupied fields
	for _, other := range others {
		if other != nil {

			if other == a {
				continue
			}
			//remove this field
			validFields = remove(validFields, other.current_field)

			//top blocks
			if isInRoom(other.current_field, 0) && other.current_field+1 == a.current_field {
				return []int{}
			}

			//remove blocked ones when on the floor
			if !isInRoom(other.current_field, 0) && !isInRoom(other.current_field, 1) {
				//everything on the left
				if other.current_field < a.current_field {
					for i := 0; i < other.current_field; i++ {
						validFields = remove(validFields, i)
					}
				} else {
					for i := 18; i > other.current_field; i-- {
						validFields = remove(validFields, i)
					}
				}
			}
		}
	}

	//when bottom of sideroom is not filled top is not an option
	if contains(validFields, f2) {
		validFields = remove(validFields, f1)
	}

	return validFields
}

func (a *amphipod) setOnPlace(others []*amphipod) *amphipod {
	//is it already on place?
	if a.OnPlace {
		return a
	}

	//is it not the right column?
	if !isFieldForColor(a.color, a.current_field) {
		return a
	}
	//is on the first level?
	if isInRoom(a.current_field, 0) {
		//are the others here, which are not correct?
		for _, other := range others {
			if other != nil {
				if other.current_field == a.current_field+1 && other.color != a.color {
					return a
				}
			}
		}
	}
	a.OnPlace = true
	return a
}

func isInRoom(field, lvl int) bool {
	for i := 3; i <= 15; i += 4 {
		if field == i+lvl {
			return true
		}
	}
	return false
}

func fieldForColor(color string) (int, int) {
	switch color {
	case "A":
		return 3, 4
	case "B":
		return 7, 8
	case "C":
		return 11, 12
	case "D":
		return 15, 16
	}
	return 0, 0
}

func isFieldForColor(color string, field int) bool {
	f1, f2 := fieldForColor(color)
	return isField(field, f1, f2)
}

func isField(field, a, b int) bool {
	if field == a || field == b {
		return true
	}
	return false
}

func remove(s []int, in int) []int {
	if contains(s, in) {
		i := sort.SearchInts(s, in)
		s2 := append(s[0:i], s[i+1:]...)
		sort.Ints(s2)
		return s2
	}
	return s
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
