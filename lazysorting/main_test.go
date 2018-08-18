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
		n      []int32
		expect float64
	}{
		{n: []int32{5, 2}, expect: 2.000000},
		{n: []int32{4, 1, 1, 2, 2, 2, 3, 3}, expect: 2.000000},
	}

	for _, test := range tests {
		start := time.Now()
		response := solve(test.n)

		elapsed := time.Since(start)
		log.Printf("time: %s", elapsed)
		// fmt.Printf("%d => %v => %d\n", test.n, r, len(r))
		// assert.Equal(t, test.expect, response, fmt.Sprintf("%d => %v => %d", test.n, r, len(r)))
		assert.Equal(t, test.expect, response, fmt.Sprintf("%v", test.n))
	}
}
