package game

import (
	"fmt"
)

var playerDiceMap = make(map[int]int)
var playerScoreMap = make(map[int]int)

func InputFirstTime () {
	var (
		n int
		diceThrow int
	)

	fmt.Println("Please enter the number of players")
	fmt.Scan(&n)

	fmt.Println("player order will be decided according to the score in the first dice round")
	for i:=0; i<n; i++ {
		fmt.Println("Player ", i+1, "rolls")
		fmt.Scan(&diceThrow)
		playerDiceMap[i] = diceThrow
	}

	positionArray := makePositionArray(playerDiceMap)
	fmt.Println("game starts")
	winningPlayer := startGame(positionArray,n)
	if winningPlayer != 0 {
		fmt.Println("The winning player is ", winningPlayer)
		fmt.Println("Game Over !!")
	} else {
		fmt.Println("No winner ")
	}

}

func startGame(positionArray []int, n int)int {
	var (
		diceThrow int
		probableScore int
		score int
		gameEnd bool
		winningPlayer int
	)
	fmt.Println("The Game has finally started")
	for {
		fmt.Println("Next Chance")
		for _, playerNumber := range positionArray {
			fmt.Println("enter the dice roll for player ", playerNumber+1)
			fmt.Scan(&diceThrow)
			playerDiceMap[playerNumber] = diceThrow
			probableScore = playerScoreMap[playerNumber] + diceThrow
			if probableScore == 100 {
				gameEnd = true
				winningPlayer = playerNumber
				score = probableScore
				//cout for all players
			} else if dest, found := snakeMap[probableScore]; found {
				score = dest
			}else if dest, found := ladderMap[probableScore]; found {
				score = dest
			}else if probableScore > 100 {
				score = playerScoreMap[playerNumber]
			} else {
				score = probableScore
			}
			playerScoreMap[playerNumber] = score
		}
		printForAllPlayers(n)
		if gameEnd {
			break
		}
	}
	return winningPlayer+1
}

func printForAllPlayers(n int)  {
	for i := 0; i < n; i++ {
		fmt.Println("Player ", i+1, " is at ", playerScoreMap[i])
	}
}