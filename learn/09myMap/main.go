package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["PY"] = "Python"

	fmt.Println("List of all languages: ", languages)

	delete(languages, "PY")
	fmt.Println("List of all languages: ", languages)

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
