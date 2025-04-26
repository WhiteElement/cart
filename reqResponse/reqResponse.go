package reqResponse

import (
	"io"
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

func VerifyBody(w http.ResponseWriter, r *http.Request) ([]byte, error) {
	payload, err := io.ReadAll(r.Body)
	if err != nil {
		WriteErr(w, 400, []byte("Error reading Body of Request"))
		return nil, err
	}

	if len(payload) == 0 {
		WriteErr(w, 400, []byte("No Body provided"))
		return nil, err
	}

	return payload, nil
}
