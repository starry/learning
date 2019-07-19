package leetcode

// 自下向上
func dailyTemperatures(T []int) []int {
	var length = len(T)
	var result = make([]int, len(T))
	var stack []int
	for i := length - 1; i >= 0; i-- {
		for len(stack) != 0 && T[i] >= T[stack[len(stack)-1]] {
			stack = stack[0 : len(stack)-1]
		}

		if len(stack) == 0 {
			result[i] = 0
		} else {
			result[i] = stack[len(stack)-1] - i
		}

		stack = append(stack, i)
	}
	return result
}
