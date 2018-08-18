package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the towerBreakers function below.
func towerBreakers(n int32, m int32) int32 {

	if n == 1 {
		return 1
	}

	if m == 1 {
		return 2
	}

	if n%2 == 0 {
		return 2
	}

	return 1

}

func getFactors(n int32) []int32 {

	var r []int32

	if n < 1 {
		return r
	}

	if n%2 == 0 {
		r = append(r, 2)
		r = append(r, getFactors(n/2)...)
	} else if n%3 == 0 {
		r = append(r, 3)
		r = append(r, getFactors(n/3)...)
	} else if n%5 == 0 {
		r = append(r, 5)
		r = append(r, getFactors(n/5)...)
	} else if n%7 == 0 {
		r = append(r, 7)
		r = append(r, getFactors(n/7)...)
	} else if n%11 == 0 {
		r = append(r, 11)
		r = append(r, getFactors(n/11)...)
	} else if n%13 == 0 {
		r = append(r, 13)
		r = append(r, getFactors(n/13)...)
	} else if n%17 == 0 {
		r = append(r, 17)
		r = append(r, getFactors(n/17)...)
	}

	return r
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nm := strings.Split(readLine(reader), " ")

		nTemp, err := strconv.ParseInt(nm[0], 10, 64)
		checkError(err)
		n := int32(nTemp)

		mTemp, err := strconv.ParseInt(nm[1], 10, 64)
		checkError(err)
		m := int32(mTemp)

		result := towerBreakers(n, m)

		fmt.Fprintf(writer, "%d\n", result)
	}

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
