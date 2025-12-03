package main

import (
	"fmt"
	"strconv"
	"strings"
)

type idRange struct {
	lower int
	upper int
}

func parse(input string) ([]idRange, error) {
	entries := strings.Split(input, ",")
	ranges := []idRange{}

	for _, entry := range entries {
		numbers := strings.Split(entry, "-")
		if len(numbers) != 2 {
			return nil, fmt.Errorf("malformed line: %v", entry)
		}
		lower, err := strconv.Atoi(numbers[0])
		if err != nil {
			return nil, err
		}

		upper, err := strconv.Atoi(numbers[1])
		if err != nil {
			return nil, err
		}

		ranges = append(ranges, idRange{lower, upper})
	}

	return ranges, nil
}
