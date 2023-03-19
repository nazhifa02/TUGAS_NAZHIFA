package main

import (
	"fmt"
	"encoding/json"
	"net/http"
)

type Product struct {
	title string 'json:"title"'
	Category string 'json:"category"'
	Price float64 'json:"price"'
	Rating struct { Rate float64 'json:"rate"'} 'json:"rating"'
}

func getProducts(url string, ch chan<- []Products) {
	res, err := http.Get(url)
	if err =/= nil {
		fmt.Println("error get products :", err)
		return
	}
	if len(products) > 5 {
		products = products[:5]
	}
	ch <- products
}

func main()  {
	url := "https://fakestoreapi.com/products"
	ch := make(chan []Product)

	go getProducts(url, ch)

	products := <-ch
	fmt.Println("Data Produk")
	for _, p := range products {
		fmt.Println("===")
		fmt.Println("Title :", p.Title)
		fmt.Println("Price : $", p.Price)
		fmt.Println("Category :", p.Category)
		fmt.Println("Rating : *", p.Rating)
		fmt.Println("===")
	}
}