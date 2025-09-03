// ! Assignment 9 (Functions-2)
package main

import "fmt"

func main() {
	firstName, lastName := getName()
	fmt.Println("Welcome", firstName, lastName+"!")
}
func getName() (string, string) {
	return "Joe", "Dane"
}

// ! Assignment 8 (Functions)
// package main

// import "fmt"

// func main() {
// 	x := 5
// 	increment(x)
// 	fmt.Println(x)
// }

// func increment(x int) {
// 	x++
// }

// package main

// import "fmt"

// func sub(x int, y int) int {
// 	return x - y
// }

// func main() {
// 	fmt.Println("Subract result : ", sub(4, 3))
// }

// func concat(s1 string, s2 string) string {
// 	return s1 + s2
// }

// func main() {
// 	test("hello", "world!")
// }

// func test(s1 string, s2 string) {
// 	fmt.Println(concat(s1, s2))
// }

// ! Assignment 7 (Switch statement)
// package main

// import "fmt"

// func billing(plan string) float64 {
// 	switch plan {
// 	case "basic":
// 		return 10.0
// 	case "pro":
// 		return 20.0
// 	case "enterprise":
// 		return 50.0
// 	default:
// 		return 0.00
// 	}
// }

// func main() {
// 	plan := "basic"
// 	fmt.Printf("the plan is %s plan cost is $%.2f\n", plan, billing(plan))
// 	plan = "pro"
// 	fmt.Printf("the plan is %s plan cost is $%.2f\n", plan, billing(plan))
// 	plan = "free"
// 	fmt.Printf("the plan is %s plan cost is $%.2f\n", plan, billing(plan))
// 	plan = "enterprise"
// 	fmt.Printf("the plan is %s plan cost is $%.2f\n", plan, billing(plan))

// }

//! Assignment 6 (if statement)
// package main

// import "fmt"

// func main() {
// 	messageMax := 20
// 	message := 10

// 	if messageMax > message {
// 		fmt.Println("Sent")
// 	} else {
// 		fmt.Println("Not sent")
// 	}
// }

// ! Assignment 5
// package main

// import "fmt"

// func main() {
// 	averageOpenRate, displayMessage := .23, "hello gooo"
// 	fmt.Println(averageOpenRate, displayMessage)
// }

// ! Assignment 4
// package main

// import "fmt"

// func main() {
// 	var userName string = "amitSharma"
// 	var password string = "12345"

// 	fmt.Println("account login details : ", userName+":"+password)
// }

// ! Assignment 3
// package main

// import "fmt"

// func main() {
// 	accountAgeFloat := 2.6
// 	accountAgeInt := int64(accountAgeFloat)
// 	fmt.Println("this is float", accountAgeFloat, "and this is int", accountAgeInt)
// }

// ! Assignment 2
// package main

// import "fmt"

// func main() {
// 	var messageStart string
// 	messageStart = "Happy birthday! you are now"

// 	var age int
// 	age = 21

// 	var messageEnd string
// 	messageEnd = "years old!"

// 	fmt.Println(messageStart, age, messageEnd)
// }

//! Assignment 1
// package main

// import "fmt"

// func main() {
// 	fmt.Println("hello there!")
// }
