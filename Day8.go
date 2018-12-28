package main

import (
	"fmt"
	"github.com/collinzh/AdventOfCode2018/util"
	"os"
)

type IntegerReader struct {
	Numbers []int
	Cursor  int
}

func NewIntegerReader(numbers []int) *IntegerReader {
	return &IntegerReader{Numbers: numbers, Cursor: 0}
}

func (r *IntegerReader) Next() int {
	ret := r.Numbers[r.Cursor]
	r.Cursor++
	return ret
}

type Node struct {
	Meta     []int
	Children []*Node
	Parent   *Node
}

func ParseNodes(parent *Node, reader *IntegerReader) *Node {
	numChildren := reader.Next()
	numMeta := reader.Next()
	myself := &Node{Meta: make([]int, 0), Children: make([]*Node, 0), Parent: parent}

	for c := 0; c < numChildren; c++ {
		ParseNodes(myself, reader)
	}

	for m := 0; m < numMeta; m++ {
		myself.Meta = append(myself.Meta, reader.Next())
	}

	if parent != nil {
		parent.Children = append(parent.Children, myself)
	}

	return myself
}

func Day8P1(numbers []int) {
	root := ParseNodes(nil, NewIntegerReader(numbers))

	sum := 0
	toCheck := []*Node{root}

	for len(toCheck) > 0 {
		next := toCheck[0]
		toCheck = toCheck[1:]
		toCheck = append(toCheck, next.Children...)
		for _, meta := range next.Meta {
			sum += meta
		}
	}

	fmt.Printf("Total metadata: %d\n", sum)
}

func NodeValue(node *Node) int {
	sum := 0
	numChildren := len(node.Children)

	if numChildren == 0 {
		for _, val := range node.Meta {
			sum += val
		}
	} else {
		for _, val := range node.Meta {
			idx := val - 1
			if idx >= numChildren {
				continue
			}

			sum += NodeValue(node.Children[idx])
		}
	}

	return sum
}

func Day8P2(numbers []int) {
	root := ParseNodes(nil, NewIntegerReader(numbers))

	val := NodeValue(root)

	fmt.Printf("Root value is %d\n", val)
}

func main() {
	if f, err := os.Open("Day8.txt"); err == nil {
		numbers := util.ScanToIntSlice(f)
		Day8P1(numbers)
		Day8P2(numbers)
	} else {
		panic(err)
	}
}
