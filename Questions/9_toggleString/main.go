package main

import (
	"fmt"
)

func toggle(s string) string {
	bytes := []byte(s)
	for i, ch := range bytes {
		// Upper Case
		if ch >= 'A' && ch <= 'Z' {
			bytes[i] = ch + 32
		} else if ch >= 'a' && ch <= 'z' {
			bytes[i] = ch - 32
		}
	}
	return string(bytes)
}

func main() {
	var s string
	fmt.Scan(&s)

	result := toggle(s)
	fmt.Println(result)
}
