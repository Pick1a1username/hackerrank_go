package main

import (
	"testing"
)

func Test_lca(t *testing.T) {
	testCases := []struct {
		description string
		root        *Node
		answer      int32
	}{
		{
			description: "TestCase 1",
			root:        genBinaryTree([]int32{4, 2, 6, 1, 3, 5, 7}),
			answer:      2,
		},
		{
			description: "TestCase 2",
			root:        genBinaryTree([]int32{3, 2, 5, 1, 4, 6, 7}),
			answer:      3,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			result := getHeight(testCase.root)
			if result != testCase.answer {
				t.Errorf("result: %d answer: %d", result, testCase.answer)
			}
		})
	}
}
