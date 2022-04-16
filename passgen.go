package main

import (
	"errors"
	"fmt"
	"regexp"
)

func main() {
	// var length, digitsCount, symbolsCount int
	// var isUpperCaseAllowed string

	// fmt.Print("Enter password length(default=16): ")
	// fmt.Scanf("%d\n", &length)
	// if length == 0 {
	// 	length = 16
	// }

	// fmt.Print("Enter digits count(default=10): ")
	// fmt.Scanf("%d\n", &digitsCount)
	// if digitsCount == 0 {
	// 	digitsCount = 10
	// }

	// fmt.Print("Enter symbols count(default=10): ")
	// fmt.Scanf("%d\n", &symbolsCount)
	// if symbolsCount == 0 {
	// 	symbolsCount = 10
	// }

	checkInput("")
	// fmt.Print("Is uppercase allowed?([y]es/[n]o. default=yes): ")
	// fmt.Scanf("%s\n", &isUpperCaseAllowed)
	// println(len(isUpperCaseAllowed))

	// res, err := password.Generate(length)
}

func checkInput(input string) (output string, err error) {
	match, err := regexp.MatchString(`(?i)^y$|^yes$|^n$|^no$`, input)
	if match {
		output = input
	} else {
		err = errors.New(fmt.Sprintf("Not valid input: %s", input))
	}
	return
}
