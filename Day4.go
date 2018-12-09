package main

import (
	"./util"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Event struct {
	timestamp time.Time
	content   string
}

type Events []*Event

func (e Events) Len() int {
	return len(e)
}

func (e Events) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (e Events) Less(i, j int) bool {
	return e[i].timestamp.Before(e[j].timestamp)
}

var lineRegex = regexp.MustCompile("^\\[(.*)] (.*)$")

//var nonNumeric = regexp.MustCompile("[^0-9]")
var guardRegex = regexp.MustCompile("Guard #([0-9]+) begins shift")

func ParseEvents(lines []string) []*Event {
	events := make([]*Event, 0)
	for _, line := range lines {
		if !lineRegex.MatchString(line) {
			fmt.Printf("Cannot parse line %s\n", line)
			continue
		}

		parts := lineRegex.FindStringSubmatch(line)

		ts, err := time.Parse("2006-01-02 15:04", parts[1])
		if err != nil {
			panic(err)
		}
		events = append(events, &Event{timestamp: ts, content: parts[2]})
	}
	sort.Sort(Events(events))
	return events
}

func AnalyzeEvents(lines []string) (map[int]int, map[int]map[int]int) {
	events := ParseEvents(lines)
	currentGuard := -1
	sleepAt := time.Now()
	sleepTimes := make(map[int]int)
	sleepMinutes := make(map[int]map[int]int)

	for _, event := range events {

		if guardRegex.MatchString(event.content) {
			g, err := strconv.Atoi(guardRegex.FindStringSubmatch(event.content)[1])
			if err != nil {
				panic(err)
			}
			currentGuard = g

		} else if strings.Compare("wakes up", event.content) == 0 {
			sleepTimes[currentGuard] += int(event.timestamp.Sub(sleepAt).Minutes())
			if sleepMinutes[currentGuard] == nil {
				sleepMinutes[currentGuard] = make(map[int]int)
			}
			for i := sleepAt.Minute(); i < event.timestamp.Minute(); i++ {
				sleepMinutes[currentGuard][i]++
			}

		} else if strings.Compare("falls asleep", event.content) == 0 {
			sleepAt = event.timestamp

		} else {
			fmt.Printf("Unrecognized event %s\n", event.content)

		}
	}
	return sleepTimes, sleepMinutes
}

func Day4P1(lines []string) {
	sleepTimes, sleepMinutes := AnalyzeEvents(lines)

	counter := 0
	longestId := 0
	for id, d := range sleepTimes {
		if d > counter {
			longestId = id
			counter = d
		}
	}
	counter = 0
	minute := 0
	for min, c := range sleepMinutes[longestId] {
		if c > counter {
			minute = min
			counter = c
		}
	}
	fmt.Printf("Guard: %d x minute: %d = %d\n", longestId, counter, longestId*minute)
}

func Day4P2(lines []string) {
	_, sleepMinutes := AnalyzeEvents(lines)
	theGuard := -1
	theMinute := -1
	theCount := -1
	for id, minutes := range sleepMinutes {
		for minute, count := range minutes {
			if count > theCount {
				theCount = count
				theMinute = minute
				theGuard = id
			}
		}
	}
	fmt.Printf("Guard %d x minute %d = %d", theGuard, theMinute, theGuard*theMinute)
}

func main() {
	file, _ := os.Open("Day4.txt")
	lines := util.ScanToStringSlices(file)
	Day4P1(lines)
	Day4P2(lines)
}
