package main

import "fmt"

/*
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.
示例 2：

输入：l1 = [0], l2 = [0]
输出：[0]
示例 3：

输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
输出：[8,9,9,9,0,0,0,1]


提示：

每个链表中的节点数在范围 [1, 100] 内
0 <= Node.val <= 9
题目数据保证列表表示的数字不含前导零
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// // 定义一个虚拟头节点，方便处理链表操作
	// dummy := &ListNode{}
	// // 当前节点指针，初始指向虚拟头节点
	// curr := dummy
	// // 进位值，初始为 0
	// carry := 0

	// // 当 l1、l2 还有节点，或者还有进位时，继续循环
	// for l1 != nil || l2 != nil || carry != 0 {
	// 	sum := carry // 先加上上一位的进位

	// 	// 如果 l1 还有节点，加上其值并移动指针
	// 	if l1 != nil {
	// 		sum += l1.Val
	// 		l1 = l1.Next
	// 	}
	// 	// 如果 l2 还有节点，加上其值并移动指针
	// 	if l2 != nil {
	// 		sum += l2.Val
	// 		l2 = l2.Next
	// 	}

	// 	// 计算当前位的值和新的进位
	// 	curr.Val = sum % 10
	// 	carry = sum / 10

	// 	// 如果还有后续节点要处理，创建新节点并移动指针
	// 	if l1 != nil || l2 != nil || carry != 0 {
	// 		curr.Next = &ListNode{}
	// 		curr = curr.Next
	// 	}
	// }

	dummy := &ListNode{}
	cur := dummy
	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		cur.Val = sum%10 + carry
		carry = sum / 10
		if l1 != nil || l2 != nil || carry != 0 {
			cur.Next = &ListNode{}
			cur = cur.Next
		}

	}

	return dummy
}

func addListNode(head *ListNode, val int) *ListNode {
	node := head
	for node.Next != nil {
		node = node.Next
	}
	node.Next = &ListNode{val, nil}
	return head
}
func main() {
	list1 := &ListNode{
		Val:  0,
		Next: nil,
	}
	list2 := &ListNode{
		Val:  0,
		Next: nil,
	}
	for i := 1; i <= 10; i++ {
		addListNode(list1, i)
		addListNode(list2, i)
	}
	dummy := addTwoNumbers(list1, list2)
	for dummy != nil {
		fmt.Println(dummy)
		dummy = dummy.Next
	}
}
