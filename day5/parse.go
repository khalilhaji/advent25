package main

import (
	"errors"
	"strconv"
	"strings"
)

type freshRange struct {
	low int
	up  int
}

func parse(input []string) ([]freshRange, []int, error) {
	freshRanges := []freshRange{}
	available := []int{}
	fresh := true

	for _, line := range input {
		if line == "" {
			fresh = false
			continue
		}

		if !fresh {
			a, err := strconv.Atoi(line)
			if err != nil {
				return nil, nil, errors.New("invalid character")
			}

			available = append(available, a)
			continue
		}

		parts := strings.Split(line, "-")
		if len(parts) < 2 {
			return nil, nil, errors.New("malformed range")
		}

		low, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, nil, errors.New("invalid numeric input")
		}
		high, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, nil, errors.New("invalid numeric input")
		}

		freshRanges = append(freshRanges, freshRange{low, high})
	}

	return freshRanges, available, nil
}
