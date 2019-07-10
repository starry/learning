package recursive

// Factorial 编写一个函数求n的阶乘
// 在数学中，正整数的阶乘（英语：factorial）是所有小于及等于该数的正整数的积，计为n!，例如5的阶乘计为5!，其值为120：
// 并定义，1的阶乘1!为1、0的阶乘0!亦为1，其中，0的阶乘表示一个空积[2]。
func Factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * Factorial(n-1)
}

// PrintNumber 编写一个函数，可以分别打印一个整数十进制的每一位
func PrintNumber(n int) []int {
	var result []int
	if n > 9 {
		PrintNumber(n / 10)
	}
	result = append(result, n)
	return result
}

// Pow 编写一个函数实现n^k，使用递归实现
func Pow(n, k int) int {
	if k == 0 {
		return 1
	} else if k == 1 {
		return n
	} else {
		return n * Pow(n, k-1)
	}
}

// Strlen 不允许创建临时变量求字符串长度，实现strlen的模拟
func Strlen(str string) int {
	runes := []rune(str)
	if str == "" {
		return 0
	}
	substring := string(runes[1:])
	return Strlen(substring) + 1
}
