package main

import (
	"fmt"
)

// func main() {
// 	s := []int{5, 7, 1, 3, 8, 5, 9}
// 	//sorted in the order of the index
// 	sort.Ints(s)
// 	fmt.Println(s)
// 	//to find the index
// 	fmt.Println(strings.Index("Geeksforgeeks", "ks"))

// 	//to find the time
// 	fmt.Println("time:", time.Now().Unix())
// }

//init function

func init() {
	fmt.Println("hello init function1")
}

func init()  {
	fmt.Println("hello init function2")
}

func main() {
	fmt.Println("hello i am a main function")
}
