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
// func solve(n int64) (result int64, r []int64) {
func solve(n int64) int64 {

	if n == 0 {
		return 1
	}

	b := strconv.FormatInt(n, 2)
	r := strings.Count(b, "0")
	e := math.Pow(2, float64(r))

	// if n == 0 {
	// 	result = 1
	// 	return
	// }

	// var i int64
	// for ; i <= n; i++ {
	// 	if i+n == n^i {
	// 		result++
	// 		r = append(r, i)
	// 	}
	// }
	return int64(e)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	n, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	result := solve(n)

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
