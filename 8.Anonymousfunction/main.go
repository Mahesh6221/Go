package main

import "fmt"

// func main() {
// 	//Anonymous function: a function which does not have any name is called as anonymous function
// 	// Anonymous function is also called as function literal
// 	func() {
// 		fmt.Println("hello i am a anonymous function")
// 	}()
// }

// In Go language, you are allowed to assign an anonymous function to a variable. When you assign a function to a variable, then the type of the variable is of function type and you can call that variable like a function call as shown in the below

// func main() {
// 	val := func() {
// 		fmt.Println("This is second anonymous function in golang")
// 	}
// 	val()
// }

// You can also pass arguments in the anonymous function

func main() {
	func(ele string) {
		fmt.Println(ele)
	}("Geeks for geeks")
}
