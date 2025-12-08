package main

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type operator int

const (
	OperatorMultiply operator = iota
	OperatorAdd
)

type problem struct {
	o        operator
	args     []int
	maxWidth int
}

func parse(input []string) ([]problem, error) {
	res := []problem{}
	lastLine := input[len(input)-1]
	i := 0
	for numStr := range strings.SplitSeq(lastLine, " ") {
		if numStr == "" {
			continue
		}
		p := problem{}
		switch numStr {
		case "*":
			p.o = OperatorMultiply
		case "+":
			p.o = OperatorAdd
		}

		res = append(res, p)
		i++
	}

	for _, line := range input[:len(input)-1] {
		i = 0
		for numStr := range strings.SplitSeq(line, " ") {
			if numStr == "" {
				continue
			}

			num, err := strconv.Atoi(numStr)
			if err != nil {
				return nil, errors.New("invalid number")
			}

			res[i].args = append(res[i].args, num)
			if len(numStr) > res[i].maxWidth {
				res[i].maxWidth = len(numStr)
			}

			i++
		}
	}

	return res, nil
}

type problemV2 struct {
	o        operator
	args     []string
	maxWidth int
}

func parseV2(input []string) []problemV2 {
	res := []problemV2{}
	lastLine := input[len(input)-1]
	r := regexp.MustCompile("([+|*] +)")
	operators := r.FindAllString(lastLine, -1)
	for i, oper := range operators {
		p := problemV2{}
		if i == len(operators)-1 {
			p.maxWidth = len(oper)
		} else {
			p.maxWidth = len(oper) - 1
		}
		switch oper[0] {
		case '+':
			p.o = OperatorAdd
		case '*':
			p.o = OperatorMultiply
		}

		res = append(res, p)
	}

	for _, line := range input[:len(input)-1] {
		start := 0
		for i, p := range res {
			res[i].args = append(res[i].args, line[start:start+p.maxWidth])
			start = start + p.maxWidth + 1
		}
	}

	return res
}
