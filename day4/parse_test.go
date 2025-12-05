package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse_Happy(t *testing.T) {
	in := []string{
		"..@@.@@@@.",
		"@@@.@.@.@@",
		"@@@@@.@.@@",
	}

	out := diagram{
		{SpaceEmpty, SpaceEmpty, SpacePaper, SpacePaper, SpaceEmpty, SpacePaper, SpacePaper, SpacePaper, SpacePaper, SpaceEmpty},
		{SpacePaper, SpacePaper, SpacePaper, SpaceEmpty, SpacePaper, SpaceEmpty, SpacePaper, SpaceEmpty, SpacePaper, SpacePaper},
		{SpacePaper, SpacePaper, SpacePaper, SpacePaper, SpacePaper, SpaceEmpty, SpacePaper, SpaceEmpty, SpacePaper, SpacePaper},
	}

	d, err := parse(in)
	assert.NoError(t, err)
	assert.Equal(t, out, d)
}

func TestParse_Say(t *testing.T) {
	in := []string{
		"foawiejfwefjowe",
	}

	d, err := parse(in)

	assert.Error(t, err)
	assert.Nil(t, d)
}
