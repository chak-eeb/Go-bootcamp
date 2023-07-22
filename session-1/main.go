package main

import (
	"fmt"
	"unicode/utf8"
)

func stringLength(str string) int {
	return utf8.RuneCountInString(str)
}
func main() {
	var str1 string = "apple"
	var str2 string = "go"
	str3 := "cat"
	fmt.Println(stringLength(str1))
	fmt.Println(stringLength(str2))
	fmt.Println(stringLength(str3))

}
