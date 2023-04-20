package main

import "fmt"

func main() {
	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[3] = "Banana"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Size of fruit list is: ", len(fruitList))

	// var vegList = [3]string{"potato, tomato, spinach"}
	var vegList = [5]string{"potato, tomato, spinach"}

	fmt.Println("Size of veggie list is: ", len(vegList))

}
