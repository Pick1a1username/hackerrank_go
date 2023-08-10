package main

const INF int32 = 10 ^ 9 // a number larger than all ratings
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

	// populate 'rises'
	for i := 1; i < int(n)+1; i++ {
		if (arr[i-1] < arr[i]) && (arr[i] <= arr[i+1]) {
			candies[i] = candies[i-1] + 1
		}
	}

	// populate 'falls'
	for i := 1; i < int(n)+1; i++ {
		if (arr[i-1] >= arr[i]) && (arr[i] > arr[i+1]) {
			candies[i] = candies[i+1] + 1
		}
	}

	// populate 'peaks'
	for i := 1; i < int(n)+1; i++ {
		if (arr[i-1] < arr[i]) && (arr[i] > arr[i+1]) {
			candies[i] = maxInt64(candies[i-1], candies[i+1]) + 1
		}
	}

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
