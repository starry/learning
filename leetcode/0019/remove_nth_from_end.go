package leetcode

import "learning/util"

// 给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
func removeNthFromEnd(head *util.ListNode, n int) *util.ListNode {
	var fast = head

	for i := n; i > 0; i-- {
		fast = fast.Next
	}

	if fast == nil {
		return head.Next
	}

	var slow = head

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return head
}
