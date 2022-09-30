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

func _sort_input(input []string) (sorted_input []TimeStamp) {
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
		for i := 0; i < len(sorted_input); i++ {
			// the entry at i has a higher time than T:
			var new []TimeStamp
			if sorted_input[i].time > T.time {
				for k := 0; k < i; k++ {
					new = append(new, sorted_input[k])
				}
				new = append(new, T)
				for k := i; k < len(sorted_input); k++ {
					new = append(new, sorted_input[k])
				}
				f = true
				sorted_input = new
				break
			}
		}
		// not found a higher time, so this one is the highest time and goes at the end
		if !f {
			sorted_input = append(sorted_input, T)
		}
	}
	return
}

func part_1() (out int) {
	sorted_input := _sort_input(input)

	// Find the guard that has the most minutes asleep
	totals := map[string]int{}
	var current_guard string
	for i := 0; i < len(sorted_input); i++ {
		T := sorted_input[i]
		s := strings.Split(T.desc, " ")
		if s[0] == "Guard" {
			current_guard = s[1]
		} else if T.desc == "wakes up" {
			sleeping_time := T.time - sorted_input[i-1].time
			totals[current_guard] += sleeping_time
		}
	}
	var most_sleepy_guard string
	most_sleep_time := 0
	for guard, time := range totals {
		if time > most_sleep_time {
			most_sleepy_guard = guard
			most_sleep_time = time
		}
	}
	fmt.Printf("Most sleepy guard: %v (%v mins)\n", most_sleepy_guard, most_sleep_time)

	// What minute does that guard spend asleep the most?
	tally := map[int]int{}
	for i := 0; i < len(sorted_input); i++ {
		T := sorted_input[i]
		s := strings.Split(T.desc, " ")
		if s[0] == "Guard" {
			current_guard = s[1]
		} else if current_guard == most_sleepy_guard && T.desc == "wakes up" {
			wakes := T.minutes
			S := sorted_input[i-1]
			sleeps := S.minutes
			for j := sleeps; j < wakes; j++ {
				tally[j]++
			}
		}
	}
	var most_sleepy_minute int
	var most_sleepy_minute_amount int
	for minute, amount := range tally {
		if amount > most_sleepy_minute_amount {
			most_sleepy_minute_amount = amount
			most_sleepy_minute = minute
		}
	}
	fmt.Printf("The most sleepy minute is: %v (%v)\n", most_sleepy_minute, most_sleepy_minute_amount)

	// What is the ID of the guard you chose multiplied by the minute you chose?
	id, _ := strconv.Atoi(most_sleepy_guard[1:])
	return id * most_sleepy_minute
}

func part_2() (out int) {
	sorted_input := _sort_input(input)

	// Of all guards, which guard is most frequently asleep on the same minute?
	// so this means, go through all the minutes 00 to 59 and see which guard is asleep the most on that minute (and for how long)
	// then we look at which "how long" is the longest

	// this is a map of guards -> minutes {0: 0, 1: 0, 2: 0 ... 59: 0}
	// so that we can see, for each guard, how long they slept in each minute overall
	// then we can find which minute any guard slept the longest in
	guards := map[string]map[int]int{}

	var current_guard string
	for i := 0; i < len(sorted_input); i++ {
		T := sorted_input[i]
		s := strings.Split(T.desc, " ")
		if s[0] == "Guard" {
			current_guard = s[1]
			if len(guards[current_guard]) == 0 {
				guards[current_guard] = map[int]int{}
			}
		} else if T.desc == "wakes up" {
			wakes := T.minutes
			S := sorted_input[i-1]
			sleeps := S.minutes
			for m := sleeps; m < wakes; m++ {
				guards[current_guard][m]++
			}
		}
	}

	var most_popular_minute int
	most_popular_minute_amount := 0
	var most_popular_minute_guard string
	for guard, times := range guards {
		//fmt.Printf("Report for guard: %v\n", guard)
		for min, num := range times {
			//fmt.Printf("  Minute: %v => %v\n", min, num)
			if num > most_popular_minute_amount {
				most_popular_minute = min
				most_popular_minute_amount = num
				most_popular_minute_guard = guard
				//fmt.Printf("   -> New most popular minute!\n")
			}
		}
	}
	fmt.Printf("The most sleepy minute is: %v => %v (%v)\n", most_popular_minute, most_popular_minute_amount, most_popular_minute_guard)

	// What is the ID of the guard you chose multiplied by the minute you chose?
	id, _ := strconv.Atoi(most_popular_minute_guard[1:])
	return id * most_popular_minute
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
