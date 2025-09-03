package main

import "fmt"

func main() {
	// fmt.Println("About slice")
	// var numbers = []int{1, 2, 3, 4, 5}
	// numbers = append(numbers, 3, 10, 11, 12, 13, 14, 15)
	// fmt.Println("Number is : ", numbers)

	// names := []string{}
	// fmt.Println("names is :", names)

	num := make([]int, 3, 5)
	num = append(num, 4, 5, 6)
	fmt.Println("Slice is : ", num)
	fmt.Println("Length is : ", len(num))
	fmt.Println("Capacity is : ", cap(num)) // capacity double

}
