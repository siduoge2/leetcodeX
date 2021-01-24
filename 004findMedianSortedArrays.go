package sort

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
	return float64(nums1[i-1] + nums1[i]) / 2
}
// 时间复杂度log(m+n)
/*
二分查找
1、查找num1满足条件的mid1，使得
num2[mid2]<=nums1[mid1+1]
num1[mid1]<=nums2[mid2+1]
mid1和mid2属于左侧
2、奇数时，mid为mid1和mid2较大的
偶数时，取mid1\mid2较大的;mid1+1\mid2+1较小的
*/
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	if len(nums1)>len(nums2) {
		nums1,nums2=nums2,nums1
	}
	if len(nums1)==0 {
		
	}
	mid:=(len(nums1)+len(nums2)+1)/2
	mid1,mid2:=0,0
	for mid1<len(nums1) {
		mid2=mid-mid1
		if mid1>0&&nums1[mid1]<nums2[mid2-1] {//mid1 too little
			mid1++
		}else if mid1>0&&nums1[mid1-1]>nums2[mid2]  {//mid1 too large
			mid1--
		}else {
			break
		}
	}
	if 2*mid==len(nums1)+len(nums2) {//偶数
		return float64(max(nums1[mid1], nums2[mid2]) + min(nums1[mid1+1], nums2[mid2+1]))
	}else {
		return float64(max(nums1[mid1], nums2[mid2]))
	}
}

func max(n1 int,n2 int)int  {
	if n1>n2 {
		return	n1
	}
	return n2
}

func min(n1 int,n2 int)int  {
	if n1<n2 {
		return	n1
	}
	return n2
}