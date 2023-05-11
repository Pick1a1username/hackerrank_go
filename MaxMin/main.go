package main

import "sort"

func maxMin(k int32, arr []int32) int32 {
	// Sort arr.
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	// Get two indice whose numbers' difference is smallest.
	indice := []int{}
	lastDiff := int32(0)
	for i := 1; i < len(arr); i++ {
		if len(indice) == 0 {
			indice = append(indice, 0, 1)
			lastDiff = arr[indice[1]] - arr[indice[0]]
			continue
		}
		currDiff := arr[i] - arr[i-1]
		if currDiff < lastDiff {
			lastDiff = currDiff
			indice[0] = i - 1
			indice[1] = i
		}
	}
	// If the number of indice is smaller than k, check adjacent numbers.
	for len(indice) < int(k) {
		minIdx := indice[0]
		maxIdx := indice[len(indice)-1]
		// What if minIdx is zero?
		if minIdx == 0 {
			indice = append(indice, maxIdx+1)
			continue
		}
		// What if maxIdx is equal to len(indice)-1?
		if maxIdx == len(indice)-1 {
			indice = append([]int{minIdx - 1}, indice...)
			continue
		}
		minIdxLeftDiff := arr[minIdx] - arr[minIdx-1]
		maxIdxRightDiff := arr[maxIdx+1] - arr[minIdx]
		if minIdxLeftDiff < maxIdxRightDiff {
			indice = append([]int{minIdx - 1}, indice...)
		} else {
			indice = append(indice, maxIdx+1)
		}
	}
	return int32(arr[indice[len(indice)-1]] - arr[indice[0]])
}
