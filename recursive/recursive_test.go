package recursive_test

import (
	"fmt"
	"learning/recursive"
)

func ExampleFibonacci() {
	fmt.Println(recursive.Fibonacci(10))
	// Output: 55
}

func ExamplePrintNumber() {
	result := recursive.PrintNumber(12345)
	fmt.Println(result)
	// Output: [12345]
}

func ExamplePow() {
	fmt.Println(recursive.Pow(2, 10))
	// Output: 1024
}

func ExampleFactorial() {
	fmt.Println(recursive.Factorial(5))
	// Output: 120
}

func ExampleStrlen() {
	fmt.Println(recursive.Strlen("abcdef"))
	// Output: 6
}

func ExampleReverseString() {
	fmt.Println(recursive.ReverseString("abcdefghijklmn"))
	// Output: nmlkjihgfedcba
}

func ExampleDigitSum() {
	fmt.Println(recursive.DigitSum(12345))
	// Output: 15
}
