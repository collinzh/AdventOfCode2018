package main

import (
	"./util"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Claim struct {
	id, width, height int
	position          *Position
}

type Position struct {
	x, y int
}

var lineExp = regexp.MustCompile("^#([0-9]+) @ ([0-9]+),([0-9]+): ([0-9]+)x([0-9]+)$")

func ParseClaim(input string) *Claim {
	if !lineExp.MatchString(input) {
		fmt.Println("Cannot parse string " + input)
	}
	found := lineExp.FindStringSubmatch(input)
	id, _ := strconv.Atoi(found[1])
	x, _ := strconv.Atoi(found[2])
	y, _ := strconv.Atoi(found[3])
	width, _ := strconv.Atoi(found[4])
	height, _ := strconv.Atoi(found[5])

	return &Claim{id: id, position: &Position{x: x, y: y}, width: width, height: height}
}

func DrawCanvas(claims []*Claim) map[Position][]*Claim {
	canvas := make(map[Position][]*Claim)

	for _, claim := range claims {
		for x := claim.position.x; x < claim.position.x+claim.width; x++ {
			for y := claim.position.y; y < claim.position.y+claim.height; y++ {
				pos := Position{x: x, y: y}
				if canvas[pos] == nil {
					canvas[pos] = make([]*Claim, 0)
				}
				canvas[pos] = append(canvas[pos], claim)
			}
		}
	}

	return canvas
}

func Day3P1(file *os.File) {
	lines := util.ScanToStringSlices(file)
	claims := make([]*Claim, 0)
	for _, line := range lines {
		claims = append(claims, ParseClaim(line))
	}

	canvas := DrawCanvas(claims)

	counter := 0
	for _, c := range canvas {
		if len(c) > 1 {
			counter++
		}
	}
	fmt.Printf("Found %d conflicts\n", counter)
}

func Day3P2(file *os.File) {
	lines := util.ScanToStringSlices(file)
	claims := make([]*Claim, 0)
	tracker := make(map[int]*Claim)
	for _, line := range lines {
		claim := ParseClaim(line)
		tracker[claim.id] = claim
		claims = append(claims, claim)
	}

	canvas := DrawCanvas(claims)

	for _, c := range canvas {
		if len(c) > 1 {
			for _, cl := range c {
				delete(tracker, cl.id)
			}
		}
	}

	for _, c := range tracker {
		fmt.Printf("Remaining claim: %d\n", c.id)
	}
}

func main() {
	if file, err := os.Open("Day3.txt"); err == nil {
		Day3P1(file)
		defer file.Close()
	}

	if file, err := os.Open("Day3.txt"); err == nil {
		Day3P2(file)
		defer file.Close()
	}
}
