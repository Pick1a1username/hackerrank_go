package main

import (
	"reflect"
	"testing"
)

func Test_maximumToysImpl_1(t *testing.T) {
	prices := []int32{1, 2, 3, 4}
	k := int32(7)
	answer := int32(3)

	result := maximumToysImpl(prices, k)

	if result != answer {
		t.Errorf("result: %d answer : %d", result, answer)
	}
}

func Test_maximumToysImpl_2(t *testing.T) {
	prices := []int32{1, 12, 5, 111, 200, 1000, 10}
	k := int32(30)
	answer := int32(4)

	result := maximumToysImpl(prices, k)

	if result != answer {
		t.Errorf("result: %d answer : %d", result, answer)
	}
}

func Test_genPossibleCombinations_1(t *testing.T) {
	prices := []int32{1, 2, 3}
	answer := [][]int32{
		{1},
		{1, 2},
		{1, 3},
		{1, 2, 3},
		{1, 3, 2},
		{2},
		{2, 1},
		{2, 3},
		{2, 1, 3},
		{2, 3, 1},
		{3},
		{3, 1},
		{3, 2},
		{3, 1, 2},
		{3, 2, 1},
	}

	result := genPossibleCombinations(prices)

	if !equalArraysInt32(result, answer) {
		t.Errorf("result: %v answer: %v", result, answer)
	}
}

func Test_genPossibleCombinations_2(t *testing.T) {
	prices := []int32{1, 1, 2}
	answer := [][]int32{
		{1},
		{1, 1},
		{1, 2},
		{1, 1, 2},
		{1, 2, 1},
		{1},
		{1, 1},
		{1, 2},
		{1, 1, 2},
		{1, 2, 1},
		{2},
		{2, 1},
		{2, 1},
		{2, 1, 1},
		{2, 1, 1},
	}

	result := genPossibleCombinations(prices)

	if !equalArraysInt32(result, answer) {
		t.Errorf("result: %v answer: %v", result, answer)
	}
}

func equalArraysInt32(a, b [][]int32) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		found := false
		for j := 0; j < len(a); j++ {
			if reflect.DeepEqual(sortArrayAscInt32(a[i]), sortArrayAscInt32(b[j])) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	return true
}

func Test_removeDup_1(t *testing.T) {
	arr := [][]int32{
		{2},
		{1},
		{2, 1},
		{1, 2},
	}
	answer := [][]int32{
		{1},
		{1, 2},
		{2},
	}

	result := removeDup(arr)

	if !equalArraysInt32(result, answer) {
		t.Errorf("result: %v answer: %v", result, answer)
	}
}

// func Test_maximumToysImpl_1(t *testing.T) {

// }
