package itemhandler

import (
	"cartv2/cart/internal/db"
	"cartv2/cart/internal/item/item"
	"cartv2/cart/internal/reqResponse"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
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

//
// GET
//

func (i Itemhandler) getAllHandler(w http.ResponseWriter, r *http.Request) {
	its := i.Conn.QueryAllItems()

	content, err := json.Marshal(its)
	if err != nil {
		reqResponse.WriteErr(w, 400, err.Error())
	}

	reqResponse.Write(w, 200, content)
}

//
// POST
//

func (i Itemhandler) newHandler(w http.ResponseWriter, r *http.Request) {
	it, err := itemFromBody(w, r)
	if err != nil {
		reqResponse.WriteErr(w, 400, err.Error())
		return
	}

	if it.ListId == 0 {
		reqResponse.WriteErr(w, 400, fmt.Sprintf("No ListId for corresponding Shoppinglist provided"))
		return
	}

	if len(it.Name) == 0 {
		reqResponse.WriteErr(w, 400, fmt.Sprintf("No Name for Item provided"))
		return
	}

	_, err = i.Conn.Conn.Exec("INSERT INTO public.\"Items\" (\"Name\", \"Checked\", \"ListId\", \"Updated\") VALUES ($1, $2, $3, $4)", it.Name, it.Checked, it.ListId, time.Now())

	if err != nil {
		reqResponse.WriteErr(w, 500, err.Error())
		return
	}

	reqResponse.Write(w, 201, []byte("Created"))
}

//
// PATCH
//

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

//
// AUX
//

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
