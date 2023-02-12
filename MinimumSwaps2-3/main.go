package main

// Complete the minimumSwaps function below.
func minimumSwaps(arr []int32) int32 {
	count := int32(0)
	// go through arr
	for i := 0; i < len(arr); i++ {
		// if value is not supposed to be there, swap it with appropriate one.
		if arr[i] != int32(i+1) {
			// https://www.hackerrank.com/challenges/minimum-swaps-2/forum/comments/458759?isFullScreen=true&h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=arrays
			swapWith := i
			for arr[swapWith] != int32(i+1) {
				swapWith++
			}
			arr[i], arr[swapWith] = arr[swapWith], arr[i]
			count++
		}
	}

	// return the count
	return count
}
