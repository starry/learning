package main

import "fmt"

// 编写一个函数求n的阶乘
// 在数学中，正整数的阶乘（英语：factorial）是所有小于及等于该数的正整数的积，计为n!，例如5的阶乘计为5!，其值为120：
// 并定义，1的阶乘1!为1、0的阶乘0!亦为1，其中，0的阶乘表示一个空积[2]。
func factorial(n int64) int64 {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

// 编写一个函数，可以分别打印一个整数十进制的每一位
func printNumber(n int64) {
	if n > 9 {
		printNumber(n / 10)
	}
	fmt.Println(n % 10)
}

func main() {
	// printNumber(12345)

	// fmt.Print("Enter a number: ")
	// var n int64
	// fmt.Scanln(&n)
	// f := factorial(n)
	// fmt.Printf("%d的阶乘是%d\n", n, f)
}
