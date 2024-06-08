package main

import "fmt"

func simplefunction() {
	fmt.Println("hello from simple function")
}


// Functions are generally the block of codes or statements in a program that gives the user the ability to reuse the same code which ultimately saves the excessive use of memory

//function with parameters and return type
func add(x, y int) int {
	//input parameters are x and y and output parameter
	//is int because the output in the
	//integer format.
	return x + y
}

//function with named return variable
func multiply(x, y int) (result int) {
	//i already said that the variable is result which gets returned
	//the result is written for the output
	// result is the return type we return the result variable
	result = x * y //here we are performing some operation
	return         //here the returned value should be a result having datatype int
}

func subtract(x, y int) (result int) {
	result = x - y
	return
}

func main() {
	fmt.Println("welcome to function in golang")
	fmt.Println(add(20, 30))
	simplefunction()
	ans := add(3, 4)
	fmt.Println(ans)
	product := multiply(4, 6)
	fmt.Println(product)
	minus := subtract(5, 3)
	fmt.Println(minus) 
}
