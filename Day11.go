package main

import (
	"fmt"
	"math"
)

func GenerateGrid(serialNbr int) map[int]map[int]int {
	grid := make(map[int]map[int]int)
	for x := 1; x <= 300; x++ {
		grid[x] = make(map[int]int)
		for y := 1; y <= 300; y++ {
			rackID := x + 10
			power := rackID*y + serialNbr
			power = (power*rackID)/100%10 - 5

			grid[x][y] = power
		}
	}
	return grid
}

func GenerateSumGrid(serialNumber int) map[int]map[int]int {
	grid := make(map[int]map[int]int)
	for x := 1; x <= 300; x++ {
		grid[x] = make(map[int]int)
		for y := 1; y <= 300; y++ {
			rackID := x + 10
			power := rackID*y + serialNumber
			power = (power*rackID)/100%10 - 5

			grid[x][y] = power + grid[x-1][y] + grid[x][y-1] - grid[x-1][y-1]
		}
	}
	return grid
}

func Day11P1(serialNumber int) {
	grid := GenerateGrid(serialNumber)
	maxPower := math.MinInt32
	cornerX := 0
	cornerY := 0
	for x := 1; x <= 300-3; x++ {
		for y := 1; y <= 300-3; y++ {
			sum := 0
			for cX := x; cX < x+3; cX++ {
				for cY := y; cY < y+3; cY++ {
					sum += grid[cX][cY]
				}
			}

			if sum > maxPower {
				maxPower = sum
				cornerX = x
				cornerY = y
			}
		}
	}

	fmt.Printf("Serial number %d, Top-left corner: %d,%d with total power of %d\n", serialNumber, cornerX, cornerY, maxPower)
}

func Day11P2(serialNumber int) {
	grid := GenerateSumGrid(serialNumber)
	maxPower := math.MinInt64
	maxSize := 0
	cornerX := 0
	cornerY := 0

	for size := 1; size <= 300; size++ {
		for x := size; x <= 300; x++ {
			for y := size; y <= 300; y++ {
				sum := grid[x][y] - grid[x-size][y] - grid[x][y-size] + grid[x-size][y-size]

				if sum > maxPower {
					maxPower = sum
					maxSize = size
					cornerX = x
					cornerY = y
				}
			}
		}
	}

	// Convert bottom-right corner to top-left corner
	cornerX = cornerX - maxSize + 1
	cornerY = cornerY - maxSize + 1

	fmt.Printf("Serial number %d, Top-left corner: %d,%d,%d, with total power of %d\n", serialNumber, cornerX, cornerY, maxSize, maxPower)
}

func main() {
	Day11P1(42)
	Day11P1(4455)
	Day11P2(18)
	Day11P2(42)
	Day11P2(4455)
}
