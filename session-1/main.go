package main

import (
	"fmt"
	"unicode/utf8"
)

type Day int

const maxStringLength = 10
const (
	Monday Day = iota + 1
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

func printDay(day Day) {
	fmt.Println("--- Task 3: Constants and IOTA ---")
	fmt.Println("Max string length is: ", maxStringLength)
	fmt.Println("--- Days of the Week ---")
	if day == Monday {
		fmt.Println("Monday: ", Monday)
	} else if day == Tuesday {
		fmt.Println("Tuesday: ", Tuesday)
	} else if day == Tuesday {
		fmt.Println("Wednesday: ", Wednesday)
	} else if day == Tuesday {
		fmt.Println("Thursday: ", Thursday)
	} else if day == Tuesday {
		fmt.Println("Friday: ", Friday)
	} else if day == Tuesday {
		fmt.Println("Saturday: ", Saturday)
	} else {
		fmt.Println("Sunday: ", Sunday)
	}
}

func main() { // main function is the entry point of the program. everything in this function will be executed
	var str1 string = "dog"
	var str2 string = "apple"
	str3 := "cat"
	fmt.Println("Length of string 1: ", stringLength(str1))
	fmt.Println("Length of string 2: ", stringLength(str2))
	fmt.Println("Length of string 3: ", stringLength(str3))

	compareStrings(str1, str2)
	compareStrings(str2, str3)
	compareStrings(str1, str3)

	printDay(Tuesday)

}
