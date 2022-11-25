package day_7

import (
	"adventofcode2018/util"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var input []string

func getDependencies() map[string][]string {
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

	return dependencies
}

// this function removes a step from the dependencies map
// it also removes that step from each other step's dependency list
func removeDependency(oldDependencies map[string][]string, step string) map[string][]string {
	dependencies := map[string][]string{}
	for s, oldDeps := range oldDependencies {
		if s != step {
			deps := []string{}
			for _, d := range oldDeps {
				if d != step {
					deps = append(deps, d)
				}
			}
			dependencies[s] = deps
		}
	}
	return dependencies
}

func partOne() int {
	dependencies := getDependencies()
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
		dependencies = removeDependency(dependencies, found[0])
	}

	fmt.Printf("Solution: %v\n", strings.Join(solution, ""))
	return 0
}

// how long a step is going to take to complete = 60 + (A=1,...)
func calculateStepTime(step string) int {
	return 60 + (int(step[0])) - 64
}

func partTwo() int {
	dependencies := getDependencies()

	// each worker will have a current step, and a time remaining until that step is complete, and status of working
	type Worker struct {
		step          string
		timeRemaining int
		working       bool
	}

	// set up the required number of workers
	const NUM_WORKERS = 5
	workers := []Worker{}
	for i := 0; i < NUM_WORKERS; i++ {
		workers = append(workers, Worker{
			step:          "",
			timeRemaining: 0,
			working:       false,
		})
	}

	// tick a clock
	clock := 0
	for {
		// increment the clock at the end, after we check if there is any work being done

		isThereAWorkerWorking := false
		// each worker that is working -> reduce time by 1, check if they are still working
		for w := range workers {
			if workers[w].working {
				workers[w].timeRemaining -= 1
				// if the time is 0, remove that work, and remove the dependency from the main dependencies list
				if workers[w].timeRemaining == 0 {
					dependencies = removeDependency(dependencies, workers[w].step)
					workers[w].working = false
					workers[w].step = ""
				} else {
					isThereAWorkerWorking = true
				}
			}
		}

		// if there are no workers working, and no steps left in the dependencies list, we are finished!
		if !isThereAWorkerWorking && len(dependencies) == 0 {
			break
		}

		// for all steps that have no dependencies and are not already being worked on, start workers going
		// (I assume there will always be workers available to do all the available steps on this tick - no tie breaker unlike in part 1...)
		// first find steps that have no dependencies
		for step, deps := range dependencies {
			if len(deps) == 0 {
				// is the step already being worked on?
				stepInProgress := false
				for w := range workers {
					if workers[w].step == step {
						stepInProgress = true
						break
					}
				}
				if !stepInProgress {
					// assign the step to an available worker
					for w := range workers {
						if !workers[w].working {
							workers[w].step = step
							workers[w].timeRemaining = calculateStepTime(step)
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

func Call(part string, inputFile string) string {
	input = util.ParseInputIntoLines(inputFile)
	var r int
	if part == "1" {
		r = partOne()
	} else {
		r = partTwo()
	}
	return strconv.Itoa(r)
}
