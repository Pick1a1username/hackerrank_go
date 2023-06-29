package main

import (
	"reflect"
	"testing"
)

func Test_whatFlavorsImpl(t *testing.T) {
	testCases := []struct {
		description string
		money       int32
		cost        []int32
		answer      []int
	}{
		// {
		// 	description: "TestCase 14-1",
		// 	money:       4,
		// 	cost:        []int32{1, 4, 5, 3, 2},
		// 	answer:      []int{1, 4},
		// },
		// {
		// 	description: "TestCase 14-2",
		// 	money:       4,
		// 	cost:        []int32{2, 2, 4, 3},
		// 	answer:      []int{1, 2},
		// },
		// {
		// 	description: "TestCase 15",
		// 	money:       5,
		// 	cost:        []int32{1, 2, 3, 5, 6},
		// 	answer:      []int{2, 3},
		// },
		{
			description: "TestCase 16-1",
			money:       8,
			cost:        []int32{4, 3, 2, 5, 7},
			answer:      []int{2, 4},
		},
		// {
		// 	description: "TestCase 16-2",
		// 	money:       12,
		// 	cost:        []int32{7, 2, 5, 4, 11},
		// 	answer:      []int{1, 3},
		// },
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			result := whatFlavorsImpl(testCase.cost, testCase.money)
			if !reflect.DeepEqual(result, testCase.answer) {
				t.Errorf("result: %v answer: %v", result, testCase.answer)
			}
		})
	}
}
