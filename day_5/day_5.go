package day_5

import (
	"adventofcode2018/util"
	"fmt"
	"strconv"
	"strings"
)

var input string

func part_1() int {
	// recursively until there are no changes
	// remove all pairs of lowercase and uppercase that are next to each other
	// e.g. Aa or Bb or cC

	alpha := "abcdefghijklmnopqrstuvwxyz"
	for {
		changed := false
		for i := 0; i < len(alpha); i++ {
			combo := alpha[i:i+1] + strings.ToUpper(alpha[i:i+1])
			s := strings.Split(input, combo)
			j := strings.Join(s, "")
			combo = strings.ToUpper(alpha[i:i+1]) + alpha[i:i+1]
			s = strings.Split(j, combo)
			new_input := strings.Join(s, "")
			if len(new_input) != len(input) {
				changed = true
				input = new_input
			}
		}
		if changed {
			changed = false
		} else {
			break
		}
	}

	// How many units remain after fully reacting the polymer you scanned?
	return len(input)
}

func _reduce(sequence string) string {
	alpha := "abcdefghijklmnopqrstuvwxyz"

	for {
		changed := false
		for i := 0; i < len(alpha); i++ {
			combo := alpha[i:i+1] + strings.ToUpper(alpha[i:i+1])
			s := strings.Split(sequence, combo)
			j := strings.Join(s, "")
			combo = strings.ToUpper(alpha[i:i+1]) + alpha[i:i+1]
			s = strings.Split(j, combo)
			new_input := strings.Join(s, "")
			if len(new_input) != len(sequence) {
				changed = true
				sequence = new_input
			}
		}
		if changed {
			changed = false
		} else {
			return sequence
		}
	}
}

func part_2() int {
	// This time we want to reduce all a/A, see what that comes to, compare to removing b/B, etc
	best := 100000
	alpha := "abcdefghijklmnopqrstuvwxyz"
	for i := 0; i < len(alpha); i++ {
		letter := alpha[i : i+1]
		s := strings.Split(input, letter)
		j := strings.Join(s, "")
		uc_letter := strings.ToUpper(letter)
		s = strings.Split(j, uc_letter)
		j = strings.Join(s, "")
		res := _reduce(j)
		if len(res) < best {
			best = len(res)
			fmt.Printf("New best: %v\n", best)
		}
	}

	return best
}

func Call(part string, input_file string) string {
	input = util.Parse_single_line_input(input_file)
	var r int
	if part == "1" {
		r = part_1()
	} else {
		r = part_2()
	}
	return strconv.Itoa(r)
}
