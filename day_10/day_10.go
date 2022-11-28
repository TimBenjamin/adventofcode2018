package day_10

import (
	"adventofcode2018/util"
	"fmt"
	"regexp"
	"strconv"
)

type Point struct {
	posX      int
	posY      int
	velocityX int
	velocityY int
}

func (point *Point) Move() {
	point.posX += point.velocityX
	point.posY += point.velocityY
}

func (point *Point) MoveBack() {
	point.posX -= point.velocityX
	point.posY -= point.velocityY
}

var points []Point
var minX int
var maxX int
var minY int
var maxY int

// To solve this I will look for when the difference between min and max Y is least
var minDiffY int

func parseInput() []Point {
	for _, line := range input {
		point := Point{}
		// position=< 9,  1> velocity=< 0,  2>
		re := regexp.MustCompile(`position=\<\s*(\-?\d+),\s+(\-?\d+)\>\s+velocity=\<\s*(\-?\d+),\s+(\-?\d+)`)
		matches := re.FindStringSubmatch(line)
		point.posX, _ = strconv.Atoi(matches[1])
		point.posY, _ = strconv.Atoi(matches[2])
		point.velocityX, _ = strconv.Atoi(matches[3])
		point.velocityY, _ = strconv.Atoi(matches[4])
		points = append(points, point)
	}
	setMinMax()
	minDiffY = maxY - minY
	return points
}

func setMinMax() {
	minX = 0
	maxX = 0
	minY = 0
	maxY = 0
	for _, point := range points {
		if point.posX < minX {
			minX = point.posX
		}
		if point.posX > maxX {
			maxX = point.posX
		}
		if point.posY < minY {
			minY = point.posY
		}
		if point.posY > maxY {
			maxY = point.posY
		}
	}
}

func visualisePoints() {
	// need to make a 2d array that has maxX-minX columns and maxY-minY rows
	setMinMax()
	sizeY := maxY - minY + 1
	sizeX := maxX - minX + 1
	grid := make([][]string, sizeY)
	for i := range grid {
		grid[i] = make([]string, sizeX)
		for j := range grid[i] {
			grid[i][j] = "."
		}
	}
	adjX := -minX
	adjY := -minY
	for _, point := range points {
		x := point.posX + adjX
		y := point.posY + adjY
		grid[y][x] = "#"
	}
	for _, row := range grid {
		fmt.Printf("%v\n", row)
	}
}

func checkForCondensedPoints() {
	setMinMax()
	diffY := maxY - minY
	if diffY < minDiffY {
		minDiffY = diffY
	}
}

func partOne() int {
	points := parseInput()
	// need to keep iterating until minDiffY hits a minimum and starts increasing again.
	for {
		for p := range points {
			points[p].Move()
		}
		checkForCondensedPoints()
		if (maxY - minY) > minDiffY {
			for p := range points {
				points[p].MoveBack()
			}
			visualisePoints()
			break
		}
	}
	fmt.Printf("min diff Y: %v\n", minDiffY)
	return 0
}

func partTwo() int {
	count := 0
	minDiffY = 100000
	points := parseInput()
	for i := 0; i < 250000; i++ {
		for p := range points {
			points[p].Move()
		}
		checkForCondensedPoints()
		if (maxY - minY) > minDiffY {
			return count
		}
		count++
	}
	return 0
}

var input []string

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
