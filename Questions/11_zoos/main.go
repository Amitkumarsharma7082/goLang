package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var s string
	fmt.Fscan(reader, &s)

	countZ, countO := 0, 0
	for _, ch := range s {
		if ch == 'z' {
			countZ++
		} else if ch == 'o' {
			countO++
		}
	}

	if countO == 2*countZ {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
