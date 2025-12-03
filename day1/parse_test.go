package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse_Happy(t *testing.T) {
	instrs := []string{
		"L22",
		"R20",
		"L109",
	}

	parsed := []instruction{
		{DIRECTION_LEFT, 22},
		{DIRECTION_RIGHT, 20},
		{DIRECTION_LEFT, 109},

	}

	actual, err := Parse(instrs)

	assert.NoError(t, err)
	assert.Equal(t, parsed, actual)
}

func TestParse_Sad_Malformed(t *testing.T) {
	instrs := []string{
		"L22",
		"R20",
		"L10A",
		"R43",
	}

	actual, err := Parse(instrs)
	assert.ErrorContains(t, err, "error parsing numerical part of instruction:")
	assert.Nil(t, actual)
}

func TestParse_Sad_Negative(t *testing.T) {
	instrs := []string{
		"L-22",
	}

	actual, err := Parse(instrs)
	assert.ErrorContains(t, err, "cannot parse negative instruction:")
	assert.Nil(t, actual)
}
