package leecode

func main() {

}

func relativeSortArray(arr1 []int, arr2 []int) []int {
	var numarr [1001]int
	var res []int
	for _, val := range arr1 {
		if val > 1000 {
			break
		}
		numarr[val]++
	}
	for _, val := range arr2 {
		for numarr[val] > 0 {
			res = append(res, val)
			numarr[val]--
		}
	}
	for i, val := range numarr {
		for val > 0 {
			res = append(res, i)
			val--
		}
	}
	return res
}
