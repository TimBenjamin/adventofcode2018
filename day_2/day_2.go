package day_2

import (
	"adventofcode2018/util"
	"fmt"
	"strconv"
	"strings"
)

var input []string

func part_1() int {
	// count how many have exactly 2 of any letter
	// count how many have exactly 3 of any letter
	// multiply the two numbers
	count_2 := 0
	count_3 := 0

	for _, idstring := range input {
		m := map[string]int{}
		chars := strings.Split(idstring, "")
		for _, ch := range chars {
			if m[ch] > 0 {
				m[ch]++
			} else {
				m[ch] = 1
			}
		}
		found_2 := false
		found_3 := false
		for _, count := range m {
			if count == 2 {
				found_2 = true
			} else if count == 3 {
				found_3 = true
			}
		}
		if found_2 {
			count_2++
		}
		if found_3 {
			count_3++
		}
	}

	return count_2 * count_3
}

func part_2() (out int) {
	// find two strings that differ only by one character
	// abc = 1a + 2b + 3c
	for idx_1, str_1 := range input {
		for idx_2, str_2 := range input {
			if idx_1 == idx_2 {
				continue
			}
			str_1_split := strings.Split(str_1, "")
			str_2_split := strings.Split(str_2, "")
			diffs := 0
			common := ""
			for i := 0; i < len(str_1_split); i++ {
				if str_1_split[i] != str_2_split[i] {
					diffs++
				} else {
					common += str_1_split[i]
				}
				if diffs > 1 {
					break
				}
			}
			if diffs == 1 {
				fmt.Printf("%v vs %v has %v diffs\n", str_1, str_2, diffs)
				// What letters are common between the two correct box IDs?
				println(common)
				return
			}
		}
	}
	return
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
