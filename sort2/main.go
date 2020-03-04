package main

import (
	"fmt"
	"sort"
)

func main() {
	groups := []*Article{
		{"Golang", "Hi, golang.", 30},
		{"Python", "Oh, python.", 20},
		{"Java", "SHit, java.", 19},
		{"Python", "Oh, python, It is good.", 43},
		{"Java", "SHit, java, it is bad.", 21},
		{"Golang", "Hi, golang, my love.", 13},
		{"Ruby", "Fuck, ruby.", 25}}

	sort.Sort(ArticleSort{groups, func(x, y *Article) bool {
		if x.Title != y.Title {
			return x.Title < y.Title
		}
		if x.Visited != y.Visited {
			return x.Visited < y.Visited
		}
		return false
	}})

	for _, g := range groups {
		fmt.Printf("%s\t%d\t%s\n", g.Title, g.Visited, g.Desc)
	}
}

type Article struct {
	Title   string
	Desc    string
	Visited int
}

type ArticleSort struct {
	g    []*Article
	less func(x, y *Article) bool
}

func (s ArticleSort) Len() int           { return len(s.g) }
func (s ArticleSort) Less(i, j int) bool { return s.less(s.g[i], s.g[j]) }
func (s ArticleSort) Swap(i, j int)      { s.g[i], s.g[j] = s.g[j], s.g[i] }
