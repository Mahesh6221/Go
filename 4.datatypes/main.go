package main

import (
	"fmt"
)

func main() {
	var x int = 32
	fmt.Println(x)
	fmt.Printf("the type of x is %T\n", x)
	var f float64 = 33.44333
	fmt.Println(f)
	fmt.Printf("the type of f is %T\n", f)
	var s string = "hello"
	fmt.Println(s)
	fmt.Printf("the type of f is %T\n", s)
}
