package main

import (
	"./util"
	"fmt"
	"math"
	"os"
	"unicode"
)

func DifferentPolarity(a, b rune) bool {
	if unicode.ToLower(a) == unicode.ToLower(b) {
		aLower := unicode.IsLower(a)
		bLower := unicode.IsLower(b)
		return aLower != bLower
	}
	return false
}

func Reduction(chars []rune) []rune {
	results := make([]rune, 0)
	for i := 0; i < len(chars); i++ {
		if i+1 < len(chars) {
			if DifferentPolarity(chars[i], chars[i+1]) {
				i++
				continue
			}
		}
		results = append(results, chars[i])
	}
	return results
}

func Filter(chars []rune, except rune) []rune {
	results := make([]rune, 0)
	except = unicode.ToLower(except)
	for _, r := range chars {
		if unicode.ToLower(r) != except {
			results = append(results, r)
		}
	}
	return results
}

func Day5P1(chars []rune) {
	before := chars
	for true {
		after := Reduction(before)
		if len(before) == len(after) {
			fmt.Printf("Final result: %d\n", len(after))
			return
		}
		before = after
	}
}

func Day5P2(chars []rune) {
	alphabet := []rune(util.Letters)
	theRune := ' '
	theLength := math.MaxInt32

	for _, a := range alphabet {
		filtered := Filter(chars, a)
		before := filtered
		for true {
			after := Reduction(before)
			if len(before) == len(after) {
				if theLength > len(after) {
					theLength = len(after)
					theRune = a
				}
				break
			}
			before = after
		}
	}
	fmt.Printf("Shortest length %d rune %c\n", theLength, theRune)
}

func main() {
	if file, err := os.Open("Day5.txt"); err == nil {
		defer file.Close()
		runes := util.ScanToRuneSlices(file)
		Day5P1(runes)
		Day5P2(runes)
	} else {
		panic(err)
	}
}
