package main

import "testing"

// Table Driven Test
func Test_maxMin(t *testing.T) {
	testCases := []struct {
		description string
		k           int32
		arr         []int32
		answer      int32
	}{
		{
			description: "TestCase 1",
			k:           3,
			arr:         []int32{10, 100, 300, 200, 1000, 20, 30},
			answer:      20,
		},
		{
			description: "TestCase 2",
			k:           4,
			arr:         []int32{1, 2, 3, 4, 10, 20, 30, 40, 100, 200},
			answer:      8,
		},
		{
			description: "TestCase 3",
			k:           2,
			arr:         []int32{1, 2, 1, 2, 1},
			answer:      0,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			result := maxMin(testCase.k, testCase.arr)
			if result != testCase.answer {
				t.Errorf("result: %d answer: %d", result, testCase.answer)
			}
		})
	}
}
