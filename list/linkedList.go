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

	list = addNodeEnd(v2, list)
	list = addNodeEnd(v3, list)
	list = addNodeEnd(v4, list)
	list = reverseList(list)
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

// func reverseList(head *ListNode) *ListNode {
// 	if head == nil {
// 		return head
// 	}

// 	cur := head

// 	if cur.Next == nil {
// 		return cur
// 	}
// 	newHead := reverseList(cur.Next)
// 	cur.Next.Next = cur 翻转链表的指向
// 	cur.Next = nil 记得赋值NULL，防止链表错乱
// 	return newHead 新链表头永远指向的是原链表的链尾
// }

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
