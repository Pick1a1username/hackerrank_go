package main

import (
	"testing"
)

func Test_activityNotifications(t *testing.T) {
	testCases := []struct {
		description string
		expenditure []int32
		d           int32
		answer      int32
	}{
		{
			description: "TestCase 1",
			expenditure: []int32{2, 3, 4, 2, 3, 6, 8, 4, 5},
			d:           5,
			answer:      2,
		},
		{
			description: "TestCase 2",
			expenditure: []int32{1, 2, 3, 4, 4},
			d:           4,
			answer:      0,
		},
		{
			description: "TestCase 3",
			expenditure: []int32{10, 20, 30, 40, 50},
			d:           3,
			answer:      1,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			result := activityNotifications(testCase.expenditure, testCase.d)
			if result != testCase.answer {
				t.Errorf("result: %d answer: %d", result, testCase.answer)
			}
		})
	}
}
