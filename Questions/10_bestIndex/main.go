package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int64, n)
	prefix := make([]int64, n+1)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		prefix[i+1] = prefix[i] + a[i]
	}

	best := int64(math.MinInt64)

	for i := 0; i < n; i++ {
		var sum int64 = 0
		length := 1
		start := i
		for start+length <= n {
			sum += prefix[start+length] - prefix[start]
			start += length
			length++
		}
		if sum > best {
			best = sum
		}
	}

	fmt.Println(best)
}
