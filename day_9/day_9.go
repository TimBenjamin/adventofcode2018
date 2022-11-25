package day_9

import (
	"adventofcode2018/util"
	"fmt"
	"strconv"
	"strings"
)

func play() int {
	scores := make([]int, num_players)
	current_player := -1
	marbles := []int{0}
	new_marble := 0
	current_marble_position := 0

	for i := 0; i <= final_marble_value; i++ {
		current_player = (current_player + 1) % num_players
		new_marble++
		if new_marble%23 == 0 {
			scores[current_player] += new_marble
			remove_marble_position := current_marble_position - 7
			if remove_marble_position < 0 {
				remove_marble_position = len(marbles) + remove_marble_position
			}
			scores[current_player] += marbles[remove_marble_position]
			if remove_marble_position == len(marbles)-1 {
				current_marble_position = 0
			} else {
				current_marble_position = remove_marble_position
			}
			marbles = append(marbles[0:remove_marble_position], marbles[remove_marble_position+1:]...)
		} else {
			new_marble_position := (current_marble_position + 2) % len(marbles)
			if new_marble_position == 0 {
				marbles = append([]int{new_marble}, marbles...)
			} else {
				marbles = append(marbles[0:new_marble_position], marbles[new_marble_position-1:]...)
				marbles[new_marble_position] = new_marble
			}
			current_marble_position = new_marble_position

		}
		// NB, when the new marble looks like it goes at the end in the example, I'm putting it at the beginning
		// This should be OK as it's a circular structure
	}
	//fmt.Printf("Last marble used - scores are: %v\n", scores)
	high_score := 0
	winner := 0
	for k := 0; k < len(scores); k++ {
		if scores[k] > high_score {
			high_score = scores[k]
			winner = k + 1
		}
	}
	fmt.Printf("Winner is player %v with score of %v\n", winner, high_score)
	return high_score
}

func part_1() int {
	return play()
}

type Marble struct {
	value    int
	next     *Marble
	previous *Marble
}

func create_marble() *Marble {
	marble := &Marble{}
	marble.next = marble
	marble.previous = marble
	return marble
}

func (marble *Marble) Clockwise(distance int) *Marble {
	for ; distance > 0; distance-- {
		marble = marble.next
	}
	return marble
}

func (marble *Marble) Anticlockwise(distance int) *Marble {
	for ; distance > 0; distance-- {
		marble = marble.previous
	}
	return marble
}

func (marble *Marble) Insert(marble_value int) *Marble {
	new_marble := &Marble{
		value:    marble_value,
		next:     marble,
		previous: marble.previous,
	}
	marble.previous = new_marble
	new_marble.previous.next = new_marble
	return new_marble
}

func (marble *Marble) Remove() *Marble {
	marble.previous.next = marble.next
	marble.next.previous = marble.previous
	return marble.next
}

func play_with_linked_lists() int {
	marble := create_marble()
	scores := make([]int, num_players)
	current_player := 0

	for marble_value := 1; marble_value < final_marble_value; marble_value++ {
		if marble_value%23 == 0 {
			marble_to_remove := marble.Anticlockwise(7)
			marble = marble_to_remove.Remove()
			scores[current_player] += marble_value + marble_to_remove.value
		} else {
			marble = marble.Clockwise(2).Insert(marble_value)
		}
		current_player = (current_player + 1) % num_players
	}

	high_score := 0
	winner := 0
	for k := 0; k < len(scores); k++ {
		if scores[k] > high_score {
			high_score = scores[k]
			winner = k + 1
		}
	}
	fmt.Printf("Winner is player %v with score of %v\n", winner, high_score)
	return high_score
}

func part_2() int {
	final_marble_value = final_marble_value * 100
	return play_with_linked_lists()
}

var num_players int
var final_marble_value int

func Call(part string, input_file string) string {
	input := util.Parse_single_line_input(input_file)
	split_input := strings.Split(input, " players; last marble is worth ")
	num_players, _ = strconv.Atoi(split_input[0])
	final_marble_value, _ = strconv.Atoi(strings.Split(split_input[1], " ")[0])
	fmt.Printf("num_players: %v\n", num_players)
	fmt.Printf("last_marble_value: %v\n", final_marble_value)
	var r int
	if part == "1" {
		r = part_1()
	} else {
		r = part_2()
	}
	return strconv.Itoa(r)
}
