package main

import "testing"

func Test_minimumAbsoluteDifference(t *testing.T) {
	testCases := []struct {
		description string
		arr         []int64
		r           int64
		answer      int64
	}{
		{
			description: "TestCase 1",
			arr:         []int64{1, 2, 2, 4},
			r:           2,
			answer:      2,
		},
		{
			description: "TestCase 2",
			arr:         []int64{1, 3, 9, 9, 27, 81},
			r:           3,
			answer:      6,
		},
		{
			description: "TestCase 3",
			arr:         []int64{1, 5, 5, 25, 125},
			r:           5,
			answer:      4,
		},
		{
			description: "TestCase 4",
			arr:         []int64{1, 3, 3, 9, 9, 27},
			r:           3,
			answer:      8,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			result := countTriplets(testCase.arr, testCase.r)
			if result != testCase.answer {
				t.Errorf("result: %d answer: %d", result, testCase.answer)
			}
		})
	}
}
