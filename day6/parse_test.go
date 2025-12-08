package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse_happy(t *testing.T) {
	in := []string{
		"123 328  51 64 0",
		"45 64  387 23 4321",
		"6 98  215 314 1234",
		"*   +   *   +  +",
	}

	expected := []problem{
		{OperatorMultiply, []int{123, 45, 6}, 3},
		{OperatorAdd, []int{328, 64, 98}, 3},
		{OperatorMultiply, []int{51, 387, 215}, 3},
		{OperatorAdd, []int{64, 23, 314}, 3},
		{OperatorAdd, []int{0, 4321, 1234}, 4},
	}

	out, err := parse(in)
	assert.NoError(t, err)
	assert.Equal(t, expected, out)
}

func TestParseV2_happy(t *testing.T) {
	in := []string{
		"123 328  51 64 ",
		" 45 64  387 23 ",
		"  6 98  215 314",
		"*   +   *   +  ",
	}

	expected := []problemV2{
		{OperatorMultiply, []string{"123", " 45", "  6"}, 3},
		{OperatorAdd, []string{"328", "64 ", "98 "}, 3},
		{OperatorMultiply, []string{" 51", "387", "215"}, 3},
		{OperatorAdd, []string{"64 ", "23 ", "314"}, 3},
	}

	out := parseV2(in)
	assert.Equal(t, expected, out)
}
