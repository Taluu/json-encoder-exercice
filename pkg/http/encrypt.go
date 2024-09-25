package http

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	domain "github.com/Taluu/json-encoder-exercise/pkg"
	"github.com/Taluu/json-encoder-exercise/pkg/di"
)

func init() {
	di.ProvideHTTP("/encrypt", func() http.Handler {
		return &encryptServer{
			encoder: di.Invoke[domain.Encoder](),
		}
	})
}

type encryptServer struct {
	encoder domain.Encoder
}

func (s *encryptServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	log.Printf("Received HTTP %s /encrypt\n", r.Method)

	if r.Method != http.MethodPost {
		log.Println("unsupported supported method")
		jsonError(w, "unsupported method", http.StatusMethodNotAllowed)
		return
	}

	request := map[string]any{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&request); err != nil && err != io.EOF {
		log.Printf("could not deserialize body into proper json : %s", err)
		jsonError(w, "json error", 400)
		return
	}

	for k := range request {
		var serializedValue string
		// if it's a string, we don't have to reencode as json as it will add
		// "s around the value, which will not encrypt the proper value.
		if _, ok := request[k].(string); !ok {
			jsonReencodedValue, _ := json.Marshal(request[k])
			serializedValue = string(jsonReencodedValue)
		} else {
			serializedValue = request[k].(string)
		}

		encrypted, err := s.encoder.Encode(serializedValue)

		if err != nil {
			log.Printf("could not reencrypt value for %s : %s\n", k, err)
			continue
		}

		request[k] = encrypted
	}

	w.WriteHeader(http.StatusOK)

	encoder := json.NewEncoder(w)
	encoder.Encode(request)

	log.Println("HTTP POST /encrypt : 200")
}
