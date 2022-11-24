package util

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Parse_input_into_lines(input_file string) (input []string) {
	f, err := os.Open(input_file)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var line string
	for scanner.Scan() {
		line = scanner.Text()
		input = append(input, line)
	}

	return input
}

func Parse_single_line_input(input_file string) string {
	f, err := os.Open(input_file)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	return scanner.Text()
}

func Parse_single_line_input_into_ints(input_file string) []int {
	s := Parse_single_line_input(input_file)
	words := strings.Fields(s)
	nums := []int{}
	for _, w := range words {
		i, _ := strconv.Atoi(w)
		nums = append(nums, i)
	}
	return nums
}
