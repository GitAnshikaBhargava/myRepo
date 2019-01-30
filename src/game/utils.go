package game

import (
	"fmt"
	"sort"
)

func makePositionArray(playerDiceMap map[int]int) []int{
	positionArray := make([]int, len(playerDiceMap))
	pl := rankByDiceThrow(playerDiceMap)
	for index, pair := range pl {
		fmt.Println("index :", index)
		positionArray[index] = pair.PlayerNumber
	}
	return positionArray
}

func rankByDiceThrow(playerDiceMap map[int]int) PairList{
	pl := make(PairList, len(playerDiceMap))
	i := 0
	for k, v := range playerDiceMap {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

type Pair struct {
	PlayerNumber int
	DiceThrowScore int
}

type PairList []Pair

func (p PairList) Len() int { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].DiceThrowScore < p[j].DiceThrowScore }
func (p PairList) Swap(i, j int){ p[i], p[j] = p[j], p[i] }
