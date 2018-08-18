package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

// Complete the solve function below.
func solve(P []int32) float64 {
	if isOrdered(P) {
		return 0
	}

	m := getMap(P)

	// math.Round(sig*1000000) / 1000000
	// return math.Round(float64(factor(len(P))) / float64(factors(m)))

	i := 1000
	var sig float64
	n := float64(factor(len(P))) / float64(factors(m))
	for i >= 0 {
		sig += sum(i, n)
		i--
	}

	return math.Round(sig*1000000) / 1000000
}

func sum(i int, n float64) float64 {
	return float64(i) * (math.Pow(float64(n), -float64(i)))
}

func factor(n int) int {

	if n == 1 {
		return 1
	}

	return n * factor(n-1)
}

func factors(m map[int32]int32) int {

	f := 1
	for v := range m {
		f *= factor(int(m[v]))
	}

	return f
}

func getMap(P []int32) map[int32]int32 {
	m := make(map[int32]int32)

	for i := 0; i < len(P); i++ {

		_, ok := m[P[i]]

		if ok {
			m[P[i]]++
			continue
		}

		m[P[i]] = 1
	}

	return m
}

func isOrdered(P []int32) bool {
	for i := 0; i < len(P)-1; i++ {
		if P[i] < P[i+1] {
			continue
		}
		return false
	}

	return true
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	PCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	PTemp := strings.Split(readLine(reader), " ")

	var P []int32

	for i := 0; i < int(PCount); i++ {
		PItemTemp, err := strconv.ParseInt(PTemp[i], 10, 64)
		checkError(err)
		PItem := int32(PItemTemp)
		P = append(P, PItem)
	}

	result := solve(P)

	fmt.Fprintf(writer, "%f\n", result)

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
