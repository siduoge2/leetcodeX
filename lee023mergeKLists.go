package leecode

import (
	"fmt"
	"math"
)

func main() {
	l1 := &ListNode{
		Val:  1,
		Next: nil,
	}
	h1 := l1
	h1.Next = &ListNode{
		Val:  4,
		Next: nil,
	}
	h1 = h1.Next
	h1.Next = &ListNode{
		Val:  5,
		Next: nil,
	}

	l2 := &ListNode{
		Val:  1,
		Next: nil,
	}
	h2 := l2
	h2.Next = &ListNode{
		Val:  3,
		Next: nil,
	}
	h2 = h2.Next
	h2.Next = &ListNode{
		Val:  4,
		Next: nil,
	}

	l3 := &ListNode{
		Val:  2,
		Next: nil,
	}
	h3 := l3
	h3.Next = &ListNode{
		Val:  6,
		Next: nil,
	}
	// res := mergeTwoLists(l1, l2)
	var lists []*ListNode
	lists = append(lists, l1)
	lists = append(lists, l2)
	lists = append(lists, l3)
	res := mergeKLists(lists)
	fmt.Println(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	// if len(lists) == 1 {
	// 	return lists[0]
	// }
	// newlists := lists[1:]
	res := &ListNode{
		Val: int(math.MinInt16),
	}
	for _, list := range lists {
		res = mergeTwoLists(res, list)
	}
	return res.Next
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	head := res
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			res.Next = &ListNode{
				Val:  l1.Val,
				Next: nil,
			}
			res = res.Next
			l1 = l1.Next
		} else {
			l1, l2 = l2, l1
		}
	}
	if l1 == nil && l2 != nil {
		res.Next = l2
		return head.Next
	}
	if l2 == nil && l1 != nil {
		res.Next = l1
		return head.Next
	}
	return head.Next
}
