
package sort

import (
	"fmt"
	"testing"
)

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{
			args: args{
				nums1: []int{1,3},
				nums2: []int{2},
			},
			want: float64(2),
		},
		{
			args: args{
				nums1: []int{1,3},
				nums2: []int{2,4},
			},
			want: 2.5,
		},
		{
			args: args{
				nums1: []int{1,3},
				nums2: []int{},
			},
			want: 2,
		},
		{
			args: args{
				nums1: []int{3},
				nums2: []int{},
			},
			want: 1.5,
		},
	}
	for _, tt := range tests {
		tt.name=fmt.Sprintf("%v %v",tt.args.nums1,tt.args.nums2)
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays2(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getmeida(t *testing.T) {
	type args struct {
		nums1 []int
		i     int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getmeida(tt.args.nums1, tt.args.i); got != tt.want {
				t.Errorf("getmeida() = %v, want %v", got, tt.want)
			}
		})
	}
}