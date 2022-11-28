package day_12

import (
	"adventofcode2018/util"
	"fmt"
	"strconv"
	"strings"
)

func advance(state string, leftPot int) (string, int) {
	state = "......" + state + "......" // add sufficient empty pots to the left and right
	leftPot -= 6
	newState := []string{}
	for k := 0; k < len(state); k++ {
		newState = append(newState, ".")
	}
	// take 5-length fragments... (starting from the old 1st element which is now the 3rd element)
	for p := 2; p < len(state)-5; p++ {
		fragment := state[p-2 : p+3]
		//... and apply the instruction ...
		for _, generationData := range instructionsData {
			instruction := strings.Split(generationData, " => ")
			if fragment == instruction[0] {
				// ... then update the corresponding elemment of newState
				newState[p] = instruction[1]
				break
			}
		}
	}
	state = strings.Join(newState, "")
	// for vis / efficiency remove all the empty pots to the left and right:
	left := 0
	for i, n := range state {
		if n == '#' {
			left = i
			leftPot += i
			break
		}
	}
	state = state[left:]
	right := 0
	for r := len(state) - 1; r > 0; r-- {
		if state[r] == '#' {
			right = r
			break
		}
	}
	state = state[0 : right+1]
	return state, leftPot
}

func calculateSum(state string, leftPot int) int {
	// calculate the sum of all pots with # in
	sum := 0
	for i, n := range state {
		if n == '#' {
			sum += i + leftPot
		}
	}
	return sum
}

func partOne() int {
	state := initialState
	leftPot := 0
	for i := 0; i < 20; i++ {
		state, leftPot = advance(state, leftPot)
	}
	fmt.Printf("The final state is: %v\n", state)
	fmt.Printf("leftPot is: %v\n", leftPot)
	sum := calculateSum(state, leftPot)
	return sum
}

func partTwo() int {
	state := initialState
	leftPot := 0
	// need to do 50000000000 generations!
	// the pattern stays the same after a short time, but just shifts left
	// I need to know how far left by 50000000000 rather than the actual pattern
	oldState := state
	generation := 0
	for ; generation < 500; generation++ {
		state, leftPot = advance(state, leftPot)
		if state == oldState {
			fmt.Printf("The old state is the same as the new state after generation %v with leftPot %v\n", generation, leftPot)
			break
		}
		oldState = state
	}
	// the pattern just seems to shift right by 1 every time
	// i.e. leftPot increases by 1 each time
	shift := (50000000000 - 1) - generation
	leftPot += shift
	fmt.Printf("The final state is: %v\n", state)
	fmt.Printf("leftPot is: %v\n", leftPot)
	sum := calculateSum(state, leftPot)
	return sum
}

var initialState string
var instructionsData []string

func Call(part string, inputFile string) string {
	initialState = util.ParseSingleLineInput(inputFile)[15:]
	instructionsData = util.ParseInputIntoLines(inputFile)[2:]
	var r int
	if part == "1" {
		r = partOne()
	} else {
		r = partTwo()
	}
	return strconv.Itoa(r)
}
