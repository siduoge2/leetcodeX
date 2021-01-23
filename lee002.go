package leecode

import "fmt"

// 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
// 输出：7 -> 0 -> 8
// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// 初始化数组
	l0, l1, l2 := new(ListNode), new(ListNode), new(ListNode)
	l3 := new(ListNode)
	l0.Val = 2
	l0.Next = l1
	l1.Val = 4
	l1.Next = l2
	l2.Val = 3
	l2.Next = l3
	l3.Val = 5

	l00, l01, l02 := new(ListNode), new(ListNode), new(ListNode)
	l00.Val = 8
	l00.Next = l01
	l01.Val = 7
	l01.Next = l02
	l02.Val = 7
	r := addTwoNumbers(l0, l00)

	for l0 != nil {
		fmt.Print(l0.Val)
		l0 = l0.Next
	}
	fmt.Println("")
	for l00 != nil {
		fmt.Print(l00.Val)
		l00 = l00.Next
	}
	fmt.Println("")
	for r != nil {
		fmt.Print(r.Val)
		r = r.Next
	}
	fmt.Println("")

	fmt.Println("")

}

/*func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	num1 := l1
	num2 := l2
	l3 := new(ListNode)
	num3 := l3
	jinWei := 0
	for num1 != nil || num2 != nil || jinWei != 0 {
		sum := 0
		if num1 == nil && num2 != nil {
			sum = num2.Val + jinWei
		} else if num2 == nil && num1 != nil {
			sum = num1.Val + jinWei
		} else if num2 == nil && num1 == nil {
			sum = jinWei
		} else {
			sum = num1.Val + num2.Val + jinWei
		}
		num3.Val = sum % 10
		jinWei = sum / 10
		if num1 != nil {
			num1 = num1.Next
		}
		if num2 != nil {
			num2 = num2.Next
		}

		if num1 != nil || num2 != nil || jinWei != 0 {
			num3.Next = new(ListNode)
			num3 = num3.Next
		}
	}
	return l3
}*/

/*
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	node1 := l1
	node2 := l2
	l3 := new(ListNode)
	node3 := l3
	a := 0
	for node1 != nil || node2 != nil || a > 0 {
		node3.Next = new(ListNode)
		node3 = node3.Next
		b := 0
		c := 0
		if node1 != nil {
			b = node1.Val
		}
		if node2 != nil {
			c = node2.Val
		}
		node3.Val = (a + b + c) % 10
		a = (a + b + c) / 10
		if node1 != nil {
			node1 = node1.Next
		}
		if node2 != nil {
			node2 = node2.Next
		}
	}
	return l3.Next
}*/

func addNode(l *ListNode, v int) *ListNode {
	l.Next = &ListNode{Val: v}
	return l.Next
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	head := &ListNode{}
	cur := head
	jz := 0
	for l1 != nil || l2 != nil || jz != 0 {
		if l1 != nil {
			jz += l1.Val
		}

		if l2 != nil {
			jz += l2.Val
		}

		cur = addNode(cur, jz%10)
		jz = jz / 10

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}

	}

	return head.Next
}
