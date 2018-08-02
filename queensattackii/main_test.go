package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAttack1(t *testing.T) {
	total := queensAttack(5, 3, 4, 3, [][]int32{
		[]int32{5, 5},
		[]int32{4, 2},
		[]int32{2, 3},
	})
	if total != 10 {
		t.Errorf("queensAttack was incorrect, got: %d, want: %d.", total, 10)
	}
}

func TestAttack2(t *testing.T) {
	total := queensAttack(1, 0, 1, 1, [][]int32{})
	if total != 0 {
		t.Errorf("queensAttack was incorrect, got: %d, want: %d.", total, 0)
	}
}

func TestAttack3(t *testing.T) {
	total := queensAttack(4, 0, 4, 4, [][]int32{})
	if total != 9 {
		t.Errorf("queensAttack was incorrect, got: %d, want: %d.", total, 9)
	}
}

func TestAttack4(t *testing.T) {
	total := queensAttack(0, 0, 0, 0, [][]int32{})
	if total != 0 {
		t.Errorf("queensAttack was incorrect, got: %d, want: %d.", total, 0)
	}
}

func TestAttack5(t *testing.T) {
	total := queensAttack(1, 0, 1, 1, [][]int32{})
	if total != 0 {
		t.Errorf("queensAttack was incorrect, got: %d, want: %d.", total, 0)
	}
}

func TestAttack6(t *testing.T) {
	total := queensAttack(5, 0, 0, 0, [][]int32{})
	if total != 0 {
		t.Errorf("queensAttack was incorrect, got: %d, want: %d.", total, 0)
	}
}

func TestAttack7(t *testing.T) {
	total := queensAttack(2, 0, 1, 1, [][]int32{})
	if total != 3 {
		t.Errorf("queensAttack was incorrect, got: %d, want: %d.", total, 3)
	}
}

func TestAttack8(t *testing.T) {
	total := queensAttack(2, 4, 1, 1, [][]int32{{2, 1}, {2, 1}, {2, 2}, {1, 2}})
	if total != 0 {
		t.Errorf("queensAttack was incorrect, got: %d, want: %d.", total, 0)
	}
}

func TestAttack9(t *testing.T) {
	total := queensAttack(4, 0, 4, 4, [][]int32{{2, 1}, {2, 1}, {2, 2}, {1, 2}})
	if total != 7 {
		t.Errorf("queensAttack was incorrect, got: %d, want: %d.", total, 7)
	}
}

