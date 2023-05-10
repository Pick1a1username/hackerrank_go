package main

func isValid(s string) string {
	// Put the number of characters of s into a hashmap.
	charNum := map[string]int{}
	for _, c := range s {
		charNum[string(c)]++
	}
	// Put the number of the number of characters of s into a hashmap.
	numNum := map[int]int{}
	for _, n := range charNum {
		numNum[n]++
	}
	// If all of the numbers are same, return YES.
	if len(numNum) == 1 {
		return "YES"
	}
	// If all of the numbers except one, try to remove one character.
	if len(numNum) == 2 {
		// If one of them is 1, return YES.
		for _, n := range numNum {
			if n == 1 {
				return "YES"
			}
		}
	}
	// Return NO.
	return "NO"
}
