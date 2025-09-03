package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	orignal := make([]int, n)
	converted := make([]int, n)
	sum := 0

	for i := 0; i < n; i++ {
		fmt.Scan(&orignal[i])

		converted[i] = -orignal[i]
		sum += orignal[i] + converted[i]
	}
	fmt.Println(sum)
}
