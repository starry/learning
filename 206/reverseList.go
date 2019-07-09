package reverselist

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
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
