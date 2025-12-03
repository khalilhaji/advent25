package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestParse_Happy(t *testing.T) {
	in := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224"
	ranges, err := parse(in)

	expected := []idRange{
		{11, 22},
		{95, 115},
		{998, 1012},
		{1188511880, 1188511890},
		{222220, 222224},

	}

	assert.NoError(t, err)
	assert.Equal(t, expected, ranges)
}

func TestParse_Sad_MalformedLine(t *testing.T) {
	in := "11-22,113242342"
	ranges, err := parse(in)

	assert.ErrorContains(t, err, "malformed line")
	assert.Nil(t, ranges)
}

func TestParse_Sad_NonIntValues(t *testing.T) {
	in := "11-22,a-b"

	ranges, err := parse(in)
	assert.Error(t, err)
	assert.Nil(t, ranges)
}
