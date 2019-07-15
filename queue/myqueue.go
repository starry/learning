package queue

// MyQueue 是一个简单的队列
type MyQueue struct {
	Data  []int
	Start int
}

// EnQueue 队列末尾添加元素
func (q *MyQueue) EnQueue(x int) bool {
	q.Data = append(q.Data, x)
	return true
}

// DeQueue 队列头删除元素
func (q *MyQueue) DeQueue() bool {
	if q.IsEmpty() == true {
		return false
	}
	q.Start++
	return true
}

// Font 返回队列的头
func (q *MyQueue) Font() int {
	return q.Data[q.Start]
}

// IsEmpty 队列是否为空
func (q *MyQueue) IsEmpty() bool {
	return q.Start >= len(q.Data)
}

// func main() {
// 	var myQueue MyQueue

// 	myQueue.EnQueue(5)
// 	myQueue.EnQueue(3)

// 	if myQueue.IsEmpty() == false {
// 		fmt.Println(myQueue.Font())
// 	}

// 	myQueue.DeQueue()

// 	if myQueue.IsEmpty() == false {
// 		fmt.Println(myQueue.Font())
// 	}

// 	myQueue.DeQueue()

// 	if myQueue.IsEmpty() == false {
// 		fmt.Println(myQueue.Font())
// 	}
// }
