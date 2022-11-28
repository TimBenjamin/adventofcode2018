package day_9

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	tables := []struct {
		solution   string
		input_file string
	}{
		{"32", "./test.txt"},
		{"37305", "./test2.txt"},
	}

	for _, table := range tables {
		got := Call("1", table.input_file)
		if got != table.solution {
			t.Fatalf("Test failed - wanted %v, got %v for input file %v\n", table.solution, got, table.input_file)
		}
	}
}
