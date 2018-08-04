package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {

	price := []int64{5, 10, 3}

	r := minimumLoss(price)

	if r != 2 {
		t.Errorf("Test1 was incorrect, got: %d, want: %d.", r, 2)
	}
}

func Test2(t *testing.T) {

	price := []int64{20, 7, 8, 2, 5}

	r := minimumLoss(price)

	if r != 2 {
		t.Errorf("Test2 was incorrect, got: %d, want: %d.", r, 2)
	}
}

func TestMinimumLoss(t *testing.T) {
	tests := []struct {
		price  []int64
		expect int32
	}{
		// {price: []int64{}, expect: 0},
		// {expect: 0},
		{price: []int64{5, 10, 3}, expect: 2},
		{price: []int64{20, 7, 8, 2, 5}, expect: 2},
		{price: []int64{1, 7, 8, 2, 5}, expect: 2},
		// {price: []int64{9, 7, 8, 2, 5}, expect: 1},
		// {price: []int64{9, 8, 7, 2, 5}, expect: 1},
		// {price: []int64{9, 8, 7, 5, 2}, expect: 1},
		{price: []int64{2, 5, 7, 8, 6, 9}, expect: 1},
		// {price: []int64{2, 5, 7, 8, 9, 6}, expect: 1},
		// {price: []int64{2, 5, 7, 6, 8, 9}, expect: 1},
		// {price: []int64{100, 90, 80, 70, 60, 50}, expect: 10},
		// {price: GenerateTest(10000), expect: 10},
		{price: []int64{5577006791947779410, 8674665223082153551, 6129484611666145821, 4037200794235010051, 3916589616287113937, 6334824724549167320, 605394647632969758, 1443635317331776148, 894385949183117216}, expect: 37190629},
	}

	for _, test := range tests {
		response := minimumLoss(test.price)
		assert.Equal(t, test.expect, response, fmt.Sprintf("%v", test.price))
	}
}
