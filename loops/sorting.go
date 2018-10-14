package main

import (
	"fmt"
	"sort"
)

func main() {

	s := []int{5, 2, 6, 3, 1, 4}
	fmt.Println(s)
	fmt.Printf("%t \n", s)
	n := sort.IntSlice(s)
	fmt.Println(s)

	sort.Sort(n)
	fmt.Println(n)
	fmt.Printf("%t \n", s)
	fmt.Println(s)

	// reverse sort
	sort.Sort(sort.Reverse(n))
	fmt.Println(n)

	// what is the length of n
	nLength := len(n)
	fmt.Println(nLength)

	// Can (len) be calculated on a variable that is not a (IntSlice)?
	sTwo := []int{2, 4, 6, 8, 10}
	fmt.Println(sTwo)
	sTwoLength := len(sTwo)
	fmt.Println(sTwoLength)

	// Can (len) be used an an integer or string

	i := 7
	j := "yellow"
	fmt.Println(i)
	fmt.Println(j)
	//k := len(i)
	l := len(j)
	//fmt.Println(k)
	fmt.Println(l)
	fmt.Println(len(j))

}
