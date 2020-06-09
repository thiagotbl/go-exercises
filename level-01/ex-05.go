package main

import "fmt"

type myType int

var x myType
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 42

	fmt.Printf("%v\n", x)
	fmt.Printf("%T\n", x)

	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
