package main

import "testing"

func Test_alternatingCharacters(t *testing.T) {
	testCases := []struct {
		description string
		s           string
		answer      int32
	}{
		{
			description: "AAAA",
			s:           "AAAA",
			answer:      3,
		},
		{
			description: "BBBBB",
			s:           "BBBBB",
			answer:      4,
		},
		{
			description: "ABABABAB",
			s:           "ABABABAB",
			answer:      0,
		},
		{
			description: "BABABA",
			s:           "BABABA",
			answer:      0,
		},
		{
			description: "AAABBB",
			s:           "AAABBB",
			answer:      4,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			result := alternatingCharacters(testCase.s)
			if result != testCase.answer {
				t.Errorf("result: %d answer: %d", result, testCase.answer)
			}
		})
	}
}
