package main

import (
	"reflect"
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

func Test_getParentData(t *testing.T) {
	testCases := []struct {
		description string
		root        *Node
		data        int32
		answer      []int32
	}{
		{
			description: "TestCase 1",
			root:        genBinaryTree([]int32{4, 2, 7, 1, 3, 6}),
			data:        1,
			answer:      []int32{4, 2, 1},
		},
		{
			description: "TestCase 2",
			root:        genBinaryTree([]int32{4, 2, 7, 1, 3, 6}),
			data:        7,
			answer:      []int32{4, 7},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			result := getParentData(testCase.root, testCase.data)
			if !reflect.DeepEqual(result, testCase.answer) {
				t.Errorf("result: %v answer: %v", result, testCase.answer)
			}
		})
	}
}

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

// {
// 	description: "TestCase 1",
// 	root:        genBinaryTree([]int32{2, 1, 3, 4, 5, 6}),
// 	v1:          4,
// 	v2:          6,
// 	answer: &Node{
// 		Data: 2,
// 		Left: &Node{Data: 1},
// 		Right: &Node{
// 			Data: 3,
// 			Left: &Node{Data: 4},
// 			Right: &Node{
// 				Data:  5,
// 				Right: &Node{Data: 6},
// 			}},
// 	},
// },
// {
// 	description: "TestCase 2",
// 	root:        genBinaryTree([]int32{4, 2, 7, 1, 3, 6}),
// 	v1:          1,
// 	v2:          7,
// 	answer: &Node{
// 		Data: 4,
// 		Left: &Node{
// 			Data:  2,
// 			Left:  &Node{Data: 1},
// 			Right: &Node{Data: 3},
// 		},
// 		Right: &Node{
// 			Data: 7,
// 			Left: &Node{Data: 6},
// 		},
// 	},
// },
// }
