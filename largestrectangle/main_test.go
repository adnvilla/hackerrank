package main

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLargestRectangle(t *testing.T) {
	tests := []struct {
		n      []int32
		expect int64
	}{
		{n: []int32{2, 3, 1, 4, 5, 4, 2}, expect: 12},
		{n: []int32{1, 2, 3, 4, 5}, expect: 9},
		{n: []int32{1, 3, 5, 9, 11}, expect: 18},
		{n: []int32{11, 11, 10, 10, 10}, expect: 50},
	}

	for _, test := range tests {
		start := time.Now()
		response := largestRectangle(test.n)

		elapsed := time.Since(start)
		log.Printf("time: %s", elapsed)
		// fmt.Printf("%d => %v => %d\n", test.n, r, len(r))
		// assert.Equal(t, test.expect, response, fmt.Sprintf("%d => %v => %d", test.n, r, len(r)))
		assert.Equal(t, test.expect, response, fmt.Sprintf("%v", test.n))
	}
}

func TestLargestRectangle2(t *testing.T) {
	tests := []struct {
		n      []int32
		expect int64
	}{
		{n: []int32{2, 3, 1, 4, 5, 4, 2}, expect: 12},
		// {n: []int32{1, 2, 3, 4, 5}, expect: 9},
		// {n: []int32{1, 3, 5, 9, 11}, expect: 18},
		// {n: []int32{11, 11, 10, 10, 10}, expect: 50},
	}

	for _, test := range tests {
		start := time.Now()
		response := largestRectangle2(test.n)

		elapsed := time.Since(start)
		log.Printf("time: %s", elapsed)
		// fmt.Printf("%d => %v => %d\n", test.n, r, len(r))
		// assert.Equal(t, test.expect, response, fmt.Sprintf("%d => %v => %d", test.n, r, len(r)))
		assert.Equal(t, test.expect, response, fmt.Sprintf("%v", test.n))
	}
}
