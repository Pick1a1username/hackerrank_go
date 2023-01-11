package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'twoStrings' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING s1
 *  2. STRING s2
 */

func twoStrings(s1 string, s2 string) string {
	// Write your code here
	return twoStringsImpl(s1, s2)
}

func twoStringsImpl(s1 string, s2 string) string {
	s1Map := convertStringToMap(s1)
	s2Map := convertStringToMap(s2)

	checkResult := isMatchRune(s1Map, s2Map)

	if checkResult {
		return "YES"
	} else {
		return "NO"
	}
}

func isMatchRune(m1, m2 map[rune]int) bool {
	if len(m1) > len(m2) {
		return isMatchRuneImpl(m1, m2)
	} else {
		return isMatchRuneImpl(m2, m1)
	}
}

func isMatchRuneImpl(m1, m2 map[rune]int) bool {
	for k1, _ := range m1 {
		if _, ok := m2[k1]; ok {
			return true
		}
	}
	return false
}

func convertStringToMap(s string) map[rune]int {
	result := map[rune]int{}
	for _, c := range s {
		result[c] += 1
	}
	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		s1 := readLine(reader)

		s2 := readLine(reader)

		result := twoStrings(s1, s2)

		fmt.Fprintf(writer, "%s\n", result)
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
