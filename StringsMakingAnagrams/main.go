package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
 * Complete the 'makeAnagram' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. STRING a
 *  2. STRING b
 */

func makeAnagram(a string, b string) int32 {
	// Write your code here
	aMap := convertStringToMap(a)
	bMap := convertStringToMap(b)

	deleteCount := 0
	for k, v := range aMap {
		bValue := bMap[k]
		if bValue > v {
			deleteCount += bValue - v
			bMap[k] = v
		} else if bValue < v {
			deleteCount += v - bValue
			aMap[k] = bValue
		}
	}

	for k, v := range bMap {
		aValue := aMap[k]
		if aValue > v {
			deleteCount += aValue - v
			aMap[k] = v
		} else if aValue < v {
			deleteCount += v - aValue
			bMap[k] = aValue
		}
	}

	return int32(deleteCount)
}

func convertStringToMap(s string) map[rune]int {
	result := map[rune]int{}
	for _, v := range s {
		result[v] += 1
	}
	return result
}
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	a := readLine(reader)

	b := readLine(reader)

	res := makeAnagram(a, b)

	fmt.Fprintf(writer, "%d\n", res)

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
