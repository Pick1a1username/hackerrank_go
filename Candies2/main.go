package main

import (
	"math"
	"os"
	"strconv"
	"strings"
)

var INF int32 = int32(math.Pow(10, 9)) // a number larger than all ratings
/*
 * Complete the 'candies' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER_ARRAY arr
 */
func candies(n int32, arr []int32) int64 {
	// Write your code here
	// https://www.hackerrank.com/challenges/candies/editorial?isFullScreen=true&h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=dynamic-programming
	candies := make([]int64, n+1)
	arr = append([]int32{INF}, arr...)
	arr = append(arr, INF)

	// populate 'valleys'
	for i := 1; i < int(n)+1; i++ {
		if (arr[i-1] >= arr[i]) && (arr[i] <= arr[i+1]) {
			candies[i] = 1
		}
	}
	saveCandies(candies, "after_valleys.txt")

	// populate 'rises'
	for i := 1; i < int(n)+1; i++ {
		if (arr[i-1] < arr[i]) && (arr[i] <= arr[i+1]) {
			candies[i] = candies[i-1] + 1
		}
	}
	saveCandies(candies, "after_rises.txt")

	// populate 'falls'
	for i := n; i > 0; i-- {
		if (arr[i-1] >= arr[i]) && (arr[i] > arr[i+1]) {
			candies[i] = candies[i+1] + 1
		}
	}
	saveCandies(candies, "after_falls.txt")

	// populate 'peaks'
	for i := 1; i < int(n)+1; i++ {
		if (arr[i-1] < arr[i]) && (arr[i] > arr[i+1]) {
			candies[i] = maxInt64(candies[i-1], candies[i+1]) + 1
		}
	}
	saveCandies(candies, "after_peaks.txt")

	return sumInt64(candies)
}

func maxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	if a < b {
		return b
	}
	return a
}

func sumInt64(arr []int64) int64 {
	sum := int64(0)
	for _, i := range arr {
		sum = sum + i
	}
	return sum
}

func saveCandies(candies []int64, savePath string) {
	candiesStrArr := make([]string, len(candies))
	for i := 0; i < len(candies); i++ {
		candiesStrArr[i] = strconv.Itoa(int(candies[i]))
	}
	candiesStr := strings.Join(candiesStrArr, "\n")
	f, err := os.Create(savePath)
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()

	_, err = f.WriteString(candiesStr)
	if err != nil {
		panic(err.Error())
	}

	err = f.Sync()
	if err != nil {
		panic(err.Error())
	}
}
