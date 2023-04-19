package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("02-01-2006 Monday"))

	createdDate := time.Date(2019, time.March, 9, 19, 69, 69, 69, time.UTC)

	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("02-01-2006 Monday"))
}
