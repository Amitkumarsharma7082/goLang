package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	lastDigit := arr[n-1] % 10
	fmt.Println("lastDigit : ", lastDigit)
	if lastDigit == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
