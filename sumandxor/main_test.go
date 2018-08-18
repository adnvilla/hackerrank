package main

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		n      int64
		expect int64
	}{
		{n: 0, expect: 1},
		{n: 1, expect: 1},
		{n: 2, expect: 2},
		{n: 3, expect: 1},
		{n: 4, expect: 4},
		{n: 5, expect: 2},
		{n: 6, expect: 2},
		{n: 7, expect: 1},
		{n: 8, expect: 8},
		{n: 9, expect: 4},
		{n: 10, expect: 4},
		{n: 11, expect: 2},
		{n: 12, expect: 4},
		{n: 13, expect: 2},
		{n: 14, expect: 2},
		{n: 15, expect: 1},
		{n: 16, expect: 16},
		{n: 17, expect: 8},
		{n: 18, expect: 8},
		{n: 19, expect: 4},
		// {n: 20, expect: 16},
		// {n: 21, expect: 16},
		// {n: 22, expect: 16},
		// {n: 23, expect: 16},
		// {n: 24, expect: 16},
		// {n: 25, expect: 16},
		// {n: 26, expect: 16},
		// {n: 27, expect: 16},
		// {n: 28, expect: 16},
		// {n: 29, expect: 16},
		// {n: 30, expect: 16},
		// {n: 31, expect: 16},
		// {n: 32, expect: 16},
		// {n: 33, expect: 16},
		// {n: 34, expect: 16},
		// {n: 35, expect: 16},
		// {n: 36, expect: 16},
		// {n: 37, expect: 16},
		// {n: 38, expect: 16},
		// {n: 39, expect: 16},
		// {n: 40, expect: 16},
		// {n: 41, expect: 16},
		// {n: 42, expect: 16},
		// {n: 43, expect: 16},
		// {n: 44, expect: 16},
		// {n: 45, expect: 16},
		// {n: 46, expect: 16},
		// {n: 47, expect: 16},
		// {n: 48, expect: 16},
		// {n: 49, expect: 16},
		// {n: 50, expect: 16},
		// {n: 51, expect: 16},
		// {n: 52, expect: 16},
		// {n: 53, expect: 16},
		// {n: 54, expect: 16},
		// {n: 55, expect: 16},
		// {n: 56, expect: 16},
		// {n: 57, expect: 16},
		// {n: 58, expect: 16},
		// {n: 59, expect: 16},
		// {n: 60, expect: 16},
		// {n: 61, expect: 16},
		// {n: 62, expect: 16},
		// {n: 63, expect: 16},
		// {n: 64, expect: 16},
		// {n: 65, expect: 16},
		// {n: 66, expect: 16},
		// {n: 67, expect: 16},
		// {n: 68, expect: 16},
		// {n: 69, expect: 16},
		// {n: 70, expect: 16},

		// {n: 100000, expect: 2},
		// {n: 10000000, expect: 2},
		// {n: 10000000000, expect: 2},
	}

	for _, test := range tests {
		start := time.Now()
		response := solve(test.n)

		elapsed := time.Since(start)
		log.Printf("time: %s", elapsed)
		// fmt.Printf("%d => %v => %d\n", test.n, r, len(r))
		// assert.Equal(t, test.expect, response, fmt.Sprintf("%d => %v => %d", test.n, r, len(r)))
		assert.Equal(t, test.expect, response, fmt.Sprintf("%d ", test.n))
	}
}
