package main

import (
	"fmt"
	"math/rand"
	"time"
)

const code = "0123456789ABCDEFGHIJKLMNOPQRSTUVXWYZabcdefghijklmnopqrstuvxwyz-*"

func main() {
	fmt.Println(Random62String(20))
	fmt.Println("")
	fmt.Println(Random62String(20))
	fmt.Println("")
	fmt.Println(Random62String(20))
	fmt.Println("")
	fmt.Println(Random62String(20))
	fmt.Println("")
	fmt.Println(Random62String(20))
}

func randomString(size int, max int) string {
	seed := time.Now().UnixNano()
	rand.Seed(seed)
	fmt.Println(seed)
	buffer := make([]byte, size, size)
	for i := 0; i < size; i++ {
		buffer[i] = code[rand.Intn(max)]
	}
	return string(buffer[:size])
}

func Random62String(size int) string {
	return randomString(size, 62)
}
