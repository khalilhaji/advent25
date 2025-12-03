package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvalidIDs_Happy(t *testing.T) {
	tests := []struct{
		rng idRange
		expected []int
	}{
		{
			rng: idRange{11, 22},
			expected: []int{11, 22},
		},
		{
			rng: idRange{95, 115},
			expected: []int{99},
		},
		{
			rng: idRange{998, 1012},
			expected: []int{1010},
		},
		{
			rng: idRange{1188511880, 1188511890},
			expected: []int{1188511885},
		},
		{
			rng: idRange{222220, 222224},
			expected: []int{222222},
		},
		{
			rng: idRange{1698522, 1698528},
			expected: []int{},
		},
		{
			rng: idRange{446443, 446449},
			expected: []int{446446},
		},
		{
			rng: idRange{38593856, 38593862},
			expected: []int{38593859},
		},
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := invalidIDs(test.rng)

			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestInvalidIDsV2_Happy(t *testing.T) {
	tests := []struct{
		rng idRange
		expected []int
	}{
		{
			rng: idRange{11, 22},
			expected: []int{11, 22},
		},
		{
			rng: idRange{95, 115},
			expected: []int{99, 111},
		},
		{
			rng: idRange{998, 1012},
			expected: []int{999, 1010},
		},
		{
			rng: idRange{1188511880, 1188511890},
			expected: []int{1188511885},
		},
		{
			rng: idRange{222220, 222224},
			expected: []int{222222},
		},
		{
			rng: idRange{1698522, 1698528},
			expected: []int{},
		},
		{
			rng: idRange{446443, 446449},
			expected: []int{446446},
		},
		{
			rng: idRange{38593856, 38593862},
			expected: []int{38593859},
		},
		{
			rng: idRange{565653, 565659},
			expected: []int{565656},
		},
		{
			rng: idRange{824824821, 824824827},
			expected: []int{824824824},
		},
		{
			rng: idRange{2121212118, 2121212124},
			expected: []int{2121212121},
		},
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := invalidIDsV2(test.rng)

			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestSumInvalid_Happy(t *testing.T) {
	ranges := []idRange{
		{11, 22},
		{95, 115},
		{998, 1012},
		{1188511880, 1188511890},
		{222220, 222224},
		{1698522, 1698528},
		{446443, 446449},
		{38593856, 38593862},
	}

	res := sumInvalid(ranges)

	assert.Equal(t, 1227775554, res)
}
