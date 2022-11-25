package day_8

import (
	"adventofcode2018/util"
	"strconv"
)

var input []int

type Node struct {
	NumChildNodes      int
	NumMetadataEntries int
	ChildNodes         []Node
	Metadata           []int
}

// parse the input into an easier tree structure
func parse(input []int) ([]Node, []int) {
	nodes := []Node{}
	node := Node{}
	node.NumChildNodes = input[0]
	node.NumMetadataEntries = input[1]

	if node.NumChildNodes == 0 {
		// 0 3 10 11 12 1 1 0 1 99 2
		node.Metadata = input[2 : 2+node.NumMetadataEntries]
		node.ChildNodes = []Node{}
		nodes = append(nodes, node)
		remains := input[2+node.NumMetadataEntries:]
		return nodes, remains
	} else {
		// 2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2
		// 1 1 0 1 99 2
		node.ChildNodes = []Node{}
		remains := input[2:]
		var child_nodes []Node
		for i := 0; i < node.NumChildNodes; i++ {
			child_nodes, remains = parse(remains)
			node.ChildNodes = append(node.ChildNodes, child_nodes...)
		}
		node.Metadata = remains[0:node.NumMetadataEntries]
		remains = remains[node.NumMetadataEntries:]
		nodes = append(nodes, node)
		return nodes, remains
	}
}

// sum a slice of ints
func sumMetadata(metadata []int) (metadataSum int) {
	for i := 0; i < len(metadata); i++ {
		metadataSum += metadata[i]
	}
	return
}

// part 1 calculation
func calculateMetadataSum(node Node) {
	s := sumMetadata(node.Metadata)
	metadataSum += s
	for _, c := range node.ChildNodes {
		calculateMetadataSum(c)
	}
}

// part 2 calculation
func calculateRootValue(node Node) {
	if node.NumChildNodes == 0 {
		s := sumMetadata(node.Metadata)
		rootValue += s
	} else {
		for _, m := range node.Metadata {
			if m > node.NumChildNodes || m == 0 {
				// ignore as it is out of bounds
			} else {
				calculateRootValue(node.ChildNodes[m-1])
			}
		}
	}
}

var metadataSum int // part 1 answer
func partOne() int {
	metadataSum = 0
	nodes, _ := parse(input)
	root_node := nodes[0]
	calculateMetadataSum(root_node)
	return metadataSum
}

var rootValue int // part 2 answer
func partTwo() int {
	rootValue = 0
	nodes, _ := parse(input)
	rootNode := nodes[0]
	calculateRootValue(rootNode)
	return rootValue
}

func Call(part string, inputFile string) string {
	input = util.ParseSingleLineInputIntoInts(inputFile)
	var r int
	if part == "1" {
		r = partOne()
	} else {
		r = partTwo()
	}
	return strconv.Itoa(r)
}
