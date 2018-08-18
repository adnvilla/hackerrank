package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestQueueUsingTowStacks(t *testing.T) {

	file, err := os.Open("steps.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var queries [][]int64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() { // internally, it advances token based on sperator
		// fmt.Println(scanner.Text()) // token in unicode-char
		queriesRowTemp := strings.Split(scanner.Text(), " ")

		var queriesRow []int64
		for _, queriesRowItem := range queriesRowTemp {
			queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
			checkError(err)
			queriesItem := int64(queriesItemTemp)
			queriesRow = append(queriesRow, queriesItem)
		}

		queries = append(queries, queriesRow)
	}

	// nqueries := queries[0][0]

	queue := make([]int64, 0)

	for i := int64(1); i < int64(len(queries)); i++ {

		switch query := queries[i]; query[0] {
		case 1:
			// enqueue X
			queue = append(queue, query[1])
		case 2:
			// dequeue
			queue = queue[1:]
		case 3:
			// print
			fmt.Println(queue[0])
		}
	}

	// total := queensAttack(queries[0][0], queries[0][1], queries[1][0], queries[1][1], queries[2:])
	// if total != 307303 {
	// 	t.Errorf("queensAttack was incorrect, got: %d, want: %d.", total, 307303)
	// }
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
