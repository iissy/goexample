package main

import "fmt"

func main() {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	list = remove(list, 5)
	fmt.Println(list)

	list = remove2(list, 3)
	fmt.Println(list)
}

// 去掉一个元素，顺序保持不变
func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

// 如果不考虑去掉一个元素后的顺序
func remove2(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
