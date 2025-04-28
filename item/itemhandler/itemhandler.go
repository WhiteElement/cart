package itemhandler

import (
	"cartv2/cart/db"
	"cartv2/cart/item/item"
	"cartv2/cart/reqResponse"
	"encoding/json"
	"fmt"
	"net/http"
)

type Itemhandler struct {
	Conn db.DB
}

func (i Itemhandler) Choose(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		i.getAllHandler(w, r)
	case "POST":
		i.newHandler(w, r)
	case "PATCH":
		i.updateHandler(w, r)
	}
}

func (i Itemhandler) getAllHandler(w http.ResponseWriter, r *http.Request) {
	//TODO: get All from database
}

func (i Itemhandler) newHandler(w http.ResponseWriter, r *http.Request) {
	item, err := itemFromBody(w, r)
	if err != nil {
		reqResponse.WriteErr(w, 400, err.Error())
		return
	}

	if item.Id == 0 {
		reqResponse.WriteErr(w, 400, fmt.Sprintf("No Id for Item provided"))
		return
	}

	if len(item.Name) == 0 {
		reqResponse.WriteErr(w, 400, fmt.Sprintf("No Name for Shoppinglist provided"))
		return
	}

	// TODO: Put into database

}

func (i Itemhandler) updateHandler(w http.ResponseWriter, r *http.Request) {
	item, err := itemFromBody(w, r)
	if err != nil {
		reqResponse.WriteErr(w, 400, err.Error())
		return
	}

	if item.Id == 0 {
		reqResponse.WriteErr(w, 400, fmt.Sprintf("No Id for Item provided"))
		return
	}

	// TODO:
	// Get from database, update, save again
}

func itemFromBody(w http.ResponseWriter, r *http.Request) (item.Item, error) {
	payload, err := reqResponse.VerifyBody(w, r)
	if err != nil {
		return item.Item{}, err
	}

	var i item.Item
	err = json.Unmarshal(payload, &i)
	if err != nil {
		return item.Item{}, err
	}

	return i, nil
}
