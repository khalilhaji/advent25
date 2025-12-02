package main

import (
	"fmt"
	"strconv"
)

type direction int

const (
	DIRECTION_LEFT direction = iota
	DIRECTION_RIGHT
)

type instruction struct {
	dir direction
	num int
}

func Parse(input []string) ([]instruction, error) {
	instructions := make([]instruction, 0)

	for _, instr := range input {
		var dir direction
		if instr[0] == 'L' {
			dir = DIRECTION_LEFT
		} else {
			dir = DIRECTION_RIGHT
		}

		num, err := strconv.Atoi(instr[1:])
		if err != nil {
			return nil, fmt.Errorf("error parsing numerical part of instruction: %w", err)
		}

		if num < 0 {
			return nil, fmt.Errorf("cannot parse negative instruction: %v", instr)
		}

		instructions = append(instructions, instruction{
			dir: dir,
			num: num,
		})
	}

	return instructions, nil
}
