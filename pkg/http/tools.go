package http

import (
	"encoding/json"
	"net/http"
)

func jsonError(w http.ResponseWriter, error string, code int) {
	HTTPError := map[string]any{
		"code":  code,
		"error": error,
	}

	result, _ := json.Marshal(HTTPError)

	w.WriteHeader(code)
	w.Write(result)
}
