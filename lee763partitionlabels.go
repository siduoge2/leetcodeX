package leecode

import "fmt"

func main() {
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
}
func partitionLabels(S string) []int {
	lastPosition := make(map[int32]int)
	for index, byt := range S {
		lastPosition[byt] = index
	}
	var ans []int
	var start int
	var end int
	for index, val := range S {
		end = max(lastPosition[val], end)
		if index == end {
			ans = append(ans, end-start+1)
			start = end + 1
		}
	}
	return ans
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
