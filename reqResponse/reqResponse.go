package reqResponse

import (
	"log"
	"net/http"
)

func Write(w http.ResponseWriter, statusCode int, body []byte) {
	w.WriteHeader(statusCode)
	w.Write(body)
}

func WriteErr(w http.ResponseWriter, statusCode int, body []byte) {
	log.Println(body)
	w.WriteHeader(statusCode)
	w.Write(body)
}
