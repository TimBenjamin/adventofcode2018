package day_4

import (
	"adventofcode2018/util"
	"fmt"
	"strconv"
	"strings"
)

var input []string

type TimeStamp struct {
	original string
	stamp    string
	year     int
	month    int
	day      int
	hours    int
	minutes  int
	time     int
	desc     string
}

func sortInput(input []string) (sortedInput []TimeStamp) {
	for _, value := range input {
		ts := value[1:17]
		desc := value[19:]
		// 1518-11-01 00:05
		// I'll turn it into the number of minutes since 1518-01-01 00:00, assuming each month has 31 days.
		dt := strings.Split(ts, " ")
		ymd := strings.Split(dt[0], "-")
		hm := strings.Split(dt[1], ":")
		y, _ := strconv.Atoi(ymd[0])
		mo, _ := strconv.Atoi(ymd[1])
		d, _ := strconv.Atoi(ymd[2])
		h, _ := strconv.Atoi(hm[0])
		min, _ := strconv.Atoi(hm[1])
		t := min + (h * 60) + ((31*mo + d) * 60 * 24)
		T := TimeStamp{
			original: value,
			stamp:    ts,
			year:     y,
			month:    mo,
			day:      d,
			hours:    h,
			minutes:  min,
			time:     t,
			desc:     desc,
		}
		f := false
		for i := 0; i < len(sortedInput); i++ {
			// the entry at i has a higher time than T:
			var new []TimeStamp
			if sortedInput[i].time > T.time {
				for k := 0; k < i; k++ {
					new = append(new, sortedInput[k])
				}
				new = append(new, T)
				for k := i; k < len(sortedInput); k++ {
					new = append(new, sortedInput[k])
				}
				f = true
				sortedInput = new
				break
			}
		}
		// not found a higher time, so this one is the highest time and goes at the end
		if !f {
			sortedInput = append(sortedInput, T)
		}
	}
	return
}

func partOne() (out int) {
	sortedInput := sortInput(input)

	// Find the guard that has the most minutes asleep
	totals := map[string]int{}
	var currentGuard string
	for i := 0; i < len(sortedInput); i++ {
		T := sortedInput[i]
		s := strings.Split(T.desc, " ")
		if s[0] == "Guard" {
			currentGuard = s[1]
		} else if T.desc == "wakes up" {
			sleepingTime := T.time - sortedInput[i-1].time
			totals[currentGuard] += sleepingTime
		}
	}
	var mostSleepyGuard string
	mostSleepTime := 0
	for guard, time := range totals {
		if time > mostSleepTime {
			mostSleepyGuard = guard
			mostSleepTime = time
		}
	}
	fmt.Printf("Most sleepy guard: %v (%v mins)\n", mostSleepyGuard, mostSleepTime)

	// What minute does that guard spend asleep the most?
	tally := map[int]int{}
	for i := 0; i < len(sortedInput); i++ {
		T := sortedInput[i]
		s := strings.Split(T.desc, " ")
		if s[0] == "Guard" {
			currentGuard = s[1]
		} else if currentGuard == mostSleepyGuard && T.desc == "wakes up" {
			wakes := T.minutes
			S := sortedInput[i-1]
			sleeps := S.minutes
			for j := sleeps; j < wakes; j++ {
				tally[j]++
			}
		}
	}
	var mostSleepyMinute int
	var mostSleepyMinuteAmount int
	for minute, amount := range tally {
		if amount > mostSleepyMinuteAmount {
			mostSleepyMinuteAmount = amount
			mostSleepyMinute = minute
		}
	}
	fmt.Printf("The most sleepy minute is: %v (%v)\n", mostSleepyMinute, mostSleepyMinuteAmount)

	// What is the ID of the guard you chose multiplied by the minute you chose?
	id, _ := strconv.Atoi(mostSleepyGuard[1:])
	return id * mostSleepyMinute
}

func partTwo() (out int) {
	sortedInput := sortInput(input)

	// Of all guards, which guard is most frequently asleep on the same minute?
	// so this means, go through all the minutes 00 to 59 and see which guard is asleep the most on that minute (and for how long)
	// then we look at which "how long" is the longest

	// this is a map of guards -> minutes {0: 0, 1: 0, 2: 0 ... 59: 0}
	// so that we can see, for each guard, how long they slept in each minute overall
	// then we can find which minute any guard slept the longest in
	guards := map[string]map[int]int{}

	var currentGuard string
	for i := 0; i < len(sortedInput); i++ {
		T := sortedInput[i]
		s := strings.Split(T.desc, " ")
		if s[0] == "Guard" {
			currentGuard = s[1]
			if len(guards[currentGuard]) == 0 {
				guards[currentGuard] = map[int]int{}
			}
		} else if T.desc == "wakes up" {
			wakes := T.minutes
			S := sortedInput[i-1]
			sleeps := S.minutes
			for m := sleeps; m < wakes; m++ {
				guards[currentGuard][m]++
			}
		}
	}

	var mostPopularMinute int
	mostPopularMinuteAmount := 0
	var mostPopularMinuteGuard string
	for guard, times := range guards {
		//fmt.Printf("Report for guard: %v\n", guard)
		for min, num := range times {
			//fmt.Printf("  Minute: %v => %v\n", min, num)
			if num > mostPopularMinuteAmount {
				mostPopularMinute = min
				mostPopularMinuteAmount = num
				mostPopularMinuteGuard = guard
				//fmt.Printf("   -> New most popular minute!\n")
			}
		}
	}
	fmt.Printf("The most sleepy minute is: %v => %v (%v)\n", mostPopularMinute, mostPopularMinuteAmount, mostPopularMinuteGuard)

	// What is the ID of the guard you chose multiplied by the minute you chose?
	id, _ := strconv.Atoi(mostPopularMinuteGuard[1:])
	return id * mostPopularMinute
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
