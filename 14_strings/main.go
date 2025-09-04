package main

import (
	"fmt"
	"strings"
)

func main() {
	// Split(,)
	fruits := "apple,orange,mango"
	result := strings.Split(fruits, ",")
	fmt.Println("Fruits :", result)

	// Count
	counting := "one two three four two five two six"
	countTwo := strings.Count(counting, "two")
	fmt.Println("How may two's :", countTwo)

	// TrimSpace
	str := "      Hello Go!     "
	fmt.Println("Print string with space :", str)
	printStringWithoutSpace := strings.TrimSpace(str)
	fmt.Println("Print string without space :", printStringWithoutSpace)

	// Join ([]string{str1, str2}, " ")
	firstName := "Amit"
	lastName := "Sharma"
	fullName := strings.Join([]string{firstName, "Kumar", lastName}, " ")
	fmt.Println("Full name :", fullName)

}
