package main

import (
	"flag"
	"fmt"

	"github.com/lenahoinkis/AdventOfCode2021/utils"
)

var inputFile = flag.String("inputFile", "ex1.input", "Relative file path to use as input.")

type Player struct {
	score    int
	startPos int
}

func main() {
	flag.Parse()
	lines, _ := utils.ReadLines(*inputFile)
	fmt.Println(lines)

	// part1
	players := make([]Player, 2)
	playersCount := 0
	fmt.Sscanf(lines[0], "Player %d starting position: %d", &playersCount, &players[0].startPos)
	fmt.Sscanf(lines[1], "Player %d starting position: %d", &playersCount, &players[1].startPos)

	diceValue := 0
	currentPlayer := 0
	rolled := 0
	for players[currentPlayer].score < 1000 {
		for i, player := range players {
			diced := diceValue + 1 + diceValue + 2 + diceValue + 3 + player.startPos
			stoppedField := getStoppedField(diced)

			player.score += stoppedField
			player.startPos = stoppedField

			diceValue += 3
			rolled += 3
			currentPlayer = i
			players[i] = player
			if player.score >= 1000 {
				break
			}
		}
	}
	if players[0].score > players[1].score {
		fmt.Println(players[1].score * rolled)
	} else {
		fmt.Println(players[0].score * rolled)
	}

	//part 2
	playersCount = 0
	fmt.Sscanf(lines[0], "Player %d starting position: %d", &playersCount, &players[0].startPos)
	fmt.Sscanf(lines[1], "Player %d starting position: %d", &playersCount, &players[1].startPos)

	diceValue = 0
	currentPlayer = 0
	rolled = 0
	for players[currentPlayer].score < 24 {
		for i, player := range players {
			for i := 1; i <= 3; i++ {
				diced := diceValue + i + player.startPos
				stoppedField := getStoppedField(diced)
				player.score += stoppedField
				player.startPos = stoppedField
			}

			diceValue += 3
			rolled += 3
			currentPlayer = i
			players[i] = player
			if player.score >= 24 {
				break
			}
		}
	}
	if players[0].score > players[1].score {
		fmt.Println(players[1].score * rolled)
	} else {
		fmt.Println(players[0].score * rolled)
	}
}

func getStoppedField(diced int) int {
	stoppedField := diced % 10
	if stoppedField == 0 {
		stoppedField = 10
	}
	return stoppedField
}

//1
// 3
//	6
// 	7
// 	8

// 4
// 	...
// 5
// 	...
//2
// 3
// 4
// 5
//3
// 3
// 4
// 5
