package main

import (
	"cartv2/cart/item"
	"cartv2/cart/shoppinglist"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

//go:embed static/*
var staticFolder embed.FS

func main() {
	args := os.Args

	if len(args) < 2 {
		log.Fatalln("No .env file as argument provided")
	}

	err := godotenv.Load(args[1])
	if err != nil {
		log.Fatal(err)
	}

	staticFS, _ := fs.Sub(staticFolder, "static")

	http.Handle("/", http.FileServer(http.FS(staticFS)))
	http.HandleFunc("/shoppinglist", shoppinglist.ChooseHandler)
	http.HandleFunc("/item", item.ChooseHandler)

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Println("No PORT Env variable found, setting 420 as a default value")
		port = 420
	}

	conn := NewConn()
	_, err = conn.Exec("INSERT INTO public.\"Items\" (\"Name\") VALUES ('Test')")
	if err != nil {
		log.Printf("Error inserting: %s\n", err.Error())
	}

	rows, err := conn.Query("select * from public.\"Items\"")
	if err != nil {
		log.Printf("Error querying: %s\n", err.Error())
	}

	for rows.Next() {
		var item item.Item
		rows.Scan(&item.Id, &item.Name)
		log.Printf("Item: +%v\n", item)
	}

	// TODO: postgres nur für bestimmte IPs freigeben
	log.Printf("Listening on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
