package main

import (
	"fmt"
)

func main() {

	//var i int
	var a
	for i := 0; i < 0; i++ { // commented out by setting condition to less than zero
		fmt.Println(i)
		fmt.Println(a)
		a = &i
		i++
	}
	a := "hello "
	a=a+a
	letter := rune("ABC"[2])
	fmt.Println(letter)
	fmt.Println(letter, string(letter)) // convert rune to string
	fmt.Println(a)


}



