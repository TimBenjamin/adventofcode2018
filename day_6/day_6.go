package day_6

import (
	"adventofcode2018/util"
	"fmt"
	"strconv"
	"strings"
)

var input []string

func partOne() int {

	// any coordinate that does not have another coordinate further NSEW can be disregarded
	// as it will lead to an infinite area

	// aaaaa.cccc
	// aAaaa.cccc
	// aaaddecccc
	// aadddeccCc
	// ..dDdeeccc
	// bb.deEeecc
	// bBb.eeee..
	// bbb.eeefff
	// bbb.eeffff
	// bbb.ffffFf

	// 1, 1
	// 1, 6
	// 8, 3
	// 3, 4
	// 5, 5
	// 8, 9

	// A (1,1) is discounted because there is nothing lower than 1 in the x-axis
	// and nothing lower than 1 in the y-axis
	// F (8,9) is discounted because there is nothing higher than 8 in the x-axis
	// and nothing higher than 9 in the y-axis

	coords := [][]int{}
	xMin := 1000
	yMin := 1000
	xMax := 0
	yMax := 0
	for _, line := range input {
		pair := strings.Split(line, ", ")
		x, _ := strconv.Atoi(pair[0])
		y, _ := strconv.Atoi(pair[1])
		coords = append(coords, []int{x, y})
		if x < xMin {
			xMin = x
		}
		if x > xMax {
			xMax = x
		}
		if y < yMin {
			yMin = y
		}
		if y > yMax {
			yMax = y
		}
	}

	reducedCoords := [][]int{}
	for _, coord := range coords {
		if coord[0] > xMin && coord[0] < xMax && coord[1] > yMin && coord[1] < yMax {
			reducedCoords = append(reducedCoords, coord)
		}
	}

	// The reduced_coords are the ones that do not lead to an infinite area
	for _, coord := range reducedCoords {
		fmt.Printf("x: %v / y: %v\n", coord[0], coord[1])
	}

	// need to find the Manhattan distance to each other point...
	// Manhattan between (x1, y1) and (x2, y2) is |(x1-x2)| + |(y1-y2)|
	// so for point D (3, 4) and A (1, 1) =
	// (3-1) + (4-1) = 2 + 3 = 5

	// For each of those points in the possible space (1,1 etc)
	// work out the distance from A, B, C, D etc to that point
	// whichever is the lowest, we save that letter as in the rubric.
	// However rather than saving a letter, I'll save the coords of that letter (I am not using letters)

	distances := map[string]string{}
	for i := xMin; i <= xMax; i++ {
		for j := yMin; j <= yMax; j++ {
			// compute distance from each coord to (i,j)
			// then the least is stored like:
			// distances["i,j"] = "coord[0],coord[1]"
			// ignore any i,j that is a coord
			// if there is a tie for an i,j we ignore that too
			if distances[fmt.Sprintf("%d,%d", i, j)] == "." {
				continue
			}

			lowestDistance := 1000
			for _, coord := range coords {
				if i == coord[0] && j == coord[1] {
					// not interested in the coord itself except to count it
					distances[fmt.Sprintf("%d,%d", i, j)] = fmt.Sprintf("%d,%d", coord[0], coord[1])
					break
				}
				distance := manhattan(i, j, coord[0], coord[1])
				if distance < lowestDistance {
					// new lowest distance to save
					distances[fmt.Sprintf("%d,%d", i, j)] = fmt.Sprintf("%d,%d", coord[0], coord[1])
					lowestDistance = distance
				} else if distance == lowestDistance {
					// it's a tie, dot it - but we carry on as we might get a lower distance that is not a tie!
					distances[fmt.Sprintf("%d,%d", i, j)] = "."
				}
			}
		}
	}

	// Now we can look at reduced_coords and see which one gives the bigger "area":
	biggestArea := 0
	for _, coord := range reducedCoords {
		c := fmt.Sprintf("%d,%d", coord[0], coord[1])
		fmt.Printf("Calculate area for %v:\n", c)
		cScore := 0
		for _, b := range distances {
			if b == c {
				cScore++
			}
		}
		fmt.Printf("  score for %v is: %v\n", c, cScore)
		if cScore > biggestArea {
			biggestArea = cScore
		}
	}

	return biggestArea
}

func partTwo() int {
	// What is the size of the region containing all locations which have a total distance to all given coordinates of less than 10000?

	coords := [][]int{}
	xMin := 1000
	yMin := 1000
	xMax := 0
	yMax := 0
	for _, line := range input {
		pair := strings.Split(line, ", ")
		x, _ := strconv.Atoi(pair[0])
		y, _ := strconv.Atoi(pair[1])
		coords = append(coords, []int{x, y})
		if x < xMin {
			xMin = x
		}
		if x > xMax {
			xMax = x
		}
		if y < yMin {
			yMin = y
		}
		if y > yMax {
			yMax = y
		}
	}

	// Now this time we want to consider each i,j point
	// And work out the sum of Manhattan distances from i,j to each of the coords

	distances := map[string]int{}
	for i := xMin; i <= xMax; i++ {
		for j := yMin; j <= yMax; j++ {
			total := 0
			for _, coord := range coords {
				total += manhattan(i, j, coord[0], coord[1])
			}
			distances[fmt.Sprintf("%d,%d", i, j)] = total
		}
	}

	count := 0
	for _, t := range distances {
		if t < 10000 { // NB, use 32 when testing!
			count++
		}
	}

	return count
}

func manhattan(x1 int, y1 int, x2 int, y2 int) int {
	d1 := x1 - x2
	if d1 < 0 {
		d1 = -d1
	}
	d2 := y1 - y2
	if d2 < 0 {
		d2 = -d2
	}
	return d1 + d2
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
