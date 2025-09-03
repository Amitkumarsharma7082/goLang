package main

import "fmt"

func main() {
	// name <-> grade || key <-> value
	studentGrade := make(map[string]int)
	studentGrade["Alice"] = 90
	studentGrade["Jane"] = 78
	studentGrade["Era"] = 80
	studentGrade["Bob"] = 67

	fmt.Println("Bob grade :", studentGrade["Bob"])
	delete(studentGrade, "Bob")

	// Print all
	for name, grade := range studentGrade {
		fmt.Println(name, "<->", grade)
	}

	// check if exists
	grades, exists := studentGrade["David"]
	fmt.Println("Grades david :", grades)
	fmt.Println("Exists or not :", exists)

	Grade, Exists := studentGrade["Jane"]
	fmt.Println("Grades Jane :", Grade)
	fmt.Println("Exists or not :", Exists)

	person := map[string]string{
		"Sita":   "O",
		"Ram":    "A+",
		"Laxman": "A",
	}

	for key, value := range person {
		fmt.Printf("__key is %s and value is %s\n", key, value)
	}

}
