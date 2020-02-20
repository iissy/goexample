package main

import "fmt"

func main() {
	var x uint8 = 1<<1 | 1<<4 | 1<<7
	var y uint8 = 1<<4 | 1<<6

	fmt.Printf("%08b\n", x)    // 10010010，代表集合 {1, 4, 7}
	fmt.Printf("%08b\n", y)    // 01010000，代表集合 {4, 6}
	fmt.Printf("%08b\n", x&y)  // 00010000，代表交集 {4}
	fmt.Printf("%08b\n", x|y)  // 11010010，代表并集 {1, 4, 6, 7}
	fmt.Printf("%08b\n", x^y)  // 11000010，代表对称差 {1, 6, 7}
	fmt.Printf("%08b\n", x&^y) // 10000010，代表差集 {1, 7}

	// 打印集合 x 将输出 1 4 7
	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 {
			fmt.Println(i)
		}
	}
	// 打印集合 y 将输出 4 6
	for i := uint(0); i < 8; i++ {
		if y&(1<<i) != 0 {
			fmt.Println(i)
		}
	}
}
