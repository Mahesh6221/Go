package main

import "fmt"

func divide(a, b float64) (float64, error) { //error is the standard format which here i write in this line which return error
	if b == 0 {
		return 0, fmt.Errorf("denominator must not be zero")
	}
	return a / b, nil
}

// func main() {
// 	//In error handling we normally handle the errors as we do in other programming,but
// 	//in go we use _ underscore to handle the error it is called as blank identifier.
// 	fmt.Println("error handling in golang")
// 	div, _ := divide(10, 5)
// 	fmt.Println("division of two numbers", div)
// }

func main() {
	fmt.Println("error handling in golang")
	div, err := divide(10, 4)
	if err != nil {
		fmt.Println("error handling")
	}
	fmt.Println("division of two numbers", div)
} 
