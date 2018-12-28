package main

import "fmt"

type Marble struct {
	Score    uint64
	Previous *Marble
	Next     *Marble
}

func (myself *Marble) InsertNext(newMarble *Marble) {
	newMarble.Previous = myself
	newMarble.Next = myself.Next

	oldNext := myself.Next
	myself.Next = newMarble
	oldNext.Previous = newMarble
}

func (myself *Marble) RemovePrevious() *Marble {
	toRemove := myself.Previous

	toRemove.Previous.Next = myself
	myself.Previous = toRemove.Previous

	toRemove.Previous = nil
	toRemove.Next = nil
	return toRemove
}

const SpecialScore = 23

func Day9(numMarbles uint64, numPlayers int) {
	currentMarble := &Marble{Score: 0}
	currentMarble.Next = currentMarble
	currentMarble.Previous = currentMarble

	scoreBoard := make(map[int]uint64)
	player := 1

	for nextMarble := uint64(1); nextMarble <= numMarbles; nextMarble++ {
		if nextMarble%SpecialScore == 0 {
			// Take the score
			scoreBoard[player] += nextMarble

			// Rotate counter-clockwise by 6
			for i := 0; i < 6; i++ {
				currentMarble = currentMarble.Previous
			}
			// Remove previous and take the score
			removing := currentMarble.RemovePrevious()
			scoreBoard[player] += removing.Score
		} else {
			// Rotate clockwise by 1
			currentMarble = currentMarble.Next

			// Insert new marble
			currentMarble.InsertNext(&Marble{Score: nextMarble})

			// Set current marble to the newly inserted marble
			currentMarble = currentMarble.Next
		}

		player++
		if player > numPlayers {
			player = 1
		}
	}

	highestPlayer := 0
	highestScore := uint64(0)
	for p, s := range scoreBoard {
		if s > highestScore {
			highestScore = s
			highestPlayer = p
		}
	}
	fmt.Printf("%d players, last marble is %d, winning player %d, score %d\n", numPlayers, numMarbles, highestPlayer, highestScore)
}

func main() {
	Day9(25, 9)
	Day9(1618, 10)
	Day9(7999, 13)
	Day9(1104, 17)
	Day9(6111, 21)
	Day9(5807, 30)
	Day9(70833, 486)

	Day9(7083300, 486)
}
