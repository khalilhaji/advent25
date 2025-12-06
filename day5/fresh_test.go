package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountFresh_Happy(t *testing.T) {
	ranges := []freshRange{
		{3, 5},
		{10, 14},
		{16, 20},
		{12, 18},
	}

	avail := []int{1, 5, 8, 11, 17, 32}

	assert.Equal(t, 3, countFresh(ranges, avail))
}

func TestCountAllFresh_Happy(t *testing.T) {
	ranges := []freshRange{
		{3, 5},   // 3
		{10, 14}, // 5
		{16, 20}, // 5
		{12, 18}, // 1
		{16, 20}, // 0
		{12, 13}, // 0
		{12, 13}, // 0
		{18, 22}, // 2
		{18, 22}, // 0
		{16, 20}, // 0
		{16, 25}, // 3
		{8, 22},  // 2
	}

	assert.Equal(t, 21, countAllFresh(ranges))
}
