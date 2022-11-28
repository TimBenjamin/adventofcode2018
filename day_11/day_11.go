package day_11

import (
	"adventofcode2018/util"
	"fmt"
	"strconv"
)

func makeGrid(gridSize int) [][]int {
	grid := make([][]int, gridSize)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]int, gridSize)
	}
	return grid
}

func getPowerLevel(x int, y int) int {
	rackId := x + 10
	powerLevel := rackId * y
	powerLevel += gridSerialNumber
	powerLevel = powerLevel * rackId
	powerLevel = getHundreds(powerLevel)
	powerLevel -= 5
	return powerLevel
}

func getHundreds(powerLevel int) int {
	// get the hundreds figure...
	powerLevelString := strconv.Itoa(powerLevel)
	if len(powerLevelString) >= 3 {
		hundreds := powerLevelString[len(powerLevelString)-3 : len(powerLevelString)-2]
		hundredsInt, _ := strconv.Atoi(hundreds)
		return hundredsInt
	} else {
		return 0
	}
}

func partOne() int {
	gridSize := 300
	grid := makeGrid(gridSize)
	for y := 0; y < gridSize; y++ {
		for x := 0; x < gridSize; x++ {
			grid[y][x] = getPowerLevel(x+1, y+1)
		}
	}

	// Your goal is to find the 3x3 square which has the largest total power.
	// we can fit 300 - 2 x 300 - 2 squares in the grid
	maxSum := 0
	maxX := 0
	maxY := 0
	for y := 0; y < gridSize-2; y++ {
		for x := 0; x < gridSize-2; x++ {
			sum := grid[y][x] + grid[y+1][x] + grid[y+2][x] + grid[y][x+1] + grid[y+1][x+1] + grid[y+2][x+1] + grid[y][x+2] + grid[y+1][x+2] + grid[y+2][x+2]
			if sum > maxSum {
				maxSum = sum
				maxX = x
				maxY = y
			}
		}
	}
	//fmt.Printf("Found max power level of %v at: %v,%v\n", maxSum, maxX+1, maxY+1)
	fmt.Printf("Solution: %v,%v\n", maxX+1, maxY+1)
	return 0
}

func partTwo() int {
	// repeatedly try grids to find the maximum 3x3 in any size grid
	gridSize := 300
	grid := makeGrid(gridSize)
	for y := 0; y < gridSize; y++ {
		for x := 0; x < gridSize; x++ {
			grid[y][x] = getPowerLevel(x+1, y+1)
		}
	}
	maxMaxX := 0
	maxMaxY := 0
	maxMaxSum := 0
	maxSquareSize := 0
	for squareSize := 2; squareSize <= gridSize; squareSize++ {
		maxSum := 0
		maxX := 0
		maxY := 0
		fmt.Printf("... trying square size %v\n", squareSize)
		for y := 0; y < gridSize-squareSize-1; y++ {
			for x := 0; x < gridSize-squareSize-1; x++ {
				sum := 0
				for i := 0; i < squareSize; i++ {
					for j := 0; j < squareSize; j++ {
						sum += grid[y+i][x+j]
					}
				}
				if sum > maxSum {
					maxSum = sum
					maxX = x
					maxY = y
				}
			}
		}
		//fmt.Printf("Found max power level of %v at: %v,%v in square size %vx%v\n", maxSum, maxX+1, maxY+1, squareSize, squareSize)
		if maxSum > maxMaxSum {
			maxMaxSum = maxSum
			maxSquareSize = squareSize
			maxMaxX = maxX
			maxMaxY = maxY
		}
		if maxSum == 0 {
			//fmt.Println("- ending as power has dropped to 0")
			break
		}
	}
	//fmt.Printf("Found max power level of %v at: %v,%v with square size: %vx%v\n", maxMaxSum, maxMaxX+1, maxMaxY+1, maxSquareSize, maxSquareSize)
	fmt.Printf("Solution: %v,%v,%v\n", maxMaxX+1, maxMaxY+1, maxSquareSize)
	return 0
}

var gridSerialNumber int

func Call(part string, inputFile string) string {
	gridSerialNumber, _ = strconv.Atoi(util.ParseSingleLineInput(inputFile))
	fmt.Printf("gridSerialNumber: %v\n", gridSerialNumber)
	var r int
	if part == "1" {
		r = partOne()
	} else {
		r = partTwo()
	}
	return strconv.Itoa(r)
}
