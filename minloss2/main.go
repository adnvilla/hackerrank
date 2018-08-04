package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func sliceToMap(slice []int64) map[int64]int64 {
	mapSlice := make(map[int64]int64)
	for i := int64(0); i < int64(len(slice)); i++ {
		mapSlice[slice[i]] = i
	}
	return mapSlice
}

// Complete the minimumLoss function below.
func minimumLoss(price []int64) int32 {

	mapPrice := sliceToMap(price)
	order := QuickSort(price)
	mapOrder := sliceToMap(order)
	minloss := order[len(order)-1] - order[0]
	for i := int64(0); i < int64(len(price)); i++ {

		priceVal := price[i]
		posOrder := mapOrder[priceVal]

		if posOrder-1 < 0 {
			continue
		}

		lowValue := order[posOrder-1]
		posPrice := mapPrice[lowValue]

		if i > posPrice {
			continue
		}

		if priceVal-lowValue < minloss {
			minloss = priceVal - lowValue
		}

		if minloss == 1 {
			break
		}
	}
	return int32(minloss)
}

func QuickSort(slice []int64) []int64 {
	length := len(slice)

	if length <= 1 {
		sliceCopy := make([]int64, length)
		copy(sliceCopy, slice)
		return sliceCopy
	}

	m := slice[rand.Intn(length)]

	less := make([]int64, 0, length)
	middle := make([]int64, 0, length)
	more := make([]int64, 0, length)

	for _, item := range slice {
		switch {
		case item < m:
			less = append(less, item)
		case item == m:
			middle = append(middle, item)
		case item > m:
			more = append(more, item)
		}
	}

	less, more = QuickSort(less), QuickSort(more)

	less = append(less, middle...)
	less = append(less, more...)

	return less
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 2048*2048)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	priceTemp := strings.Split(readLine(reader), " ")

	var price []int64

	for i := 0; i < int(n); i++ {
		priceItem, err := strconv.ParseInt(priceTemp[i], 10, 64)
		checkError(err)
		price = append(price, priceItem)
	}

	result := minimumLoss(price)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
