package synchandler

import (
	"cartv2/cart/internal/db"
	"cartv2/cart/internal/item/item"
	"cartv2/cart/internal/reqResponse"
	"cartv2/cart/internal/shoppinglist/shoppinglist"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type SyncHandler struct {
	Conn db.DB
}

type SyncRequest struct {
	LastWrite time.Time
}

type SyncResponse struct {
	SyncNeeded bool
	SyncType   string
	DataLists  []byte
	DataItems  []byte
}

func (h SyncHandler) Choose(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.GetSyncState(w, r)
	case "POST":
		h.ReceiveClientSync(w, r)
	}
}

func (h SyncHandler) GetSyncState(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	param := queryParams.Get("ts")
	ts, err := time.Parse(time.RFC3339, param)

	if err != nil {
		reqResponse.WriteErr(w, 400, err.Error())
		return
	}

	req := SyncRequest{LastWrite: ts}

	serverLastWrite := h.Conn.QueryLastWrite()

	var res SyncResponse
	if serverLastWrite == req.LastWrite {
		res.SyncNeeded = false
	}

	if serverLastWrite.After(req.LastWrite) {
		res.SyncNeeded = true
		res.SyncType = "ServerToClient"

		lChan := make(chan []shoppinglist.List)
		iChan := make(chan []item.Item)
		go func() {
			lists, err := h.Conn.QueryAllLists()
			if err != nil {
				log.Printf("Error querying Lists: %s\n", err.Error())
			}
			lChan <- lists
		}()

		go func() {
			items := h.Conn.QueryAllItems()
			iChan <- items
		}()

		lists := <-lChan
		items := <-iChan

		listByte, err := json.Marshal(lists)
		if err != nil {
			reqResponse.WriteErr(w, 400, "Cannot Marhsal listByte")
			return
		}
		itemByte, err := json.Marshal(items)
		if err != nil {
			reqResponse.WriteErr(w, 400, "Cannot Marhsal itemByte")
			return
		}

		res.DataLists = listByte
		res.DataItems = itemByte
	}

	if req.LastWrite.After(serverLastWrite) {
		res.SyncNeeded = true
		res.SyncType = "ClientToServer"
	}

	json, err := json.Marshal(res)
	if err != nil {
		reqResponse.WriteErr(w, 400, "Could not marshal SyncResponse")
		return
	}

	reqResponse.Write(w, 200, []byte(json))
}

func (h SyncHandler) ReceiveClientSync(w http.ResponseWriter, r *http.Request) {
}
