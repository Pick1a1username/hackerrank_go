package main

import (
	"reflect"
	"sort"
	"testing"
)

func Test_minimumAbsoluteDifference(t *testing.T) {
	testCases := []struct {
		description string
		arr         []int32
		answer      int32
	}{
		{
			description: "TestCase 1",
			arr:         []int32{-2, 2, 4},
			answer:      2,
		},
		{
			description: "TestCase 2",
			arr:         []int32{3, 7, 0},
			answer:      3,
		},
		{
			description: "TestCase 3",
			arr:         []int32{-59, -36, -13, 1, -53, -92, -2, -96, -54, 75},
			answer:      1,
		},
		{
			description: "TestCase 4",
			arr:         []int32{1, -3, 71, 68, 17},
			answer:      3,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			result := minimumAbsoluteDifference(testCase.arr)
			if result != testCase.answer {
				t.Errorf("result: %d answer: %d", result, testCase.answer)
			}
		})
	}
}

func Test_genTwoPairs(t *testing.T) {
	testCases := []struct {
		description string
		arr         []int32
		answer      [][]int32
	}{
		{
			description: "TestCase 1",
			arr:         []int32{-2, 2, 4},
			answer: [][]int32{
				{-2, 2},
				{-2, 4},
				{2, 4},
			},
		},
		{
			description: "TestCase 2",
			arr:         []int32{3, 7, 0},
			answer: [][]int32{
				{3, 7},
				{3, 0},
				{7, 0},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			result := genTwoPairs(testCase.arr)
			if !equalArraysInt32(result, testCase.answer) {
				t.Errorf("result: %v answer: %v", result, testCase.answer)
			}
		})
	}
}

func sortArrayAscInt32(arr []int32) []int32 {
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	return arr
}

// Check if two arrays have same elements.
// The order is ignored.
func equalArraysInt32(a, b [][]int32) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		found := false
		for j := 0; j < len(a); j++ {
			if reflect.DeepEqual(sortArrayAscInt32(a[i]), sortArrayAscInt32(b[j])) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	return true
}
