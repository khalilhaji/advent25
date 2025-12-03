package main

import (
	"strconv"
	"strings"
)

func invalidIDs(idRange idRange) []int {
	results := []int{}
	for i := idRange.lower; i <= idRange.upper; i++ {
		str := strconv.Itoa(i)
		if len(str) % 2 == 0 && str[:len(str)/2] == str[len(str)/2:] {
			results = append(results, i)
		}
	}

	return results
}

func invalidIDsV2(idRange idRange) []int {
	results := []int{}
	for i := idRange.lower; i <= idRange.upper; i++ {
		str := strconv.Itoa(i)
		for j := 1; j <= len(str) / 2; j++ {
			if strings.Count(str, str[:j]) * j == len(str) {
				results = append(results, i)
				break
			}
		}
	}

	return results
}


func sumInvalid(ranges []idRange, identifier func(idRange) []int) int {
	total := 0
	for _, rng := range ranges {
		ids := identifier(rng)
		for _, id := range ids {
			total += id
		}
	}

	return total
}
