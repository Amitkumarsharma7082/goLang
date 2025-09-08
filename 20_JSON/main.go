package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_Adult"`
}

func main() {
	fmt.Println("JSON")
	/*
		package : encoded/json
		Marshalling(encoded) : conver Go struct to Json.encoded : json.Marshal
		Unmarshalling(decoded) : convert JSON-encoded to Go struct : json.Unmarshal
	*/

	person := Person{
		Name:    "Jhon",
		Age:     34,
		IsAdult: true,
	}
	// fmt.Println("Person Details :", person) // Person Details : {Jhon 34 true}

	//! convert person details : JSON encoding(Marshalling)
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error occurred while marshalling", err)
		return
	}
	fmt.Println("Person Details in JSON Format :", string(jsonData))
	//Output : Person Details in JSON Format : {"name":"Jhon","age":34,"is_Adult":true}

	//! Convert JSON encoding(Marshalling) : Json decoding(Unmarshalling)
	var personData Person
	err = json.Unmarshal(jsonData, &personData)
	if err != nil {
		fmt.Println("Error occurred unmarshalling", err)
		return
	}
	fmt.Println("data :", personData)

}
