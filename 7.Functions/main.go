package main

import "fmt"

func main() {
	fmt.Println("welcome to function in golang")
	fmt.Println(add(20, 30))
	simplefunction()
}

func simplefunction() {
	fmt.Println("hello from simple function")
}

func add(x, y int) int {
	return x + y
}
