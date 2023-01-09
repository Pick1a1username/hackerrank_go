package main

import (
	"reflect"
	"testing"
)

func Test_minimumBribesImpl_1(t *testing.T) {
	q := []int32{2, 1, 5, 3, 4}
	answer := "3"
	result := minimumBribesImpl(q)
	if !reflect.DeepEqual(result, answer) {
		t.Errorf("result: %v answer: %v\n", result, answer)
	}
}

func Test_minimumBribesImpl_2(t *testing.T) {
	q := []int32{1, 5, 2, 3, 4}
	answer := "Too chaotic"
	result := minimumBribesImpl(q)
	if !reflect.DeepEqual(result, answer) {
		t.Errorf("result: %v answer: %v\n", result, answer)
	}
}

func Test_minimumBribesImpl_3(t *testing.T) {
	q := []int32{1, 2, 3, 5, 4, 6, 7}
	answer := "1"
	result := minimumBribesImpl(q)
	if !reflect.DeepEqual(result, answer) {
		t.Errorf("result: %v answer: %v\n", result, answer)
	}
}

func Test_minimumBribesImpl_4(t *testing.T) {
	q := []int32{4, 1, 2, 3}
	answer := "Too chaotic"
	result := minimumBribesImpl(q)
	if !reflect.DeepEqual(result, answer) {
		t.Errorf("result: %v answer: %v\n", result, answer)
	}
}
func Test_addMaps_1(t *testing.T) {
	a := map[int]int{
		1: 1,
		2: 2,
		3: 3,
	}
	b := map[int]int{
		2: 4,
		3: 5,
		4: 6,
	}
	answer := map[int]int{
		1: 1,
		2: 6,
		3: 8,
	}

	result := addMaps(a, b)
	if !reflect.DeepEqual(result, answer) {
		t.Errorf("result: %v answer: %v", result, answer)
	}
}

func Test_countBribeNum_ok_1(t *testing.T) {
	currQ := []int{1, 2, 3}
	targetQ := []int{3, 1, 2}
	briberOrder := []int{1, 2, 3}
	answerMap := map[int]int{
		1: 0,
		2: 0,
		3: 2,
	}

	resultMap := countBribeNum(currQ, targetQ, briberOrder)

	if !reflect.DeepEqual(resultMap, answerMap) {
		t.Errorf("map result: %v answer: %v", resultMap, answerMap)
	}
}

func Test_countBribeNum_ok_2(t *testing.T) {
	currQ := []int{1, 2, 3, 4}
	targetQ := []int{4, 3, 1, 2}
	briberOrder := []int{1, 2, 3, 4}
	answerMap := map[int]int{
		1: 0,
		2: 0,
		3: 2,
		4: 3,
	}

	resultMap := countBribeNum(currQ, targetQ, briberOrder)

	if !reflect.DeepEqual(resultMap, answerMap) {
		t.Errorf("map result: %v answer: %v", resultMap, answerMap)
	}
}

func Test_removeElemArrayInt_1(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	i := 2
	answer := []int{1, 2, 4, 5}

	result := removeElemArrayInt(&arr, i)
	if !reflect.DeepEqual(result, answer) {
		t.Errorf("result: %v answer: %v\n", result, answer)
	}
}

func Test_insertElemArrayInt_1(t *testing.T) {
	arr := []int{1, 2, 4, 5}
	i := 2
	v := 3
	answer := []int{1, 2, 3, 4, 5}

	result := insertElemArrayInt(&arr, i, v)
	if !reflect.DeepEqual(result, answer) {
		t.Errorf("result: %v answer: %v\n", result, answer)
	}
}

func Test_findIdxByValInt_ok_1(t *testing.T) {
	sl := []int{1, 2, 3, 4, 5}
	v := 4
	answerIdx := 3

	resultIdx, resultErr := findIdxByValInt(&sl, v)
	if resultIdx != answerIdx {
		t.Errorf("index result: %d answer: %d\n", resultIdx, answerIdx)
	}
	if resultErr != nil {
		t.Errorf("error result: %v answer: nil\n", resultErr)
	}
}

func Test_findIdxByValInt_ng_1(t *testing.T) {
	sl := []int{1, 2, 3, 4, 5}
	v := 6
	answer := "not found"

	_, result := findIdxByValInt(&sl, v)
	if result == nil {
		t.Errorf("result: nil answer: %v\n", result)
	}
	if result != nil && result.Error() != answer {
		t.Errorf("result: %s answer: %s\n", result.Error(), answer)
	}

}

// func Test_genNoBribeQ_1(t *testing.T) {}
