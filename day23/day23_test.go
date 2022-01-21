package main

import (
	"testing"
)

func Test_Solve(t *testing.T) {
	var tests = []struct {
		name string
		as   []*amphipod
		want int
	}{
		{"Everything already in Place",
			[]*amphipod{
				NewAmphipod(3, "A"),
				NewAmphipod(4, "A"),
				NewAmphipod(7, "B"),
				NewAmphipod(8, "B"),
				NewAmphipod(11, "C"),
				NewAmphipod(12, "C"),
				NewAmphipod(15, "D"),
				NewAmphipod(16, "D"),
			}, 0,
		},
		{"Top A and B switched",
			[]*amphipod{
				NewAmphipod(3, "B"),
				NewAmphipod(4, "A"),
				NewAmphipod(7, "A"),
				NewAmphipod(8, "B"),
				NewAmphipod(11, "C"),
				NewAmphipod(12, "C"),
				NewAmphipod(15, "D"),
				NewAmphipod(16, "D"),
			}, 46,
		},
		{"Top A and B switched",
			[]*amphipod{
				NewAmphipod(3, "B"),
				NewAmphipod(4, "A"),
				NewAmphipod(7, "C"),
				NewAmphipod(8, "D"),
				NewAmphipod(11, "B"),
				NewAmphipod(12, "C"),
				NewAmphipod(15, "D"),
				NewAmphipod(16, "A"),
			}, 12521,
		},
	}
	for _, tt := range tests {
		result := solve(tt.as, 999999)
		if result != tt.want {
			t.Errorf("Error in %s: \n Expected %d got: %d", tt.name, tt.want, result)
		}
	}
}

func Test_possibleMoves(t *testing.T) {
	t.SkipNow()
}

//Here we need to execute level2 first
//level 2 and right column
//level 1, level 2 is on place and on right column
func Test_isOnPlace(t *testing.T) {
	var tests = []struct {
		name     string
		a, other *amphipod
		want     bool
	}{
		{"A and A are correct in 11 and 15",
			NewAmphipod(4, "A"), NewAmphipod(3, "A"), true},
		{"A is on the bottom of B",
			NewAmphipod(4, "A"), NewAmphipod(3, "B"), true},
		{"B is on the bottom of A",
			NewAmphipod(3, "A"), NewAmphipod(4, "B"), false},
		{"We do not have others and A is correctly on 4",
			NewAmphipod(4, "A"), nil, true},
		{"We do not have others and A is on 1",
			NewAmphipod(1, "A"), nil, false},
	}
	for _, tt := range tests {
		tt.a.setOnPlace([]*amphipod{tt.other})
		if tt.want != tt.a.OnPlace {
			t.Errorf("Error in %s: \n Expected %t got: %t", tt.name, tt.want, tt.a.OnPlace)
		}
	}
}

func Test_calcCost(t *testing.T) {
	var tests = []struct {
		name  string
		a     *amphipod
		field int
		want  int
	}{
		{"A moves from 5 to 3",
			NewAmphipod(5, "A"), 3, 2},
		{"A moves from 7 to 5",
			NewAmphipod(7, "A"), 5, 2},
		{"A moves from 3 to 0",
			NewAmphipod(3, "A"), 0, 3},
		{"A moves from 3 to 9",
			NewAmphipod(3, "A"), 9, 4},
		{"A moves from 8 to 0",
			NewAmphipod(8, "A"), 0, 6},
		{"A moves from 17 to 4",
			NewAmphipod(4, "A"), 17, 9},
		{"A moves from 8 to 18",
			NewAmphipod(8, "A"), 18, 8},
		{"B moves from 8 to 18",
			NewAmphipod(8, "B"), 18, 80},
		{"C moves from 13 to 11",
			NewAmphipod(13, "C"), 11, 200},
	}
	for _, tt := range tests {
		tt.a.calcCost(tt.field)
		if tt.want != tt.a.cost {
			t.Errorf("Error in %s: \n Expected %d got: %d", tt.name, tt.want, tt.a.cost)
		}
	}
}

// is somewhere it beetween alreadz occupied?
func Test_isWayBlocked(t *testing.T) {
	t.SkipNow()
}

func Test_getValidFields(t *testing.T) {
	{
		var tests = []struct {
			name     string
			a, other *amphipod
			want     []int
		}{
			{"only A at place 0 and everthing empty",
				&amphipod{0, "A", false, false, 10}, nil, []int{1, 4, 5, 9, 13, 17, 18}},
			{"only A at place 16 (B room) and everthing empty",
				&amphipod{16, "A", false, false, 10}, nil, []int{0, 1, 4, 5, 9, 13, 17, 18}},
			{"A is at place 19 and B on bottom in A room",
				&amphipod{19, "A", false, false, 10},
				&amphipod{4, "B", false, false, 10},
				[]int{0, 1, 3, 5, 9, 13, 17, 18}},
			{"A is at place 16 and B on 13, so it blocks the left side",
				&amphipod{16, "A", false, false, 10},
				&amphipod{13, "B", false, false, 10},
				[]int{10, 11}},
			{"A is at place 4 and B on 3, so all is blocked",
				&amphipod{4, "A", false, false, 10},
				&amphipod{3, "B", false, false, 10},
				[]int{}},
		}
		for _, tt := range tests {
			result := tt.a.getValidFields([]*amphipod{tt.other})
			if len(tt.want) != len(result) {
				t.Errorf("Error in %s: \n Expected %d got: %d", tt.name, tt.want, result)
			}
		}
	}
}
