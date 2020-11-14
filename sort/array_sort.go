package sort

import "sort"

/* 1122
给你两个数组，arr1 和 arr2，

arr2 中的元素各不相同
arr2 中的每个元素都出现在 arr1 中
对 arr1 中的元素进行排序，使 arr1 中项的相对顺序和 arr2 中的相对顺序相同。未在 arr2 中出现过的元素需要按照升序放在 arr1 的末尾。
*/
func relativeSortArray(arr1 []int, arr2 []int) []int {
	cMap := make(map[int]int)
	for i, v := range arr2 {
		cMap[v] = i
	}

	sort.Slice(arr1, func(i, j int) bool {
		indexI, okI := cMap[arr1[i]]
		indexJ, okJ := cMap[arr1[j]]
		if !okI && !okJ {
			return arr1[i] < arr1[j]
		}

		if !okI {
			return 10000 < indexJ
		}

		if !okJ {
			return indexI < 10000

		}

		return indexI < indexJ
	})

	return arr1
}
