package sort

import (
	"fmt"
	"sort"
	"testing"
)

func TestReconstructQueue(t *testing.T) {
	people := [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}
	reconstructQueue(people)
	//copy(people[p[1]+1:i+1], people[p[1]:i+1])
	for i,p:=range people{
		if i==p[1] {
			continue
		}
		var res [][]int
		res=append(res,people[:p[1]]...)
		res=append(res,p)
		res=append(res,people[p[1]:i]...)
		if i<len(people)-1 {
			res=append(res,people[i+1:]...)
		}
		people=res
	}
	fmt.Println(people)
}

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] > people[j][0] {
			return true
		}else if people[i][0] == people[j][0]&& people[i][1] < people[j][1]{
			return true
		}
		return false
	})
	fmt.Println(people)
	return people
}

func sortSlice(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		return people[i][0] > people[j][0]||(people[i][0] == people[j][0]&&people[i][1] < people[j][1])
	})
	fmt.Println(people)
	return people
}