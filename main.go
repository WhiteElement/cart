package main

import (
	"cartv2/cart/item"
	"cartv2/cart/shoppinglist"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
)

//go:embed static/*
var staticFolder embed.FS

func main() {
	// TODO: .env file(Port)

	staticFS, _ := fs.Sub(staticFolder, "static")
	port := 420

	http.Handle("/", http.FileServer(http.FS(staticFS)))
	http.HandleFunc("/shoppinglist", shoppinglist.ChooseHandler)
	http.HandleFunc("/item", item.ChooseHandler)

	fmt.Printf("Listening on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}
