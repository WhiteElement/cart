package shoppinglist

import (
	"cartv2/cart/item/item"
	"cartv2/cart/reqResponse"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Shoppinglist struct {
	Id      int
	Name    string
	Items   []item.Item
	Created time.Time
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

func updateHandler(w http.ResponseWriter, r *http.Request) {
	shoppinglist, err := listFromBody(w, r)
	if err != nil {
		reqResponse.WriteErr(w, 400, fmt.Sprintf(err.Error()))
		return
	}

	if shoppinglist.Id == 0 {
		reqResponse.WriteErr(w, 400, fmt.Sprintf("No Id for Shoppinglist provided"))
	}

	if len(shoppinglist.Name) == 0 {
		reqResponse.WriteErr(w, 400, fmt.Sprintf("No Name for Shoppinglist provided"))
	}

	//TODO:
	// 1. Get from Database
	// 2. Update
	// 3. Write to Database

}

func getAllHandler(w http.ResponseWriter, r *http.Request) {
	//TODO: get all from database
}

func newHandler(w http.ResponseWriter, r *http.Request) {
	shoppinglist, err := listFromBody(w, r)
	if err != nil {
		reqResponse.WriteErr(w, 400, fmt.Sprintf(err.Error()))
		return
	}

	if len(shoppinglist.Name) == 0 {
		reqResponse.WriteErr(w, 400, fmt.Sprintf("No Name for Shoppinglist provided"))
	}

	//TODO: save to database
}

func listFromBody(w http.ResponseWriter, r *http.Request) (Shoppinglist, error) {
	payload, err := reqResponse.VerifyBody(w, r)
	if err != nil {
		return Shoppinglist{}, err
	}

	var list Shoppinglist
	err = json.Unmarshal(payload, &list)
	if err != nil {
		return Shoppinglist{}, err
	}

	return list, nil
}
