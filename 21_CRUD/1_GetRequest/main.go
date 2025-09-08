package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func performGetRequest() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error occurred while getting response", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error occuured status code :", res.Status)
		return
	}
	// Read Data
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error occurred while reading", err)
		return
	}
	// fmt.Println("Data :", string(data))

	var todo Todo
	err = json.Unmarshal(data, &todo)
	if err != nil {
		fmt.Println("Error unmarshalling", err)
		return
	}
	fmt.Println("data :", todo)
}
func main() {
	fmt.Println("Get Request......")
	// Call Get Request
	performGetRequest()

}
