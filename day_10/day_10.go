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

var points []Point
var minX int
var maxX int
var minY int
var maxY int
var numLines int

func parseInput() []Point {
	numLines = len(input)
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
	fmt.Printf("minX: %v, maxX: %v, minY: %v, maxY: %v\n", minX, maxX, minY, maxY)
}

func visualisePoints() {
	// need to make a 2d array that has maxX-minX columns and maxY-minY rows
	setMinMax()
	sizeY := maxY - minY + 1
	sizeX := maxX - minX + 1
	fmt.Printf("sizeY: %v, sizeX: %v\n", sizeY, sizeX)
	grid := make([][]string, sizeY)
	fmt.Println("makegrid")
	for i := range grid {
		grid[i] = make([]string, sizeX)
		for j := range grid[i] {
			grid[i][j] = "."
		}
	}
	fmt.Println("done initialising grid")
	// -15 is the min
	// want that to be 0
	// x += minX
	// -15 => 0
	// 4 => 19
	adjX := -minX
	adjY := -minY
	for _, point := range points {
		x := point.posX + adjX
		y := point.posY + adjY
		grid[y][x] = "#"
	}
	fmt.Println("done populating grid")
	for _, row := range grid {
		fmt.Printf("%v\n", row)
	}
	fmt.Println("-----------------------")
}

// look for when the difference between min and max Y is least?
var minDiffY int

func checkForCondensedPoints() {
	setMinMax()
	diffY := maxY - minY
	diffX := maxX - minX
	fmt.Printf("diffY: %v\n", diffY)
	fmt.Printf("diffX: %v\n", diffX)
	fmt.Println("--------")
	if diffY < minDiffY {
		minDiffY = diffY
	}
}

func partOne() int {
	minDiffY = 100000
	points := parseInput()
	for i := 0; i < numLines; i++ {
		for p := range points {
			points[p].Move()
		}
		checkForCondensedPoints()
	}
	fmt.Printf("min diff Y: %v\n", minDiffY)
	return 0
}

func partTwo() int {
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
