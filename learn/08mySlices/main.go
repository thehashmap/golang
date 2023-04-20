package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList = []string{}
	fmt.Printf("Type of fruitlist is %T\n", fruitList)

	// fruitList = append(fruitList, "Mango", "Kiwi", "Apple")
	fruitList = append(fruitList, "Mango", "Kiwi", "Apple")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 534
	highScores[2] = 264
	highScores[3] = 238
	// highScores[4] = 232 // gives error

	highScores = append(highScores, 267, 455, 564) // doesn't give error as append reallocates the memory so everything will be accomodated

	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(highScores)

}
