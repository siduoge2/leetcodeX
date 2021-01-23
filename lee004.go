package leecode

import (
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	for _, num := range nums2 {
		nums1 = append(nums1, num)
	}
	sort.Ints(nums1)
	var result float64
	if len(nums1)%2 == 0 {
		result = getmeida(nums1, len(nums1)/2)
	} else {
		result = float64(nums1[len(nums1)/2])
	}
	return result
}

func getmeida(nums1 []int, i int) float64 {
	return float64((nums1[i-1] + nums1[i])) / 2
}
