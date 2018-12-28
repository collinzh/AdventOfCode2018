package main

import (
	"fmt"
	"github.com/collinzh/AdventOfCode2018/util"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

type Step struct {
	Name          string
	Unlocks       []*Step
	Prerequisites []*Step
}

type Steps []*Step

func (s Steps) Less(a, b int) bool {
	return s[a].Name < s[b].Name
}

func (s Steps) Swap(a, b int) {
	s[a], s[b] = s[b], s[a]
}

func (s Steps) Len() int {
	return len(s)
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

	unlocked := make([]*Step, 0)
	for _, step := range steps {
		if len(step.Prerequisites) == 0 {
			unlocked = append(unlocked, step)
		}
	}

	for len(unlocked) > 0 {
		sort.Sort(Steps(unlocked))
		next := unlocked[0]
		unlocked = unlocked[1:]
		completed[next.Name] = true
		fmt.Printf("%s", next.Name)
		for _, candidate := range next.Unlocks {
			satisfied := true
			for _, p := range candidate.Prerequisites {
				if completed[p.Name] != true {
					satisfied = false
				}
			}
			if satisfied {
				unlocked = append(unlocked, candidate)
			}
		}
	}
	fmt.Println()
}

const StepSecond int = 61
const NumWorkers int = 5

func Day7P2(lines []string) {
	steps := ParseSteps(lines)
	completed := make(map[string]bool)
	progress := make(map[string]int)

	unlocked := make([]*Step, 0)
	for _, step := range steps {
		if len(step.Prerequisites) == 0 {
			progress[step.Name] = StepSecond + strings.Index(util.Letters, strings.ToLower(step.Name))
		}
	}

	solution := ""
	cycle := 0
	for len(progress) > 0 {
		cycle++
		for name, counter := range progress {
			newCounter := counter - 1
			if newCounter == 0 {
				delete(progress, name)
				completed[name] = true
				solution = solution + name

				next := steps[name]
				for _, candidate := range next.Unlocks {
					satisfied := true
					for _, p := range candidate.Prerequisites {
						if completed[p.Name] != true {
							satisfied = false
						}
					}
					if satisfied {
						unlocked = append(unlocked, candidate)
					}
				}
			} else {
				progress[name] = newCounter
			}
		}

		for i := len(progress); i < NumWorkers && len(unlocked) > 0; i++ {
			sort.Sort(Steps(unlocked))
			next := unlocked[0]
			unlocked = unlocked[1:]
			progress[next.Name] = StepSecond + strings.Index(util.Letters, strings.ToLower(next.Name))
		}
	}
	fmt.Println(solution)
	fmt.Printf("Total cycles: %d\n", cycle)
}

func main() {
	if f, err := os.Open("Day7.txt"); err == nil {
		lines := util.ScanToStringSlices(f)
		Day7P1(lines)
		Day7P2(lines)
	}
}
