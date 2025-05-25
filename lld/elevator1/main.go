package main

import (
	"fmt"
)

func main() {
	fmt.Println("Elevator system design")

	for {

		var number int
		fmt.Println("Enter a number 1 for External button")
		fmt.Println("Enter a number 2 for Internal button")
		fmt.Println("Enter a number 3 for Step up")
		fmt.Println("Enter a number 4 for exit")
		fmt.Println("Enter a number...")
		if _, err := fmt.Scan(&number); err != nil {
			panic(err)
		}
		fmt.Println("number ", number)

		switch number {
		case 1:
			fmt.Println("EExternal button")
		case 2:
			fmt.Println("Internal button")
		case 3:
			fmt.Println("Stepn")
		case 4:
			fmt.Println("Exit button")
		default:
			fmt.Println("Retry")
		}

		if number == 4 {
			break
		}

	}

	fmt.Println("Basic Iteration complete")
}
