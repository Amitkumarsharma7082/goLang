//! Unmarshalling data from URL

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
URL : https://fakestoreapi.com/products
response : [
{
"id": 1,
"title": "Fjallraven - Foldsack No. 1 Backpack, Fits 15 Laptops",
"price": 109.95,
"description": "Your perfect pack for everyday use and walks in the forest. Stash your laptop (up to 15 inches) in the padded sleeve, your everyday",
"category": "men's clothing",
"image": "https://fakestoreapi.com/img/81fPKd-2AYL._AC_SL1500_t.png",
"rating": {
"rate": 3.9,
"count": 120
}
},
]
*/

type Product struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Image       string  `json:"image"`
	Rating      Rating  `json:"rating"`
}
type Rating struct {
	Rate  float64 `json:"rate"`
	Count int     `json:"count"`
}

func products() {
	fmt.Println("json products")
	// Make http Get Request
	Url := "https://fakestoreapi.com/products"
	res, err := http.Get(Url)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer res.Body.Close()

	// Check the status code
	if res.StatusCode != http.StatusOK {
		fmt.Println("HTTP error:", res.Status)
		return
	}

	// Read the data
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// json unmarshalling
	var products []Product
	err = json.Unmarshal(data, &products)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	// Print the results
	fmt.Printf("Fetched %d products:\n\n", len(products))

	for i, product := range products {
		fmt.Printf("Product %d:\n", i+1)
		fmt.Printf("  Id :  %d\n", product.Id)
		fmt.Printf("  Title: %s\n", product.Title)
		fmt.Printf("  Price: $%.2f\n", product.Price)
		fmt.Printf("  Category: %s\n", product.Category)
		fmt.Printf("  Rating: %.1f/5 (%d reviews)\n", product.Rating.Rate, product.Rating.Count)
		fmt.Println("----------------------------------------")
	}
}
