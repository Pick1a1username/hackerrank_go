package main

import "testing"

func Test_genBinaryTree(t *testing.T) {
	testCases := []struct {
		description string
		data        []int32
		answer      *Node
	}{
		{
			description: "TestCase 1",
			data:        []int32{2, 1, 3},
			answer: &Node{
				Data:  2,
				Left:  &Node{Data: 1},
				Right: &Node{Data: 3},
			},
		},
		{
			description: "TestCase 2",
			data:        []int32{4, 2, 3},
			answer: &Node{
				Data: 4,
				Left: &Node{
					Data:  2,
					Left:  nil,
					Right: &Node{Data: 3},
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			result := genBinaryTree(testCase.data)
			if !equalBinaryTree(result, testCase.answer) {
				t.Errorf("result: %v answer: %v", result, testCase.answer)
			}
		})
	}
}

func equalBinaryTree(r1 *Node, r2 *Node) bool {
	if r1.Data != r2.Data {
		return false
	}
	if (r1.Left == nil && r2.Left != nil) || (r1.Right == nil && r2.Right != nil) {
		return false
	}
	if r1.Left != nil && !equalBinaryTree(r1.Left, r2.Left) {
		return false
	}
	if r1.Right != nil && !equalBinaryTree(r1.Right, r2.Right) {
		return false
	}
	return true
}
