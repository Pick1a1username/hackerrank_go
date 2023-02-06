package main

import "sort"

// Complete the countTriplets function below.
func countTriplets(arr []int64, r int64) int64 {
	// get one and numbers which could be divided by r.
	// save the numbers to a map.
	m := map[int64]int64{}

	for _, n := range arr {
		m[n]++
	}

	// loop by the map
	keys := []int64{}

	for k, _ := range m {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	keyLen := len(keys)
	sum := int64(0)

	for i := 0; i < keyLen; i++ {
		// if there are less than two keys and values in front of current key. break.
		// keyLen = 3 -> i < 1
		// keyLen = 5 -> i < 2
		if i >= keyLen-2 {
			break
		}

		// get the first three keys and values
		firstKey := keys[i]
		firstVal := m[firstKey]
		secondKey := keys[i+1]
		secondVal := m[secondKey]
		thirdKey := keys[i+2]
		thirdVal := m[thirdKey]

		// multiply three of values and add it to the sum variable.
		sum += firstVal * secondVal * thirdVal
	}

	// return value.
	return sum
}
