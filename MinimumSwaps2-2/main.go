package main

// Complete the minimumSwaps function below.
func minimumSwaps(arr []int32) int32 {
	// https://www.hackerrank.com/challenges/minimum-swaps-2/forum/comments/473577?isFullScreen=true&h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=arrays
	indexMap := map[int32]int{}
	for i, n := range arr {
		indexMap[n] = i
	}

	swaps := int32(0)

	for i := 0; i < len(arr); i++ {
		correctValue := int32(i + 1)
		currentValue := arr[i]
		if currentValue != correctValue {
			toSwapIndex := indexMap[correctValue]
			arr[toSwapIndex], arr[i] = arr[i], arr[toSwapIndex]
			indexMap[currentValue] = toSwapIndex
			indexMap[correctValue] = i
			swaps += 1
		}
	}

	return swaps
}
