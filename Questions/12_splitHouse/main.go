package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	var s string
	fmt.Fscan(in, &s)

	for i := 0; i < n-1; i++ {
		if s[i] == 'H' && s[i+1] == 'H' {
			fmt.Println("NO")
			return
		}
	}

	result := []rune(s)
	for i := 0; i < n; i++ {
		if result[i] == '.' {
			result[i] = 'B'
		}
	}

	fmt.Println("YES")
	fmt.Println(string(result))
}
