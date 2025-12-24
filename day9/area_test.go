package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArea(t *testing.T) {
	tests := []struct {
		a   coordinate
		b   coordinate
		out int
	}{
		{
			coordinate{2, 5},
			coordinate{9, 7},
			24,
		},
		{
			coordinate{7, 3},
			coordinate{2, 3},
			6,
		},
		{
			coordinate{0, 0},
			coordinate{0, 0},
			1,
		},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := area(tt.a, tt.b)
			assert.Equal(t, tt.out, got)
		})
	}
}

func TestMaxArea(t *testing.T) {
	in, _ := parse([]string{
		"7,1",
		"11,1",
		"11,7",
		"9,7",
		"9,5",
		"2,5",
		"2,3",
		"7,3",
	})

	out, out2 := maxArea(in)
	assert.Equal(t, 50, out)
	assert.Equal(t, 24, out2)
}

func TestIsValid(t *testing.T) {
}
