package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
)

/*
 * Complete the 'minimumBribes' function below.
 *
 * The function accepts INTEGER_ARRAY q as parameter.
 */

const (
	MaxBribe = 2
)

func minimumBribes(q []int32) {
	// Write your code here
	fmt.Printf("%s\n", minimumBribesImpl(q))
}

func minimumBribesImpl(q []int32) string {
	// Find persons who bribed.
	bribers := genNoBribeQ(len(q)) // Contains values
	qInt := convertInt32ToIntSlice(&q)

	// Check how many times each person bribed.
	bribeNum := countBribeNum(genNoBribeQ(len(qInt)), qInt, bribers)

	// If there's no non-chaotic situation, return.
	if isChaotic(bribeNum) {
		return "Too chaotic"
	}

	// Return the minimum number of bribing.
	return strconv.Itoa(sumValMap(bribeNum))
}

func sumValMap(m map[int]int) int {
	result := 0
	for _, v := range m {
		result += v
	}
	return result
}

func isChaotic(m map[int]int) bool {
	for _, v := range m {
		if v > MaxBribe {
			return true
		}
	}
	return false
}

func convertInt32ToIntSlice(sl *[]int32) []int {
	result := []int{}
	for i := 0; i < len(*sl); i++ {
		result = append(result, int((*sl)[i]))
	}
	return result
}

func findIdxByValInt(slPtr *[]int, v int) (int, error) {
	for i := 0; i < len(*slPtr); i++ {
		if (*slPtr)[i] == v {
			return i, nil
		}
	}
	return -1, errors.New("not found")
}

func countBribeNum(currQ []int, targetQ []int, bribeOrder []int) map[int]int {
	result := map[int]int{}

	for _, b := range bribeOrder {
		result[b] = 0
		beforeIdx, err := findIdxByValInt(&currQ, b)
		if err != nil {
			panic(err.Error())
		}
		afterIdx, err := findIdxByValInt(&targetQ, b)
		if err != nil {
			panic(err.Error())
		}
		bribeNum := beforeIdx - afterIdx
		if bribeNum < 0 {
			// Does nothing if a briber is trying to go back.
			continue
		}
		// Remove the briber temporarily.
		currQ = removeElemArrayInt(&currQ, beforeIdx)
		// Insert the briber.
		currQ = insertElemArrayInt(&currQ, afterIdx, b)

		result[b] += bribeNum
	}

	if !reflect.DeepEqual(currQ, targetQ) {
		tmpResult := countBribeNum(currQ, targetQ, bribeOrder)
		result = addMaps(result, tmpResult)
	}

	return result
}

// // Only add keys in a.
func addMaps(a, b map[int]int) map[int]int {
	result := map[int]int{}
	for k, v := range a {
		result[k] = v + b[k]
	}
	return result
}

func removeElemArrayInt(arr *[]int, i int) []int {
	tmpLeftQ := append([]int{}, (*arr)[:i]...)
	tmpRightQ := append([]int{}, (*arr)[i+1:]...)
	return append(tmpLeftQ, tmpRightQ...)
}

func insertElemArrayInt(arr *[]int, i, v int) []int {
	result := []int{}
	tmpLeftQ := append([]int{}, (*arr)[:i]...)
	tmpRightQ := append([]int{}, (*arr)[i:]...)
	result = append(tmpLeftQ, v)
	result = append(result, tmpRightQ...)
	return result
}

// // Bribers are expected to bribe after someone's bribing finished.
// // It doesn't return by index.
// func genPossibleSituations(bribers []int) [][]int {
// 	briberNum := len(bribers)
// 	result := [][]int{}
// 	for i, b := range bribers {
// 		subResultBase := []int{}
// 		subResultBase = append(subResultBase, b)
// 		if briberNum == 1 {
// 			result = append(result, subResultBase)
// 			return result
// 		}
// 		// append() modify source slice.
// 		tmpBriber := append([]int{}, bribers...)
// 		subBribers := append(tmpBriber[:i], tmpBriber[i+1:]...)
// 		for _, subResultOne := range genPossibleSituations(subBribers) {
// 			subResult := append(subResultBase, subResultOne...)
// 			result = append(result, subResult)
// 		}

// 	}
// 	return result
// }

func genNoBribeQ(n int) []int {
	result := []int{}
	for i := 1; i <= n; i++ {
		result = append(result, i)
	}
	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int32(nTemp)

		qTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		var q []int32

		for i := 0; i < int(n); i++ {
			qItemTemp, err := strconv.ParseInt(qTemp[i], 10, 64)
			checkError(err)
			qItem := int32(qItemTemp)
			q = append(q, qItem)
		}

		minimumBribes(q)
	}
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
