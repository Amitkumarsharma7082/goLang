package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func performUpdateRequest() {

	todo := Todo{
		UserId:    23,
		Title:     "Amit Sharma Golang",
		Completed: true,
	}

	// Convert the data struct to JSON
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error occurred in JSON format", err)
		return
	}

	// Json into string
	jsonString := string(jsonData)

	// Convert string into Reader
	jsonReader := strings.NewReader(jsonString)

	// myURL
	myURL := "https://jsonplaceholder.typicode.com/todos/1"

	// Create PUT Request
	req, err := http.NewRequest(http.MethodPut, myURL, jsonReader)
	if err != nil {
		fmt.Println("Error occurred in Put methods", err)
		return
	}
	req.Header.Set("Content-type", "application/json")

	// Send Request
	client := http.Client{}

	// res
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error occurred when sending request", err)
		return
	}
	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Response :", string(data))

	fmt.Println("Status code :", res.Status)

}

func main() {
	fmt.Println("Update Request.....")
	performUpdateRequest()
}
