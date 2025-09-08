package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("")
	myURL := "https://fakestoreapi.com/products"
	fmt.Printf("Type : %T\n", myURL) // String

	paresed, err := url.Parse(myURL)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Printf("Type parsedURL : %T\n", paresed) // *url.URL
	fmt.Println("Scheme :", paresed.Scheme)
	fmt.Println("HOST :", paresed.Host)
	fmt.Println("PATH :", paresed.Path)
	fmt.Println("RAWQUERY :", paresed.RawQuery)
	/*
			Scheme : https
		HOST : fakestoreapi.com
		PATH : /products
		RAWQUERY :
	*/
}
