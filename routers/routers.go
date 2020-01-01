package routers

import (
	"fmt"
	. "goexample/intset"
)

// 位运算求交集、并集
func ExecIntSet() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(88)
	fmt.Printf("集合x：\t%s\n", x.String())

	y.Add(9)
	y.Add(42)
	fmt.Printf("集合y：\t%s\n", y.String())

	x.UnionWith(&y)
	fmt.Printf("xy的并集：\t%s\n", x.String())

	fmt.Printf("xy的并集是否包含9：%t；xy的并集是否包含123：%t", x.Has(9), x.Has(123))
}
