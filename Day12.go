package main

import (
	"fmt"
	"github.com/collinzh/AdventOfCode2018/util"
	"os"
	"regexp"
)

func ParsePots(line string) map[int]bool {
	pots := make(map[int]bool)
	reg := regexp.MustCompile("[^.#]")
	line = reg.ReplaceAllString(line, "")
	for idx, r := range line {
		pots[idx] = r == '#'
	}
	return pots
}

func ParseRules(lines []string) map[int]bool {
	rules := make(map[int]bool)
	for _, line := range lines {
		rule := 0
		for i := 0; i <= 4; i++ {
			if line[i] == '#' {
				rule++
			}
			rule *= 10
		}
		rules[rule] = line[len(line)-1] == '#'
	}
	return rules
}

func Day12P1(lines []string, generations int64) {
	pots := ParsePots(lines[0])
	rules := ParseRules(lines[2:])
	count := 0

	for _, status := range pots {
		if status {
			count++
		}
	}

	prevGen := int64(0)

	for generation := int64(1); generation <= generations; generation++ {
		newPots := make(map[int]bool)
		for number := -10; number <= len(pots)+20; number++ {
			rule := 0
			for id := number - 2; id <= number+2; id++ {
				if pots[id] {
					rule++
				}
				rule *= 10
			}
			newPots[number] = rules[rule]
		}
		pots = newPots
		for _, status := range pots {
			if status {
				count++
			}
		}

		sum := int64(0)
		for id, status := range pots {
			if status {
				sum += int64(id)
			}
		}
		fmt.Printf("Result of gen %d is %d, delta %d\n", generation, sum, sum-prevGen)
		prevGen = sum
	}

	fmt.Printf("Number of plants: %d\n", count)
	fmt.Printf("Sum is %d\n", prevGen)
}

func Day12P2(gen int64) {
	// The solution is specifically tailored for my puzzle input.
	// Observe your output pattern to conceive your formula
	result := 11697 + (gen-195)*53
	fmt.Printf("Sum is %d\n", result)
}

func main() {
	if f, err := os.Open("Day12.txt"); err == nil {
		lines := util.ScanToStringSlices(f)
		Day12P1(lines, 20)
		Day12P2(50000000000)
	} else {
		panic(err)
	}
}
