package main

import "sort"

/*
 * Complete the 'pairs' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER k
 *  2. INTEGER_ARRAY arr
 */

func pairs(k int32, arr []int32) int32 {
	// Write your code here
	// Copy arr.
	sortedArr := append([]int32{}, arr...)

	// Sort the array.
	sort.Slice(sortedArr, func(i, j int) bool { return sortedArr[i] < sortedArr[j] })

	count := int32(0)
	// for loop through the array.
	for i := 0; i < len(sortedArr)-1; i++ {
		// for loop through the rest of the array
		for j := i + 1; j < len(sortedArr); j++ {
			diff := sortedArr[j] - sortedArr[i]
			// If the difference is bigger than k, there's no pair.
			if diff > k {
				break
			}
			// If the difference is equal to k, count it.
			if diff == k {
				count++
			}

		}
	}
	// return the result.
	return count
}
