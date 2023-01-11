package main

import "testing"

func Test_twoStringsImpl_yes_1(t *testing.T) {
	s1 := "hello"
	s2 := "world"
	answer := "YES"

	result := twoStringsImpl(s1, s2)
	if result != answer {
		t.Errorf("result: %s answer: %s", result, answer)
	}
}

func Test_twoStringsImpl_no_1(t *testing.T) {
	s1 := "hi"
	s2 := "world"
	answer := "NO"

	result := twoStringsImpl(s1, s2)
	if result != answer {
		t.Errorf("result: %s answer: %s", result, answer)
	}
}
