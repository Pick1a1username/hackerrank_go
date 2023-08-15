package main

import (
	"os"
	"strconv"
	"strings"
)

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
	// Initialize the number of candies each student should get.
	students := make([]int64, n)

	// This loop doesn't work properly if there's consecutive incline or decline.
	for i := 0; i < int(n)-1; i++ {
		if students[i] == 0 {
			students[i]++
		}
		if students[i+1] == 0 {
			students[i+1]++
		}
		if arr[i] < arr[i+1] && students[i] >= students[i+1] {
			students[i+1] = students[i] + 1
		}
		if arr[i] > arr[i+1] && students[i] <= students[i+1] {
			students[i] = students[i+1] + 1
		}
	}
	saveCandies(students, "result.txt")

	count := int64(0)
	for _, n := range students {
		count = count + n
	}
	return count
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
