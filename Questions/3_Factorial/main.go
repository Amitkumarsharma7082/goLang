package main

import "fmt"

func fact(n int) uint64 {
	if n < 0 {
		return 0
	}
	var fac uint64 = 1
	for i := 1; i <= n; i++ {
		fac *= uint64(i)
	}
	return fac
}

func main() {
	var n int
	fmt.Print("Enter a non-negative integer: ")
	fmt.Scan(&n)

	if n < 0 {
		fmt.Println("Negative numbers are not allowed.")
		return
	}

	ans := fact(n)
	fmt.Println("Factorial of", n, "is", ans)
}
