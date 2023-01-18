package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
 * Complete the 'minimumAbsoluteDifference' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func minimumAbsoluteDifference(arr []int32) int32 {
	// Write your code here
	// Generate pairs.
	pairs := genTwoPairs(arr)

	// Calculate differences.
	pairsLen := len(pairs)
	diffs := make([]int32, pairsLen)
	for pi := 0; pi < pairsLen; pi++ {
		diffs[pi] = absInt32(pairs[pi][0], pairs[pi][1])
	}

	// Get the minimum difference.
	sort.Sort(Int32List(diffs))

	// Return.
	return diffs[0]
}

type Int32List []int32

func (l Int32List) Len() int           { return len(l) }
func (l Int32List) Less(i, j int) bool { return l[i] < l[j] }
func (l Int32List) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }

func absInt32(a, b int32) int32 {
	if a > b {
		return a - b
	}
	return b - a
}

func genTwoPairs(arr []int32) [][]int32 {
	if len(arr) < 1 {
		return [][]int32{}
	}

	if len(arr) < 3 {
		return [][]int32{arr}
	}

	// Create players and set start positions.
	player1 := 0
	player2 := 1
	player1FinishLine := len(arr) - 2
	player2FinishLine := len(arr) - 1

	// Calculate the number of pairs.
	pairNum := calculatePairNum(len(arr))

	// Make the players move and get possible pairs.
	result := make([][]int32, pairNum)
	for pi := 0; pi < pairNum; pi++ {
		result[pi] = []int32{arr[player1], arr[player2]}

		if player1 == player1FinishLine && player2 == player2FinishLine {
			break
		}

		if player1+1 == player2 {
			player1 = 0
			player2 += 1
		} else {
			player1 += 1
		}
	}

	// Return.
	return result
}

func calculatePairNum(n int) int {
	/*
	   len(arr) = len(result)
	   1 = 1
	   2 = 1
	   3 = 3
	   4 = 6   3 + 3
	   5 = 10   6 + 4
	   6 = 15   10 + 5
	   (...)

	*/
	if n < 3 {
		return 1
	}
	currN := 4
	prevNum := 3
	pairNum := 3

	if n > 3 {
		for c := 4; c <= n; c++ {
			pairNum = prevNum + currN - 1
			prevNum = pairNum
			currN += 1
		}
	}
	return pairNum
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	result := minimumAbsoluteDifference(arr)

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
