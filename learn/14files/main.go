package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "This needs to go in a file"

	file, err := os.Create("./letsgo.txt")

	if err != nil {
		panic(err)
	}

	length, error := io.WriteString(file, content)
	if error != nil {
		panic(err)
	}

	fmt.Println("length is: ", length)
	defer file.Close()
	readFile("./letsgo.txt")
}

func readFile(fileName string) {
	databyte, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	fmt.Println("Text data inside the file is:\n", string(databyte))
}
