package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountAdjacent_Happy(t *testing.T) {
	d := diagram{
		{SpaceEmpty, SpacePaper, SpacePaper, SpaceEmpty},
		{SpaceEmpty, SpacePaper, SpacePaper, SpacePaper},
		{SpaceEmpty, SpacePaper, SpacePaper, SpaceEmpty},
	}

	tests := []struct {
		i int
		j int
		expected int
	}{
		{0, 0, 2},
		{0, 1, 3},
		{0, 2, 4},
		{0, 3, 3},
		{1, 0, 3},
		{1, 1, 5},
		{1, 2, 6},
		{1, 3, 3},
		{2, 0, 2},
		{2, 1, 3},
		{2, 2, 4},
		{2, 3, 3},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := countAdjacentRolls(tt.i, tt.j, d)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestTraverse_Happy(t *testing.T) {
	in := []string{
		"..@@.@@@@.",
		"@@@.@.@.@@",
		"@@@@@.@.@@",
		"@.@@@@..@.",
		"@@.@@@@.@@",
		".@@@@@@@.@",
		".@.@.@.@@@",
		"@.@@@.@@@@",
		".@@@@@@@@.",
		"@.@.@@@.@.",
	}

	d, err := parse(in)
	assert.NoError(t, err)

	res := traverse(d)
	assert.Equal(t, 13, res)
}

func TestTraverseAndRemove_Happy(t *testing.T) {
in := []string{
		"..@@.@@@@.",
		"@@@.@.@.@@",
		"@@@@@.@.@@",
		"@.@@@@..@.",
		"@@.@@@@.@@",
		".@@@@@@@.@",
		".@.@.@.@@@",
		"@.@@@.@@@@",
		".@@@@@@@@.",
		"@.@.@@@.@.",
	}

	d, err := parse(in)
	assert.NoError(t, err)

	assert.Equal(t, 43, traverseAndRemove(d))
}
