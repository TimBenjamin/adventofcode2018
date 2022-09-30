package day_6

import (
	"adventofcode2018/util"
	"fmt"
	"strconv"
	"strings"
)

var input []string

func part_1() int {

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
	x_min := 1000
	y_min := 1000
	x_max := 0
	y_max := 0
	for _, line := range input {
		pair := strings.Split(line, ", ")
		x, _ := strconv.Atoi(pair[0])
		y, _ := strconv.Atoi(pair[1])
		coords = append(coords, []int{x, y})
		if x < x_min {
			x_min = x
		}
		if x > x_max {
			x_max = x
		}
		if y < y_min {
			y_min = y
		}
		if y > y_max {
			y_max = y
		}
	}

	reduced_coords := [][]int{}
	for _, coord := range coords {
		if coord[0] > x_min && coord[0] < x_max && coord[1] > y_min && coord[1] < y_max {
			reduced_coords = append(reduced_coords, coord)
		}
	}

	// The reduced_coords are the ones that do not lead to an infinite area
	for _, coord := range reduced_coords {
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
	for i := x_min; i <= x_max; i++ {
		for j := y_min; j <= y_max; j++ {
			// compute distance from each coord to (i,j)
			// then the least is stored like:
			// distances["i,j"] = "coord[0],coord[1]"
			// ignore any i,j that is a coord
			// if there is a tie for an i,j we ignore that too
			if distances[fmt.Sprintf("%d,%d", i, j)] == "." {
				continue
			}

			lowest_distance := 1000
			for _, coord := range coords {
				if i == coord[0] && j == coord[1] {
					// not interested in the coord itself except to count it
					distances[fmt.Sprintf("%d,%d", i, j)] = fmt.Sprintf("%d,%d", coord[0], coord[1])
					break
				}
				distance := _manhattan(i, j, coord[0], coord[1])
				if distance < lowest_distance {
					// new lowest distance to save
					distances[fmt.Sprintf("%d,%d", i, j)] = fmt.Sprintf("%d,%d", coord[0], coord[1])
					lowest_distance = distance
				} else if distance == lowest_distance {
					// it's a tie, dot it - but we carry on as we might get a lower distance that is not a tie!
					distances[fmt.Sprintf("%d,%d", i, j)] = "."
				}
			}
		}
	}

	// Now we can look at reduced_coords and see which one gives the bigger "area":
	biggest_area := 0
	for _, coord := range reduced_coords {
		c := fmt.Sprintf("%d,%d", coord[0], coord[1])
		fmt.Printf("Calculate area for %v:\n", c)
		c_score := 0
		for _, b := range distances {
			if b == c {
				c_score++
			}
		}
		fmt.Printf("  score for %v is: %v\n", c, c_score)
		if c_score > biggest_area {
			biggest_area = c_score
		}
	}

	return biggest_area
}

func part_2() int {
	// What is the size of the region containing all locations which have a total distance to all given coordinates of less than 10000?

	coords := [][]int{}
	x_min := 1000
	y_min := 1000
	x_max := 0
	y_max := 0
	for _, line := range input {
		pair := strings.Split(line, ", ")
		x, _ := strconv.Atoi(pair[0])
		y, _ := strconv.Atoi(pair[1])
		coords = append(coords, []int{x, y})
		if x < x_min {
			x_min = x
		}
		if x > x_max {
			x_max = x
		}
		if y < y_min {
			y_min = y
		}
		if y > y_max {
			y_max = y
		}
	}

	// Now this time we want to consider each i,j point
	// And work out the sum of Manhattan distances from i,j to each of the coords

	distances := map[string]int{}
	for i := x_min; i <= x_max; i++ {
		for j := y_min; j <= y_max; j++ {
			total := 0
			for _, coord := range coords {
				total += _manhattan(i, j, coord[0], coord[1])
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

func _manhattan(x1 int, y1 int, x2 int, y2 int) int {
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
