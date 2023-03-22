package main

const MaxExpenditure int32 = 200

func activityNotifications(expenditure []int32, d int32) int32 {
	// Write your code here
	// https://www.hackerrank.com/challenges/fraudulent-activity-notifications/editorial?isFullScreen=true&h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=sorting

	expenditureMap := map[int32]int{}
	ans := int32(0)

	for i := 0; i < len(expenditure); i++ {
		val := expenditure[i]

		if i >= int(d) {
			med := findInMap(expenditureMap, int(d/2+d%2))
			if d%2 == 0 {
				ret := findInMap(expenditureMap, int(d/2+1))
				if val >= med+ret {
					ans++
				}
			} else {
				if val >= med*2 {
					ans++
				}
			}
		}

		expenditureMap[val]++

		if i >= int(d) {
			prev := expenditure[int32(i)-d]
			expenditureMap[prev]++
		}
	}

	return ans
}

func findInMap(m map[int32]int, idx int) int32 {
	s := 0
	for i := int32(0); i < MaxExpenditure; i++ {
		freq := 0
		if v, ok := m[i]; ok {
			freq = v
		}
		s = s + freq
		if s >= idx {
			return i
		}
	}
	// ??
	return 0
}
