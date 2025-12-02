package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecute_Happy(t *testing.T) {
	tests := []struct{
		init state
		instr instruction
		expected state
	}{
		{
			init: state{0},
			instr: instruction{DIRECTION_RIGHT, 5},
			expected: state{5},
		},
		{
			init: state{10},
			instr: instruction{DIRECTION_LEFT, 4},
			expected: state{6},
		},
		{
			init: state{0},
			instr: instruction{DIRECTION_RIGHT, 500},
			expected: state{0},
		},
		{
			init: state{0},
			instr: instruction{DIRECTION_RIGHT, 150},
			expected: state{50},
		},
		{
			init: state{0},
			instr: instruction{DIRECTION_LEFT, 150},
			expected: state{50},
		},
		{
			init: state{99},
			instr: instruction{DIRECTION_RIGHT, 10},
			expected: state{9},
		},
		{
			init: state{50},
			instr: instruction{DIRECTION_LEFT, 68},
			expected: state{82},
		},
	}


	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			newState, _ := test.init.Execute(test.instr)
			assert.Equal(t, test.expected, newState)
		})
	}
}

func TestCountZeros_Happy(t *testing.T) {
	instrs := []instruction{
		{DIRECTION_LEFT, 68},
		{DIRECTION_LEFT, 30},
		{DIRECTION_RIGHT, 48},
		{DIRECTION_LEFT, 5},
		{DIRECTION_RIGHT, 60},
		{DIRECTION_LEFT, 55},
		{DIRECTION_LEFT, 1},
		{DIRECTION_LEFT, 99},
		{DIRECTION_RIGHT, 14},
		{DIRECTION_LEFT, 82},
	}


	res := CountZeroEndStates(instrs)

	assert.Equal(t, 3, res)
}

func TestCountZeroPasses_Happy(t *testing.T) {
	tests := []struct{
		instructions []instruction
		expected int
	}{
		{
			[]instruction{
				{DIRECTION_LEFT, 68},
				{DIRECTION_LEFT, 30},
				{DIRECTION_RIGHT, 48},
				{DIRECTION_LEFT, 5},
				{DIRECTION_RIGHT, 60},
				{DIRECTION_LEFT, 55},
				{DIRECTION_LEFT, 1},
				{DIRECTION_LEFT, 99},
				{DIRECTION_RIGHT, 14},
				{DIRECTION_LEFT, 82},
			},
			6,
		},
		{
			[]instruction{
				{DIRECTION_LEFT, 50},
			},
			1,
		},
		{
			[]instruction{
				{DIRECTION_LEFT, 50},
				{DIRECTION_LEFT, 95},
			},
			1,
		},
		{
			[]instruction{
				{DIRECTION_LEFT, 50},
				{DIRECTION_LEFT, 100},
			},
			2,
		},
		{
			[]instruction{
				{DIRECTION_RIGHT, 50},
				{DIRECTION_RIGHT, 200},
			},
			3,
		},
		{
			[]instruction{
				{DIRECTION_LEFT, 50},
				{DIRECTION_LEFT, 100},
				{DIRECTION_RIGHT, 200},
			},
			4,
		},
		{
			[]instruction{
				{DIRECTION_RIGHT, 45},
				{DIRECTION_RIGHT, 15},
			},
			1,
		},
		{
			[]instruction{
				{DIRECTION_RIGHT, 1000},
			},
			10,
		},
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			res := CountZeroPasses(test.instructions)
			assert.Equal(t, test.expected, res)
		})
	}
}

