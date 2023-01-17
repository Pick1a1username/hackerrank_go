package main

import "testing"

func TestLuckBalance(t *testing.T) {
	testCases := []struct {
		description string
		k           int32
		contests    [][]int32
		answer      int32
	}{
		{
			description: "TestCase 1",
			k:           2,
			contests: [][]int32{
				{5, 1},
				{1, 1},
				{4, 0},
			},
			answer: 10,
		},
		{
			description: "TestCase 2",
			k:           1,
			contests: [][]int32{
				{5, 1},
				{1, 1},
				{4, 0},
			},
			answer: 8,
		},
		{
			description: "TestCase 3",
			k:           3,
			contests: [][]int32{
				{5, 1},
				{2, 1},
				{1, 1},
				{8, 1},
				{10, 0},
				{5, 0},
			},
			answer: 29,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			result := luckBalance(testCase.k, testCase.contests)
			if result != testCase.answer {
				t.Errorf("result: %d answer: %d", result, testCase.answer)
			}
		})
	}
}
