// 2937429 将输出 2,937,429
package main

import "fmt"

func main() {
	fmt.Println(comma("2937429"))
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	return comma(s[:n-3]) + "," + s[n-3:]
}
