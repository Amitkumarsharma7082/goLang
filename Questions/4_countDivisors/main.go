package main

import "fmt"

func main() {
	var l, r, k, i int

	fmt.Scanf("%d %d %d\n", &l, &r, &k)
	var count int
	for i = l; i <= r; i++ {
		if i%k == 0 {
			count = count + 1
		}
	}
	fmt.Println(count)
}
