package main

import "fmt"

func main() {

	i := 0
	for {
		i++
		if i%1000 != 0 {
			continue
		}
		fmt.Println(i)
		fmt.Println("remainder", i%1000)
		if i >= 50000 {
			break
		}

	}

}
