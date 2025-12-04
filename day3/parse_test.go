package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse_Happy(t *testing.T) {
	lines := []string{
		"123",
		"234",
		"457567657",
	}

	actual, err := parse(lines)

	expected := []bank{
		[]int{1, 2, 3},
		[]int{2, 3, 4},
		[]int{4, 5, 7, 5, 6, 7, 6, 5, 7},
	}

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestParse_Sad(t *testing.T) {
	lines := []string{
		"123",
		"abc",
	}

	actual, err := parse(lines)
	assert.Error(t, err)
	assert.Equal(t, []bank(nil), actual)
}
