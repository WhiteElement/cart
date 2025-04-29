package main

import (
	"cartv2/cart/db"
	"cartv2/cart/item/itemhandler"
	"cartv2/cart/shoppinglist/listhandler"
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

	conn := db.NewConn()
	itemhandler := itemhandler.Itemhandler{Conn: conn}
	listhandler := listhandler.Listhandler{Conn: conn}

	staticFS, _ := fs.Sub(staticFolder, "static")

	http.Handle("/", http.FileServer(http.FS(staticFS)))
	http.HandleFunc("/shoppinglist", listhandler.Choose)
	http.HandleFunc("/shoppinglist/{id}", listhandler.GetOneList)
	http.HandleFunc("/item", itemhandler.Choose)

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Println("No PORT Env variable found, setting 420 as a default value")
		port = 420
	}

	// TODO: postgres nur für bestimmte IPs freigeben
	log.Printf("Listening on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
