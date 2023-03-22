package main

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
	// Bubble sort
	arrLen := len(arr)
	bubble(&arr)

	if arrLen == 0 {
		return 0
	} else if arrLen%2 == 0 {
		return (float32(arr[arrLen/2-1]) + float32(arr[arrLen/2])) / 2
	} else {
		return float32(arr[arrLen/2])
	}
}

func bubble(arr *[]int32) {
	// Bubble sort
	for i := 0; i < len(*arr)-1; i++ {
		swaped := false
		for j := 0; j < len(*arr)-1; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
				swaped = true
			}
		}
		if !swaped {
			break
		}
	}
}
