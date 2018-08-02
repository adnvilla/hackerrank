package main

import (
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"
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
