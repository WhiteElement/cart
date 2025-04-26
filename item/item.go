package item

import "net/http"

type Item struct {
	Id   int
	Name string
}

func ChooseHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getAllHandler(w, r)
	case "POST":
		newHandler(w, r)
	case "PATCH":
		updateHandler(w, r)
	}
}

func getAllHandler(w http.ResponseWriter, r *http.Request) {}

func newHandler(w http.ResponseWriter, r *http.Request) {}

func updateHandler(w http.ResponseWriter, r *http.Request) {}
