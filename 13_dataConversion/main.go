package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 42
	fmt.Println("Value of num :", num)
	fmt.Printf("Type of num : %T\n", num)

	var data float64 = float64(num)
	fmt.Print("data is :", data)
	fmt.Printf(" and type is : %T\n", data)

	fmt.Println("________________________________")

	// data convert int to string
	number := 18
	str := strconv.Itoa(number)
	fmt.Println("str is :", str)
	fmt.Printf("str type is : %T\n", str)

	// data convert string to int
	number_string := "5"
	number_string1 := "5"
	addTwoString := number_string + number_string1
	number_int, _ := strconv.Atoi(addTwoString) // Atoi is return two values
	number_int = number_int + 1
	fmt.Println("number_int is :", number_int)
	fmt.Printf("number_int type is : %T\n", number_int)

	// data convert string to float
	number_float := "18.0"
	str1, _ := strconv.ParseFloat(number_float, 64) // use ParseFloat and pass bit(64)
	fmt.Printf("float is : %.1f\n", str1)
	fmt.Printf("float type is : %T\n", str1)

}
