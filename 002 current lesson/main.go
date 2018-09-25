package main

import "fmt"

const p string = "death and taxes"
const ( _ = iota
		kb = 1 <<(iota *10)
		mb = 1 << (iota *10)

)

func main() {

	const q = 42

	fmt.Println("p - ", p)
	fmt.Println("q - ", q)

	fmt.Println("binary\t\tdecimal")
	fmt.Printf("%b\t", kb)
	fmt.Printf("%d\n",kb)
	fmt.Printf("%b\t",mb)
	fmt.Printf("%d\n",mb)




}
