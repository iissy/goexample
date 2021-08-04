package main

import (
	"fmt"
	"goexample/types"
)

func main() {
	p := types.Point{Dian: types.Dian{X: 1, Y: 2}}
	q := types.Point{Dian: types.Dian{X: 4, Y: 6}}

	fmt.Println(types.Point.Distance(p, q))
}
