package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	total := queensAttack(5, 10, 1, 2, [][]int32{
		[]int32{5, 5},
		[]int32{4, 2},
		[]int32{2, 3},
	})
	if total != 10 {
		t.Errorf("queensAttack was incorrect, got: %d, want: %d.", total, 10)
	}
}

func TestGetFinalPoints(t *testing.T) {
	tests := []struct {
		n      int32
		x      int32
		y      int32
		expect [][]int32
	}{
		{
			n: 5,
			x: 4,
			y: 3,
			expect: [][]int32{
				[]int32{5, 2},
				[]int32{4, 1},
				[]int32{2, 1},
				[]int32{1, 3},
				[]int32{2, 5},
				[]int32{4, 5},
				[]int32{5, 4},
				[]int32{5, 3},
			},
		},
		{
			n: 4,
			x: 4,
			y: 3,
			expect: [][]int32{
				[]int32{4, 3},
				[]int32{4, 1},
				[]int32{2, 1},
				[]int32{1, 3},
				[]int32{3, 4},
				[]int32{4, 4},
				[]int32{4, 3},
				[]int32{4, 3},
			},
		},
		{
			n: 5,
			x: 1,
			y: 1,
			expect: [][]int32{
				[]int32{1, 1},
				[]int32{1, 1},
				[]int32{1, 1},
				[]int32{1, 1},
				[]int32{1, 1},
				[]int32{1, 5},
				[]int32{5, 5},
				[]int32{5, 1},
			},
		},
	}

	for _, test := range tests {
		response := getFinalPoints(test.n, test.x, test.y)
		assert.Equal(t, test.expect, response)
	}
}
