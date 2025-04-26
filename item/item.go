package item

import (
	"cartv2/cart/reqResponse"
	"encoding/json"
	"fmt"
	"net/http"
)

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

func getAllHandler(w http.ResponseWriter, r *http.Request) {
	//TODO: get All from database
}

func newHandler(w http.ResponseWriter, r *http.Request) {
	item, err := itemFromBody(w, r)
	if err != nil {
		reqResponse.WriteErr(w, 400, []byte(err.Error()))
		return
	}

	if item.Id == 0 {
		reqResponse.WriteErr(w, 400, []byte(fmt.Sprintf("\nNo Id for Item provided")))
		return
	}

	if len(item.Name) == 0 {
		reqResponse.WriteErr(w, 400, []byte(fmt.Sprintf("\nNo Name for Shoppinglist provided")))
		return
	}

	// TODO: Put into database

}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	item, err := itemFromBody(w, r)
	if err != nil {
		reqResponse.WriteErr(w, 400, []byte(err.Error()))
		return
	}

	if item.Id == 0 {
		reqResponse.WriteErr(w, 400, []byte(fmt.Sprintf("\nNo Id for Item provided")))
		return
	}

	// TODO:
	// Get from database, update, save again
}

func itemFromBody(w http.ResponseWriter, r *http.Request) (Item, error) {
	payload, err := reqResponse.VerifyBody(w, r)
	if err != nil {
		return Item{}, err
	}

	var item Item
	err = json.Unmarshal(payload, &item)
	if err != nil {
		return Item{}, err
	}

	return item, nil
}
