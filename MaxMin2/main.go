package main

import "sort"

func maxMin(k int32, arr []int32) int32 {
	// https://www.hackerrank.com/challenges/angry-children/editorial?isFullScreen=true&h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=greedy-algorithms
	kInt := int(k)
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	ans := arr[k-1] - arr[0]
	for i := 1; i < len(arr)-kInt+1; i++ {
		diff := arr[i+kInt-1] - arr[i]
		if diff < ans {
			ans = diff
		}
	}
	return ans
}
