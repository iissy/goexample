// 本样例测试循环中，开期携程的，携程调用的函数里面使用的变量是由函数本事传入，还是直接使用循环变量
// 在函数体中直接使用循环变量，可能用的是瞬时的循环变量，而非对应的那次循环对应的变量值
// 所以，如何函数体内需要使用循环变量，必须通过函数的实参带入
package main

import (
	"fmt"
	"time"
)

func main() {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, i := range list {
		fun := func(o int) {
			fmt.Printf("Worker start process task: %d\n", o)
		}

		go fun(i)
	}

	ticker := time.NewTicker(3 * time.Second)
	for d := range ticker.C {
		fmt.Println(d)
	}
}
