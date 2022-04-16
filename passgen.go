package main

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/sethvargo/go-password/password"
)

func main() {
start:
	length := 16
	digitsCount := 4
	symbolsCount := 2
	isUpperCaseAllowed := true
	isRepeatAllowed := false

	fmt.Print("Enter password length(default=16): ")
	fmt.Scanf("%d\n", &length)

	fmt.Print("Enter digits count(default=4,max=10): ")
	fmt.Scanf("%d\n", &digitsCount)
	if digitsCount > 10 {
		digitsCount = 10
	}

	fmt.Print("Enter symbols count(default=2): ")
	fmt.Scanf("%d\n", &symbolsCount)

	for {
		var ucErr error
		var isUcAllowed string
		fmt.Print("Is uppercase allowed?([y]es/[n]o. default=yes): ")
		fmt.Scanf("%s\n", &isUcAllowed)
		if len(isUcAllowed) == 0 {
			break
		}
		isUpperCaseAllowed, ucErr = isYes(isUcAllowed)
		if ucErr == nil {
			break
		} else {
			fmt.Println(ucErr)
		}
	}

	for {
		var ira string
		var rErr error
		fmt.Print("Are repeated chars allowed?([y]es/[n]o. default=no): ")
		fmt.Scanf("%s\n", &ira)
		if len(ira) == 0 {
			break
		}
		isRepeatAllowed, rErr = isYes(ira)
		if rErr == nil {
			break
		} else {
			fmt.Println(rErr)
		}
	}

	res, err := password.Generate(16, digitsCount, symbolsCount, !isUpperCaseAllowed, isRepeatAllowed)

	if err == nil {
		fmt.Println(res)
	} else {
		fmt.Printf("error: %v\nPress enter to restart.", err)
		fmt.Scanln()
		fmt.Println()
		goto start
	}

}

func isYes(input string) (bool, error) {
	match, err := regexp.MatchString(`(?i)^y$|^yes$|^n$|^no$`, input)
	if match {
		return (input[0] == 'y' || input[0] == 'Y'), err
	} else {
		err = errors.New(fmt.Sprintf("Not valid input: %s", input))
	}
	return false, err
}
