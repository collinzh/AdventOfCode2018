package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Day1P1(f *os.File) {
	s := bufio.NewScanner(f)
	total := 0
	for s.Scan() {
		line := s.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			continue
		}
		total += num
	}
	fmt.Printf("Total is %d\n", total)
}

func ScanToSlice(f *os.File) []int {
	numbers := make([]int, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(nil)
		}
		numbers = append(numbers, num)
	}
	return numbers
}

func Day1P2(f *os.File) {
	numbers := make(map[int]bool)
	list := ScanToSlice(f)
	counter := 0
	total := 0

	for ; counter < len(list); counter++ {
		next := list[counter]
		if numbers[total] == false {
			numbers[total] = true
		} else {
			fmt.Printf("Found %d\n", total)
			return
		}
		total += next
		if counter == len(list)-1 {
			counter = -1
		}
	}
	fmt.Printf("Didn't find anything\n")
}

func main() {
	if file, err := os.Open("Day1.txt"); err == nil {
		Day1P1(file)
		defer file.Close()
	}

	if file, err := os.Open("Day1.txt"); err == nil {
		Day1P2(file)
		defer file.Close()
	}
}
