package main

import (
	"errors"
	"fmt"
	"github.com/collinzh/AdventOfCode2018/util"
	"math"
	"os"
	"regexp"
	"strconv"
)

type Star struct {
	Position *util.Position
	Velocity *util.Position
}

func (s *Star) Next() *Star {
	return &Star{
		Position: &util.Position{X: s.Position.X + s.Velocity.X, Y: s.Position.Y + s.Velocity.Y},
		Velocity: s.Velocity,
	}
}

func ParseStars(lines []string) []*Star {
	lineRegex := regexp.MustCompile(".*<\\s*([-0-9]+),\\s*([-0-9]+)>.*<\\s*([-0-9]+),\\s*([-0-9]+)>")
	stars := make([]*Star, 0)
	for _, line := range lines {
		if lineRegex.MatchString(line) {
			parts := lineRegex.FindStringSubmatch(line)
			x, _ := strconv.Atoi(parts[1])
			y, _ := strconv.Atoi(parts[2])
			vx, _ := strconv.Atoi(parts[3])
			vy, _ := strconv.Atoi(parts[4])
			stars = append(stars, &Star{Position: &util.Position{X: x, Y: y}, Velocity: &util.Position{X: vx, Y: vy}})
		} else {
			panic(errors.New("cannot parse " + line))
		}
	}
	return stars
}

func Day10(lines []string) {
	stars := ParseStars(lines)

	var minSize int64 = math.MaxInt64
	var minSky []*Star
	var minRound int

	// Hopefully we can get a good enough answer in 30000 seconds
	for i := 0; i < 30000; i++ {
		size := SkySize(stars)
		// Assuming the sky with the smallest total size would be the one with message
		if size < minSize {
			minSize = size
			minSky = stars
			minRound = i
		}
		newSky := make([]*Star, 0)
		for _, star := range stars {
			newSky = append(newSky, star.Next())
		}
		stars = newSky
	}

	fmt.Printf("Message at round %d\n", minRound)
	if minSky != nil {
		PrintSky(minSky)
	}
}

func SkySize(stars []*Star) int64 {
	var minX int64 = math.MaxInt64
	var maxX int64 = math.MinInt64
	var minY int64 = math.MaxInt64
	var maxY int64 = math.MinInt64

	for _, star := range stars {
		x := int64(star.Position.X)
		y := int64(star.Position.Y)
		if minX > x {
			minX = x
		} else if maxX < x {
			maxX = x
		}
		if minY > y {
			minY = y
		} else if maxY < y {
			maxY = y
		}
	}

	// I know it's supposed to be x*y, but I'm gonna go with x+y to avoid overflow.
	// I'm only using this for comparison, and the results should be the same as long as both x and y are positive numbers
	// which should be true in this case. Considering it's size we're measuring
	return (maxX - minX) + (maxY - minY)
}

func PrintSky(stars []*Star) {
	minX := math.MaxInt32
	maxX := math.MinInt32
	minY := math.MaxInt32
	maxY := math.MinInt32

	sky := make(map[int]map[int]bool)

	for _, star := range stars {
		x := star.Position.X
		y := star.Position.Y
		if minX > x {
			minX = x
		} else if maxX < x {
			maxX = x
		}
		if minY > y {
			minY = y
		} else if maxY < y {
			maxY = y
		}

		if sky[x] == nil {
			sky[x] = make(map[int]bool)
		}
		sky[x][y] = true
	}

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if sky[x] != nil && sky[x][y] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
}

func main() {
	if f, err := os.Open("Day10.txt"); err == nil {
		lines := util.ScanToStringSlices(f)
		Day10(lines)
	} else {
		panic(err)
	}
}
