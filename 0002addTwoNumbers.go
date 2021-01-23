package leecode

//给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
//
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
//
// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
//
//
// 示例 1：
//
//
//输入：l1 = [2,4,3], l2 = [5,6,4]
//输出：[7,0,8]
//解释：342 + 465 = 807.
//
//
// 示例 2：
//
//
//输入：l1 = [0], l2 = [0]
//输出：[0]
//
//
// 示例 3：
//
//
//输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
//输出：[8,9,9,9,0,0,0,1]
//
//
//
//
// 提示：
//
//
// 每个链表中的节点数在范围 [1, 100] 内
// 0 <= Node.val <= 9
// 题目数据保证列表表示的数字不含前导零
//
// Related Topics 递归 链表 数学
// 👍 5534 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
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
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
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
}

func addNode(l *ListNode, v int) *ListNode {
	l.Next = &ListNode{Val: v}
	return l.Next
}

func addTwoNumbers3(l1 *ListNode, l2 *ListNode) *ListNode {
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
//leetcode submit region end(Prohibit modification and deletion)
