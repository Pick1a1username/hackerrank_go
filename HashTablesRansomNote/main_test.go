package main

import "testing"

func Test_checkMagazineImpl_yes_1(t *testing.T) {
	magazine := []string{"give", "me", "one", "grand", "today", "night"}
	note := []string{"give", "one", "grand", "today"}
	answer := "Yes"

	result := checkMagazineImpl(magazine, note)
	if result != answer {
		t.Errorf("result: %s answer: %s", result, answer)
	}
}

func Test_checkMagazineImpl_no_1(t *testing.T) {
	magazine := []string{"two", "times", "three", "is", "not", "four"}
	note := []string{"two", "times", "two", "is", "four"}
	answer := "No"

	result := checkMagazineImpl(magazine, note)
	if result != answer {
		t.Errorf("result: %s answer: %s", result, answer)
	}
}
