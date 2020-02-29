package main

import (
	"fmt"
	"sort"
)

func main() {
	groups := OneGroups{
		{"d8f4f", "Golang", 30},
		{"d9f4a", "Python", 20},
		{"u8f3n", "Java", 10},
		{"t7y68", "Ruby", 25}}

	// 正序
	sort.Sort(groups)
	for _, g := range groups {
		fmt.Printf("%v\t", g)
	}

	fmt.Println("")
	// 反转
	sort.Sort(sort.Reverse(groups))
	for _, g := range groups {
		fmt.Printf("%v\t", g)
	}
}

type OneGroup struct {
	CatId string
	Name  string
	Size  int
}

type OneGroups []*OneGroup

func (s OneGroups) Len() int           { return len(s) }
func (s OneGroups) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s OneGroups) Less(i, j int) bool { return s[i].Size < s[j].Size }
