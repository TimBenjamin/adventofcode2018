package day_8

import (
	"adventofcode2018/util"
	"strconv"
)

var input []int

type Node struct {
	num_child_nodes      int
	num_metadata_entries int
	child_nodes          []Node
	metadata             []int
}

// parse the input into an easier tree structure
func parse(input []int) ([]Node, []int) {
	nodes := []Node{}
	node := Node{}
	node.num_child_nodes = input[0]
	node.num_metadata_entries = input[1]

	if node.num_child_nodes == 0 {
		// 0 3 10 11 12 1 1 0 1 99 2
		node.metadata = input[2 : 2+node.num_metadata_entries]
		node.child_nodes = []Node{}
		nodes = append(nodes, node)
		remains := input[2+node.num_metadata_entries:]
		return nodes, remains
	} else {
		// 2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2
		// 1 1 0 1 99 2
		node.child_nodes = []Node{}
		remains := input[2:]
		var child_nodes []Node
		for i := 0; i < node.num_child_nodes; i++ {
			child_nodes, remains = parse(remains)
			node.child_nodes = append(node.child_nodes, child_nodes...)
		}
		node.metadata = remains[0:node.num_metadata_entries]
		remains = remains[node.num_metadata_entries:]
		nodes = append(nodes, node)
		return nodes, remains
	}
}

// sum a slice of ints
func sum_metadata(metadata []int) (metadata_sum int) {
	for i := 0; i < len(metadata); i++ {
		metadata_sum += metadata[i]
	}
	return
}

// part 1 calculation
func calculate_metadata_sum(node Node) {
	s := sum_metadata(node.metadata)
	metadata_sum += s
	for _, c := range node.child_nodes {
		calculate_metadata_sum(c)
	}
}

// part 2 calculation
func calculate_root_value(node Node) {
	if node.num_child_nodes == 0 {
		s := sum_metadata(node.metadata)
		root_value += s
	} else {
		for _, m := range node.metadata {
			if m > node.num_child_nodes || m == 0 {
				// ignore as it is out of bounds
			} else {
				calculate_root_value(node.child_nodes[m-1])
			}
		}
	}
}

var metadata_sum int // part 1 answer
func part_1() int {
	metadata_sum = 0
	nodes, _ := parse(input)
	root_node := nodes[0]
	calculate_metadata_sum(root_node)
	return metadata_sum
}

var root_value int // part 2 answer
func part_2() int {
	root_value = 0
	nodes, _ := parse(input)
	root_node := nodes[0]
	calculate_root_value(root_node)
	return root_value
}

func Call(part string, input_file string) string {
	input = util.Parse_single_line_input_into_ints(input_file)
	var r int
	if part == "1" {
		r = part_1()
	} else {
		r = part_2()
	}
	return strconv.Itoa(r)
}
