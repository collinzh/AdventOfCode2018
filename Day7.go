package main

import (
	"container/list"
	"fmt"
	"github.com/collinzh/AdventOfCode2018/util"
	"log"
	"os"
	"regexp"
	"sort"
)

type Step struct {
	Name          string
	Unlocks       []*Step
	Prerequisites []*Step
}

var (
	LineRegex = regexp.MustCompile("^Step ([A-Z]) must be finished before step ([A-Z]) can begin")
)

func ParseSteps(lines []string) map[string]*Step {
	steps := make(map[string]*Step)
	for _, line := range lines {
		if !LineRegex.MatchString(line) {
			log.Fatalf("Cannot parse line %s\n", line)
			continue
		}

		parts := LineRegex.FindStringSubmatch(line)
		name := parts[1]
		prereq := parts[2]

		if steps[name] == nil {
			steps[name] = &Step{Name: name, Unlocks: make([]*Step, 0), Prerequisites: make([]*Step, 0)}
		}
		if steps[prereq] == nil {
			steps[prereq] = &Step{Name: prereq, Unlocks: make([]*Step, 0), Prerequisites: make([]*Step, 0)}
		}
		steps[name].Unlocks = append(steps[name].Unlocks, steps[prereq])
		steps[prereq].Prerequisites = append(steps[prereq].Prerequisites, steps[name])
	}
	return steps
}

func Day7P1(lines []string) {
	steps := ParseSteps(lines)
	completed := make(map[string]bool)

	toCheck := make([]*Step, 0)
	for _, step := range steps {
		if len(step.Prerequisites) == 0 {
			toCheck = append(toCheck, step)
		}
	}

	for len(toCheck) > 0 {

		completable := make([]string, 0)
		for _, next := range toCheck {
			if completed[next.Name] {
				continue
			}
			unlocked := true
			for _, preReq := range next.Prerequisites {
				unlocked = unlocked && completed[preReq.Name]
			}
			if unlocked {
				completable = append(completable, next.Name)
				completed[next.Name] = true
			}
		}

		sort.Strings(completable)
		nextRound := make([]*Step, 0)
		for _, name := range completable {
			fmt.Print(name)
			nextRound = append(nextRound, steps[name].Unlocks...)
		}
		toCheck = nextRound
	}
}

func AddAll(lst *list.List, elements []*Step) {
	for _, elem := range elements {
		lst.PushBack(elem)
	}
}

func main() {
	if f, err := os.Open("Day7_test.txt"); err == nil {
		lines := util.ScanToStringSlices(f)
		Day7P1(lines)
	}
}
