package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_solve(t *testing.T) {
	tests := []struct {
		name string
		p    problem
		want int
	}{
		{
			"1+1+1",
			problem{o: OperatorAdd, args: []int{1, 1, 1}, maxWidth: 1},
			3,
		},
		{
			"2*2*2",
			problem{o: OperatorMultiply, args: []int{2, 2, 2}, maxWidth: 1},
			8,
		},
		{
			"empty multiply",
			problem{o: OperatorMultiply, args: []int{}},
			1,
		},
		{
			"empty add",
			problem{o: OperatorAdd, args: []int{}},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := solve(tt.p)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_sumSolutions(t *testing.T) {
	problems := []problem{
		{OperatorMultiply, []int{123, 45, 6}, 3},
		{OperatorAdd, []int{328, 64, 98}, 3},
		{OperatorMultiply, []int{51, 387, 215}, 3},
		{OperatorAdd, []int{64, 23, 314}, 3},
	}

	res := sumSolutions(problems)
	assert.Equal(t, 4277556, res)
}

func Test_solveV2(t *testing.T) {
	tests := []struct {
		name string
		p    problemV2
		want int
	}{
		{
			"123 45 6",
			problemV2{OperatorMultiply, []string{"123", " 45", "  6"}, 3},
			8544,
		},
		{
			"328 64 98",
			problemV2{OperatorAdd, []string{"328", "64 ", "98 "}, 3},
			625,
		},
		{
			"51 387 215",
			problemV2{OperatorMultiply, []string{" 51", "387", "215"}, 3},
			3253600,
		},
		{
			"64 23 314",
			problemV2{OperatorAdd, []string{"64 ", "23 ", "314"}, 3},
			1058,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := solveV2(tt.p)
			assert.Equal(t, tt.want, got)
		})
	}
}
