package leetcode

import "learning/util"

func reverseList(head *util.ListNode) *util.ListNode {
	if head == nil {
		return head
	}

	cur := head

	if cur.Next == nil {
		return cur
	}
	newHead := reverseList(cur.Next)
	cur.Next.Next = cur
	cur.Next = nil
	return newHead
}
