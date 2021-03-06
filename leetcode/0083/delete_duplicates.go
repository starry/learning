package leetcode

import "learning/util"

// 给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。
// head.next = head.next.next
func deleteDuplicates(head *util.ListNode) *util.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	head.Next = deleteDuplicates(head.Next)

	if head.Val == head.Next.Val {
		return head.Next
	}

	return head
}

func deleteDuplicates2(head *util.ListNode) *util.ListNode {
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == cur.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}
