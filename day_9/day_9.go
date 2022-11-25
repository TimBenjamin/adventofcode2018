package day_9

import (
	"adventofcode2018/util"
	"fmt"
	"strconv"
	"strings"
)

func play() {
	scores := make([]int, num_players)
	current_player := -1
	marbles := []int{0}
	current_marble := 0
	marble_counter := 0
	current_marble_position := 0

	for {
		current_player += 1
		if current_player >= num_players {
			current_player = 0
		}
		new_marble := marble_counter + 1
		marble_counter += 1
		if new_marble%23 == 0 {
			//fmt.Printf("!! new marble %v is a multiple of 23!\n", new_marble)
			scores[current_player] += new_marble
			remove_marble_position := current_marble_position - 7
			if remove_marble_position < 0 {
				remove_marble_position = len(marbles) + remove_marble_position
			}
			scores[current_player] += marbles[remove_marble_position]
			if remove_marble_position == len(marbles)-1 {
				current_marble = marbles[0]
				current_marble_position = 0
			} else {
				current_marble = marbles[remove_marble_position+1]
				current_marble_position = remove_marble_position
			}
			//fmt.Printf("current marble is now: %v\n", current_marble)
			_marbles := []int{}
			_marbles = append(_marbles, marbles[0:remove_marble_position]...)
			_marbles = append(_marbles, marbles[remove_marble_position+1:]...)
			marbles = _marbles
		} else {
			new_marble_position := (current_marble_position + 2) % len(marbles)
			_marbles := []int{}
			_marbles = append(_marbles, marbles[0:new_marble_position]...)
			_marbles = append(_marbles, new_marble)
			current_marble_position = len(_marbles) - 1
			_marbles = append(_marbles, marbles[new_marble_position:]...)
			marbles = _marbles
			current_marble = new_marble
		}
		//fmt.Printf("[%v] %v\n", current_player+1, marbles)
		if current_marble == last_marble_value {
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
			break
		}
		// NB, when the new marble looks like it goes at the end in the example, I'm putting it at the beginning
		// This should be OK as it's a circular structure
	}
}

func part_1() int {
	play()
	return 0
}

func part_2() int {
	last_marble_value = last_marble_value * 100
	play()
	return 0
}

var num_players int
var last_marble_value int

func Call(part string, input_file string) string {
	input := util.Parse_single_line_input(input_file)
	split_input := strings.Split(input, " players; last marble is worth ")
	num_players, _ = strconv.Atoi(split_input[0])
	last_marble_value, _ = strconv.Atoi(strings.Split(split_input[1], " ")[0])
	fmt.Printf("num_players: %v\n", num_players)
	fmt.Printf("last_marble_value: %v\n", last_marble_value)
	var r int
	if part == "1" {
		r = part_1()
	} else {
		r = part_2()
	}
	return strconv.Itoa(r)
}
