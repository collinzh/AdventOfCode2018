package main

import (
	"./util"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type CoordinateExpansion struct {
	Id       int
	Origin   *util.Position
	Boundary []*util.Position
	Area     int
}

type ClaimedPosition struct {
	position *util.Position
	owner    int
	inRound  int
}

type Canvas map[int]map[int]*ClaimedPosition

func (c *Canvas) Prepare(position *util.Position) {
	if (*c)[position.X] == nil {
		(*c)[position.X] = make(map[int]*ClaimedPosition)
	}
}

func ParsePositions(lines []string) []*util.Position {
	positions := make([]*util.Position, 0)
	regex := regexp.MustCompile("^([0-9]+), ([0-9]+)$")
	for _, line := range lines {
		if regex.MatchString(line) {
			parts := regex.FindStringSubmatch(line)
			x, _ := strconv.Atoi(parts[1])
			y, _ := strconv.Atoi(parts[2])
			positions = append(positions, &util.Position{X: x, Y: y})
		}
	}
	return positions
}

func Day6P1(lines []string) {
	centers := ParsePositions(lines)
	origins := make(map[int]*CoordinateExpansion)
	claims := make(Canvas)

	for id, center := range centers {
		boundary := make([]*util.Position, 0)
		boundary = append(boundary, center)
		origins[id] = &CoordinateExpansion{Id: id, Origin: center, Boundary: boundary, Area: 1}
		claims.Prepare(center)
		claims[center.X][center.Y] = &ClaimedPosition{position: center, owner: id, inRound: 0}
	}

	for round := 1; round < 3000; round++ {
		if round%1000 == 0 {
			fmt.Printf("Calculating round %d\n", round)
		}

		for _, origin := range origins {
			newBoundary := make([]*util.Position, 0)

			for _, next := range origin.Boundary {
				// Check surrounding 4 coordinates
				expandToCoordinate := func(x, y int) {
					claims.Prepare(&util.Position{X: x, Y: y})

					coordinate := claims[x][y]

					if coordinate != nil {
						// If already claimed
						if coordinate.inRound == round && coordinate.owner != -1 && coordinate.owner != origin.Id {
							// If the coordinate is equally distant from two origins
							origins[coordinate.owner].Area--
							coordinate.owner = -1
						}
					} else {
						// If not claimed yet
						claims[x][y] = &ClaimedPosition{position: &util.Position{X: x, Y: y}, owner: origin.Id, inRound: round}
						newBoundary = append(newBoundary, &util.Position{X: x, Y: y})
						origin.Area++
					}
				}

				expandToCoordinate(next.X, next.Y+1)
				expandToCoordinate(next.X, next.Y-1)
				expandToCoordinate(next.X+1, next.Y)
				expandToCoordinate(next.X-1, next.Y)
			}

			origin.Boundary = newBoundary
		}

	}

	biggestArea := -1
	biggestOrigin := 0
	for _, coord := range origins {
		if len(coord.Boundary) == 0 {
			if biggestArea < coord.Area {
				biggestArea = coord.Area
				biggestOrigin = coord.Id
			}
		}
	}
	fmt.Printf("Largest area %d belongs to %d\n", biggestArea, biggestOrigin)
}

func Day6P2(lines []string) {
	coordinates := ParsePositions(lines)
	minX := math.MaxInt32
	minY := math.MaxInt32
	maxX := math.MinInt32
	maxY := math.MinInt32
	for _, coordinate := range coordinates {
		if coordinate.X < minX {
			minX = coordinate.X
		} else if coordinate.X > maxX {
			maxX = coordinate.X
		}
		if coordinate.Y < minY {
			minY = coordinate.Y
		} else if coordinate.Y > maxY {
			maxY = coordinate.Y
		}
	}

	counter := 0
	for x := minX; x <= maxX; x++ {
		for y := minY; y < maxY; y++ {
			sum := 0
			position := &util.Position{X: x, Y: y}
			for _, coordinate := range coordinates {
				sum += util.Distance(position, coordinate)
			}
			if sum < 10000 {
				counter++
			}
		}
	}
	fmt.Printf("Counter is %d\n", counter)
}

func main() {
	if f, err := os.Open("Day6.txt"); err == nil {
		lines := util.ScanToStringSlices(f)
		Day6P1(lines)
		Day6P2(lines)
	}
}
