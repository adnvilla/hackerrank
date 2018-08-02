package main

import (
	"fmt"
	"log"
	"math/rand"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {

	rand.Seed(time.Now().UTC().UnixNano())
	//...

	price := make([]int64, 100)
	for i := 0; i < len(price)-1; i++ {
		r := rand.Int63()

		price[i] = r
	}

	start := time.Now()
	r := MinimumLoss(price)
	elapsed := time.Since(start)
	log.Printf("time: %s", elapsed)

	fmt.Println(r)

	if r != 12 {
		t.Errorf("TestMin was incorrect, got: %d, want: %d.", 12, r)
	}
}

func Test1(t *testing.T) {

	price := []int64{5, 10, 3}

	r := MinimumLoss(price)

	if r != 2 {
		t.Error("Is not expected")
	}
}

func Test2(t *testing.T) {

	price := []int64{20, 7, 8, 2, 5}

	r := MinimumLoss(price)

	if r != 2 {
		t.Error("Is not expected")
	}
}

func TestMinimumLoss(t *testing.T) {
	tests := []struct {
		price  []int64
		expect int32
	}{
		// {price: []int64{}, expect: 0},
		// {expect: 0},
		// {price: []int64{5, 10, 3}, expect: 2},
		// {price: []int64{20, 7, 8, 2, 5}, expect: 2},
		// {price: []int64{1, 7, 8, 2, 5}, expect: 2},
		// {price: []int64{9, 7, 8, 2, 5}, expect: 1},
		// {price: []int64{9, 8, 7, 2, 5}, expect: 1},
		// {price: []int64{9, 8, 7, 5, 2}, expect: 1},
		// {price: []int64{2, 5, 7, 8, 6, 9}, expect: 1},
		// {price: []int64{2, 5, 7, 8, 9, 6}, expect: 1},
		// {price: []int64{2, 5, 7, 6, 8, 9}, expect: 1},
		// {price: []int64{100, 90, 80, 70, 60, 50}, expect: 10},
		{price: GenerateTest(10000), expect: 10},
		// {price: []int64{5577006791947779410, 8674665223082153551, 6129484611666145821, 4037200794235010051, 3916589616287113937, 6334824724549167320, 605394647632969758, 1443635317331776148, 894385949183117216}, expect: 37190629},
	}

	for _, test := range tests {
		response := MinimumLoss(test.price)
		assert.Equal(t, test.expect, response, fmt.Sprintf("%v", test.price))
	}
}

func TestMinimumLoss2(t *testing.T) {
	tests := []struct {
		price  []int64
		expect int64
	}{
		{price: []int64{}, expect: 0},
		{expect: 0},
		{price: []int64{5, 10, 3}, expect: 2},
		{price: []int64{20, 7, 8, 2, 5}, expect: 2},
		{price: []int64{1, 7, 8, 2, 5}, expect: 2},
		{price: GenerateTest(1000), expect: 2},
		// {price: []int64{9, 7, 8, 2, 5}, expect: 1},
		// {price: []int64{9, 8, 7, 2, 5}, expect: 1},
		// {price: []int64{9, 8, 7, 5, 2}, expect: 1},
	}

	for _, test := range tests {
		response := MinimumLoss2(test.price)
		assert.Equal(t, test.expect, response, fmt.Sprintf("%v", test.price))
	}
}

func TestSearch(t *testing.T) {
	tests := []struct {
		price  []int64
		value  int64
		expect int64
	}{
		// {price: []int64{}, expect: 0},
		// {expect: 0},
		// {value: 6, expect: 5, price: []int64{10, 5, 3}},
		// {value: 6, expect: 5, price: []int64{5, 10, 3}},
		// {value: 17, expect: 8, price: []int64{20, 7, 8, 2, 5}},
		// {value: 6, expect: 5, price: []int64{1, 7, 8, 2, 5}},
		// {value: 1, expect: 2, price: []int64{9, 7, 8, 2, 5}},
		{price: GenerateTest(100000), expect: 1},
		// {price: []int64{9, 8, 7, 5, 2}, expect: 1},
	}

	for _, test := range tests {
		// fmt.Println(QuickSort(test.price))
		// Quicksort2(test.price)
		r := QuickSort(test.price)
		response := Search(test.value, r)
		assert.Equal(t, test.expect, response, fmt.Sprintf("%v", test.price))
	}
}

var testData = []struct {
	input          []int64
	expectedOutput []int64
}{
	{[]int64{}, []int64{}},
	{[]int64{42}, []int64{42}},
	{[]int64{42, 23}, []int64{23, 42}},
	{[]int64{23, 42, 32, 64, 12, 4}, []int64{4, 12, 23, 32, 42, 64}},
}

func TestQuickSort(t *testing.T) {
	for _, testCase := range testData {
		actual := QuickSort(testCase.input)
		expected := testCase.expectedOutput

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("%v != %v\n", actual, expected)
		}
	}
}
