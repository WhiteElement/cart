package itemhandler

import (
	"cartv2/cart/internal/db"
	"cartv2/cart/internal/item/item"
	"cartv2/cart/internal/reqResponse"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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

func (i Itemhandler) ChooseSingle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "DELETE":
		i.deleteHandler(w, r)
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

	item.Updated = time.Now()
	err = i.updateItem(item)
	if err != nil {
		reqResponse.WriteErr(w, 400, err.Error())
		return
	}

	reqResponse.Write(w, 200, []byte("Updated"))

}

//
// AUX
//

func (i Itemhandler) deleteHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	convId, err := strconv.Atoi(id)
	if err != nil {
		reqResponse.WriteErr(w, 500, err.Error())
		return
	}

	// TODO: Update Updated bei der Liste
	_, err = i.Conn.Conn.Exec("DELETE FROM public.\"Items\" WHERE \"Id\" = $1", convId)
	if err != nil {
		reqResponse.WriteErr(w, 500, err.Error())
		return
	}

	reqResponse.Write(w, 200, []byte("Deleted"))
}

//
// AUX
//

// TODO: Update Updated bei der Liste
func (i Itemhandler) updateItem(item item.Item) error {
	_, err := i.Conn.Conn.Exec("UPDATE public.\"Items\" SET (\"Name\", \"Checked\", \"Updated\") = ($1, $2, $3) WHERE \"Id\" = $4", item.Name, item.Checked, item.Updated, item.Id)

	if err != nil {
		return err
	}

	return nil

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
