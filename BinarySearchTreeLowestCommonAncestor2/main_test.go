package main

import (
	"testing"
)

func Test_lca(t *testing.T) {
	testCases := []struct {
		description string
		root        *Node
		v1          int32
		v2          int32
		answer      int32
	}{
		{
			description: "TestCase 1",
			root:        genBinaryTree([]int32{4, 2, 7, 1, 3, 6}),
			v1:          1,
			v2:          7,
			answer:      4,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			result := lca(testCase.root, testCase.v1, testCase.v2)
			if result != testCase.answer {
				t.Errorf("result: %d answer: %d", result, testCase.answer)
			}
		})
	}
}
