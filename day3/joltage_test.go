package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumJoltage_Happy(t *testing.T) {
	banks := []bank{
		{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1},
		{8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 9},
		{2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 7, 8},
		{8, 1, 8, 1, 8, 1, 9, 1, 1, 1, 1, 2, 1, 1, 1},
	}

	assert.Equal(t, 357, sumJoltage(banks, 2))
}

func TestMaxJoltage_Happy(t *testing.T) {
	tests := []struct{
		in bank
		n int
		expected int
	}{
		{
			bank{1, 2, 3, 4},
			2,
			34,
		},
		{
			bank{9, 9, 9, 8},
			2,
			99,
		},
		{
			bank{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1},
			12,
			987654321111,
		},
		{
			bank{8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 9},
			12,
			811111111119,
		},
		{
			bank{2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 7, 8},
			12,
			434234234278,
		},
		{
			bank{8, 1, 8, 1, 8, 1, 9, 1, 1, 1, 1, 2, 1, 1, 1},
			12,
			888911112111,
		},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := maxJoltage(tt.in, tt.n)

			assert.Equal(t, tt.expected, actual)
		})
	}
}
