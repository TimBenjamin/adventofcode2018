package day_5

import (
	"adventofcode2018/util"
	"fmt"
	"strconv"
	"strings"
)

var input string

func partOne() int {
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
			newInput := strings.Join(s, "")
			if len(newInput) != len(input) {
				changed = true
				input = newInput
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
			newInput := strings.Join(s, "")
			if len(newInput) != len(sequence) {
				changed = true
				sequence = newInput
			}
		}
		if changed {
			changed = false
		} else {
			return sequence
		}
	}
}

func partTwo() int {
	// This time we want to reduce all a/A, see what that comes to, compare to removing b/B, etc
	best := 100000
	alpha := "abcdefghijklmnopqrstuvwxyz"
	for i := 0; i < len(alpha); i++ {
		letter := alpha[i : i+1]
		s := strings.Split(input, letter)
		j := strings.Join(s, "")
		ucLetter := strings.ToUpper(letter)
		s = strings.Split(j, ucLetter)
		j = strings.Join(s, "")
		res := _reduce(j)
		if len(res) < best {
			best = len(res)
			fmt.Printf("New best: %v\n", best)
		}
	}

	return best
}

func Call(part string, inputFile string) string {
	input = util.ParseSingleLineInput(inputFile)
	var r int
	if part == "1" {
		r = partOne()
	} else {
		r = partTwo()
	}
	return strconv.Itoa(r)
}
