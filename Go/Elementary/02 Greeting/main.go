package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print("Enter your name: ")
	var input string
	fmt.Scanf("%s", &input)

	var output string

	if strings.ToLower(input) == "alice" || strings.ToLower(input) == "bob" {
		output = "Hello, " + input + "!"
	} else {
		output = "Who are you?"
	}

	fmt.Println(output)

}
