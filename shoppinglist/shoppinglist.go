package shoppinglist

import (
	"cartv2/cart/item"
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
	payload, err := reqResponse.VerifyBody(w, r)
	if err != nil {
		return
	}

	var shoppinglist Shoppinglist
	err = json.Unmarshal(payload, &shoppinglist)
	if err != nil {
		reqResponse.WriteErr(w, 400, []byte(fmt.Sprintf("\nCould not unmarshal payload: '%s'", payload)))
	}

	if shoppinglist.Id == 0 {
		reqResponse.WriteErr(w, 400, []byte(fmt.Sprintf("\nNo Name for Shoppinglist provided")))
	}

	if len(shoppinglist.Name) == 0 {
		reqResponse.WriteErr(w, 400, []byte(fmt.Sprintf("\nNo Name for Shoppinglist provided")))
	}

	//TODO:
	// 1. Get from Database
	// 2. Update
	// 3. Write to Database

}

func getAllHandler(w http.ResponseWriter, r *http.Request) {
}

func newHandler(w http.ResponseWriter, r *http.Request) {
	payload, err := reqResponse.VerifyBody(w, r)
	if err != nil {
		return
	}

	var shoppinglist Shoppinglist
	err = json.Unmarshal(payload, &shoppinglist)
	if err != nil {
		reqResponse.WriteErr(w, 400, []byte(fmt.Sprintf("\nCould not unmarshal payload: '%s'", payload)))
	}

	if len(shoppinglist.Name) == 0 {
		reqResponse.WriteErr(w, 400, []byte(fmt.Sprintf("\nNo Name for Shoppinglist provided")))
	}

	//TODO: save to database
}
