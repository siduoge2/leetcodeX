package leecode

import (
	"math"
)

/*func main() {
	fmt.Println("length:", shortestSubarray([]int{1}, 1))
	fmt.Println("length:", shortestSubarray([]int{1, 2}, 4))
	fmt.Println("length:", shortestSubarray([]int{2, -1, 2}, 3))
	fmt.Println("length:", shortestSubarray([]int{48, 99, 37, 4, -31}, 140))
}*/

//返回 A 的最短的非空连续子数组的长度，该子数组的和至少为 K 。
//
// 如果没有和至少为 K 的非空子数组，返回 -1 。
//
//
//
//
//
//
// 示例 1：
//
// 输入：A = [1], K = 1
//输出：1
//
//
// 示例 2：
//
// 输入：A = [1,2], K = 4
//输出：-1
//
//
// 示例 3：
//
// 输入：A = [2,-1,2], K = 3
//输出：3
//
//
//
//
// 提示：
//
//
// 1 <= A.length <= 50000
// -10 ^ 5 <= A[i] <= 10 ^ 5
// 1 <= K <= 10 ^ 9
//
// Related Topics 队列 二分查找

//leetcode submit region begin(Prohibit modification and deletion)
func shortestSubarray(A []int, K int) int {
	var sum int
	var length int = math.MaxInt32
	for i := len(A) - 1; i >= 0; i-- {
		sum += A[i]
		sum1 := sum
		if A[i] >= K {
			return 1
		}
		for j := len(A) - 1; j > i; j-- {
			if sum1 >= K {
				tempLength := j - i + 1
				if tempLength < length && tempLength != 0 {
					length = tempLength
				}
			}
			sum1 -= A[j]
			if sum1 >= K {
				tempLength := j - i
				if tempLength < length && tempLength != 0 {
					length = tempLength
				}
			}
		}
	}
	if length == math.MaxInt32 || length == 0 {
		return -1
	}
	return length
}

func subarraySum(tempSlicee []int) int {
	var sum int
	for _, val := range tempSlicee {
		sum += val
	}
	return sum
}

//leetcode submit region end(Prohibit modification and deletion)
