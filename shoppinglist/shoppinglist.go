package shoppinglist

import (
	"cartv2/cart/item"
	"cartv2/cart/reqResponse"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Shoppinglist struct {
	Name    string
	Items   []item.Item
	Created time.Time
}

func DecideShoppingHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getAllHandler(w, r)
	case "POST":
		createNewHandler(w, r)
	}
}

func getAllHandler(w http.ResponseWriter, r *http.Request) {
}

func createNewHandler(w http.ResponseWriter, r *http.Request) {
	payload, err := io.ReadAll(r.Body)
	if err != nil {
		reqResponse.WriteErr(w, 400, []byte("Error reading Body of Request"))
		return
	}

	if len(payload) == 0 {
		reqResponse.WriteErr(w, 400, []byte("No Body provided"))
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
}
