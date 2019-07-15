package queue_test

import (
	"fmt"
	"learning/queue"
)

func ExampleMyQueue() {
	var myQueue queue.MyQueue

	myQueue.EnQueue(5)
	myQueue.EnQueue(3)

	if myQueue.IsEmpty() == false {
		fmt.Println(myQueue.Font())
	}

	myQueue.DeQueue()

	if myQueue.IsEmpty() == false {
		fmt.Println(myQueue.Font())
	}

	myQueue.DeQueue()

	if myQueue.IsEmpty() == false {
		fmt.Println(myQueue.Font())
	}
	// Output:
	// 5
	// 3
}
