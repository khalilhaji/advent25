package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse_happy(t *testing.T) {
	expected := []coordinate{{7, 1}, {11, 1}, {8, 1}}

	got, err := parse([]string{"7,1", "11,1", "8,1"})
	assert.NoError(t, err)
	assert.Equal(t, expected, got)
}

func TestParse_error(t *testing.T) {
	_, err := parse([]string{"a, b"})
	assert.Error(t, err, errInvalidInput)
}
