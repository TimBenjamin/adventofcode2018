package day_7

import (
	"adventofcode2018/util"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var input []string

func get_dependencies() map[string][]string {
	dependencies := map[string][]string{}
	for _, instruction := range input {
		step := string(instruction[36])
		dependency := string(instruction[5])

		// initialise the list of dependencies for this step if it doesn't already exist
		if len(dependencies[step]) == 0 {
			dependencies[step] = []string{dependency}
		} else {
			dependencies[step] = append(dependencies[step], dependency)
		}
		// create a step->dependency list for the dependency too
		// (i.e. the dependency may have dependencies, and I need to know even if it has 0 depedencies)
		if len(dependencies[dependency]) == 0 {
			dependencies[dependency] = []string{}
		}
	}

	// Check:
	// for step, requires := range dependencies {
	// 	fmt.Printf("step %v requires...\n", step)
	// 	for _, r := range requires {
	// 		fmt.Printf("  %v\n", r)
	// 	}
	// }

	return dependencies
}

// this function removes a step from the dependencies map
// it also removes that step from each other step's dependency list
func remove_dependency(old_dependencies map[string][]string, step string) map[string][]string {
	dependencies := map[string][]string{}
	for s, old_deps := range old_dependencies {
		if s != step {
			deps := []string{}
			for _, d := range old_deps {
				if d != step {
					deps = append(deps, d)
				}
			}
			dependencies[s] = deps
		}
	}
	return dependencies
}

func part_1() int {
	dependencies := get_dependencies()
	solution := []string{}

	for {
		// find steps that have no dependencies
		found := []string{}
		for step, deps := range dependencies {
			if len(deps) == 0 {
				found = append(found, step)
			}
		}

		// if there isn't any left, break
		if len(found) == 0 {
			break
		}

		// resolve ties by picking the earliest in alphabetical order
		sort.Strings(found)

		// add the step to the answer sequence
		solution = append(solution, found[0])

		// remove the step from the dependencies list
		dependencies = remove_dependency(dependencies, found[0])
	}

	fmt.Printf("Solution: %v\n", strings.Join(solution, ""))
	return 0
}

// how long a step is going to take to complete = 60 + (A=1,...)
func calculate_step_time(step string) int {
	return 60 + (int(step[0])) - 64
}

func part_2() int {
	dependencies := get_dependencies()

	// each worker will have a current step, and a time remaining until that step is complete, and status of working
	type Worker struct {
		step           string
		time_remaining int
		working        bool
	}

	// set up the required number of workers
	const NUM_WORKERS = 5
	workers := []Worker{}
	for i := 0; i < NUM_WORKERS; i++ {
		workers = append(workers, Worker{
			step:           "",
			time_remaining: 0,
			working:        false,
		})
	}

	// tick a clock
	clock := 0
	for {
		// increment the clock at the end, after we check if there is any work being done

		is_there_a_worker_working := false
		// each worker that is working -> reduce time by 1, check if they are still working
		for w := range workers {
			if workers[w].working {
				workers[w].time_remaining -= 1
				// if the time is 0, remove that work, and remove the dependency from the main dependencies list
				if workers[w].time_remaining == 0 {
					dependencies = remove_dependency(dependencies, workers[w].step)
					workers[w].working = false
					workers[w].step = ""
				} else {
					is_there_a_worker_working = true
				}
			}
		}

		// if there are no workers working, and no steps left in the dependencies list, we are finished!
		if !is_there_a_worker_working && len(dependencies) == 0 {
			break
		}

		// for all steps that have no dependencies and are not already being worked on, start workers going
		// (I assume there will always be workers available to do all the available steps on this tick - no tie breaker unlike in part 1...)
		// first find steps that have no dependencies
		for step, deps := range dependencies {
			if len(deps) == 0 {
				// is the step already being worked on?
				step_in_progress := false
				for w := range workers {
					if workers[w].step == step {
						step_in_progress = true
						break
					}
				}
				if !step_in_progress {
					// assign the step to an available worker
					for w := range workers {
						if !workers[w].working {
							workers[w].step = step
							workers[w].time_remaining = calculate_step_time(step)
							workers[w].working = true
							break
						}
					}
				}
			}
		}

		// clock ticks:
		clock++
	}

	// the solution is how many ticks of the clock there were:
	return clock
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
