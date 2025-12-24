package main

import (
	"errors"
	"strconv"
	"strings"
)

type coordinate struct {
	row, col int
}

var errInvalidInput = errors.New("")

func parse(input []string) ([]coordinate, error) {
	res := []coordinate{}

	for _, line := range input {
		coordinates := strings.Split(line, ",")
		if len(coordinates) != 2 {
			return nil, errInvalidInput
		}
		row, err := strconv.Atoi(coordinates[0])
		if err != nil {
			return nil, errInvalidInput
		}
		col, err := strconv.Atoi(coordinates[1])
		if err != nil {
			return nil, errInvalidInput
		}

		res = append(res, coordinate{row, col})
	}

	return res, nil
}
