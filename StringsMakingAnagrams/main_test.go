package main

import "testing"

func Test_makeAnagram(t *testing.T) {
	testCases := []struct {
		description string
		a           string
		b           string
		answer      int32
	}{
		{
			description: "cde, abc",
			a:           "cde",
			b:           "abc",
			answer:      4,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			result := makeAnagram(testCase.a, testCase.b)
			if result != testCase.answer {
				t.Errorf("result: %d answer: %d", result, testCase.answer)
			}
		})
	}
}
