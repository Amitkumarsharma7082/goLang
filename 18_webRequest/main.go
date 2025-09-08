package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// web request package : net/http import
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/2")
	if err != nil {
		fmt.Println("Error occurred", err)
		return
	}
	// res type : *http.Response
	// data : data response Body
	defer res.Body.Close()

	//Read data with help of ioutil.Read
	data, err := ioutil.ReadAll(res.Body) // data : byte return : string
	if err != nil {
		fmt.Println("Error occurred when read the data", err)
		return
	}
	fmt.Println("Response :", string(data))

}
