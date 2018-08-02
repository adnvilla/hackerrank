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
	fmt.Println(price)
	if len(price) == 0 {
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

func MinimumLoss2(price []int64) int64 {

	if len(price) == 0 {
		return 0
	}

	if len(price) == 1 {
		return 0
	}

	value := price[0]

	search := price[1:]

	var order []int64

	if len(search) == 1 {
		order = search
	} else {
		// order = append(order, search...)
		// Quicksort2(order)
		order = QuickSort(search)
		fmt.Println(order)
	}

	if value-order[0] == 1 {
		return 1
	}

	i := Search(value, order)

	var actual = order[len(order)-1]
	if value-order[i] > 0 {
		actual = value - order[i]
	}

	if value-order[i] == 1 {
		return 1
	}

	next := MinimumLoss2(search)

	if next == 0 {
		return actual
	}

	if actual > next {
		actual = next
	}

	return actual
}

func Search(value int64, price []int64) int64 {
	if value < price[0] {
		return 0
	}
	if value > price[len(price)-1] {
		return int64(len(price) - 1)
	}

	lo := int64(0)
	hi := int64(len(price) - 1)

	for lo <= hi {
		mid := (hi + lo) / 2

		if value < price[mid] {
			hi = mid - 1
		} else if value > price[mid] {
			lo = mid + 1
		} else {
			return mid
		}
	}
	// lo == hi + 1
	// return price[hi]
	return hi
	// if (price[lo] - value) > (value - price[hi]) {
	// 	return price[lo]
	// } else {
	// 	return price[hi]
	// }
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

//  Quicksort sorts the elements of v in ascending order.
func Quicksort2(v []int64) {
	if int64(len(v)) < 20 {
		InsertionSort(v)
		return
	}
	p := Pivot(v)
	low, high := Partition(v, p)
	Quicksort2(v[:low])
	Quicksort2(v[high:])
}

func Pivot(v []int64) int64 {
	n := int64(len(v))
	return Median(v[rand.Int63n(n)],
		v[rand.Int63n(n)],
		v[rand.Int63n(n)])
}

func Median(a, b, c int64) int64 {
	if a < b {
		switch {
		case b < c:
			return b
		case a < c:
			return c
		default:
			return a
		}
	}
	switch {
	case a < c:
		return a
	case b < c:
		return c
	default:
		return b
	}
}

// Partition reorders the elements of v so that:
// - all elements in v[:low] are less than p,
// - all elements in v[low:high] are equal to p,
// - all elements in v[high:] are greater than p.
func Partition(v []int64, p int64) (low, high int64) {
	low, high = 0, int64(len(v))
	for mid := int64(0); mid < high; {
		// Invariant:
		//  - v[:low] < p
		//  - v[low:mid] = p
		//  - v[mid:high] are unknown
		//  - v[high:] > p
		//
		//         < p       = p        unknown       > p
		//     -----------------------------------------------
		// v: |          |          |a            |           |
		//     -----------------------------------------------
		//                ^          ^             ^
		//               low        mid           high
		switch a := v[mid]; {
		case a < p:
			v[mid] = v[low]
			v[low] = a
			low++
			mid++
		case a == p:
			mid++
		default: // a > p
			v[mid] = v[high-1]
			v[high-1] = a
			high--
		}
	}
	return
}

func InsertionSort(v []int64) {
	for j := int64(1); j < int64(len(v)); j++ {
		// Invariant: v[:j] contains the same elements as
		// the original slice v[:j], but in sorted order.
		key := v[j]
		i := j - 1
		for i >= 0 && v[i] > key {
			v[i+1] = v[i]
			i--
		}
		v[i+1] = key
	}
}
