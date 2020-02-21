package main

import "fmt"

func main() {
	s := []int{5, 6, 7, 8, 9}
	fmt.Println(remove(s, 2))
}

func remove(slice []int, i int) []int {
	a := slice[i:]
	b := slice[i+1:]
	u := copy(a, b)
	fmt.Println(u)
	return slice[:len(slice)-1]
}
