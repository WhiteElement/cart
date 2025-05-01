package listhandler

import (
	"cartv2/cart/internal/db"
	"cartv2/cart/internal/item/item"
	"cartv2/cart/internal/reqResponse"
	"cartv2/cart/internal/shoppinglist/shoppinglist"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Listhandler struct {
	Conn db.DB
}

func (l Listhandler) Choose(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		l.getAllHandler(w, r)
	case "POST":
		l.newHandler(w, r)
	case "PATCH":
		l.patchHandler(w, r)
	}
}

func (l Listhandler) patchHandler(w http.ResponseWriter, r *http.Request) {
	shoppinglist, err := listFromBody(w, r)
	if err != nil {
		reqResponse.WriteErr(w, 400, fmt.Sprintf(err.Error()))
		return
	}

	if shoppinglist.Id == 0 {
		reqResponse.WriteErr(w, 400, fmt.Sprintf("No Id for Shoppinglist provided"))
	}

	args := make(map[string]interface{})

	if len(shoppinglist.Name) != 0 {
		fmt.Println("Name is nil")
		args[fmt.Sprintf("\"Name\" = $%d", len(args)+1)] = shoppinglist.Name
	}
	if shoppinglist.Archived {
		fmt.Println("Archived is true")
		args[fmt.Sprintf("\"Archived\" = $%d", len(args)+1)] = shoppinglist.Archived
	}

	shoppinglist.Updated = time.Now()
	args[fmt.Sprintf("\"Updated\" = $%d", len(args)+1)] = shoppinglist.Updated

	err = l.updateList(args, shoppinglist.Id)
	if err != nil {
		reqResponse.WriteErr(w, 400, err.Error())
		return
	}

	reqResponse.Write(w, 200, []byte("Updated"))
}

func (l Listhandler) updateList(args map[string]interface{}, listId int) error {
	var keys []string
	values := []interface{}{}

	for key, value := range args {
		keys = append(keys, key)
		values = append(values, value)
	}

	sql := fmt.Sprintf("UPDATE public.\"Lists\" SET %s WHERE \"Id\" = $%d\n", strings.Join(keys, ", "), len(values)+1)
	values = append(values, listId)
	_, err := l.Conn.Conn.Exec(sql, values...)

	if err != nil {
		return err
	}

	return nil
}

func (l Listhandler) GetOneList(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		reqResponse.WriteErr(w, 400, err.Error())
	}

	var list shoppinglist.List
	var items []item.Item
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		list = l.Conn.QueryList(id)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		items = l.Conn.QueryItemsFromList(id)
	}()

	wg.Wait()

	list.Items = items
	content, err := json.Marshal(list)
	if err != nil {
		reqResponse.WriteErr(w, 500, err.Error())
		return
	}

	reqResponse.Write(w, 200, content)
}

func (l Listhandler) getAllHandler(w http.ResponseWriter, r *http.Request) {
	lists, err := l.Conn.QueryAllLists()

	content, err := json.Marshal(lists)
	if err != nil {
		reqResponse.WriteErr(w, 500, err.Error())
		return
	}
	reqResponse.Write(w, 200, content)
}

func (l Listhandler) newHandler(w http.ResponseWriter, r *http.Request) {
	list, err := listFromBody(w, r)
	if err != nil {
		reqResponse.WriteErr(w, 400, fmt.Sprintf(err.Error()))
		return
	}

	if len(list.Name) == 0 {
		reqResponse.WriteErr(w, 400, fmt.Sprintf("No Name for Shoppinglist provided"))
		return
	}

	_, err = l.Conn.Conn.Exec("INSERT INTO public.\"Lists\" (\"Name\", \"Created\", \"Updated\") VALUES ($1, $2, $2)", list.Name, time.Now())

	if err != nil {
		reqResponse.WriteErr(w, 500, err.Error())
		return
	}

	reqResponse.Write(w, 201, []byte("Created"))
}

func listFromBody(w http.ResponseWriter, r *http.Request) (shoppinglist.List, error) {
	payload, err := reqResponse.VerifyBody(w, r)
	if err != nil {
		return shoppinglist.List{}, err
	}

	var list shoppinglist.List
	err = json.Unmarshal(payload, &list)
	if err != nil {
		return shoppinglist.List{}, err
	}

	return list, nil
}
