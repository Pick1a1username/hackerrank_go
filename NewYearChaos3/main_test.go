package main

import (
	"errors"
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
	answer := "-1"
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
	answer := "-1"
	result := minimumBribesImpl(q)
	if !reflect.DeepEqual(result, answer) {
		t.Errorf("result: %v answer: %v\n", result, answer)
	}
}

func Test_getBack(t *testing.T) {
	testCases := []struct {
		description string
		q           *[]int32
		p           int
		answerQ     *[]int32
		answerErr   error
	}{
		{
			description: "TestCase 1",
			q:           &[]int32{1, 2, 3},
			p:           1,
			answerQ:     &[]int32{1, 3, 2},
			answerErr:   nil,
		},
		{
			description: "TestCase 2",
			q:           &[]int32{1, 2, 3},
			p:           2,
			answerQ:     &[]int32{1, 2, 3},
			answerErr:   errors.New("cannot get back anymore"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			resultErr := getBack(testCase.q, testCase.p)
			if testCase.answerErr == nil && resultErr != nil {
				t.Errorf("resultErr: %v answerErr: nil", resultErr)
			}
			if testCase.answerErr != nil && resultErr.Error() != testCase.answerErr.Error() {
				t.Errorf("resultErr: %v answerErr: %v", resultErr, testCase.answerErr)
			}
			if !reflect.DeepEqual(testCase.q, testCase.answerQ) {
				t.Errorf("q: %v answerQ: %v", testCase.q, testCase.answerQ)
			}
		})
	}
}
