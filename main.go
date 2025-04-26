package main

import (
	"cartv2/cart/item"
	"cartv2/cart/shoppinglist"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 420

	http.HandleFunc("/shoppinglist", shoppinglist.ChooseHandler)
	http.HandleFunc("/item", item.ChooseHandler)

	fmt.Printf("Listening on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
