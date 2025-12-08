package main

import "strings"

type diagram struct {
	start     int
	splitters []map[int]bool
}

func parse(input []string) diagram {
	start := strings.Index(input[0], "S")
	splitters := []map[int]bool{}

	for _, line := range input[1:] {
		splittersForLine := map[int]bool{}
		for i, char := range line {
			if char == '^' {
				splittersForLine[i] = true
			}
		}
		splitters = append(splitters, splittersForLine)
	}

	return diagram{
		start,
		splitters,
	}
}
