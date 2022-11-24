package day_8

import (
	"adventofcode2018/util"
	"fmt"
	"strconv"
)

var input []int
var metadata_count int

func add_metadata(node []int) {
	fmt.Printf("this segment: %v\n", node)
	metadata := node[len(node)-(node[1]):]
	for i := 0; i < len(metadata); i++ {
		metadata_count += metadata[i]
	}
	fmt.Printf("metadata count after this node: %v\n", metadata_count)
	return
}

func analyse(node []int) {
	fmt.Printf("> Analysing... %v\n", node)
	num_of_subnodes := node[0]
	num_metadata := node[1]
	metadata := node[(len(node) - num_metadata):]
	fmt.Printf("num of subnodes: %v\n", num_of_subnodes)
	fmt.Printf("num of metadatas: %v\n", num_metadata)
	fmt.Printf("metadata: %v\n", metadata)

	remains := node[2:(len(node) - num_metadata - 1)]
	fmt.Printf("remains: %v\n", remains)

	if num_of_subnodes > 0 {

	} else {
		//sub_metadata := remains
	}

	analyse(remains)
}

func part_1() int {
	metadata_count = 0
	// for i := 0; i < len(input); i++ {
	// 	// first i will be a header
	// 	fmt.Printf("header: %v\n", i)
	// 	// this tells us how many nodes there are in this section
	// 	// need a function to get each node
	// 	// (that will be recursively called!)
	// 	// but for this part we can simply recursively call and simply add the metadata in each node that is found in some global variable
	// }

	//add_metadata(input)
	analyse(input)

	// maybe recurse down until we get nodes with 0 as the number of sub-nodes
	// and eliminate those from the input

	return metadata_count
}

func part_2() int {

	return 0
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
