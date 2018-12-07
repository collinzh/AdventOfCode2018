package util

import (
	"bufio"
	"os"
	"strconv"
)

func ScanToIntSlice(f *os.File) []int {
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

func ScanToStringSlices(file *os.File) []string {
	str := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str = append(str, scanner.Text())
	}
	return str
}
