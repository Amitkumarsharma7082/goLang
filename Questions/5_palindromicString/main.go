package main

import "fmt"

func reverse(s string) string {
	var emptyString string
	var lengthString = len(s) - 1

	for i := lengthString; i >= 0; i-- {
		emptyString = emptyString + string(s[i])
	}
	return emptyString
}

func main() {
	var s string
	fmt.Scan(&s)

	a := reverse(s)

	if a == s {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
