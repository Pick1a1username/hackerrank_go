package main

// Complete the countTriplets function below.
func countTriplets(arr []int64, r int64) int64 {
	// https://www.hackerrank.com/challenges/count-triplets-1/forum/comments/495008?isFullScreen=true&h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=dictionaries-hashmaps
	rightMap := getOccurenceMap(arr)
	leftMap := map[int64]int64{}
	numberOfGeometricPairs := int64(0)

	for _, val := range arr {
		countLeft := int64(0)
		countRight := int64(0)
		lhs := int64(0)
		rhs := int64(val * r)

		if val%r == 0 {
			lhs = val / r
		}

		rightMap[val]--

		if rightOccurence, ok := rightMap[rhs]; ok {
			countRight = rightOccurence
		}
		if leftOccurence, ok := leftMap[lhs]; ok {
			countLeft = leftOccurence
		}
		numberOfGeometricPairs += countLeft * countRight
		leftMap[val]++
	}
	return numberOfGeometricPairs
}

func getOccurenceMap(arr []int64) map[int64]int64 {
	occurenceMap := map[int64]int64{}
	for _, v := range arr {
		occurenceMap[v]++
	}
	return occurenceMap
}

// private static Map<Long, Long> getOccurenceMap(List<Long> test) {
//     Map<Long, Long> occurenceMap = new HashMap<>();
//     for (long val : test) {
//         insertIntoMap(occurenceMap, val);
//     }
//    return occurenceMap;
// }

// private static void insertIntoMap(Map<Long, Long> occurenceMap, Long val) {
//     if (!occurenceMap.containsKey(val)) {
//         occurenceMap.put(val, 1L);
//       } else {
//         Long occurence = occurenceMap.get(val);
//         occurenceMap.put(val, occurence + 1L);
//       }
// }
