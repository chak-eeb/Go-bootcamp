package main

import (
	"errors"
	"fmt"
	"strconv"
)

func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("can't divide by 0")
	}
	return x / y, nil
}

// always handle errors it's best practice

func convertToInt(str string) int {
	num, err := strconv.Atoi(str) // strconv is a package that deals with strings it has many methods we can use
	if err != nil {
		fmt.Println("Error: ", err)
	}
	return num
}

// short if. let's change the statement above to a short statement
// if num, err := strconv.Atoi(str); err == nil{
// 	return num
// } else{
// 	fmt.Println("Error: ", err)

// For loop

func looping() {
	for i := 1; i <= 6; i++ {
		fmt.Println("i is: ", i)
	}
}

func main() {
	div, err := divide(5, 3)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("Result: ", div)
	fmt.Println("the int value of that string is: ", convertToInt("3597"))

	// Switch statement
	fruit := "kiwi"

	switch fruit {
	case "apple":
		fmt.Println("I like apple juice")
	case "banana":
		fmt.Println("i like to eat Bananas")
	case "strawberries":
		fmt.Println("i like them sweet")
	default:
		fmt.Println("i like to eat " + fruit + " too but i prefer apples and bananas")

	}

	looping()
}
