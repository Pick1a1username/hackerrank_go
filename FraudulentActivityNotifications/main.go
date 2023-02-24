package main

import "sort"

func activityNotifications(expenditure []int32, d int32) int32 {
	// Write your code here

	// Find the start day to check.
	// The start day is eqaul to d.

	count := int32(0)
	// Loop from the start day to the end day.
	for currDay := d; currDay < int32(len(expenditure)); currDay++ {
		// Get the spendings in the trailing days.
		spendings := expenditure[currDay-d : currDay]

		// Calculate the median in the trailing days.
		median := getMedianInt32(spendings)

		// Check the spending today is greater than or equal to the median.
		// Don't consider precision
		if float32(expenditure[currDay]) >= median*2 {
			// If so, count it.
			count++
		}

	}
	// Return the count.

	return count
}

func getMedianInt32(arr []int32) float32 {
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	l := len(arr)
	if l == 0 {
		return 0
	} else if l%2 == 0 {
		return (float32(arr[l/2-1]) + float32(arr[l/2])) / 2
	} else {
		return float32(arr[l/2])
	}
}
