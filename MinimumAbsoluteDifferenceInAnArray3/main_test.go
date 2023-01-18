package main

import (
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
