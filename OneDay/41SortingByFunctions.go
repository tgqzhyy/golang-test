package main

import (
	"fmt"
	"sort"
)

type byLength []string

//func (b byLength) Len() int {
//	panic("implement me")
//}
//
//func (b byLength) Less(i, j int) bool {
//	panic("implement me")
//}
//
//func (b byLength) Swap(i, j int) {
//	panic("implement me")
//}

func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	fmt.Println(byLength(fruits))
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}
