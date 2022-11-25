package day_9

import (
	"adventofcode2018/util"
	"fmt"
	"strconv"
	"strings"
)

func play() int {
	scores := make([]int, numPlayers)
	currentPlayer := -1
	marbles := []int{0}
	newMarble := 0
	currentMarblePosition := 0

	for i := 0; i <= finalMarbleValue; i++ {
		currentPlayer = (currentPlayer + 1) % numPlayers
		newMarble++
		if newMarble%23 == 0 {
			scores[currentPlayer] += newMarble
			removeMarblePosition := currentMarblePosition - 7
			if removeMarblePosition < 0 {
				removeMarblePosition = len(marbles) + removeMarblePosition
			}
			scores[currentPlayer] += marbles[removeMarblePosition]
			if removeMarblePosition == len(marbles)-1 {
				currentMarblePosition = 0
			} else {
				currentMarblePosition = removeMarblePosition
			}
			marbles = append(marbles[0:removeMarblePosition], marbles[removeMarblePosition+1:]...)
		} else {
			newMarblePosition := (currentMarblePosition + 2) % len(marbles)
			if newMarblePosition == 0 {
				marbles = append([]int{newMarble}, marbles...)
			} else {
				marbles = append(marbles[0:newMarblePosition], marbles[newMarblePosition-1:]...)
				marbles[newMarblePosition] = newMarble
			}
			currentMarblePosition = newMarblePosition

		}
		// NB, when the new marble looks like it goes at the end in the example, I'm putting it at the beginning
		// This should be OK as it's a circular structure
	}
	//fmt.Printf("Last marble used - scores are: %v\n", scores)
	highScore := 0
	winner := 0
	for k := 0; k < len(scores); k++ {
		if scores[k] > highScore {
			highScore = scores[k]
			winner = k + 1
		}
	}
	fmt.Printf("Winner is player %v with score of %v\n", winner, highScore)
	return highScore
}

func partOne() int {
	return play()
}

type Marble struct {
	Value    int
	Next     *Marble
	Previous *Marble
}

func createMarble() *Marble {
	marble := &Marble{}
	marble.Next = marble
	marble.Previous = marble
	return marble
}

func (marble *Marble) Clockwise(distance int) *Marble {
	for ; distance > 0; distance-- {
		marble = marble.Next
	}
	return marble
}

func (marble *Marble) Anticlockwise(distance int) *Marble {
	for ; distance > 0; distance-- {
		marble = marble.Previous
	}
	return marble
}

func (marble *Marble) Insert(marble_value int) *Marble {
	new_marble := &Marble{
		Value:    marble_value,
		Next:     marble,
		Previous: marble.Previous,
	}
	marble.Previous = new_marble
	new_marble.Previous.Next = new_marble
	return new_marble
}

func (marble *Marble) Remove() *Marble {
	marble.Previous.Next = marble.Next
	marble.Next.Previous = marble.Previous
	return marble.Next
}

func playWithLinkedLists() int {
	marble := createMarble()
	scores := make([]int, numPlayers)
	currentPlayer := 0

	for marbleValue := 1; marbleValue < finalMarbleValue; marbleValue++ {
		if marbleValue%23 == 0 {
			marbleToRemove := marble.Anticlockwise(7)
			marble = marbleToRemove.Remove()
			scores[currentPlayer] += marbleValue + marbleToRemove.Value
		} else {
			marble = marble.Clockwise(2).Insert(marbleValue)
		}
		currentPlayer = (currentPlayer + 1) % numPlayers
	}

	highScore := 0
	winner := 0
	for k := 0; k < len(scores); k++ {
		if scores[k] > highScore {
			highScore = scores[k]
			winner = k + 1
		}
	}
	fmt.Printf("Winner is player %v with score of %v\n", winner, highScore)
	return highScore
}

func partTwo() int {
	finalMarbleValue = finalMarbleValue * 100
	return playWithLinkedLists()
}

var numPlayers int
var finalMarbleValue int

func Call(part string, inputFile string) string {
	input := util.ParseSingleLineInput(inputFile)
	splitInput := strings.Split(input, " players; last marble is worth ")
	numPlayers, _ = strconv.Atoi(splitInput[0])
	finalMarbleValue, _ = strconv.Atoi(strings.Split(splitInput[1], " ")[0])
	fmt.Printf("number of players: %v\n", numPlayers)
	fmt.Printf("final marble value: %v\n", finalMarbleValue)
	var r int
	if part == "1" {
		r = partOne()
	} else {
		r = partTwo()
	}
	return strconv.Itoa(r)
}
