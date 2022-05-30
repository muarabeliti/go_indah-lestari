package main

import (
	"fmt"
	"string"
)
type pair struct {
	name string
	count int
}

func MostAppearItem(items []string, f func(string) bool) []string {
    vsf := make([]string, 0)
    for _, v := range vs {
        if f(v) && len(v) > 7 {
            vsf = append(vsf, v)
        }
    }
    return vsf
}

func main() {

    var strs = []string{"foo1", "foo2", "foo3", "foo3", "foo_testfor", "_foo"}

    fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	fmt.Println(MostAppearItem([]string{"footbal", "basketball", "tenis"}))
}