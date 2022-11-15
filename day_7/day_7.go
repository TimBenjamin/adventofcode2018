package day_7

import (
	"adventofcode2018/util"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var input []string

func part_1() int {

	order := []string{} // This is our answer

	for {
		dependencies := map[string][]string{}
		for _, instruction := range input {
			step := string(instruction[36])
			dependency := string(instruction[5])

			// initialise the list of dependencies for this step if it doesn't already exist
			if len(dependencies[step]) == 0 {
				dependencies[step] = []string{}
			}
			if len(dependencies[dependency]) == 0 {
				dependencies[dependency] = []string{}
			}

			// is "dependency" already done? i.e. does "step" still require "dependency"
			done := false
			for _, o := range order {
				if o == dependency {
					done = true
					break
				}
			}
			if !done {
				dependencies[step] = append(dependencies[step], dependency)
			}

		}

		// for step, requires := range dependencies {
		// 	fmt.Printf("step %v requires...\n", step)
		// 	for _, r := range requires {
		// 		fmt.Printf("  %v\n", r)
		// 	}
		// }

		finished := true
		available := []string{} // a list of steps that have no dependencies
		for step, requires := range dependencies {
			if len(requires) == 0 {
				// check it isn't already done
				done := false
				for _, o := range order {
					if o == step {
						done = true
						break
					}
				}
				if !done {
					available = append(available, step)
				}
			}
		}

		if len(available) > 0 {
			sort.Strings(available)
			order = append(order, available[0]) // NB, we only add one of any multiple-available steps (per rubric)
			finished = false
		}

		if finished {
			break
		}

	}

	fmt.Printf("Solution: %v\n", strings.Join(order, ""))

	return 0
}

func part_2() int {
	order := []string{} // This is our answer

	for {
		dependencies := map[string][]string{}
		for _, instruction := range input {
			step := string(instruction[36])
			dependency := string(instruction[5])

			// initialise the list of dependencies for this step if it doesn't already exist
			if len(dependencies[step]) == 0 {
				dependencies[step] = []string{}
			}
			if len(dependencies[dependency]) == 0 {
				dependencies[dependency] = []string{}
			}

			// is "dependency" already done? i.e. does "step" still require "dependency"
			done := false
			for _, o := range order {
				if o == dependency {
					done = true
					break
				}
			}
			if !done {
				dependencies[step] = append(dependencies[step], dependency)
			}

		}

		finished := true
		available := []string{} // a list of steps that have no dependencies
		for step, requires := range dependencies {
			if len(requires) == 0 {
				// check it isn't already done
				done := false
				for _, o := range order {
					if o == step {
						done = true
						break
					}
				}
				if !done {
					available = append(available, step)
				}
			}
		}

		// This is different - when is there more than 1 thing available?
		fmt.Printf("Available:\n")
		for _, a := range available {
			fmt.Printf("  %v\n", a)
		}
		println("----")
		if len(available) > 0 {
			sort.Strings(available)
			order = append(order, available[0]) // NB, we only add one of any multiple-available steps (per rubric)
			finished = false
		}

		if finished {
			break
		}

	}

	fmt.Printf("Solution: %v\n", strings.Join(order, ""))

	return 0
}

func Call(part string, input_file string) string {
	input = util.Parse_input_into_lines(input_file)
	var r int
	if part == "1" {
		r = part_1()
	} else {
		r = part_2()
	}
	return strconv.Itoa(r)
}
