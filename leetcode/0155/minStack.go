package leetcode

import "math"

type MinStack struct {
	Data []int
	Help []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{[]int{}, []int{math.MaxInt32}}
}

func (this *MinStack) Push(x int) {

	this.Data = append(this.Data, x)

	if this.Help[len(this.Help)-1] >= x {
		this.Help = append(this.Help, x)
	} else {
		this.Help = append(this.Help, this.Help[len(this.Help)-1])
	}

}

func (this *MinStack) Pop() {
	this.Data = this.Data[0 : len(this.Data)-1]
	this.Help = this.Help[0 : len(this.Help)-1]
}

func (this *MinStack) Top() int {
	return this.Data[len(this.Data)-1]
}

func (this *MinStack) GetMin() int {
	return this.Help[len(this.Help)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
