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
 * Complete the 'maximumToys' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY prices
 *  2. INTEGER k
 */

func maximumToys(prices []int32, k int32) int32 {
	// Write your code here
	return maximumToysImpl(prices, k)
}

func maximumToysImpl(prices []int32, k int32) int32 {
	// generate possible combinations.
	combinations := genPossibleCombinations(prices)
	// clean up combinations
	combinations = removeDup(combinations)
	// return max items.
	return getMaxItemNum(combinations, k)
}

func getMaxItemNum(arr [][]int32, k int32) int32 {
	kUint64 := uint64(k)
	max := 0
	for _, a := range arr {
		sum := sumArrInt32(a)
		if sum <= kUint64 && max < len(a) {
			max = len(a)
		}
	}
	return int32(max)
}

func sumArrInt32(arr []int32) uint64 {
	sum := uint64(0)
	for i := 0; i < len(arr); i++ {
		sum += uint64(arr[i])
	}
	return sum
}

func genPossibleCombinations(prices []int32) [][]int32 {
	result := [][]int32{}
	for i, p := range prices {
		result = append(result, []int32{p})

		// Get the rest of prices.
		// append() modify source slice.
		tmpPrices := append([]int32{}, prices...)
		subPrices := append(tmpPrices[:i], tmpPrices[i+1:]...)
		for _, subResultOne := range genPossibleCombinations(subPrices) {
			subResult := append([]int32{p}, subResultOne...)
			result = append(result, subResult)
		}

	}
	return result
}

func sortArrayAscInt32(arr []int32) []int32 {
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	return arr
}

func arrayInt32ToString(arr []int32) []string {
	result := []string{}
	for _, v := range arr {
		result = append(result, fmt.Sprintf("%d", v))
	}
	return result
}

func removeDup(arr [][]int32) [][]int32 {
	elemMap := map[string][]int32{}
	for _, a := range arr {
		sorted := sortArrayAscInt32(a)
		elemMap[strings.Join(arrayInt32ToString(sorted), "_")] = a
	}

	result := [][]int32{}
	for _, v := range elemMap {
		result = append(result, v)
	}

	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	pricesTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var prices []int32

	for i := 0; i < int(n); i++ {
		pricesItemTemp, err := strconv.ParseInt(pricesTemp[i], 10, 64)
		checkError(err)
		pricesItem := int32(pricesItemTemp)
		prices = append(prices, pricesItem)
	}

	result := maximumToys(prices, k)

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
