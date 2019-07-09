package main

import "fmt"

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	v1 := &ListNode{0, nil}
	list := v1
	v2 := &ListNode{1, nil}
	v3 := &ListNode{2, nil}
	v4 := &ListNode{3, nil}
	v5 := &ListNode{4, nil}
	v6 := &ListNode{4, nil}
	v7 := &ListNode{5, nil}

	list = addNodeEnd(v2, list)
	list = addNodeEnd(v3, list)
	list = addNodeEnd(v4, list)
	list = addNodeEnd(v5, list)
	list = addNodeEnd(v6, list)
	list = addNodeEnd(v7, list)
	list = deleteDuplicates(list)
	printList(list)
}

func printList(list *ListNode) {
	for l := list; l != nil; l = l.Next {
		fmt.Println(l)
	}
}

func addNodeEnd(v, list *ListNode) *ListNode {
	if list == nil {
		return list
	}
	for l := list; l != nil; l = l.Next {
		if l.Next == nil {
			l.Next = v
			return list
		}
	}
	return list
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	fmt.Printf("start %v\n", head)
	head.Next = deleteDuplicates(head.Next)
	fmt.Printf("end %v\n", head)
	if head.Val == head.Next.Val {
		return head.Next
	}

	return head
}

func reverseList(head *ListNode) *ListNode {
	var result *ListNode
	for {
		if head == nil {
			break
		}
		// 暂存下一个地址，防止变化指针指向后找不到后续的数
		tmp := head.Next
		// 把指针指向指向前一个空间
		head.Next = result
		// 新链表的头移动到head，扩长一步链表
		result = head
		// head指向原始链表head指向的下一个空间
		head = tmp
	}

	return result
}
