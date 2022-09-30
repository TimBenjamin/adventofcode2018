package day_3

import (
	"adventofcode2018/util"
	"strconv"
	"strings"
)

var input []string

func _get_claim(claim string) (id string, x int, y int, w int, h int) {
	// #1261 @ 228,739: 13x18
	claim_split := strings.Split(claim, " ")
	id = claim_split[0][1:]
	xy := strings.Split(claim_split[2][0:len(claim_split[2])-1], ",")
	x, _ = strconv.Atoi(xy[0])
	y, _ = strconv.Atoi(xy[1])
	wh := strings.Split(claim_split[3], "x")
	w, _ = strconv.Atoi(wh[0])
	h, _ = strconv.Atoi(wh[1])
	// fmt.Printf("id: %v / x: %v / y: %v / w: %v / h: %v\n", id, x, y, w, h)
	return id, x, y, w, h
}

func part_1() (out int) {
	cells := map[string]int{}
	for _, claim := range input {
		_, x, y, w, h := _get_claim(claim)
		// we want to count the number of overlapping cells.
		for i := x; i < x+w; i++ {
			for j := y; j < y+h; j++ {
				coord := strconv.Itoa(i) + "," + strconv.Itoa(j)
				//println(coord)
				cells[coord] += 1
			}
		}
	}
	for cell, count := range cells {
		if count > 1 {
			println(cell)
			out++
		}
	}
	return
}

func part_2() (out int) {
	// We want to find an area with no overlaps
	// So assemble the map as in part 1
	// and for each area, go through the map again and discard all areas where there's >1 in a point
	cells := map[string]int{}
	for _, claim := range input {
		_, x, y, w, h := _get_claim(claim)
		// we want to count the number of overlapping cells.
		for i := x; i < x+w; i++ {
			for j := y; j < y+h; j++ {
				coord := strconv.Itoa(i) + "," + strconv.Itoa(j)
				//println(coord)
				cells[coord] += 1
			}
		}
	}
	for _, claim := range input {
		good_claim := true
		id, x, y, w, h := _get_claim(claim)
		// we want to count the number of overlapping cells.
		for i := x; i < x+w; i++ {
			for j := y; j < y+h; j++ {
				coord := strconv.Itoa(i) + "," + strconv.Itoa(j)
				if cells[coord] > 1 {
					// this claim is invalid
					good_claim = false
					break
				}
			}
		}
		if good_claim {
			println(id)
			return
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
