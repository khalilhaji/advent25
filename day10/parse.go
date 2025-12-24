package main

import (
	"errors"
	"strconv"
	"strings"
)

type state int

const (
	Off state = iota
	On
)

type node struct {
	target  []state
	state   []state
	joltage []int
	edges   [][]int
	costs   []int
}

var errInvalidInput = errors.New("invalid input")

func parse(input []string) ([]node, error) {
	res := []node{}
	for _, line := range input {
		newNode := node{}
		parts := strings.Split(line, " ")
		target := parts[0]
		for _, light := range target[1 : len(target)-1] {
			switch light {
			case '.':
				newNode.target = append(newNode.target, Off)
			case '#':
				newNode.target = append(newNode.target, On)
			}
		}

		// start with everything off
		newNode.state = make([]state, len(newNode.target))

		for _, part := range parts[1:] {
			numStrs := strings.Split(part[1:len(part)-1], ",")
			arr := []int{}
			for _, numStr := range numStrs {
				num, err := strconv.Atoi(numStr)
				if err != nil {
					return nil, errInvalidInput
				}
				arr = append(arr, num)
			}

			switch part[0] {
			case '(':
				newNode.edges = append(newNode.edges, arr)
			case '{':
				newNode.costs = arr
			}
		}

		newNode.joltage = make([]int, len(newNode.costs))

		res = append(res, newNode)
	}

	return res, nil
}
