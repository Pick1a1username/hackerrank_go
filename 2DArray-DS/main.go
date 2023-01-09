package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	HourGlassWidth  int = 3
	HourGlassHeight int = 3
)

/*
 * Complete the 'hourglassSum' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.
 */

func hourglassSum(arr [][]int32) int32 {
	// Write your code here
	sums := []int32{}

	arrColNum := len(arr[0])
	arrRowNum := len(arr)

	for row := 0; row < arrRowNum-HourGlassHeight+1; row++ {
		for col := 0; col < arrColNum-HourGlassWidth+1; col++ {
			hourGlassElems, err := getHourGlassElems(arr, row, col)
			if err != nil {
				return 0
			}
			sum := sumInt32Arr(hourGlassElems)
			sums = append(sums, sum)
		}
	}

	sort.Slice(sums, func(i, j int) bool { return sums[i] > sums[j] })
	return sums[0]
}

func getHourGlassElems(arr [][]int32, row, col int) ([]int32, error) {
	// check the number of columns.
	var elemNum int
	for i, row := range arr {
		if i == 0 {
			elemNum = len(row)
			continue
		}
		if elemNum != len(row) {
			return nil, errors.New("invalid array")
		}
	}

	arrColNum := elemNum
	arrRowNum := len(arr)

	// check the target hourglass is valid.
	if arrColNum-col < HourGlassWidth-1 {
		return nil, errors.New("out of range of hourglass width")
	}
	if arrRowNum-row < HourGlassHeight-1 {
		return nil, errors.New("out of range of hourglass height")
	}

	var result []int32

	result = append(result, arr[row][col])
	result = append(result, arr[row][col+1])
	result = append(result, arr[row][col+2])
	result = append(result, arr[row+1][col+1])
	result = append(result, arr[row+2][col])
	result = append(result, arr[row+2][col+1])
	result = append(result, arr[row+2][col+2])

	return result, nil
}

func sumInt32Arr(arr []int32) int32 {
	var result int32 = 0
	for _, e := range arr {
		result += e
	}
	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	var arr [][]int32
	for i := 0; i < 6; i++ {
		arrRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var arrRow []int32
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arrRow = append(arrRow, arrItem)
		}

		if len(arrRow) != 6 {
			panic("Bad input")
		}

		arr = append(arr, arrRow)
	}

	result := hourglassSum(arr)

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
