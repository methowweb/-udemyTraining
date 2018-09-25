package main

import "fmt"

const p string = "death and taxes"
const (
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
)

func main() {

	const q = 42

	fmt.Println("p - ", p)
	fmt.Println("q - ", q)

	/*fmt.Println("binary\t\tdecimal")
	fmt.Printf("%b\t", kb)
	fmt.Printf("%d\n", kb)
	fmt.Printf("%b\t", mb)
	fmt.Printf("%d\n", mb)*/

	var meters string
	var userinput string
	fmt.Print("Enter something ")
	fmt.Scan(&meters)

	fmt.Println("You typed ", meters)

	//fmt.Println(meters)
	fmt.Print("Type something and Press enter to exit ")
	fmt.Scan(&userinput)
	fmt.Print("Finished ")

}
