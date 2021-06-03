package main

import "fmt"

func main() {
	newTowerLevel := 0
	newTowerLevel |= 2
	fmt.Println(newTowerLevel)
	newTowerLevel |= 1
	fmt.Println(newTowerLevel)
}
