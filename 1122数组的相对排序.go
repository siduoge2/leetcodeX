package sort


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

func relativeSortArray2(arr1 []int, arr2 []int) []int {
	const maxLength = 1000
	arr1times:=make([]int,maxLength+1)
	result :=make([]int,0,len(arr1))
	// init
	//for i,_:=range arr1times{
	//	arr1times[i]=-1
	//}
	// 计算次数
	for _,val:=range arr1{
		arr1times[val]++
	}
	// 按照arry2顺序排序
	for _,val:=range arr2{
		times:=arr1times[val]
		if times>0 {
			result =appendXNTimes(val,times, result)
		}
		arr1times[val]=0
	}
	// 余下的数升序排列在最后
	for index,times:=range arr1times {
		if times>0 {
			result =appendXNTimes(index,times, result)
		}
		arr1times[index]=0
	}
	return result
}

func appendXNTimes(addNum int,times int ,inputarr []int)[]int  {
	if times<=0 {
		return inputarr
	}
	for times>0 {
		inputarr=append(inputarr,addNum)
		times--
	}
	return inputarr
}