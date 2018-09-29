package main

import "fmt"

func main() {

	fmt.Println("hello again")
	nextLine()
	third := thirdLine
	third()
	fourthLine(22)
}

func nextLine() {

	fmt.Println("next line")
}

func thirdLine() {

	fmt.Println("third line")
}

func fourthLine(x int) {

	fmt.Println("fourth line-- ", x)
}
