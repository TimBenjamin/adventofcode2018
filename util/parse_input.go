package util

import (
	"bufio"
	"os"
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
