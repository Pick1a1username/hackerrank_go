package main

// Complete the substrCount function below.
func substrCount(n int32, s string) int64 {
	nint := int(n)
	specialStrings := []string{}
	// loop
	for i := 0; i < nint; i++ {
		specialString := ""
		middleString := ""

		// one more loop
		for j := i; j < nint; j++ {
			curr := string(s[j])
			if specialString == "" {
				specialString = specialString + string(curr)
				continue
			}
			// check if the current character is same with previous one,
			// add to specialString and continue to check.
			if len(middleString) == 0 && specialString[len(specialString)-1:] == curr {
				specialString = specialString + string(curr)
				continue
			}
			//
			if len(middleString) == 0 && specialString[len(specialString)-1:] != curr {
				middleString = curr
				specialString = specialString + string(curr)
				continue
			}
			//
			if specialString[len(specialString)-1:] != curr {
				middleString = curr
				specialString = specialString + string(curr)
				continue
			}

		}
	}
	return 0
}
