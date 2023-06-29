package main

import (
	"fmt"
	"sort"
)

func whatFlavors(cost []int32, money int32) {
	// Write your code here
	vals := whatFlavorsImpl(cost, money)
	fmt.Printf("%d %d\n", vals[0], vals[1])
}

func whatFlavorsImpl(cost []int32, money int32) []int {
	// Push costs into a map.
	// map[cost][]index
	costMap := map[int32][]int{}
	for i, c := range cost {
		costMap[c] = append(costMap[c], i)
	}

	// For loop by the map
	for cost, indexArr := range costMap {
		remaining := money - cost
		for indexOfIndex, index := range indexArr {
			// Make a map without the current cost.
			tmpCostMap := map[int32][]int{}
			copyMap(tmpCostMap, costMap)
			tmpCostMap[cost] = removeElemArrayInt(&indexArr, indexOfIndex)
			// If a cost for remaining exists, return.
			if anotherCostIndexArr := tmpCostMap[remaining]; len(anotherCostIndexArr) > 0 {
				result := []int{index + 1, anotherCostIndexArr[0] + 1}
				sort.Slice(result, func(i, j int) bool { return result[i] < result[j] })
				return result
			}
		}
	}
	return []int{-1, -1}
}

func copyMap(dst, src map[int32][]int) {
	for k, v := range src {
		newV := make([]int, len(v))
		copy(newV, v)
		dst[k] = newV
	}
}

func removeElemArrayInt(arr *[]int, i int) []int {
	tmpLeftQ := append([]int{}, (*arr)[:i]...)
	tmpRightQ := append([]int{}, (*arr)[i+1:]...)
	return append(tmpLeftQ, tmpRightQ...)
}
