package day_4

import (
	"bufio"
	"os"
	"strconv"
)

var input []string

func get_input(input_file string) {
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

	// The input needs sorting based on the timestamp
	// [1518-11-02 00:50] wakes up
	// Always [1:17]
}

func part_1() (out int) {
	for _, value := range input {
		ts := value[1:17]
		println(ts)
	}
	return
}

func part_2() (out int) {
	return
}

func Call(part string, input_file string) string {
	get_input(input_file)
	var r int
	if part == "1" {
		r = part_1()
	} else {
		r = part_2()
	}
	return strconv.Itoa(r)
}
