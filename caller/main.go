package main

import (
	"fmt"
	"runtime"
)

func main() {
	Ping()
}

func Ping() {
	Caller()
}

func Caller() {
	_, file, line, _ := runtime.Caller(2)
	fmt.Printf("%s: %d", file, line)
}
