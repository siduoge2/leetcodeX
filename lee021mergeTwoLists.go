package leecode

//
// import "fmt"
//
// func main() {
// 	l1 := &ListNode{
// 		Val:  1,
// 		Next: nil,
// 	}
// 	h1 := l1
// 	h1.Next = &ListNode{
// 		Val:  2,
// 		Next: nil,
// 	}
// 	h1 = h1.Next
// 	h1.Next = &ListNode{
// 		Val:  4,
// 		Next: nil,
// 	}
//
// 	l2 := &ListNode{
// 		Val:  1,
// 		Next: nil,
// 	}
// 	h2 := l2
// 	h2.Next = &ListNode{
// 		Val:  3,
// 		Next: nil,
// 	}
// 	h2 = h2.Next
// 	h2.Next = &ListNode{
// 		Val:  4,
// 		Next: nil,
// 	}
// 	res := mergeTwoLists(l1, l2)
// 	fmt.Println(res)
// }
//
// /**
//  * Definition for singly-linked list.
//  * type ListNode struct {
//  *     Val int
//  *     Next *ListNode
//  * }
//  */
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }
//
// func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
// 	res := &ListNode{}
// 	head := res
// 	for l1 != nil && l2 != nil {
// 		if l1.Val <= l2.Val {
// 			res.Next = &ListNode{
// 				Val:  l1.Val,
// 				Next: nil,
// 			}
// 			res = res.Next
// 			l1 = l1.Next
// 		} else {
// 			l1, l2 = l2, l1
// 		}
// 	}
// 	if l1 == nil && l2 != nil {
// 		res.Next = l2
// 		return head.Next
// 	}
// 	if l2 == nil && l1 != nil {
// 		res.Next = l1
// 		return head.Next
// 	}
// 	return head.Next
// }
