package leetcode

// 设计循环队列

// MyCircularQueue 自制循环队列
type MyCircularQueue struct {
	Data []int
	Size int
	Head int
	Tail int
}

// Constructor Initialize your data structure here. Set the size of the queue to be k.
func Constructor(k int) MyCircularQueue {
	q := MyCircularQueue{
		Data: make([]int, k),
		Head: -1,
		Tail: -1,
		Size: k,
	}
	return q
}

// EnQueue Insert an element into the circular queue. Return true if the operation is successful.
func (q *MyCircularQueue) EnQueue(value int) bool {
	if q.IsFull() == true {
		return false
	}

	if q.IsEmpty() == true {
		q.Head = 0
	}

	q.Tail = (q.Tail + 1) % q.Size
	q.Data[q.Tail] = value
	return true
}

// DeQueue Delete an element from the circular queue. Return true if the operation is successful.
func (q *MyCircularQueue) DeQueue() bool {
	if q.IsEmpty() == true {
		return false
	}

	if q.Head == q.Tail {
		q.Head = -1
		q.Tail = -1
		return true
	}
	q.Head = (q.Head + 1) % q.Size
	return true
}

// Front Get the front item from the queue.
func (q *MyCircularQueue) Front() int {
	if q.IsEmpty() == true {
		return -1
	}
	return q.Data[q.Head]
}

// Rear Get the last item from the queue.
func (q *MyCircularQueue) Rear() int {
	if q.IsEmpty() == true {
		return -1
	}
	return q.Data[q.Tail]
}

// IsEmpty Checks whether the circular queue is empty or not.
func (q *MyCircularQueue) IsEmpty() bool {
	return q.Head == -1
}

// IsFull Checks whether the circular queue is full or not.
func (q *MyCircularQueue) IsFull() bool {
	return ((q.Tail + 1) % q.Size) == q.Head
}
