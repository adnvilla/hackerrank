package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the queensAttack function below.
func queensAttack(n int32, k int32, r_q int32, c_q int32, obstacles [][]int32) int32 {

	leftup := getleftup(n, r_q, c_q)
	left := getleft(n, r_q, c_q)
	leftdown := getleftdown(n, r_q, c_q)
	down := getdown(n, r_q, c_q)
	rigthdown := getrigthdown(n, r_q, c_q)
	rigth := getrigth(n, r_q, c_q)
	rigthup := getrigthup(n, r_q, c_q)
	up := getup(n, r_q, c_q)

	fmt.Println(leftup)
	fmt.Println(left)
	fmt.Println(leftdown)
	fmt.Println(down)
	fmt.Println(rigthdown)
	fmt.Println(rigth)
	fmt.Println(rigthup)
	fmt.Println(up)

	return 1
}

func getFinalPoints(n int32, x int32, y int32) [][]int32 {

	leftup := getleftup(n, x, y)
	left := getleft(n, x, y)
	leftdown := getleftdown(n, x, y)
	down := getdown(n, x, y)
	rigthdown := getrigthdown(n, x, y)
	rigth := getrigth(n, x, y)
	rigthup := getrigthup(n, x, y)
	up := getup(n, x, y)

	return [][]int32{
		leftup,
		left,
		leftdown,
		down,
		rigthdown,
		rigth,
		rigthup,
		up,
	}
}

func getleftup(n, x, y int32) []int32 {
	var xx int32
	var yy int32

	//SI(H2*2 > H8,H8,H2*2)
	if x == 1 {
		xx = 1
	} else if x*2 > n {
		xx = n
	} else {
		xx = y * 2
	}

	//=SI(H2=H8,I2,SI(H2*2>H8,I2-(H8-H2),I2*2))
	if y == 1 {
		yy = 1
	} else if x == n {
		yy = y
	} else if x*2 > n {
		yy = y - (n - x)
	} else {
		yy = y * 2
	}

	return []int32{xx, yy}
}

func getleft(n, x, y int32) []int32 {
	return []int32{x, 1}
}

func getleftdown(n, x, y int32) []int32 {
	//=(H2-I2)+1
	var xx int32
	if x == 1 {
		xx = 1
	} else {
		xx = x - y + 1
	}
	//=H2-I2
	var yy int32
	if y == 1 {
		yy = 1
	} else {
		yy = x - y
	}
	return []int32{xx, yy}
}

func getdown(n, x, y int32) []int32 {
	return []int32{1, y}
}

func getrigthdown(n, x, y int32) []int32 {
	//=SI(I2=H8,H2,SI(I2*2>H8,H2-(H8-I2),H2*2))
	var xx int32
	if x == 1 {
		xx = 1
	} else if y == n {
		xx = x
	} else if y*2 > n {
		xx = x - (n - y)
	} else {
		xx = x * 2
	}

	//=SI(I2*2 > H8, H8, I2*2)
	var yy int32
	if xx == 1 {
		yy = 1
	} else if y*2 > n {
		yy = n
	} else {
		yy = y * 2
	}

	return []int32{xx, yy}
}

func getrigth(n, x, y int32) []int32 {
	return []int32{x, n}
}

func getrigthup(n, x, y int32) []int32 {
	// =SI(H2*2 > H8,H8,H2*2)
	var xx int32
	if y == 1 {
		xx = n
	} else if x-y < 0 {
		xx = n - y + 1
	} else {
		xx = n
	}

	//=SI(H2=H8,I2,SI(H2*2>H8,I2+(H8-H2),H2*2))
	var yy int32
	if x == 1 {
		yy = n
	} else if x-y < 0 {
		yy = n
	} else {
		yy = y + (n - x)
	}

	return []int32{xx, yy}
}

func getup(n, x, y int32) []int32 {
	return []int32{n, y}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nk := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nk[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(nk[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	r_qC_q := strings.Split(readLine(reader), " ")

	r_qTemp, err := strconv.ParseInt(r_qC_q[0], 10, 64)
	checkError(err)
	r_q := int32(r_qTemp)

	c_qTemp, err := strconv.ParseInt(r_qC_q[1], 10, 64)
	checkError(err)
	c_q := int32(c_qTemp)

	var obstacles [][]int32
	for i := 0; i < int(k); i++ {
		obstaclesRowTemp := strings.Split(readLine(reader), " ")

		var obstaclesRow []int32
		for _, obstaclesRowItem := range obstaclesRowTemp {
			obstaclesItemTemp, err := strconv.ParseInt(obstaclesRowItem, 10, 64)
			checkError(err)
			obstaclesItem := int32(obstaclesItemTemp)
			obstaclesRow = append(obstaclesRow, obstaclesItem)
		}

		if len(obstaclesRow) != int(2) {
			panic("Bad input")
		}

		obstacles = append(obstacles, obstaclesRow)
	}

	result := queensAttack(n, k, r_q, c_q, obstacles)

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
