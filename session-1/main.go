package main

import (
	"fmt"
	"unicode/utf8"
)

const maxStringLength = 10
const (
	Monday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func stringLength(str string) int {
	fmt.Println("--- Task 1: Variable Declaration and String Operations ---")
	return utf8.RuneCountInString(str)
}

func compareStrings(a, b string) {
	fmt.Println("--- Task 2: If-Else Conditions ---")

	if stringLength(a) > stringLength(b) {
		fmt.Printf("The length of string: %s is higher than the length of string %s \n", a, b)
	} else if stringLength(b) > stringLength(a) {
		fmt.Printf("The length of string: %s is higher than the length of string %s \n ", b, a)
	} else {
		fmt.Println("The 2 strings are equal in length ")

	}

}

func printDay(day) {

}

func main() {
	var str1 string = "dog"
	var str2 string = "apple"
	str3 := "cat"
	fmt.Println("Length of string 1: ", stringLength(str1))
	fmt.Println("Length of string 2: ", stringLength(str2))
	fmt.Println("Length of string 3: ", stringLength(str3))

	compareStrings(str1, str2)
	compareStrings(str2, str3)
	compareStrings(str1, str3)

}
