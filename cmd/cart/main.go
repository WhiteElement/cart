package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strconv"

	"cartv2/cart/internal/db"
	"cartv2/cart/internal/item/itemhandler"
	"cartv2/cart/internal/shoppinglist/listhandler"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

//go:embed web
var webFolder embed.FS

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

	staticFS, _ := fs.Sub(webFolder, "web")

	http.Handle("/", http.FileServer(http.FS(staticFS)))
	http.HandleFunc("/api/shoppinglist", listhandler.Choose)
	http.HandleFunc("/api/shoppinglist/{id}", listhandler.GetOneList)
	http.HandleFunc("/api/shoppingitem", itemhandler.Choose)
	http.HandleFunc("/api/shoppingitem/{id}", itemhandler.ChooseSingle)

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Println("No PORT Env variable found, setting 420 as a default value")
		port = 420
	}

	log.Printf("Listening on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
