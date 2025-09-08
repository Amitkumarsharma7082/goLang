package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("CRUD......")
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
	fmt.Println("Data :", string(data))

}
