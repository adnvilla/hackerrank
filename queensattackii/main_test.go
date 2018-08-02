package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAttack(t *testing.T) {
	total := queensAttack(5, 10, 2, 1, [][]int32{
		[]int32{5, 5},
		[]int32{4, 2},
		[]int32{2, 3},
	})
	if total != 10 {
		t.Errorf("queensAttack was incorrect, got: %d, want: %d.", total, 10)
	}
}

func TestGetleftup(t *testing.T) {
	tests := []struct {
		n      int32
		x      int32
		y      int32
		expect []int32
	}{
		{n: 5, x: 1, y: 1, expect: []int32{1, 1}},
		{n: 5, x: 2, y: 1, expect: []int32{2, 1}},
		{n: 5, x: 3, y: 1, expect: []int32{3, 1}},
		{n: 5, x: 4, y: 1, expect: []int32{4, 1}},
		{n: 5, x: 5, y: 1, expect: []int32{5, 1}},
		{n: 5, x: 1, y: 2, expect: []int32{2, 1}},
		{n: 5, x: 2, y: 2, expect: []int32{3, 1}},
		{n: 5, x: 3, y: 2, expect: []int32{4, 1}},
		{n: 5, x: 4, y: 2, expect: []int32{5, 1}},
		{n: 5, x: 5, y: 2, expect: []int32{5, 2}},
		{n: 5, x: 1, y: 3, expect: []int32{3, 1}},
		{n: 5, x: 2, y: 3, expect: []int32{4, 1}},
		{n: 5, x: 3, y: 3, expect: []int32{5, 1}},
		{n: 5, x: 4, y: 3, expect: []int32{5, 2}},
		{n: 5, x: 5, y: 3, expect: []int32{5, 3}},
		{n: 5, x: 1, y: 4, expect: []int32{4, 1}},
		{n: 5, x: 2, y: 4, expect: []int32{5, 1}},
		{n: 5, x: 3, y: 4, expect: []int32{5, 2}},
		{n: 5, x: 4, y: 4, expect: []int32{5, 3}},
		{n: 5, x: 5, y: 4, expect: []int32{5, 4}},
		{n: 5, x: 1, y: 5, expect: []int32{5, 1}},
		{n: 5, x: 2, y: 5, expect: []int32{5, 2}},
		{n: 5, x: 3, y: 5, expect: []int32{5, 3}},
		{n: 5, x: 4, y: 5, expect: []int32{5, 4}},
		{n: 5, x: 5, y: 5, expect: []int32{5, 5}},
	}

	for _, test := range tests {
		response := getleftup(test.n, test.x, test.y)
		assert.Equal(t, test.expect, response)
	}
}

func TestGetleft(t *testing.T) {
	tests := []struct {
		n      int32
		x      int32
		y      int32
		expect []int32
	}{
		{n: 5, x: 1, y: 1, expect: []int32{1, 1}},
		{n: 5, x: 2, y: 1, expect: []int32{2, 1}},
		{n: 5, x: 3, y: 1, expect: []int32{3, 1}},
		{n: 5, x: 4, y: 1, expect: []int32{4, 1}},
		{n: 5, x: 5, y: 1, expect: []int32{5, 1}},
		{n: 5, x: 1, y: 2, expect: []int32{1, 1}},
		{n: 5, x: 2, y: 2, expect: []int32{2, 1}},
		{n: 5, x: 3, y: 2, expect: []int32{3, 1}},
		{n: 5, x: 4, y: 2, expect: []int32{4, 1}},
		{n: 5, x: 5, y: 2, expect: []int32{5, 1}},
		{n: 5, x: 1, y: 3, expect: []int32{1, 1}},
		{n: 5, x: 2, y: 3, expect: []int32{2, 1}},
		{n: 5, x: 3, y: 3, expect: []int32{3, 1}},
		{n: 5, x: 4, y: 3, expect: []int32{4, 1}},
		{n: 5, x: 5, y: 3, expect: []int32{5, 1}},
		{n: 5, x: 1, y: 4, expect: []int32{1, 1}},
		{n: 5, x: 2, y: 4, expect: []int32{2, 1}},
		{n: 5, x: 3, y: 4, expect: []int32{3, 1}},
		{n: 5, x: 4, y: 4, expect: []int32{4, 1}},
		{n: 5, x: 5, y: 4, expect: []int32{5, 1}},
		{n: 5, x: 1, y: 5, expect: []int32{1, 1}},
		{n: 5, x: 2, y: 5, expect: []int32{2, 1}},
		{n: 5, x: 3, y: 5, expect: []int32{3, 1}},
		{n: 5, x: 4, y: 5, expect: []int32{4, 1}},
		{n: 5, x: 5, y: 5, expect: []int32{5, 1}},
	}

	for _, test := range tests {
		response := getleft(test.n, test.x, test.y)
		assert.Equal(t, test.expect, response)
	}
}

func TestGetleftdown(t *testing.T) {
	tests := []struct {
		n      int32
		x      int32
		y      int32
		expect []int32
	}{
		{n: 5, x: 1, y: 1, expect: []int32{1, 1}},
		{n: 5, x: 2, y: 1, expect: []int32{2, 1}},
		{n: 5, x: 3, y: 1, expect: []int32{3, 1}},
		{n: 5, x: 4, y: 1, expect: []int32{4, 1}},
		{n: 5, x: 5, y: 1, expect: []int32{5, 1}},
		{n: 5, x: 1, y: 2, expect: []int32{1, 2}},
		{n: 5, x: 2, y: 2, expect: []int32{1, 1}},
		{n: 5, x: 3, y: 2, expect: []int32{2, 1}},
		{n: 5, x: 4, y: 2, expect: []int32{3, 1}},
		{n: 5, x: 5, y: 2, expect: []int32{4, 1}},
		{n: 5, x: 1, y: 3, expect: []int32{1, 3}},
		{n: 5, x: 2, y: 3, expect: []int32{1, 2}},
		{n: 5, x: 3, y: 3, expect: []int32{1, 1}},
		{n: 5, x: 4, y: 3, expect: []int32{2, 1}},
		{n: 5, x: 5, y: 3, expect: []int32{3, 1}},
		{n: 5, x: 1, y: 4, expect: []int32{1, 4}},
		{n: 5, x: 2, y: 4, expect: []int32{1, 3}},
		{n: 5, x: 3, y: 4, expect: []int32{1, 2}},
		{n: 5, x: 4, y: 4, expect: []int32{1, 1}},
		{n: 5, x: 5, y: 4, expect: []int32{2, 1}},
		{n: 5, x: 1, y: 5, expect: []int32{1, 5}},
		{n: 5, x: 2, y: 5, expect: []int32{1, 4}},
		{n: 5, x: 3, y: 5, expect: []int32{1, 3}},
		{n: 5, x: 4, y: 5, expect: []int32{1, 2}},
		{n: 5, x: 5, y: 5, expect: []int32{1, 1}},
	}

	for _, test := range tests {
		response := getleftdown(test.n, test.x, test.y)
		assert.Equal(t, test.expect, response, fmt.Sprintf("[%d,%d]", test.x, test.y))
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
