package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func performDeleteRequest() {
	// myURL - Delete a specific todo (let's delete todo with ID 1)
	myURL := "https://jsonplaceholder.typicode.com/todos/1"

	// Create DELETE Request
	// For DELETE requests, we typically don't send a body, so use nil
	req, err := http.NewRequest(http.MethodDelete, myURL, nil)
	if err != nil {
		fmt.Println("Error occurred in Delete method", err)
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

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println("Delete Response :", string(data))
	fmt.Println("Status code :", res.Status)
}

func main() {
	fmt.Println("Delete Request.....")
	performDeleteRequest()

}
