package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse_Happy(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		input   []string
		want    []freshRange
		want2   []int
		wantErr bool
	}{
		{
			"advent example",
			[]string{
				"3-4",
				"10-14",
				"16-20",
				"12-18",
				"",
				"1",
				"5",
				"8",
				"11",
				"17",
				"32",
			},
			[]freshRange{
				{3, 4},
				{10, 14},
				{16, 20},
				{12, 18},
			},
			[]int{
				1,
				5,
				8,
				11,
				17,
				32,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fresh, avail, err := parse(tt.input)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, fresh)
			assert.Equal(t, tt.want2, avail)
		})
	}
}
