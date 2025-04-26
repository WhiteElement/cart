package reqResponse

import (
	"errors"
	"io"
	"log"
	"net/http"
)

func Write(w http.ResponseWriter, statusCode int, body []byte) {
	w.WriteHeader(statusCode)
	w.Write(body)
}

func WriteErr(w http.ResponseWriter, statusCode int, msg string) {
	log.Printf("Error: %s\n", msg)
	w.WriteHeader(statusCode)
	w.Write([]byte(msg))
}

func VerifyBody(w http.ResponseWriter, r *http.Request) ([]byte, error) {
	payload, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	if len(payload) == 0 {
		return nil, errors.New("No Body provided")
	}

	return payload, nil
}
