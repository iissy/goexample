package main

import (
	"log"
	"time"
)

func main() {
	t := time.Now()
	time.Sleep(5 * time.Second)
	u := time.Now()

	if u.After(t) {
		log.Printf("t(%s) < u(%s)", t, u)
	}
}
