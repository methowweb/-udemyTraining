package main

import (
	"fmt"
	"sort"
)

func main() {

	s := []int{5, 2, 6, 3, 1, 4}
	fmt.Println(s)
	fmt.Printf("%t \n", s)
	n:=sort.IntSlice(s)
	fmt.Println(s)

sort.Sort(n)
	fmt.Println(n)
	fmt.Printf("%t \n", s)
	fmt.Println(s)

}git s