// 斐波那契数列（Fibonacci sequence）
// F(1)=1，F(2)=1, F(n)=F(n - 1)+F(n - 2) (n ≥ 3)
package main

import "fmt"

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}

	return x
}

func main() {
	fmt.Println(fib(1))
	fmt.Println(fib(2))
	fmt.Println(fib(3))
	fmt.Println(fib(4))
	fmt.Println(fib(5))
}
