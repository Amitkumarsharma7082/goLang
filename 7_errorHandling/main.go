package main

import "fmt"

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("denominator never zero")
	}
	return a / b, nil

}

func main() {
	// ans, _ := divide(10,0)
	ans, err := divide(10, 8)
	if err != nil {
		fmt.Println("Error handling")
	}
	fmt.Printf("Dividing two number ans : %.1f\n", ans)
}
