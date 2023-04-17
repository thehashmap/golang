package main

import "fmt"

const LoginToken string = "hemlo" // public

func main() {
	var username string = "cheems"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 253.45645349
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// no var style
	numberOfUser := 3000 // this walrus operator is not allowed main func
	fmt.Println(numberOfUser)
}
