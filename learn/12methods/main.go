package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	// no inheritance in golang; No super or parent

	cheems := User{"Cheems", "cheems@hemlo.bonk", true, 6}
	fmt.Println(cheems)
	fmt.Printf("Cheems details are: %+v\n", cheems)
	fmt.Printf("Name is: %v and email is %v\n", cheems.Name, cheems.Email)

	cheems.GetStatus()
	cheems.NewMail()
	fmt.Printf("Name is: %v and email is %v\n", cheems.Name, cheems.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "cheemsReturns@hemlo.bonk"
	fmt.Println("Email of this user is: ", u.Email)
}
