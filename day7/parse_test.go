package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse_happy(t *testing.T) {
	in := []string{
		".......S.......",
		"...............",
		".......^.......",
		"...............",
		"......^.^......",
		"...............",
		".....^.^.^.....",
		"...............",
		"....^.^...^....",
		"...............",
	}

	d := parse(in)

	expected := diagram{
		start: 7,
		splitters: []map[int]bool{
			{},
			{7: true},
			{},
			{6: true, 8: true},
			{},
			{5: true, 7: true, 9: true},
			{},
			{4: true, 6: true, 10: true},
			{},
		},
	}

	assert.Equal(t, expected, d)
}
