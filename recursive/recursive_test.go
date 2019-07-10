package recursive_test

import (
	"fmt"

	"../recursive"
)

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
