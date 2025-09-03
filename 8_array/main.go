package main

import "fmt"

func main() {
	fmt.Println("Learning array")

	var name [5]string

	name[0] = "Ram"
	name[2] = "Amit"

	// Ram,_,Amit,_,_[5]
	//fmt.Println("Names of person is : ", name)
	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Number is : ", numbers)

	var price [10]float32
	price[2] = 90
	fmt.Printf("Price is : %.2f\n", price)

	fmt.Println("Length is : ", len(price))
	fmt.Println("at 2 index is : ", price[2])

}
