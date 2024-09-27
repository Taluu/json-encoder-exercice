package http

import (
	"encoding/json"
	"log"
	"net/http"

	domain "github.com/Taluu/json-encoder-exercise/pkg"
	"github.com/Taluu/json-encoder-exercise/pkg/di"
)

func init() {
	di.ProvideHTTP("/sign", func() http.Handler {
		return &signServer{di.Invoke[domain.SignatureHandler]()}
	})
}

type signServer struct {
	handler domain.SignatureHandler
}

func (s *signServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	log.Printf("Received HTTP %s /sign\n", r.Method)

	if r.Method != http.MethodPost {
		log.Println("unsupported supported method")
		jsonError(w, "unsupported method", http.StatusMethodNotAllowed)
		return
	}

	var data json.RawMessage
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		log.Printf("could not deserialize body into proper json : %s", err)
		jsonError(w, "json error", http.StatusBadRequest)
		return
	}

	signature, err := s.handler.Create(string(data))
	if err != nil {
		log.Printf("error on creating signature : %s", err)
		jsonError(w, "could not create signature from body", http.StatusInternalServerError)
		return
	}

	response := struct {
		Signature string `json:"signature"`
		Data      any    `json:"data"`
	}{
		Signature: signature,
		Data:      data,
	}

	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(response)

	log.Println("HTTP POST /sign : 200")
}
