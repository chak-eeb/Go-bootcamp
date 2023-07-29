package main

import (
	"errors"
	"fmt"
)

func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("can't divide by 0")
	}
	return x / y, nil
}

func main() {
	div, err := divide(5, 3)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("Result: ", div)
}
