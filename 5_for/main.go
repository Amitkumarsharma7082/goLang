package main

import "fmt"

// ! Question : Life, the Universe, and Everything
func main() {
	for i := range 100 {
		if i == 42 {
			break
		}
		fmt.Println(i)
	}
}

// func main() {
// 	i := 1
// 	for i < 5 {
// 		fmt.Println(i)
// 		i = i + 1
// 	}

// 	for j := 10; j <= 15; j++ {
// 		fmt.Println(j)
// 	}

// 	for i := range 7 {
// 		fmt.Println(i)
// 	}

// 	for {
// 		fmt.Println("loop")
// 		break
// 	}

// 	for n := range 6 {
// 		if n%2 == 0 {
// 			continue
// 		}
// 		fmt.Println(n)
// 	}
// }
