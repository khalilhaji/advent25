package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse_happy(t *testing.T) {
	in := []string{
		"162,817,812",
		"57,618,57",
	}

	expected := []box{
		{162, 817, 812, 0},
		{57, 618, 57, 1},
	}

	out, err := parse(in)
	assert.NoError(t, err)
	assert.Equal(t, expected, out)
}

func TestParse_sad(t *testing.T) {
	in := []string{
		"162,817,812",
		"57,618,57",
		"a,1,2",
	}

	_, err := parse(in)
	assert.Error(t, err)
}
