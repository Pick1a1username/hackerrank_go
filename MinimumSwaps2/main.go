package main

import "errors"

// Complete the minimumSwaps function below.
func minimumSwaps(arr []int32) int32 {
	count := int32(0)
	// go through arr
	for i := 0; i < len(arr); i++ {
		// if value is not supposed to be there, swap it with appropriate one.
		if arr[i] != int32(i+1) {
			swapWith, err := findIdxByValInt32(&arr, int32(i+1))
			if err != nil {
				panic("appropriate number not found")
			}
			arr[i], arr[swapWith] = arr[swapWith], arr[i]
			count++
		}
	}

	// return the count
	return count
}

func findIdxByValInt32(slPtr *[]int32, v int32) (int32, error) {
	for i := 0; i < len(*slPtr); i++ {
		if (*slPtr)[i] == v {
			return int32(i), nil
		}
	}
	return -1, errors.New("not found")
}
