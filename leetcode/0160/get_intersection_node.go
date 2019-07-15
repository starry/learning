package leetcode

import "learning/util"

/**
 *  找出两个链表的交点
 *  要求时间复杂度为 O(N)，空间复杂度为 O(1)。如果不存在交点则返回 null。
 *  当访问 A 链表的指针访问到链表尾部时，令它从链表 B 的头部开始访问链表 B；
 *  同样地，当访问 B 链表的指针访问到链表尾部时，令它从链表 A 的头部开始访问链表 A。
 *	这样就能控制访问 A 和 B 两个链表的指针能同时访问到交点。
 */
func getIntersectionNode(headA, headB *util.ListNode) *util.ListNode {
	l1 := headA
	l2 := headB

	for l1 != l2 {
		if l1 == nil {
			l1 = headB
		} else {
			l1 = l1.Next
		}
		if l2 == nil {
			l2 = headA
		} else {
			l2 = l2.Next
		}
	}

	return l1
}
