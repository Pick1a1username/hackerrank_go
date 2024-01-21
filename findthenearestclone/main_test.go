package main

import "testing"

func Test_findShortest(t *testing.T) {
	testCases := []struct {
		description string
		graphNodes  int32
		graphFrom   []int32
		graphTo     []int32
		ids         []int64
		val         int32
		answer      int32
	}{
		{
			description: "Sample Input 0",
			graphNodes:  4,
			graphFrom:   []int32{1, 1, 4},
			graphTo:     []int32{2, 3, 2},
			ids:         []int64{1, 2, 1, 1},
			val:         1,
			answer:      1,
		},
		{
			description: "Sample Input 1",
			graphNodes:  4,
			graphFrom:   []int32{1, 1, 4},
			graphTo:     []int32{2, 3, 2},
			ids:         []int64{1, 2, 3, 4},
			val:         2,
			answer:      -1,
		},
		{
			description: "Sample Input 2",
			graphNodes:  5,
			graphFrom:   []int32{1, 1, 2, 3},
			graphTo:     []int32{2, 3, 4, 5},
			ids:         []int64{1, 2, 3, 3, 2},
			val:         2,
			answer:      3,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			result := findShortest(testCase.graphNodes, testCase.graphFrom, testCase.graphTo, testCase.ids, testCase.val)
			if result != testCase.answer {
				t.Errorf("result: %d answer: %d", result, testCase.answer)
			}
		})
	}
}
