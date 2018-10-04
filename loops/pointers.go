package main

import "fmt"

// from the example in class

func main() {

	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a

	fmt.Println(b)

	// try with a string
	var c string
	c = "a string"

	fmt.Println(c)
	fmt.Println(&c)

	var d *string
	fmt.Println(d)
	d = &c
	fmt.Println(*d)

}
