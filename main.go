package main

import (
	"cartv2/cart/shoppinglist"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 420

	http.HandleFunc("/shoppinglist", shoppinglist.DecideShoppingHandler)

	fmt.Printf("Listening on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
