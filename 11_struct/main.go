package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	Age       int
}

type Contact struct {
	phone_number int
	email        string
}

type Address struct {
	house_number int
	area         string
	state        string
}

type Employee struct {
	Person_details Person
	Person_contact Contact
	Person_address Address
}

func main() {

	Employee1 := Employee{
		Person{
			firstName: "Amit",
			lastName:  "Sharma",
			Age:       26,
		},
		Contact{
			phone_number: 98765432,
			email:        "ak@gmail.com",
		},
		Address{
			house_number: 11,
			area:         "Narnaul",
			state:        "Haryana",
		},
	}
	fmt.Println("Person age :", Employee1.Person_details.Age)
	fmt.Println("Employee1 all details :", Employee1)

	// Ram := Person{
	// 	firstName: "Ram",
	// 	lastName:  "Ji",
	// 	Age:       16,
	// }

	// fmt.Println("Ram Person :", Ram)

	// // 1st method
	// person1 := Person{
	// 	firstName: "Aakash",
	// 	lastName:  "Gupta",
	// 	Age:       26,
	// }
	// fmt.Println("Person1 :", person1)

	// // 2nd method
	// var John Person
	// John.firstName = "John"
	// John.lastName = "Doe"
	// John.Age = 25

	// fmt.Println("John Person :", John)

	// // 3rd method
	// person2 := new(Person) // new word point locate memory
	// person2.firstName = "Laxman"
	// person2.lastName = "Ji"
	// person2.Age = 15

	// fmt.Println("Person2 :", person2)

}