func TestAttack10(t *testing.T) {
	total := queensAttack(100000, 0, 1, 1, [][]int32{})
	if total != 299997 {
		t.Errorf("queensAttack was incorrect, got: %d, want: %d.", total, 299997)
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
		{n: 6, x: 3, y: 2, expect: []int32{4, 1}},
		{n: 6, x: 2, y: 3, expect: []int32{4, 1}},
		{n: 6, x: 2, y: 4, expect: []int32{5, 1}},
		{n: 6, x: 5, y: 3, expect: []int32{6, 2}},
		{n: 6, x: 4, y: 5, expect: []int32{6, 3}},
	}

	for _, test := range tests {
		response := getleftup(test.n, test.x, test.y)
		assert.Equal(t, test.expect, response, fmt.Sprintf("%d [%d,%d]", test.n, test.x, test.y))
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
		assert.Equal(t, test.expect, response, fmt.Sprintf("[%d,%d]", test.x, test.y))
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

func TestGetdown(t *testing.T) {
	tests := []struct {
		n      int32
		x      int32
		y      int32
		expect []int32
	}{
		{n: 5, x: 1, y: 1, expect: []int32{1, 1}},
		{n: 5, x: 2, y: 1, expect: []int32{1, 1}},
		{n: 5, x: 3, y: 1, expect: []int32{1, 1}},
		{n: 5, x: 4, y: 1, expect: []int32{1, 1}},
		{n: 5, x: 5, y: 1, expect: []int32{1, 1}},
		{n: 5, x: 1, y: 2, expect: []int32{1, 2}},
		{n: 5, x: 2, y: 2, expect: []int32{1, 2}},
		{n: 5, x: 3, y: 2, expect: []int32{1, 2}},
		{n: 5, x: 4, y: 2, expect: []int32{1, 2}},
		{n: 5, x: 5, y: 2, expect: []int32{1, 2}},
		{n: 5, x: 1, y: 3, expect: []int32{1, 3}},
		{n: 5, x: 2, y: 3, expect: []int32{1, 3}},
		{n: 5, x: 3, y: 3, expect: []int32{1, 3}},
		{n: 5, x: 4, y: 3, expect: []int32{1, 3}},
		{n: 5, x: 5, y: 3, expect: []int32{1, 3}},
		{n: 5, x: 1, y: 4, expect: []int32{1, 4}},
		{n: 5, x: 2, y: 4, expect: []int32{1, 4}},
		{n: 5, x: 3, y: 4, expect: []int32{1, 4}},
		{n: 5, x: 4, y: 4, expect: []int32{1, 4}},
		{n: 5, x: 5, y: 4, expect: []int32{1, 4}},
		{n: 5, x: 1, y: 5, expect: []int32{1, 5}},
		{n: 5, x: 2, y: 5, expect: []int32{1, 5}},
		{n: 5, x: 3, y: 5, expect: []int32{1, 5}},
		{n: 5, x: 4, y: 5, expect: []int32{1, 5}},
		{n: 5, x: 5, y: 5, expect: []int32{1, 5}},
	}

	for _, test := range tests {
		response := getdown(test.n, test.x, test.y)
		assert.Equal(t, test.expect, response, fmt.Sprintf("[%d,%d]", test.x, test.y))
	}
}

func TestGetrigthdown(t *testing.T) {
	tests := []struct {
		n      int32
		x      int32
		y      int32
		expect []int32
	}{
		{n: 5, x: 1, y: 1, expect: []int32{1, 1}},
		{n: 5, x: 2, y: 1, expect: []int32{1, 2}},
		{n: 5, x: 3, y: 1, expect: []int32{1, 3}},
		{n: 5, x: 4, y: 1, expect: []int32{1, 4}},
		{n: 5, x: 5, y: 1, expect: []int32{1, 5}},
		{n: 5, x: 1, y: 2, expect: []int32{1, 2}},
		{n: 5, x: 2, y: 2, expect: []int32{1, 3}},
		{n: 5, x: 3, y: 2, expect: []int32{1, 4}},
		{n: 5, x: 4, y: 2, expect: []int32{1, 5}},
		{n: 5, x: 5, y: 2, expect: []int32{2, 5}},
		{n: 5, x: 1, y: 3, expect: []int32{1, 3}},
		{n: 5, x: 2, y: 3, expect: []int32{1, 4}},
		{n: 5, x: 3, y: 3, expect: []int32{1, 5}},
		{n: 5, x: 4, y: 3, expect: []int32{2, 5}},
		{n: 5, x: 5, y: 3, expect: []int32{3, 5}},
		{n: 5, x: 1, y: 4, expect: []int32{1, 4}},
		{n: 5, x: 2, y: 4, expect: []int32{1, 5}},
		{n: 5, x: 3, y: 4, expect: []int32{2, 5}},
		{n: 5, x: 4, y: 4, expect: []int32{3, 5}},
		{n: 5, x: 5, y: 4, expect: []int32{4, 5}},
		{n: 5, x: 1, y: 5, expect: []int32{1, 5}},
		{n: 5, x: 2, y: 5, expect: []int32{2, 5}},
		{n: 5, x: 3, y: 5, expect: []int32{3, 5}},
		{n: 5, x: 4, y: 5, expect: []int32{4, 5}},
		{n: 5, x: 5, y: 5, expect: []int32{5, 5}},
	}

	for _, test := range tests {
		response := getrigthdown(test.n, test.x, test.y)
		assert.Equal(t, test.expect, response, fmt.Sprintf("[%d,%d]", test.x, test.y))
	}
}

func TestGetrigth(t *testing.T) {
	tests := []struct {
		n      int32
		x      int32
		y      int32
		expect []int32
	}{
		{n: 5, x: 1, y: 1, expect: []int32{1, 5}},
		{n: 5, x: 2, y: 1, expect: []int32{2, 5}},
		{n: 5, x: 3, y: 1, expect: []int32{3, 5}},
		{n: 5, x: 4, y: 1, expect: []int32{4, 5}},
		{n: 5, x: 5, y: 1, expect: []int32{5, 5}},
		{n: 5, x: 1, y: 2, expect: []int32{1, 5}},
		{n: 5, x: 2, y: 2, expect: []int32{2, 5}},
		{n: 5, x: 3, y: 2, expect: []int32{3, 5}},
		{n: 5, x: 4, y: 2, expect: []int32{4, 5}},
		{n: 5, x: 5, y: 2, expect: []int32{5, 5}},
		{n: 5, x: 1, y: 3, expect: []int32{1, 5}},
		{n: 5, x: 2, y: 3, expect: []int32{2, 5}},
		{n: 5, x: 3, y: 3, expect: []int32{3, 5}},
		{n: 5, x: 4, y: 3, expect: []int32{4, 5}},
		{n: 5, x: 5, y: 3, expect: []int32{5, 5}},
		{n: 5, x: 1, y: 4, expect: []int32{1, 5}},
		{n: 5, x: 2, y: 4, expect: []int32{2, 5}},
		{n: 5, x: 3, y: 4, expect: []int32{3, 5}},
		{n: 5, x: 4, y: 4, expect: []int32{4, 5}},
		{n: 5, x: 5, y: 4, expect: []int32{5, 5}},
		{n: 5, x: 1, y: 5, expect: []int32{1, 5}},
		{n: 5, x: 2, y: 5, expect: []int32{2, 5}},
		{n: 5, x: 3, y: 5, expect: []int32{3, 5}},
		{n: 5, x: 4, y: 5, expect: []int32{4, 5}},
		{n: 5, x: 5, y: 5, expect: []int32{5, 5}},
	}

	for _, test := range tests {
		response := getrigth(test.n, test.x, test.y)
		assert.Equal(t, test.expect, response)
	}
}

func TestGetrigthup(t *testing.T) {
	tests := []struct {
		n      int32
		x      int32
		y      int32
		expect []int32
	}{
		{n: 5, x: 1, y: 1, expect: []int32{5, 5}},
		{n: 5, x: 2, y: 1, expect: []int32{5, 4}},
		{n: 5, x: 3, y: 1, expect: []int32{5, 3}},
		{n: 5, x: 4, y: 1, expect: []int32{5, 2}},
		{n: 5, x: 5, y: 1, expect: []int32{5, 1}},
		{n: 5, x: 1, y: 2, expect: []int32{4, 5}},
		{n: 5, x: 2, y: 2, expect: []int32{5, 5}},
		{n: 5, x: 3, y: 2, expect: []int32{5, 4}},
		{n: 5, x: 4, y: 2, expect: []int32{5, 3}},
		{n: 5, x: 5, y: 2, expect: []int32{5, 2}},
		{n: 5, x: 1, y: 3, expect: []int32{3, 5}},
		{n: 5, x: 2, y: 3, expect: []int32{4, 5}},
		{n: 5, x: 3, y: 3, expect: []int32{5, 5}},
		{n: 5, x: 4, y: 3, expect: []int32{5, 4}},
		{n: 5, x: 5, y: 3, expect: []int32{5, 3}},
		{n: 5, x: 1, y: 4, expect: []int32{2, 5}},
		{n: 5, x: 2, y: 4, expect: []int32{3, 5}},
		{n: 5, x: 3, y: 4, expect: []int32{4, 5}},
		{n: 5, x: 4, y: 4, expect: []int32{5, 5}},
		{n: 5, x: 5, y: 4, expect: []int32{5, 4}},
		{n: 5, x: 1, y: 5, expect: []int32{1, 5}},
		{n: 5, x: 2, y: 5, expect: []int32{2, 5}},
		{n: 5, x: 3, y: 5, expect: []int32{3, 5}},
		{n: 5, x: 4, y: 5, expect: []int32{4, 5}},
		{n: 5, x: 5, y: 5, expect: []int32{5, 5}},
	}

	for _, test := range tests {
		response := getrigthup(test.n, test.x, test.y)
		assert.Equal(t, test.expect, response, fmt.Sprintf("[%d,%d]", test.x, test.y))
	}
}

func TestGetup(t *testing.T) {
	tests := []struct {
		n      int32
		x      int32
		y      int32
		expect []int32
	}{
		{n: 5, x: 1, y: 1, expect: []int32{5, 1}},
		{n: 5, x: 2, y: 1, expect: []int32{5, 1}},
		{n: 5, x: 3, y: 1, expect: []int32{5, 1}},
		{n: 5, x: 4, y: 1, expect: []int32{5, 1}},
		{n: 5, x: 5, y: 1, expect: []int32{5, 1}},
		{n: 5, x: 1, y: 2, expect: []int32{5, 2}},
		{n: 5, x: 2, y: 2, expect: []int32{5, 2}},
		{n: 5, x: 3, y: 2, expect: []int32{5, 2}},
		{n: 5, x: 4, y: 2, expect: []int32{5, 2}},
		{n: 5, x: 5, y: 2, expect: []int32{5, 2}},
		{n: 5, x: 1, y: 3, expect: []int32{5, 3}},
		{n: 5, x: 2, y: 3, expect: []int32{5, 3}},
		{n: 5, x: 3, y: 3, expect: []int32{5, 3}},
		{n: 5, x: 4, y: 3, expect: []int32{5, 3}},
		{n: 5, x: 5, y: 3, expect: []int32{5, 3}},
		{n: 5, x: 1, y: 4, expect: []int32{5, 4}},
		{n: 5, x: 2, y: 4, expect: []int32{5, 4}},
		{n: 5, x: 3, y: 4, expect: []int32{5, 4}},
		{n: 5, x: 4, y: 4, expect: []int32{5, 4}},
		{n: 5, x: 5, y: 4, expect: []int32{5, 4}},
		{n: 5, x: 1, y: 5, expect: []int32{5, 5}},
		{n: 5, x: 2, y: 5, expect: []int32{5, 5}},
		{n: 5, x: 3, y: 5, expect: []int32{5, 5}},
		{n: 5, x: 4, y: 5, expect: []int32{5, 5}},
		{n: 5, x: 5, y: 5, expect: []int32{5, 5}},
	}

	for _, test := range tests {
		response := getup(test.n, test.x, test.y)
		assert.Equal(t, test.expect, response, fmt.Sprintf("[%d,%d]", test.x, test.y))
	}
}

func TestGetCleanMoves(t *testing.T) {
	tests := []struct {
		n      int32
		x      int32
		y      int32
		expect int32
	}{
		{n: 5, x: 1, y: 1, expect: 12},
		{n: 5, x: 2, y: 1, expect: 12},
		{n: 5, x: 3, y: 1, expect: 12},
		{n: 5, x: 4, y: 1, expect: 12},
		{n: 5, x: 5, y: 1, expect: 12},
		{n: 5, x: 1, y: 2, expect: 12},
		{n: 5, x: 2, y: 2, expect: 14},
		{n: 5, x: 3, y: 2, expect: 14},
		{n: 5, x: 4, y: 2, expect: 14},
		{n: 5, x: 5, y: 2, expect: 12},
		{n: 5, x: 1, y: 3, expect: 12},
		{n: 5, x: 2, y: 3, expect: 14},
		{n: 5, x: 3, y: 3, expect: 16},
		{n: 5, x: 4, y: 3, expect: 14},
		{n: 5, x: 5, y: 3, expect: 12},
		{n: 5, x: 1, y: 4, expect: 12},
		{n: 5, x: 2, y: 4, expect: 14},
		{n: 5, x: 3, y: 4, expect: 14},
		{n: 5, x: 4, y: 4, expect: 14},
		{n: 5, x: 5, y: 4, expect: 12},
		{n: 5, x: 1, y: 5, expect: 12},
		{n: 5, x: 2, y: 5, expect: 12},
		{n: 5, x: 3, y: 5, expect: 12},
		{n: 5, x: 4, y: 5, expect: 12},
		{n: 5, x: 5, y: 5, expect: 12},
		{n: 6, x: 1, y: 6, expect: 15},
		{n: 6, x: 3, y: 2, expect: 17},
		{n: 6, x: 4, y: 3, expect: 19},
	}

	for _, test := range tests {
		response := getCleanMoves(test.n, test.x, test.y)
		assert.Equal(t, test.expect, response, fmt.Sprintf("[%d,%d]", test.x, test.y))
	}
}

func TestGetleftupObstacles(t *testing.T) {
	tests := []struct {
		n         int32
		x         int32
		y         int32
		obstacles [][]int32
		expect    []int32
	}{
		{n: 5, x: 4, y: 3, obstacles: [][]int32{{5, 2}}, expect: []int32{4, 3}},
		{n: 5, x: 1, y: 3, obstacles: [][]int32{{3, 1}}, expect: []int32{2, 2}},
		{n: 5, x: 2, y: 5, obstacles: [][]int32{{3, 1}}, expect: []int32{5, 2}},
		{n: 5, x: 2, y: 5, obstacles: [][]int32{{5, 2}}, expect: []int32{4, 3}},
		{n: 5, x: 2, y: 5, obstacles: [][]int32{{5, 2}, {3, 4}}, expect: []int32{2, 5}},
		{n: 5, x: 2, y: 5, obstacles: [][]int32{{3, 4}, {5, 2}}, expect: []int32{2, 5}},
		{n: 5, x: 2, y: 5, obstacles: [][]int32{{4, 3}, {3, 4}}, expect: []int32{2, 5}},
		{n: 5, x: 2, y: 5, obstacles: [][]int32{{3, 4}, {4, 3}}, expect: []int32{2, 5}},
		{n: 5, x: 2, y: 5, obstacles: [][]int32{{5, 2}, {4, 3}}, expect: []int32{3, 4}},
	}

	for _, test := range tests {
		response := getleftupObstacles(test.n, test.x, test.y, test.obstacles)
		assert.Equal(t, test.expect, response, fmt.Sprintf("%d [%d,%d]", test.n, test.x, test.y))
	}
}

func TestGetleftObstacles(t *testing.T) {
	tests := []struct {
		n         int32
		x         int32
		y         int32
		obstacles [][]int32
		expect    []int32
	}{
		{n: 5, x: 4, y: 3, obstacles: [][]int32{{4, 4}}, expect: []int32{4, 1}},
		{n: 5, x: 1, y: 3, obstacles: [][]int32{{1, 1}}, expect: []int32{1, 2}},
		{n: 5, x: 5, y: 5, obstacles: [][]int32{{5, 1}, {5, 3}}, expect: []int32{5, 4}},
		{n: 5, x: 5, y: 5, obstacles: [][]int32{{5, 3}, {5, 1}}, expect: []int32{5, 4}},
		{n: 5, x: 2, y: 5, obstacles: [][]int32{{3, 2}}, expect: []int32{2, 1}},
	}

	for _, test := range tests {
		response := getleftObstacles(test.n, test.x, test.y, test.obstacles)
		assert.Equal(t, test.expect, response, fmt.Sprintf("%d [%d,%d]", test.n, test.x, test.y))
	}
}

func TestGetleftdownObstacles(t *testing.T) {
	tests := []struct {
		n         int32
		x         int32
		y         int32
		obstacles [][]int32
		expect    []int32
	}{
		{n: 5, x: 4, y: 3, obstacles: [][]int32{{4, 4}}, expect: []int32{2, 1}},
		{n: 5, x: 4, y: 3, obstacles: [][]int32{{2, 1}}, expect: []int32{3, 2}},
		{n: 5, x: 5, y: 5, obstacles: [][]int32{{1, 1}, {3, 3}}, expect: []int32{4, 4}},
		{n: 5, x: 5, y: 5, obstacles: [][]int32{{3, 3}, {1, 1}}, expect: []int32{4, 4}},
	}

	for _, test := range tests {
		response := getleftdownObstacles(test.n, test.x, test.y, test.obstacles)
		assert.Equal(t, test.expect, response, fmt.Sprintf("%d [%d,%d]", test.n, test.x, test.y))
	}
}

func TestGetdownObstaclesObstacles(t *testing.T) {
	tests := []struct {
		n         int32
		x         int32
		y         int32
		obstacles [][]int32
		expect    []int32
	}{
		{n: 5, x: 4, y: 3, obstacles: [][]int32{{4, 4}}, expect: []int32{1, 3}},
		{n: 5, x: 4, y: 3, obstacles: [][]int32{{1, 3}}, expect: []int32{2, 3}},
		{n: 5, x: 5, y: 5, obstacles: [][]int32{{1, 5}, {3, 5}}, expect: []int32{4, 5}},
		{n: 5, x: 5, y: 5, obstacles: [][]int32{{3, 5}, {1, 5}}, expect: []int32{4, 5}},
	}

	for _, test := range tests {
		response := getdownObstacles(test.n, test.x, test.y, test.obstacles)
		assert.Equal(t, test.expect, response, fmt.Sprintf("%d [%d,%d]", test.n, test.x, test.y))
	}
}

func TestGetrigthdownObstaclesObstacles(t *testing.T) {
	tests := []struct {
		n         int32
		x         int32
		y         int32
		obstacles [][]int32
		expect    []int32
	}{
		{n: 5, x: 4, y: 3, obstacles: [][]int32{{4, 4}}, expect: []int32{2, 5}},
		{n: 5, x: 4, y: 3, obstacles: [][]int32{{3, 4}}, expect: []int32{4, 3}},
		{n: 5, x: 5, y: 1, obstacles: [][]int32{{3, 3}}, expect: []int32{4, 2}},
		{n: 5, x: 5, y: 1, obstacles: [][]int32{{1, 5}, {3, 3}}, expect: []int32{4, 2}},
		{n: 5, x: 5, y: 1, obstacles: [][]int32{{3, 3}, {1, 5}}, expect: []int32{4, 2}},
	}

	for _, test := range tests {
		response := getrigthdownObstacles(test.n, test.x, test.y, test.obstacles)
		assert.Equal(t, test.expect, response, fmt.Sprintf("%d [%d,%d]", test.n, test.x, test.y))
	}
}

func TestGetrigthObstaclesObstacles(t *testing.T) {
	tests := []struct {
		n         int32
		x         int32
		y         int32
		obstacles [][]int32
		expect    []int32
	}{
		{n: 5, x: 4, y: 3, obstacles: [][]int32{{5, 4}}, expect: []int32{4, 5}},
		{n: 5, x: 4, y: 3, obstacles: [][]int32{{4, 5}}, expect: []int32{4, 4}},
		{n: 5, x: 5, y: 1, obstacles: [][]int32{{5, 3}}, expect: []int32{5, 2}},
		{n: 5, x: 5, y: 1, obstacles: [][]int32{{5, 5}, {5, 3}}, expect: []int32{5, 2}},
		{n: 5, x: 5, y: 1, obstacles: [][]int32{{5, 3}, {5, 5}}, expect: []int32{5, 2}},
	}

	for _, test := range tests {
		response := getrigthObstacles(test.n, test.x, test.y, test.obstacles)
		assert.Equal(t, test.expect, response, fmt.Sprintf("%d [%d,%d]", test.n, test.x, test.y))
	}
}

func TestGetrigthupObstaclesObstacles(t *testing.T) {
	tests := []struct {
		n         int32
		x         int32
		y         int32
		obstacles [][]int32
		expect    []int32
	}{
		{n: 5, x: 4, y: 3, obstacles: [][]int32{{4, 4}}, expect: []int32{5, 4}},
		{n: 5, x: 1, y: 1, obstacles: [][]int32{{5, 5}}, expect: []int32{4, 4}},
		{n: 5, x: 1, y: 1, obstacles: [][]int32{{3, 3}}, expect: []int32{2, 2}},
		{n: 5, x: 1, y: 1, obstacles: [][]int32{{5, 5}, {3, 3}}, expect: []int32{2, 2}},
		{n: 5, x: 1, y: 1, obstacles: [][]int32{{3, 3}, {5, 5}}, expect: []int32{2, 2}},
	}

	for _, test := range tests {
		response := getrigthupObstacles(test.n, test.x, test.y, test.obstacles)
		assert.Equal(t, test.expect, response, fmt.Sprintf("%d [%d,%d]", test.n, test.x, test.y))
	}
}

func TestGetupObstaclesObstacles(t *testing.T) {
	tests := []struct {
		n         int32
		x         int32
		y         int32
		obstacles [][]int32
		expect    []int32
	}{
		{n: 5, x: 4, y: 3, obstacles: [][]int32{{4, 4}}, expect: []int32{5, 3}},
		{n: 5, x: 1, y: 1, obstacles: [][]int32{{5, 1}}, expect: []int32{4, 1}},
		{n: 5, x: 1, y: 1, obstacles: [][]int32{{3, 1}}, expect: []int32{2, 1}},
		{n: 5, x: 1, y: 1, obstacles: [][]int32{{5, 1}, {3, 1}}, expect: []int32{2, 1}},
		{n: 5, x: 1, y: 1, obstacles: [][]int32{{3, 1}, {5, 1}}, expect: []int32{2, 1}},
	}

	for _, test := range tests {
		response := getupObstacles(test.n, test.x, test.y, test.obstacles)
		assert.Equal(t, test.expect, response, fmt.Sprintf("%d [%d,%d]", test.n, test.x, test.y))
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
		{
			n: 6,
			x: 3,
			y: 2,
			expect: [][]int32{
				[]int32{4, 1},
				[]int32{3, 1},
				[]int32{2, 1},
				[]int32{1, 2},
				[]int32{1, 4},
				[]int32{3, 6},
				[]int32{6, 5},
				[]int32{6, 2},
			},
		},
	}

	for _, test := range tests {
		response := getFinalPoints(test.n, test.x, test.y)
		assert.Equal(t, test.expect, response)
	}
}
