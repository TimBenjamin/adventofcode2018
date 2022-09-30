package day_1

import (
	"adventofcode2018/util"
	"strconv"
)

var input []string

func part_1() (out int) {
	for _, value := range input {
		sign := value[0]
		num := value[1:]
		var amt int
		amt, _ = strconv.Atoi(num)
		if sign == '-' {
			out -= amt
		} else {
			out += amt
		}
	}
	return
}

func part_2() (out int) {
	// repeat the loop through the input
	// keep applying the delta to produce a result
	// stop when we have seen the result before, and return it
	seen := map[int]bool{}
	for {
		for _, value := range input {
			sign := value[0]
			num := value[1:]
			var amt int
			amt, _ = strconv.Atoi(num)
			if sign == '-' {
				out -= amt
			} else {
				out += amt
			}
			println(out)
			if seen[out] {
				return out
			}
			seen[out] = true
		}
	}
}

func Call(part string, input_file string) string {
	input = util.Parse_input_into_lines(input_file)
	var r int
	if part == "1" {
		r = part_1()
	} else {
		r = part_2()
	}
	return strconv.Itoa(r)
}
