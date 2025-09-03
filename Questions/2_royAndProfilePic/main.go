package main

import "fmt"

func main() {
	var l int // length
	var n int // no. of pic
	var w int
	var h int

	fmt.Scanf("%d\n %d\n", &l, &n)
	// fmt.Printf("%d\n%d\n", l, n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d\n", &w, &h)
		if w < l || h < l {
			fmt.Println("UPLOAD ANOTHER")
		} else {
			if w == h {
				fmt.Println("ACCEPTED")
			} else {
				fmt.Println("CROP IT")
			}
		}
		fmt.Printf("\n")
	}
}
