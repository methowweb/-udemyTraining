package funcfolder2



import "fmt"

func FirstLoop() {
	var foo int
	fmt.Println(foo)
	foo ++
	fmt.Println(foo, "end")
	for i :=1; i<=10; i++ {
		fmt.Println(i)
	}
}