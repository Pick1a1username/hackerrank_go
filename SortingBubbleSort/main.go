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
 * Complete the 'countSwaps' function below.
 *
 * The function accepts INTEGER_ARRAY a as parameter.
 */

func countSwaps(a []int32) {
	// Write your code here
	count, first, last := countSwapsImpl(a)
	fmt.Printf("%s\n", count)
	fmt.Printf("%s\n", first)
	fmt.Printf("%s\n", last)
}

func countSwapsImpl(a []int32) (string, string, string) {
	count := 0
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-1; j++ {
			// Swap adjacent elements if they are in decreasing order
			if a[j] > a[j+1] {
				// swap(a[j], a[j+1])
				tmp := a[j]
				a[j] = a[j+1]
				a[j+1] = tmp
				count++
			}
		}
	}
	return genSwapResult(count, int(a[0]), int(a[len(a)-1]))
}

func genSwapResult(count, first, last int) (string, string, string) {
	countS := fmt.Sprintf("Array is sorted in %d swaps.", count)
	firstS := fmt.Sprintf("First Element: %d", first)
	lastS := fmt.Sprintf("Last Element: %d", last)
	return countS, firstS, lastS
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var a []int32

	for i := 0; i < int(n); i++ {
		aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
		checkError(err)
		aItem := int32(aItemTemp)
		a = append(a, aItem)
	}

	countSwaps(a)
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
