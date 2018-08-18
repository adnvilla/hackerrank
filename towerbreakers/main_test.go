package main

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTowerBreakers(t *testing.T) {
	tests := []struct {
		n      int32
		m      int32
		expect int32
	}{
		{n: 2, m: 2, expect: 2},
		{n: 1, m: 4, expect: 1},
		{n: 2, m: 4, expect: 2},
		{n: 2, m: 6, expect: 2},
		{n: 2, m: 9, expect: 2},
	}

	for _, test := range tests {
		start := time.Now()
		response := towerBreakers(test.n, test.m)

		elapsed := time.Since(start)
		log.Printf("time: %s", elapsed)
		// fmt.Printf("%d => %v => %d\n", test.n, r, len(r))
		// assert.Equal(t, test.expect, response, fmt.Sprintf("%d => %v => %d", test.n, r, len(r)))
		assert.Equal(t, test.expect, response, fmt.Sprintf("%d %d", test.n, test.m))
	}
}

func TestGetFactors(t *testing.T) {
	tests := []struct {
		n      int32
		expect []int32
	}{
		{n: 2, expect: []int32{}},
		{n: 3, expect: []int32{}},
		{n: 4, expect: []int32{2, 2}},
		{n: 5, expect: []int32{5}},
		{n: 6, expect: []int32{2, 3}},
		{n: 7, expect: []int32{7}},
		{n: 8, expect: []int32{2, 2, 2}},
		{n: 9, expect: []int32{3, 3}},
		{n: 10, expect: []int32{2, 5}},
		{n: 11, expect: []int32{11}},
	}

	for _, test := range tests {
		start := time.Now()
		response := getFactors(test.n)

		elapsed := time.Since(start)
		log.Printf("time: %s", elapsed)
		fmt.Printf("%d => %v => %d\n", test.n, response, len(response))
		// assert.Equal(t, test.expect, response, fmt.Sprintf("%d => %v => %d", test.n, r, len(r)))
		assert.Equal(t, test.expect, response, fmt.Sprintf("%d ", test.n))
	}
}
