package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/profile"
)

// Complete the minimumLoss function below.
func minimumLoss(price []int64) int32 {

	f := int64(len(price))

	if f > 100000 {
		return 0
	}

	result := int32(math.MaxInt32)

	if len(price) == 2 {
		w := int32(price[0] - price[1])
		if w < 0 {
			return result
		}

		if result > w {
			result = w
		}

		return result
	}

	x := price[0]
	for i := 1; i < len(price); i++ {

		v := int32(x - price[i])

		if v < 0 {
			continue
		}

		if result > v {
			result = v
		}

		if result == 1 {
			return result
		}
	}

	next := make([]int64, int64(len(price)-1))
	for i := int64(1); i < int64(len(price)); i++ {
		next[i-1] = price[i]
	}

	c := minimumLoss(next)

	if c == 1 {
		return c
	}

	if c < 0 {
		return result
	}

	if result > c {
		result = c
	}

	return result

}

func main() {

	rand.Seed(time.Now().UTC().UnixNano())
	// CPU profiling by default
	defer profile.Start().Stop()
	start := time.Now()
	//...

	price := make([]int64, 100000)
	for i := 0; i < len(price)-1; i++ {
		r := rand.Int63()

		price[i] = r
	}

	r := minimumLoss(price)
	elapsed := time.Since(start)
	log.Printf("time: %s", elapsed)

	fmt.Println(r)
	// fmt.Println(r)
}

func main2() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

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
