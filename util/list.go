package util

import "fmt"

// ListNode finition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// PrintList 打印链表
func PrintList(list *ListNode) {
	for l := list; l != nil; l = l.Next {
		fmt.Println(l)
	}
}

// AddNodeEnd 添加节点
func AddNodeEnd(v, list *ListNode) *ListNode {
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
