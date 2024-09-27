package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	domain "github.com/Taluu/json-encoder-exercise/pkg"
	"github.com/Taluu/json-encoder-exercise/pkg/di"
)

func init() {
	di.ProvideHTTP("/verify", func() http.Handler {
		return &verifyServer{di.Invoke[domain.SignatureHandler]()}
	})
}

type verifyServer struct {
	handler domain.SignatureHandler
}

func (v *verifyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	log.Printf("Received HTTP %s /verify\n", r.Method)

	if r.Method != http.MethodPost {
		log.Println("unsupported supported method")
		jsonError(w, "unsupported method", http.StatusMethodNotAllowed)
		return
	}

	var request struct {
		Signature string          `json:"signature"`
		Data      json.RawMessage `json:"data"`
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		log.Printf("could not deserialize body into proper json : %s", err)
		jsonError(w, "json error", http.StatusBadRequest)
		return
	}

	if err := v.handler.Verify(string(request.Data), request.Signature); err != nil {
		log.Printf("error on verifying signature : %s", err)

		if errors.Is(err, domain.ErrInvalidSignature) {
			jsonError(w, "invalid signature", http.StatusBadRequest)
			return
		}

		jsonError(w, "error on checking signature", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	log.Println("HTTP POST /verify : 204")
}
