package main

import (
	"./util"
	"fmt"
	"os"
	"strings"
)

func Day2P1(file *os.File) {
	lines := util.ScanToStringSlices(file)
	countTwo := 0
	countThree := 0

	for _, line := range lines {
		hasTwo := false
		hasThree := false
		letters := make(map[rune]int)
		strings.Map(func(r rune) rune {
			letters[r]++
			return r
		}, line)

		for _, c := range letters {
			if c == 2 {
				hasTwo = true
			} else if c == 3 {
				hasThree = true
			}
		}

		if hasTwo {
			countTwo++
		}
		if hasThree {
			countThree++
		}
	}

	fmt.Printf("Check sum %d\n", countTwo*countThree)
}

func Day2P2(file *os.File) {
	lines := util.ScanToStringSlices(file)

	for idx, line := range lines {
		for j := idx + 1; j < len(lines); j++ {
			toCompare := lines[j]

			if DiffByOne(line, toCompare) {
				return
			}
		}
	}
}

func DiffByOne(a, b string) bool {
	runesA := []rune(a)
	runesB := []rune(b)

	if len(runesA) != len(runesB) {
		return false
	}

	diffs := 0
	lastDiff := 0
	for i := 0; i < len(runesA); i++ {
		if runesA[i] != runesB[i] {
			diffs++
			if diffs > 1 {
				return false
			}
			lastDiff = i
		}
	}

	if diffs == 1 {
		fmt.Printf("%s%s\n", string(runesA[0:lastDiff]), string(runesA[lastDiff+1:]))
	}

	return diffs == 1
}

func main() {
	if file, err := os.Open("Day2.txt"); err == nil {
		defer file.Close()
		Day2P1(file)
	}

	if file, err := os.Open("Day2.txt"); err == nil {
		defer file.Close()
		Day2P2(file)
	}
}
