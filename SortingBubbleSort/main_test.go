package main

import (
	"testing"
)

func Test_countSwapsImpl_1(t *testing.T) {
	a := []int32{6, 4, 1}
	answerCount, answerFirst, answerLast := genSwapResult(3, 1, 6)
	resultCount, resultFirst, resultLast := countSwapsImpl(a)

	if resultCount != answerCount {
		t.Errorf("count result: %s answer: %s", resultCount, answerCount)
	}
	if resultFirst != answerFirst {
		t.Errorf("first result: %s answer: %s", resultFirst, answerFirst)
	}
	if resultLast != answerLast {
		t.Errorf("last result: %s answer: %s", resultLast, answerLast)
	}
}
