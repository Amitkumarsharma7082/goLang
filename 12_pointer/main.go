package main

import "fmt"

func modifyValueByReference(num *int) {
	*num *= 20
}

func main() {
	// House_Name := 11
	// House_Ptr := &House_Name
	// fmt.Println("House Ptr :", House_Ptr)

	var pointer *int // by default : nil
	if pointer == nil {
		fmt.Println("Pointer is not assigned")
	}

	value := 10
	modifyValueByReference(&value)
	fmt.Println("Modify value :", value)
}
