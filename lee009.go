package leecode

func isPalindrome(x int) bool {
	xx := x
	if xx < 0 {
		return false
	}
	re := 0
	for xx > 0 {
		re = (xx % 10) + re*10
		xx /= 10
	}
	if re == x {
		return true
	}
	return false
}
/*
func main() {
	var input []int
	input = append(input, 181, 121, 1, 11, 111, -1, 0)
	for _, in := range input {
		fmt.Println(in)
		fmt.Println(isPalindrome(in))
	}
}
*/