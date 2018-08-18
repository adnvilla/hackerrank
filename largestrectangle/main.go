package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func largestRectangle(h []int32) int64 {
	maxArea := int32(0)

	for i := int32(0); i < int32(len(h)); i++ {
		x := h[i]

		if maxArea < x {
			maxArea = x
		}
		// maxArea = Max(maxArea, x)

		for j := i - 1; j >= 0; j-- {
			w := (i - j + 1)

			if x > h[j] {
				x = h[j]
			}
			//x = Min(x, h[j])

			if maxArea < x*w {
				maxArea = x * w
			}
			// maxArea = Max(maxArea, x*w)
		}
	}

	return int64(maxArea)
}

func largestRectangle2(h []int32) int64 {
	var s []int32
	var maxArea, tp, areaWithTop, i int32

	for i < int32(len(h)) {
		if len(s) == 0 || h[s[len(s)-1]] <= h[i] {
			s = append(s, i)
			i++
		} else {
			tp = s[len(s)-1]
			s = s[0 : len(s)-1]
			var w int32
			if len(s) == 0 {
				w = i
			} else {
				w = i - s[len(s)-1] - 1
			}

			areaWithTop = h[tp] * w

			if maxArea < areaWithTop {
				maxArea = areaWithTop
			}
		}
	}

	for len(s) > 0 {
		tp = s[len(s)-1]
		s = s[0 : len(s)-1]

		var w int32
		if len(s) == 0 {
			w = i
		} else {
			w = i - s[len(s)-1] - 1
		}

		areaWithTop = h[tp] * w

		if maxArea < areaWithTop {
			maxArea = areaWithTop
		}
	}

	return int64(maxArea)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	hTemp := strings.Split(readLine(reader), " ")

	var h []int32

	for i := 0; i < int(n); i++ {
		hItemTemp, err := strconv.ParseInt(hTemp[i], 10, 64)
		checkError(err)
		hItem := int32(hItemTemp)
		h = append(h, hItem)
	}

	result := largestRectangle(h)

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
