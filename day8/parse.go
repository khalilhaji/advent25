package main

import (
	"errors"
	"strconv"
	"strings"
)

type box struct {
	x, y, z int
	circuit int
}

var errMalformedInput = errors.New("malformed input")

func parse(input []string) ([]box, error) {
	res := []box{}

	for i, line := range input {
		coordinates := strings.Split(line, ",")
		if len(coordinates) != 3 {
			return nil, errMalformedInput
		}

		x, err := strconv.Atoi(coordinates[0])
		if err != nil {
			return nil, errMalformedInput
		}

		y, err := strconv.Atoi(coordinates[1])
		if err != nil {
			return nil, errMalformedInput
		}

		z, err := strconv.Atoi(coordinates[2])
		if err != nil {
			return nil, errMalformedInput
		}

		res = append(res, box{x, y, z, i})
	}

	return res, nil
}
