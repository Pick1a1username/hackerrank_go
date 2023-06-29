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
	for c, indexArr := range costMap {
		remaining := money - c
		if remaining == c && len(indexArr) > 1 {
			return []int{indexArr[0] + 1, indexArr[1] + 1}
		}
		if remaining == c && len(indexArr) == 1 {
			continue
		}
		for _, index := range indexArr {
			// If a cost for remaining exists, return.
			if anotherCostIndexArr := costMap[remaining]; len(anotherCostIndexArr) > 0 {
				result := []int{index + 1, anotherCostIndexArr[0] + 1}
				sort.Slice(result, func(i, j int) bool { return result[i] < result[j] })
				return result
			}
		}
	}
	return []int{-1, -1}
}
