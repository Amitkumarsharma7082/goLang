package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	const mod = 1000000007 // 10^9 + 7

	ans := 1
	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)
		ans = (ans * x) % mod
	}

	fmt.Println(ans)
}
