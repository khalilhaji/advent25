package main

import (
	"strconv"
	"strings"
)

type bank []int

func parse(lines []string) ([]bank, error) {
	res := []bank{}

	for _, line := range lines {
		chars := strings.Split(line, "")
		bank := []int{}
		for _, char := range chars {
			i, err := strconv.Atoi(char)
			if err != nil {
				return nil, err
			}
			bank = append(bank, i)
		}

		res = append(res, bank)
	}

	return res, nil
}
