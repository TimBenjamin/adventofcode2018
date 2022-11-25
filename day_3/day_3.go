package day_3

import (
	"adventofcode2018/util"
	"strconv"
	"strings"
)

var input []string

func getClaim(claim string) (id string, x int, y int, w int, h int) {
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

func partOne() (out int) {
	cells := map[string]int{}
	for _, claim := range input {
		_, x, y, w, h := getClaim(claim)
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

func partTwo() (out int) {
	// We want to find an area with no overlaps
	// So assemble the map as in part 1
	// and for each area, go through the map again and discard all areas where there's >1 in a point
	cells := map[string]int{}
	for _, claim := range input {
		_, x, y, w, h := getClaim(claim)
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
		goodClaim := true
		id, x, y, w, h := getClaim(claim)
		// we want to count the number of overlapping cells.
		for i := x; i < x+w; i++ {
			for j := y; j < y+h; j++ {
				coord := strconv.Itoa(i) + "," + strconv.Itoa(j)
				if cells[coord] > 1 {
					// this claim is invalid
					goodClaim = false
					break
				}
			}
		}
		if goodClaim {
			println(id)
			return
		}
	}

	return
}

func Call(part string, inputFile string) string {
	input = util.ParseInputIntoLines(inputFile)
	var r int
	if part == "1" {
		r = partOne()
	} else {
		r = partTwo()
	}
	return strconv.Itoa(r)
}
