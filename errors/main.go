package main

import (
	"errors"
	"fmt"
)

func main() {
	err1 := errors.New("error1")
	err2 := errors.New("error2")
	err3 := fmt.Errorf("%s.%w", err1, err2)
	fmt.Println(err3)
}
