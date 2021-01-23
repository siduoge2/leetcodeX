package stack

import "fmt"

func main() {
	height := []int{2, 0, 2}
	sum := trap(height)
	fmt.Println(sum)
}

func trap(height []int) int {
	var sum int
	if len(height) < 2 {
		return sum
	}
	leftMax := height[0]
	for i := 1; i < len(height)-1; i++ {
		rightMax := getMax(height[i+1:])
		heighOfI := getMin(leftMax, rightMax)
		if snglWater := heighOfI - height[i]; snglWater > 0 {
			sum += snglWater
		}
		if height[i] > leftMax {
			leftMax = height[i]
		}
	}
	return sum
}

func getMin(n1 int, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}

func getMax(arr []int) int {
	if len(arr) < 1 {
		return 0
	}
	max := arr[0]
	for _, val := range arr {
		if val > max {
			max = val
		}
	}
	return max
}
