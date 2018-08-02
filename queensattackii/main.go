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

// Complete the queensAttack function below.
func queensAttack(n int32, k int32, r_q int32, c_q int32, obstacles [][]int32) int32 {

	if k == 0 {
		return getCleanMoves(n, r_q, c_q)
	}

	// leftup := getleftup(n, r_q, c_q)
	// left := getleft(n, r_q, c_q)
	// leftdown := getleftdown(n, r_q, c_q)
	// down := getdown(n, r_q, c_q)
	// rigthdown := getrigthdown(n, r_q, c_q)
	// rigth := getrigth(n, r_q, c_q)
	// rigthup := getrigthup(n, r_q, c_q)
	// up := getup(n, r_q, c_q)

	return 1
}

func getCleanMoves(n int32, x int32, y int32) int32 {
	leftup := getleftup(n, x, y)
	left := getleft(n, x, y)
	leftdown := getleftdown(n, x, y)
	down := getdown(n, x, y)
	rigthdown := getrigthdown(n, x, y)
	rigth := getrigth(n, x, y)
	rigthup := getrigthup(n, x, y)
	up := getup(n, x, y)

	dleftup := []int32{int32(math.Abs(float64(x - leftup[0]))), int32(math.Abs(float64(y - leftup[1])))}
	dleft := []int32{int32(math.Abs(float64(x - left[0]))), int32(math.Abs(float64(y - left[1])))}
	dleftdown := []int32{int32(math.Abs(float64(x - leftdown[0]))), int32(math.Abs(float64(y - leftdown[1])))}
	ddown := []int32{int32(math.Abs(float64(x - down[0]))), int32(math.Abs(float64(y - down[1])))}
	drigthdown := []int32{int32(math.Abs(float64(x - rigthdown[0]))), int32(math.Abs(float64(y - rigthdown[1])))}
	drigth := []int32{int32(math.Abs(float64(x - rigth[0]))), int32(math.Abs(float64(y - rigth[1])))}
	drigthup := []int32{int32(math.Abs(float64(x - rigthup[0]))), int32(math.Abs(float64(y - rigthup[1])))}
	dup := []int32{int32(math.Abs(float64(x - up[0]))), int32(math.Abs(float64(y - up[1])))}

	moves := (dleftup[0] + dleftup[1]) / 2
	moves += dleft[0] + dleft[1]
	moves += (dleftdown[0] + dleftdown[1]) / 2
	moves += ddown[0] + ddown[1]
	moves += (drigthdown[0] + drigthdown[1]) / 2
	moves += drigth[0] + drigth[1]
	moves += (drigthup[0] + drigthup[1]) / 2
	moves += dup[0] + dup[1]

	return moves
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

	if x-y < 0 { //inverso
		r := x
		x = y
		y = r
	}

	//SI(H2*2 > H8,H8,H2*2)
	if y == 1 { //No hay esquina
		xx = x
	} else if x-y == 0 { //Diagonal principal
		if x+y < n {
			xx = x + y - 1
		} else {
			xx = n
		}
	} else if x*2 > n { //
		if x+y-1 > n {
			xx = n
		} else {
			xx = x + y - 1
		}
	} else {
		xx = y * 2
	}

	//=SI(H2=H8,I2,SI(H2*2>H8,I2-(H8-H2),I2*2))
	if y == 1 {
		yy = 1
	} else if x == n {
		yy = y
	} else if x-y == 0 {
		if x+y-1 <= n {
			yy = x - y + 1
		} else {
			yy = y - (n - x)
		}
	} else if x*2 > n {
		if x+y > n {
			yy = y - (n - x)
		} else {
			yy = y - (n - x) + 1
		}
	} else {
		yy = x - y
	}

	// if inverse {
	// 	return []int32{yy, xx}
	// }

	return []int32{xx, yy}
}

func getleft(n, x, y int32) []int32 {
	return []int32{x, 1}
}

func getleftdown(n, x, y int32) []int32 {
	invert := false
	if x-y < 0 { //inverso
		r := x
		x = y
		y = r
		invert = true
	}

	//=(H2-I2)+1
	var xx int32
	if x == 1 {
		xx = 1
	} else if x-y == 0 {
		xx = 1
	} else {
		xx = x - y + 1
	}
	//=H2-I2
	var yy int32
	if y == 1 {
		yy = 1
	} else if x == 1 {
		yy = y
	} else if x-y >= 0 {
		yy = 1
	}

	if invert {
		return []int32{yy, xx}
	}
	return []int32{xx, yy}
}

func getdown(n, x, y int32) []int32 {
	return []int32{1, y}
}

func getrigthdown(n, x, y int32) []int32 {
	r := getleftup(n, x, y)
	return []int32{r[1], r[0]}
}

func getrigth(n, x, y int32) []int32 {
	return []int32{x, n}
}

func getrigthup(n, x, y int32) []int32 {
	// =SI(H2*2 > H8,H8,H2*2)

	if x == n {
		return []int32{x, y}
	}

	if y == n {
		return []int32{x, y}
	}

	if y-x == 0 {
		return []int32{n, n}
	}

	invert := false
	if x-y < 0 { //inverso
		r := x
		x = y
		y = r
		invert = true
	}

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

	if invert {
		return []int32{yy, xx}
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
