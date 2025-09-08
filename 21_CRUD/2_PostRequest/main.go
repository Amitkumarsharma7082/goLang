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

func performPostRequest() {
	todo := Todo{
		UserId:    23,
		Title:     "Amit Kumar Sharma",
		Completed: true,
	}
	// Convert the Todo struct in JSON
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error occurred marshalling", err)
		return
	}

	// Convert the data Json to String
	jsonString := string(jsonData)

	// Convert String into Reader
	jsonReader := strings.NewReader(jsonString)

	// url
	myURL := "https://jsonplaceholder.typicode.com/todos"

	// Post the data
	res, err := http.Post(myURL, "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error occurred when post data", err)
		return
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body) // []byte, err
	fmt.Println("Response :", string(data))
	fmt.Println("Status code :", res.Status)
}

func main() {
	fmt.Println("Post request......")
	// Post Request
	performPostRequest()

}
