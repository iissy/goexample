package main

import (
	"bytes"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	// 定义集合 x, y，并添加整型元素
	var x, y IntSet
	x.Add(168)
	x.Add(1)
	x.Add(9)
	x.Add(88)
	x.Add(788)
	x.Add(10000)
	fmt.Printf("集合x：\t%s\n", x.String())

	y.Add(0)
	y.Add(9)
	y.Add(89)
	y.Add(168)
	y.Add(999)
	fmt.Printf("集合y：\t%s\n", y.String())

	x.DifferenceSet(&y)
	fmt.Printf("xy的并集：\t%s\n", x.String())

	fmt.Printf("xy的并集是否包含9：%t；xy的并集是否包含123：%t", x.Has(9), x.Has(123))

	fmt.Printf("%s", time.Since(start))
}

// 定义一个无符号的整型集合（小的非负整数）
type IntSet struct {
	words []uint64
}

// 集合是否包含元素
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// 添加集合元素
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// 计算两个集合的并集
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) Intersection(t *IntSet) {
	result := new(IntSet)
	for i, w := range t.words {
		if i < len(s.words) {
			result.words = append(result.words, w&s.words[i])
		}
	}

	s.words = result.words
}

func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) DifferenceSet(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		}
	}
}

// 输出集合格式这样 {1 4 7}
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
