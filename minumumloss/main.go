package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/profile"
	_ "go.uber.org/automaxprocs"
)

// Complete the minimumLoss function below.
func MinimumLoss(price []int64) int32 {
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

	c := MinimumLoss(next)

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

	// rand.Seed(time.Now().UTC().UnixNano())
	// CPU profiling by default
	defer profile.Start(profile.CPUProfile).Stop()
	// start := time.Now()
	//...

	price := GenerateTest(100)

	r := MinimumLoss(price)
	// elapsed := time.Since(start)
	// log.Printf("time: %s", elapsed)

	fmt.Println(r)
	// fmt.Println(r)
}

func GenerateTest(n int) []int64 {
	price := make([]int64, n)
	for i := 0; i < len(price)-1; i++ {
		r := rand.Int63()

		price[i] = r
	}
	return price
}

func main2() {
	reader := bufio.NewReaderSize(os.Stdin, 2048*2048)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 2048*2048)

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

	result := MinimumLoss(price)

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
