package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "errors"
)

const (
    HourGlassWidth int = 3
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
    return 0
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
    if arrColNum - col < HourGlassWidth - 1 {
        return nil, errors.New("out of range of hourglass width")
    }
    if arrRowNum - row < HourGlassHeight - 1 {
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

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    var arr [][]int32
    for i := 0; i < 6; i++ {
        arrRowTemp := strings.Split(strings.TrimRight(readLine(reader)," \t\r\n"), " ")

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