package funcfolder2



import "fmt"

func FirstLoop() {
	var foo int
	var evenOrOdd string
	fmt.Println(foo)
	foo ++
	fmt.Println(foo, "end")
	for i :=1; i<=10; i++ {
		fmt.Println(i)
		fmt.Printf("the number is %v\n",i)
		if i%2==0 {
			evenOrOdd= "Even"
		}else {
			evenOrOdd= "Odd"
		}
		fmt.Printf("%v is an %v number  and the remainder is %v\n",i,evenOrOdd, i%2)
	}
}