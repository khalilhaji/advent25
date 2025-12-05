package main

import (
	"errors"
	"fmt"
)

type symbol int

const (
	SpaceEmpty symbol = iota
	SpacePaper
)

type diagram [][]symbol

func parse(input []string) (diagram, error) {
	if len(input) == 0 {
		return diagram{}, nil
	}

	length := len(input[0])

	res := diagram{}
	for _, line := range input {
		if len(line) != length {
			return nil, errors.New("lines must have consistent lengths")
		}
		diagramLine := []symbol{}
		for _, char := range line {
			switch char {
			case '.':
				diagramLine = append(diagramLine, SpaceEmpty)
			case '@':
				diagramLine = append(diagramLine, SpacePaper)
			default:
				return nil, fmt.Errorf("invalid character: %v", char)

			}
		}
		res = append(res, diagramLine)
	}

	return res, nil
}

