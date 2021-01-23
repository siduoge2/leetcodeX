package leecode

import "fmt"

func main() {
	nums := []int{1, 2, 7, 5, 11}
	r := twoSum(nums, 17)
	fmt.Print(r)
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
