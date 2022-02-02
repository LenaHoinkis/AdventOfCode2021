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

type gameState struct {
	score    int
	startPos int
	player   int
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
	//with the help of reddit:
	//0..10 possible positions
	//0..21 possible points
	//0..2 players
	// -> 10*21^2 (44100) possible game states
	// map[gamestate]numberOfUniverses
	// map over gamestates, apply all possible outcomes, +outcome counters, -original counters

	//init
	playersCount = 0
	fmt.Sscanf(lines[0], "Player %d starting position: %d", &playersCount, &players[0].startPos)
	fmt.Sscanf(lines[1], "Player %d starting position: %d", &playersCount, &players[1].startPos)

	gameStatePlayer1 := gameState{0, players[1].startPos, 1}
	gameStatePlayer0 := gameState{0, players[0].startPos, 0}

	newGameStates := make(map[gameState]int)
	newGameStates[gameStatePlayer1]++
	newGameStates[gameStatePlayer0]++

	playerWins := []int{0, 0}

	//start game as long es there are new gamestates
	for len(newGameStates) > 0 {
		tempGameStates := make(map[gameState]int)
		for gs, count := range newGameStates {

			//see if this state is a win
			if gs.score >= 21 {
				playerWins[gs.player] += count
				continue
			}

			//calc next Scores
			//first switch the player
			var player int
			if gs.player == 0 {
				player = 1
			} else {
				player = 0
			}

			//now dice
			for a := 1; a <= 3; a++ {
				for b := 1; b <= 3; b++ {
					for c := 1; c <= 3; c++ {
						newGameState := getNewGameState(a+b+c+gs.startPos, gs.score, player)
						tempGameStates[newGameState] += count
					}
				}
			}
		}
		newGameStates = tempGameStates
	}
	fmt.Println(playerWins)
}

func getNewGameState(diced, score, player int) gameState {
	stoppedField := getStoppedField(diced)
	newScore := stoppedField + score
	return gameState{score: newScore, startPos: stoppedField, player: player}
}

func getStoppedField(diced int) int {
	stoppedField := diced % 10
	if stoppedField == 0 {
		stoppedField = 10
	}
	return stoppedField
}
