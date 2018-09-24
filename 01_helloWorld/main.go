package main

import "fmt"

var x = 8

func main() {
	y := 5
	z := 4
	fmt.Println("hello World")
	for i := 0; i < 4; i++ {
		// fmt.Printf("%d \n", i)
		// test line
		//test line 2 plus more and backcheck
		//fmt.Println(y)
		fmt.Printf("\t %d next   a     %d \n", z, y)
		fmt.Printf("\t %d next  %d \n", z, y)
		//fmt.Printf("\t %v \t %v \n", z, y)
	}
	//fmt.Println(x)
}
