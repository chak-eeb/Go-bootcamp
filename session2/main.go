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
// }

func main() {
	div, err := divide(5, 3)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("Result: ", div)
	fmt.Println("the int value of that string is: ", convertToInt("3597"))
}
