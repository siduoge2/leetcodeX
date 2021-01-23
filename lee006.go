package leecode

import (
	"strings"
)

// 输入: s = "LEETCODEISHIRING", numRows = 4
//输出: "LDREOEIIECIHNTSG"
//解释:
//
//L     D     R
//E   O E   I I
//E C   I H   N
//T     S     G
// Related Topics 字符串

//leetcode submit region begin(Prohibit modification and deletion)
func convert(s string, numRows int) string {
	result := ""
	if numRows <= 1 {
		return s
	}
	strings := strings.Split(s, "")
	arr := make([]string, numRows)
	i := 0
	stp := 2*numRows - 2 // 步长
	for n, s := range strings {
		arr[i] += s
		if (n+1)%stp < numRows && (n+1)%stp != 0 {
			i++
		} else {
			i--
		}

	}
	for n := 0; n < len(arr); n++ {
		result += arr[n]
	}
	return result
}

/*func main() {
	// s := "LEETCODEISHIRING"
	s := "LEETCODEISHIRING"
	numRows := 1
	result := convert(s, numRows)
	fmt.Println(result)
}
*/