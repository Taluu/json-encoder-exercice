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
	di.ProvideHTTP("/decrypt", func() http.Handler {
		return &decryptServer{
			encoder: di.Invoke[domain.Encoder](),
		}
	})
}

type decryptServer struct {
	encoder domain.Encoder
}

func (s *decryptServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	log.Printf("Received HTTP %s /decrypt\n", r.Method)

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

	result := s.decodeMap(request)

	w.WriteHeader(http.StatusOK)

	encoder := json.NewEncoder(w)
	encoder.Encode(result)

	log.Println("HTTP POST /decrypt : 200")
}

func (s *decryptServer) decodeMap(object map[string]any) map[string]any {
	result := make(map[string]any, len(object))

	for k := range object {
		switch object[k].(type) {

		case string:
			result[k] = s.decodeString(object[k].(string))

		case []any:
			array := object[k].([]any)

			for i, v := range array {
				if _, ok := v.(string); ok {
					v = s.decodeString(v.(string))
				}

				array[i] = v
			}

			result[k] = array

		case map[string]any:
			result[k] = s.decodeMap(object[k].(map[string]any))

		default:
			// not our problem, nothing to do
			result[k] = object[k]
		}
	}

	return result
}

func (s *decryptServer) decodeString(in string) any {
	var out any
	out = in

	// try to decode value
	if decoded, err := s.encoder.Decode(in); err == nil {
		out = decoded

		// if it was a json, decode it too
		var jsonMap map[string]any
		err := json.Unmarshal([]byte(decoded), &jsonMap)
		if err == nil {
			out = jsonMap
		}
	}

	return out
}
