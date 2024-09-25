package http

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Taluu/json-encoder-exercise/pkg/di"
	_ "github.com/Taluu/json-encoder-exercise/pkg/encoder/base64"
)

func TestEncrypt(t *testing.T) {
	server := di.InvokeNamed[http.Handler]("/encrypt")

	t.Run("failures", func(t *testing.T) {
		type test struct {
			name            string
			method          string
			body            io.Reader
			expectedCode    int
			expectedMessage string
		}

		tests := []test{
			{name: "Wrong HTTP verb", method: "GET", body: nil, expectedCode: http.StatusMethodNotAllowed, expectedMessage: "unsupported method"},
			{name: "Not a json body", method: "POST", body: strings.NewReader("{\"int2\": \"foo}"), expectedCode: http.StatusBadRequest, expectedMessage: "json error"},
		}

		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				r := httptest.NewRequest(tc.method, "/encrypt", tc.body)
				w := httptest.NewRecorder()
				server.ServeHTTP(w, r)

				resp := w.Result()

				if resp.Header.Get("Content-type") != "application/json" {
					t.Fatalf("Expected a application/json content-type, got %s", resp.Header.Get("Content-type"))
				}

				if resp.StatusCode != tc.expectedCode {
					t.Fatalf("Did not expect HTTP %d (%s)", resp.StatusCode, resp.Status)
				}

				error := struct {
					Error string `json:"error"`
				}{}

				decoder := json.NewDecoder(resp.Body)
				decoder.Decode(&error)

				if error.Error != tc.expectedMessage {
					t.Fatalf("Did not expect message %v", error.Error)
				}
			})
		}
	})

	t.Run("success", func(t *testing.T) {
		bodyRequest := strings.NewReader(
			`{
	"foo": "bar",
	"number": 1,
	"object": {
		"one":   "two",
		"three": 3
	}
}`)

		expectedResponse := map[string]string{
			"foo":    "YmFy",
			"number": "MQ==",
			"object": "eyJvbmUiOiJ0d28iLCJ0aHJlZSI6M30=",
		}

		r := httptest.NewRequest("POST", "/encrypt", bodyRequest)
		w := httptest.NewRecorder()
		server.ServeHTTP(w, r)

		resp := w.Result()

		if resp.Header.Get("Content-type") != "application/json" {
			t.Fatalf("Expected a application/json content-type, got %s", resp.Header.Get("Content-type"))
		}

		if resp.StatusCode != 200 {
			t.Fatalf("Did not expect HTTP %d (%s)", resp.StatusCode, resp.Status)
		}

		var gotResponse map[string]string
		decoder := json.NewDecoder(resp.Body)
		decoder.Decode(&gotResponse)

		if fmt.Sprint(gotResponse) != fmt.Sprint(expectedResponse) {
			t.Fatalf("not the expected response, got %v when expecting %v", gotResponse, expectedResponse)
		}
	})
}
