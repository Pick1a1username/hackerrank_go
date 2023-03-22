package main

/*
 * Complete the 'commonChild' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. STRING s1
 *  2. STRING s2
 */
func commonChild(s1 string, s2 string) int32 {
	// Write your code here
	// https://en.wikipedia.org/wiki/Longest_common_subsequence
	// Computing the length of the LCS
	c := make([][]int, len(s1)+1)
	for i := 0; i <= len(s1); i++ {
		c_inner := make([]int, len(s2)+1)
		c[i] = c_inner
	}

	// int's default value is 0.
	// for i := 0; i <= len(s1); i++ {
	// 	c[i][0] = 0
	// }
	// for j := 0; j <= len(s2); j++ {
	// 	c[0][j] = 0
	// }
	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			if s1[i-1] == s2[j-1] {
				c[i][j] = c[i-1][j-1] + 1
			} else {
				c[i][j] = maxInt(c[i][j-1], c[i-1][j])
			}
		}
	}

	return int32(c[len(s1)][len(s2)])
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	if a < b {
		return b
	}
	return a
}
